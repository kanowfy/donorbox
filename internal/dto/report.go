package dto

type CreateProjectReportRequest struct {
	Email       string  `json:"email" validate:"required,email"`
	FullName    string  `json:"full_name" validate:"required"`
	PhoneNumber string  `json:"phone_number" validate:"required,min=9,max=11"`
	Relation    *string `json:"relation,omitempty" validate:"omitnil"`
	Reason      string  `json:"reason" validate:"required"`
	Details     string  `json:"details" validate:"required"`
	ProofImage  *string `json:"proof_image,omitempty" validate:"omitnil"`
}
