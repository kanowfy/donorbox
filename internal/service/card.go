package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/models"
)

var (
	ErrInvalidCardInfo = errors.New("invalid card information")
)

// getCardToken simulates sending card info to external payment gateway then receive a tokenized card object,
// which then get converted to card database model
func (s *Service) getCardToken(ctx context.Context, cardInfo models.CardInformation) (*db.Card, error) {
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

	card, err := s.repository.CreateCard(ctx, db.CreateCardParams{
		Token:          token,
		CardOwnerName:  cardInfo.OwnerName,
		LastFourDigits: cardInfo.Number[len(cardInfo.Number)-4:],
		Brand:          brand,
	})

	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (s *Service) SetupEscrowCard(ctx context.Context, escrowID pgtype.UUID, cardInfo models.CardInformation) error {
	card, err := s.getCardToken(ctx, cardInfo)
	if err != nil {
		return err
	}

	err = s.repository.UpdateEscrowCard(ctx, db.UpdateEscrowCardParams{
		ID:     escrowID,
		CardID: card.ID,
	})

	return err
}

func (s *Service) SetupProjectCard(ctx context.Context, projectID pgtype.UUID, cardInfo models.CardInformation) error {
	card, err := s.getCardToken(ctx, cardInfo)
	if err != nil {
		return err
	}

	err = s.repository.UpdateProjectCard(ctx, db.UpdateProjectCardParams{
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
