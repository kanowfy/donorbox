package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

type Rag interface {
	AddDocuments(ctx context.Context, req dto.AddDocumentRequest) error
	Query(ctx context.Context, req dto.QueryRequest) (string, error)
}

type rag struct {
	weaviateClient *weaviate.Client
	genModel       *genai.GenerativeModel
	embedModel     *genai.EmbeddingModel
}

func NewRag(weaviateClient *weaviate.Client, genModel *genai.GenerativeModel, embedModel *genai.EmbeddingModel) Rag {
	return &rag{
		weaviateClient,
		genModel,
		embedModel,
	}
}

const promptTemplate = `
You are an assistant for question-answering tasks. 
Assume this context information is factual and correct, as part of internal
documentation.
If the question relates to the context, answer it using the context.
If the question does not relate to the context, do not refer to the context and just say that you don't know.

For example, if the context does mention minerology and I ask you about that,
provide information from the context along with general knowledge.
Use two sentences maximum and keep the answer concise.

Question:
%s

Context:
%s
`

func (r *rag) AddDocuments(ctx context.Context, req dto.AddDocumentRequest) error {
	batch := r.embedModel.NewBatch()
	for _, doc := range req.Documents {
		batch.AddContent(genai.Text(doc.Text))
	}

	res, err := r.embedModel.BatchEmbedContents(ctx, batch)
	if err != nil {
		return fmt.Errorf("adding documents: %w", err)
	}

	if len(res.Embeddings) != len(req.Documents) {
		return fmt.Errorf("mismatch length between embeddings and documents")
	}

	objects := make([]*models.Object, len(req.Documents))
	for i, doc := range req.Documents {
		objects[i] = &models.Object{
			Class: "Document",
			Properties: map[string]any{
				"text": doc.Text,
			},
			Vector: res.Embeddings[i].Values,
		}
	}

	if _, err = r.weaviateClient.Batch().ObjectsBatcher().WithObjects(objects...).Do(ctx); err != nil {
		return fmt.Errorf("storing vector embeddings: %w", err)
	}

	return nil
}

func (r *rag) Query(ctx context.Context, req dto.QueryRequest) (string, error) {
	res, err := r.embedModel.EmbedContent(ctx, genai.Text(req.Content))
	if err != nil {
		return "", fmt.Errorf("embedding content: %w", err)
	}

	gql := r.weaviateClient.GraphQL()
	result, err := gql.Get().WithNearVector(gql.NearVectorArgBuilder().WithVector(res.Embedding.Values)).WithClassName("Document").WithFields(graphql.Field{
		Name: "text",
	}).WithLimit(2).Do(ctx)
	if err != nil {
		return "", fmt.Errorf("weaviate graphql get: %w", err)
	}

	if len(result.Errors) != 0 {
		var errs []string
		for _, e := range result.Errors {
			errs = append(errs, e.Message)
		}

		return "", fmt.Errorf("weaviate graphql get: %v", errs)
	}

	contents, err := decodeWeaviateResult(result)
	if err != nil {
		return "", fmt.Errorf("reading weaviate result: %w", err)
	}

	query := fmt.Sprintf(promptTemplate,
		req.Content, strings.Join(contents, "\n"))
	ans, err := r.genModel.GenerateContent(ctx, genai.Text(query))
	if err != nil {
		log.Printf("error calling generative model: %v", err)
		return "", fmt.Errorf("generative model: %w", err)
	}

	if len(ans.Candidates) != 1 {
		return "", fmt.Errorf("generative model error")
	}

	var responseText []string
	for _, part := range ans.Candidates[0].Content.Parts {
		if pt, ok := part.(genai.Text); ok {
			responseText = append(responseText, string(pt))
		} else {
			log.Printf("invalid part type of part %q", pt)
			return "", fmt.Errorf("generative model error")
		}
	}

	return strings.Join(responseText, "\n"), nil
}

func decodeWeaviateResult(result *models.GraphQLResponse) ([]string, error) {
	data, ok := result.Data["Get"]
	if !ok {
		return nil, fmt.Errorf("no get key found in result")
	}

	doc, ok := data.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("unexpected key type")
	}

	resultList, ok := doc["Document"].([]any)
	if !ok {
		return nil, fmt.Errorf("invalid document type")
	}

	var out []string
	for _, rl := range resultList {
		m, ok := rl.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid element in list of documents")
		}
		rl, ok := m["text"].(string)
		if !ok {
			return nil, fmt.Errorf("expected string in list of documents")
		}
		out = append(out, rl)
	}

	return out, nil
}
