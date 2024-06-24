package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/model"
)

var (
	ErrInvalidCardInfo = errors.New("invalid card information")
)

type Card interface {
	RequestCardToken(ctx context.Context, cardInfo dto.CardInformation) (*model.Card, error)
	SetupEscrowCard(ctx context.Context, escrowID uuid.UUID, cardInfo dto.CardInformation) error
	SetupProjectCard(ctx context.Context, projectID uuid.UUID, cardInfo dto.CardInformation) error
	GetCardByID(ctx context.Context, cardID uuid.UUID) (*model.Card, error)
}

type card struct {
	repository db.Querier
}

func NewCard(querier db.Querier) Card {
	return &card{
		querier,
	}
}

// RequestCardToken simulates sending card info to external payment gateway then receive a tokenized card object,
// which then get converted to card database model
func (c *card) RequestCardToken(ctx context.Context, cardInfo dto.CardInformation) (*model.Card, error) {
	token, err := generateCardTokenString()
	if err != nil {
		return nil, err
	}

	var brand db.CardBrand
	switch cardInfo.Brand {
	case "VISA":
		brand = db.CardBrandVISA
	case "MASTERCARD":
		brand = db.CardBrandMASTERCARD
	default:
		return nil, ErrInvalidCardInfo
	}

	card, err := c.repository.CreateCard(ctx, db.CreateCardParams{
		Token:          token,
		CardOwnerName:  cardInfo.OwnerName,
		LastFourDigits: cardInfo.Number[len(cardInfo.Number)-4:],
		Brand:          brand,
	})

	if err != nil {
		return nil, err
	}

	return &model.Card{
		ID:             card.ID,
		Token:          card.Token,
		CardOwnerName:  card.CardOwnerName,
		LastFourDigits: card.LastFourDigits,
		CreatedAt:      card.CreatedAt,
	}, nil
}

// Probably want to take user and project ID to check permission
func (c *card) GetCardByID(ctx context.Context, cardID uuid.UUID) (*model.Card, error) {
	card, err := c.repository.GetCardByID(ctx, cardID)
	if err != nil {
		return nil, err
	}

	var brand model.CardBrand
	if card.Brand == db.CardBrandVISA {
		brand = model.VISA
	} else {
		brand = model.MASTERCARD
	}

	return &model.Card{
		ID:             card.ID,
		Token:          card.Token,
		CardOwnerName:  card.CardOwnerName,
		LastFourDigits: card.LastFourDigits,
		Brand:          brand,
		CreatedAt:      card.CreatedAt,
	}, nil
}

func (c *card) SetupEscrowCard(ctx context.Context, escrowID uuid.UUID, cardInfo dto.CardInformation) error {
	card, err := c.RequestCardToken(ctx, cardInfo)
	if err != nil {
		return err
	}

	err = c.repository.UpdateEscrowCard(ctx, db.UpdateEscrowCardParams{
		ID:     escrowID,
		CardID: card.ID,
	})

	return err
}

func (c *card) SetupProjectCard(ctx context.Context, projectID uuid.UUID, cardInfo dto.CardInformation) error {
	card, err := c.RequestCardToken(ctx, cardInfo)
	if err != nil {
		return err
	}

	err = c.repository.UpdateProjectCard(ctx, db.UpdateProjectCardParams{
		ID:     projectID,
		CardID: card.ID,
	})

	return err
}

func generateCardTokenString() (string, error) {
	buffer := make([]byte, 24)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	token := fmt.Sprintf("card_%s", base64.URLEncoding.EncodeToString(buffer)[:24])
	return token, nil

}
