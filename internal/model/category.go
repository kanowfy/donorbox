package model

type Category struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CoverPicture string `json:"cover_picture"`
}
