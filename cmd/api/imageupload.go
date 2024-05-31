package main

import (
	"fmt"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (app *application) uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10MB limit

	file, _, err := r.FormFile("file")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	defer file.Close()

	cld, err := cloudinary.NewFromURL(app.config.CloudinaryAPIUrl)
	if err != nil {
		app.serverErrorResponse(w, r, fmt.Errorf("failed to initialize Cloudinary: %v", err))
		return
	}

	res, err := cld.Upload.Upload(r.Context(), file, uploader.UploadParams{})
	if err != nil {
		app.serverErrorResponse(w, r, fmt.Errorf("failed to upload image: %v", err))
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"url": res.SecureURL,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
