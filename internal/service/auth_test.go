package service

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/kanowfy/donorbox/internal/db"
	dbmocks "github.com/kanowfy/donorbox/internal/db/mocks"
	"github.com/kanowfy/donorbox/internal/dto"
	mailmocks "github.com/kanowfy/donorbox/internal/mail/mocks"
	"github.com/kanowfy/donorbox/internal/token"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	type dbOutput struct {
		user db.User
		err  error
	}

	type expectedOutput struct {
		hasToken bool
		err      error
	}

	validHashedPassword, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	assert.NoError(t, err)

	cases := []struct {
		name     string
		req      dto.LoginRequest
		dbOut    dbOutput
		expected expectedOutput
	}{
		{
			"valid login",
			dto.LoginRequest{
				Email:    "abc@gmail.com",
				Password: "12345678",
			},
			dbOutput{
				user: db.User{ID: 1, Email: "abc@gmail.com", HashedPassword: string(validHashedPassword)},
				err:  nil,
			},
			expectedOutput{
				hasToken: true,
				err:      nil,
			},
		},
		{
			"invalid email",
			dto.LoginRequest{
				Email:    "abcd@gmail.com",
				Password: "12345678",
			},
			dbOutput{
				user: db.User{},
				err:  pgx.ErrNoRows,
			},
			expectedOutput{
				hasToken: false,
				err:      ErrUserNotFound,
			},
		},
		{
			"invalid password",
			dto.LoginRequest{
				Email:    "abc@gmail.com",
				Password: "12345678910",
			},
			dbOutput{
				user: db.User{ID: 1, Email: "abc@gmail.com", HashedPassword: string(validHashedPassword)},
				err:  nil,
			},
			expectedOutput{
				hasToken: false,
				err:      ErrWrongPassword,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockDB := dbmocks.NewQuerier(t)
			mockMailer := mailmocks.NewMailer(t)

			ctx := context.Background()

			mockDB.EXPECT().GetUserByEmail(ctx, tt.req.Email).Return(tt.dbOut.user, tt.dbOut.err)

			svc := NewAuth(mockDB, mockMailer)

			tok, err := svc.Login(ctx, tt.req)
			assert.ErrorIs(t, err, tt.expected.err)

			if tt.expected.hasToken {
				id, err := token.VerifyToken(tok)
				assert.NoError(t, err)

				assert.Equal(t, tt.dbOut.user.ID, id)
			}
		})
	}

}
