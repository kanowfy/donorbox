package dto

type Document struct {
	Text string `json:"text"`
}

type AddDocumentRequest struct {
	Documents []Document `json:"documents"`
}

type QueryRequest struct {
	Content string `json:"content"`
}
