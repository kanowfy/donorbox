package middleware

import (
	"errors"
	"net/http"

	"github.com/kanowfy/donorbox/internal/rcontext"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/internal/token"
	"github.com/kanowfy/donorbox/pkg/httperror"
)

type Auth interface {
	RequireUserAuthentication(next http.HandlerFunc) http.HandlerFunc
	RequireEscrowAuthentication(next http.HandlerFunc) http.HandlerFunc
}

type auth struct {
	userService   service.User
	escrowService service.Escrow
}

func NewAuth(userService service.User, escrowService service.Escrow) Auth {
	return &auth{
		userService,
		escrowService,
	}
}

// RequireUserAuthentication checks for valid JWT token from user request,
// it short circuit the request if the token is missing or malformed.
func (a *auth) RequireUserAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// extract and verify token
		id, err := token.VerifyRequestToken(r)
		if err != nil {
			if errors.Is(err, token.ErrMissingToken) {
				httperror.AuthenticationRequiredResponse(w, r)
			} else {
				httperror.InvalidAuthenticationTokenResponse(w, r)
			}
			return
		}

		user, err := a.userService.GetUserByID(r.Context(), id)
		if err != nil {
			httperror.InvalidCredentialsResponse(w, r)
			return
		}

		r = rcontext.SetUser(r, user)
		next.ServeHTTP(w, r)
	})
}

// RequireUserAuthentication checks for valid JWT token from escrow user request,
// it short circuit the request if the token is missing or malformed.
func (a *auth) RequireEscrowAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := token.VerifyRequestToken(r)
		if err != nil {
			if errors.Is(err, token.ErrMissingToken) {
				httperror.AuthenticationRequiredResponse(w, r)
			} else {
				httperror.InvalidAuthenticationTokenResponse(w, r)
			}
			return
		}

		user, err := a.escrowService.GetEscrowByID(r.Context(), id)
		if err != nil {
			httperror.InvalidCredentialsResponse(w, r)
			return
		}

		r = rcontext.SetEscrowUser(r, user)
		next.ServeHTTP(w, r)

	})
}
