package handler

type Handlers struct {
	Auth          Auth
	Backing       Backing
	Card          Card
	Escrow        Escrow
	Project       Project
	Transaction   Transaction
	User          User
	ImageUploader ImageUploader
}
