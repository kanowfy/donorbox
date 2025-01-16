package handler

import (
	"fmt"
	"net/http"
	"strings"

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

	files := r.MultipartForm.File["file"]
	if len(files) == 0 {
		httperror.BadRequestResponse(w, r, fmt.Errorf("no file uploaded"))
		return
	}

	cld, err := cloudinary.NewFromURL(i.cfg.CloudinaryAPIUrl)
	if err != nil {
		httperror.ServerErrorResponse(w, r, fmt.Errorf("failed to initialize Cloudinary: %v", err))
		return
	}

	uris := make([]string, len(files))

	for i, f := range files {
		res, err := cld.Upload.Upload(r.Context(), f, uploader.UploadParams{})
		if err != nil {
			httperror.ServerErrorResponse(w, r, fmt.Errorf("failed to upload image: %v", err))
			return
		}

		uris[i] = res.SecureURL
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"url": strings.Join(uris, ","),
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}
