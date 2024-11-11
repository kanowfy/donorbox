package handler

type Handlers struct {
	Auth          Auth
	Backing       Backing
	Escrow        Escrow
	Project       Project
	User          User
	ImageUploader ImageUploader
}
