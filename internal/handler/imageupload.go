package handler

import (
	"fmt"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/kanowfy/donorbox/internal/config"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type ImageUploader interface {
	UploadImage(w http.ResponseWriter, r *http.Request)
}

type imageUploader struct {
	cfg config.Config
}

func NewImageUploader(cfg config.Config) ImageUploader {
	return &imageUploader{
		cfg,
	}
}

func (i *imageUploader) UploadImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10MB limit

	file, _, err := r.FormFile("file")
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}
	defer file.Close()

	cld, err := cloudinary.NewFromURL(i.cfg.CloudinaryAPIUrl)
	if err != nil {
		httperror.ServerErrorResponse(w, r, fmt.Errorf("failed to initialize Cloudinary: %v", err))
		return
	}

	res, err := cld.Upload.Upload(r.Context(), file, uploader.UploadParams{})
	if err != nil {
		httperror.ServerErrorResponse(w, r, fmt.Errorf("failed to upload image: %v", err))
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"url": res.SecureURL,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}
