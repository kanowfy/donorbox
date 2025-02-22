package handler

import (
	"encoding/base64"
	ejson "encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/sharing"
	"github.com/go-playground/validator/v10"
	"github.com/kanowfy/donorbox/internal/config"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/rcontext"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type User interface {
	GetAuthenticatedUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	UpdateAccount(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	GetPendingVerificationUsers(w http.ResponseWriter, r *http.Request)
	UploadDocument(w http.ResponseWriter, r *http.Request)
}

type user struct {
	service      service.User
	validator    *validator.Validate
	cfg          config.Config
	dropboxToken string
}

func NewUser(service service.User, validator *validator.Validate, cfg config.Config, dropboxToken string) User {
	return &user{
		service,
		validator,
		cfg,
		dropboxToken,
	}
}

func (u *user) GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) {
	user := rcontext.GetUser(r)

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"user": user,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (u *user) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	user, err := u.service.GetUserByID(r.Context(), id)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"user": user,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (u *user) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateAccountRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = u.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	user := rcontext.GetUser(r)

	err = u.service.UpdateAccount(r.Context(), user, req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "profile updated successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (u *user) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var req dto.ChangePasswordRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = u.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	user := rcontext.GetUser(r)

	err = u.service.ChangePassword(r.Context(), user.ID, req)
	if err != nil {
		if errors.Is(err, service.ErrWrongPassword) {
			httperror.BadRequestResponse(w, r, err)
		} else {
			httperror.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "password changed successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (u *user) GetPendingVerificationUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.service.GetPendingVerificationUsers(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"users": users,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (u *user) UploadDocument(w http.ResponseWriter, r *http.Request) {
	user := rcontext.GetUser(r)

	r.ParseMultipartForm(10 << 20) // 10MB limit

	file, header, err := r.FormFile("file")
	if err != nil {
		httperror.BadRequestResponse(w, r, fmt.Errorf("error parsing file from request: %w", err))
		return
	}
	defer file.Close()

	fileNameParts := strings.Split(header.Filename, ".")
	ext := fileNameParts[len(fileNameParts)-1]

	// upload file
	dest := fmt.Sprintf("/verification_docs/userdoc_%d_%d.%s", user.ID, time.Now().Unix(), ext)
	commitInfo := files.NewCommitInfo(dest)
	commitInfo.Mode.Tag = "overwrite"

	dbx := files.New(dropbox.Config{
		Token: u.dropboxToken,
	})

	_, err = dbx.Upload(&files.UploadArg{
		CommitInfo: *commitInfo,
	}, file)
	if err != nil {
		if strings.Contains(err.Error(), "expired_access_token") {
			slog.Info("requesting new dropbox access token...")
			if err = u.requestDropboxAccessToken(); err != nil {
				httperror.ServerErrorResponse(w, r, fmt.Errorf("failed to request new access token: %w", err))
			}

			dbx = files.New(dropbox.Config{
				Token: u.dropboxToken,
			})
			_, err = dbx.Upload(&files.UploadArg{
				CommitInfo: *commitInfo,
			}, file)

			if err != nil {
				httperror.ServerErrorResponse(w, r, fmt.Errorf("error uploading document from expired token: %w", err))
			}
		} else {
			httperror.ServerErrorResponse(w, r, fmt.Errorf("error uploading document: %w", err))
			return
		}
	}

	// create share link
	shareClient := sharing.New(dropbox.Config{
		Token: u.dropboxToken,
	})
	sharedLink, err := shareClient.CreateSharedLinkWithSettings(&sharing.CreateSharedLinkWithSettingsArg{
		Path: dest,
	})
	if err != nil {
		if strings.Contains(err.Error(), "shared_link_already_exists") {
			listLinkRes, err := shareClient.ListSharedLinks(&sharing.ListSharedLinksArg{
				Path: dest,
			})

			if err == nil && len(listLinkRes.Links) > 0 {
				sharedLink = listLinkRes.Links[0]
			} else {
				httperror.ServerErrorResponse(w, r, err)
				return
			}
		} else {
			httperror.ServerErrorResponse(w, r, err)
			return
		}
	}

	link := sharedLink.(*sharing.FileLinkMetadata).Url

	if err := u.service.UploadDocument(r.Context(), user.ID, link); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "uploaded verification document",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (u *user) requestDropboxAccessToken() error {
	req, err := http.NewRequest("POST", "https://api.dropbox.com/oauth2/token", nil)
	if err != nil {
		log.Fatal(err)
	}

	base64Auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", u.cfg.DropboxAppKey, u.cfg.DropboxAppSecret)))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64Auth))

	qs := req.URL.Query()
	qs.Add("refresh_token", u.cfg.DropboxRefreshToken)
	qs.Add("grant_type", "refresh_token")

	req.URL.RawQuery = qs.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var payload struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := ejson.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return fmt.Errorf("unable to decode dropbox response: %w", err)
	}

	u.dropboxToken = payload.AccessToken
	return nil
}
