package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-pdf/fpdf"
	"github.com/go-playground/validator/v10"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/rcontext"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Escrow interface {
	Login(w http.ResponseWriter, r *http.Request)
	GetAuthenticatedEscrow(w http.ResponseWriter, r *http.Request)
	ApproveOfProject(w http.ResponseWriter, r *http.Request)
	ResolveMilestone(w http.ResponseWriter, r *http.Request)
	ApproveUserVerification(w http.ResponseWriter, r *http.Request)
	GetStatistics(w http.ResponseWriter, r *http.Request)
	ApproveSpendingProof(w http.ResponseWriter, r *http.Request)
	ReviewReport(w http.ResponseWriter, r *http.Request)
	GenerateReport(w http.ResponseWriter, r *http.Request)
	ResolveDispute(w http.ResponseWriter, r *http.Request)
}

type escrow struct {
	service   service.Escrow
	validator *validator.Validate
}

func NewEscrow(service service.Escrow, validator *validator.Validate) Escrow {
	return &escrow{
		service,
		validator,
	}
}

func (e *escrow) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	token, err := e.service.Login(r.Context(), req)
	if err != nil {
		httperror.InvalidCredentialsResponse(w, r)
		return
	}

	// return token
	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"token": token,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) GetAuthenticatedEscrow(w http.ResponseWriter, r *http.Request) {
	escrow := rcontext.GetEscrowUser(r)

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"user": escrow,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) ApproveOfProject(w http.ResponseWriter, r *http.Request) {
	var req dto.ProjectApprovalRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	escrow := rcontext.GetEscrowUser(r)

	if err := e.service.ApproveOfProject(r.Context(), escrow.ID, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "project approved successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) ResolveMilestone(w http.ResponseWriter, r *http.Request) {
	mid, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	var req dto.ResolveMilestoneRequest

	err = json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	escrow := rcontext.GetEscrowUser(r)

	if err := e.service.ResolveMilestone(r.Context(), escrow.ID, mid, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "confirming milestone resolution",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) ApproveUserVerification(w http.ResponseWriter, r *http.Request) {
	var req dto.VerificationApprovalRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	escrow := rcontext.GetEscrowUser(r)

	if err := e.service.ApproveUserVerification(r.Context(), escrow.ID, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "project approved successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) GetStatistics(w http.ResponseWriter, r *http.Request) {
	stats, err := e.service.GetStatistics(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"statistics": stats,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) ApproveSpendingProof(w http.ResponseWriter, r *http.Request) {
	var req dto.ProofApprovalRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	escrow := rcontext.GetEscrowUser(r)

	//TODO: more nuanced error response
	if err := e.service.ApproveSpendingProof(r.Context(), escrow.ID, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "spending proof approval processed",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) ReviewReport(w http.ResponseWriter, r *http.Request) {
	var req dto.ReportReviewRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	escrow := rcontext.GetEscrowUser(r)

	if err := e.service.ReviewReport(r.Context(), escrow.ID, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return

	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "report review processed",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) GenerateReport(w http.ResponseWriter, r *http.Request) {
	var data dto.DisputedProject

	err := json.ReadJSON(w, r, &data)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	var buf bytes.Buffer
	if err = generatePdfReport(&buf, data); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=Report.pdf")
	w.Header().Set("Content-Length", strconv.FormatInt(int64(buf.Len()), 10))

	if _, err = w.Write(buf.Bytes()); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) ResolveDispute(w http.ResponseWriter, r *http.Request) {
	var req dto.DisputedProjectActionRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	escrow := rcontext.GetEscrowUser(r)

	if err := e.service.ResolveDispute(r.Context(), escrow.ID, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "request processed",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func generatePdfReport(w io.Writer, item dto.DisputedProject) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	/*
		pageWidth, _ := pdf.GetPageSize()
			pdf.ImageOptions("logo.png", 10, 0, 10, 10, false, fpdf.ImageOptions{}, 0, "")
			pdf.Line(0, 10, pageWidth, 10)
	*/
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Disputed Project Report")
	pdf.Ln(15)

	// Project details
	x, y := pdf.GetX(), pdf.GetY()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(50, 10, "Project ID: ")
	pdf.SetFont("Arial", "", 12)
	pdf.SetXY(x+50, y)
	pdf.Cell(50, 10, strconv.FormatInt(item.Project.ID, 10))
	pdf.Ln(10)

	x, y = pdf.GetX(), pdf.GetY()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(50, 10, "Project Title: ")
	pdf.SetFont("Arial", "", 12)
	pdf.SetXY(x+50, y)
	pdf.Cell(50, 10, item.Project.Title)
	pdf.Ln(10)

	x, y = pdf.GetX(), pdf.GetY()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(50, 10, "Link to Project: ")
	pdf.SetFont("Arial", "", 12)
	pdf.SetXY(x+50, y)
	pdf.Cell(50, 10, fmt.Sprintf("http://localhost:4001/fundraiser/%d", item.Project.ID))
	pdf.Ln(10)

	x, y = pdf.GetX(), pdf.GetY()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(50, 10, "Owner: ")
	pdf.SetFont("Arial", "", 12)
	pdf.SetXY(x+50, y)
	pdf.Cell(50, 10, fmt.Sprintf("%s %s", item.User.FirstName, item.User.LastName))
	pdf.Ln(10)

	x, y = pdf.GetX(), pdf.GetY()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(50, 10, "Total Funds Raised: ")
	pdf.SetFont("Arial", "", 12)
	pdf.SetXY(x+50, y)
	pdf.Cell(50, 10, strconv.FormatInt(item.Project.TotalFund, 10))
	pdf.Ln(10)

	x, y = pdf.GetX(), pdf.GetY()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(50, 10, "Total Funds Disputed: ")
	pdf.SetFont("Arial", "", 12)
	pdf.SetXY(x+50, y)
	pdf.Cell(50, 10, strconv.FormatInt(item.Project.TotalFund, 10))
	pdf.Ln(10)

	x, y = pdf.GetX(), pdf.GetY()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(50, 10, "Reason For Dispute: ")
	pdf.SetFont("Arial", "", 12)
	pdf.SetXY(x+50, y)
	if item.IsReported {
		pdf.Cell(50, 10, "Fundraiser has confirmed reports")
	} else {
		pdf.Cell(50, 10, "Fundraiser has not provided valid proof of expenditure")
	}
	pdf.Ln(20)

	if item.IsReported {
		yStart := pdf.GetY()
		pdf.SetFont("Arial", "", 14)
		pdf.Cell(50, 10, "List of reports: ")
		pdf.Ln(10)
		pdf.Line(10, yStart+10, 190, yStart+10)

		for _, r := range item.Reports {
			x, y = pdf.GetX(), pdf.GetY()
			pdf.SetFont("Arial", "B", 11)
			pdf.Cell(50, 10, "Reporter Name: ")
			pdf.SetFont("Arial", "", 11)
			pdf.SetXY(x+50, y)
			pdf.Cell(50, 10, r.FullName)
			pdf.Ln(8)

			x, y = pdf.GetX(), pdf.GetY()
			pdf.SetFont("Arial", "B", 11)
			pdf.Cell(50, 10, "Reason for Report: ")
			pdf.SetFont("Arial", "", 11)
			pdf.SetXY(x+50, y)
			pdf.Cell(50, 10, r.Reason)
			pdf.Ln(8)

			x, y = pdf.GetX(), pdf.GetY()
			pdf.SetFont("Arial", "B", 11)
			pdf.Cell(50, 10, "Relation with Fundraiser: ")
			pdf.SetFont("Arial", "", 11)
			pdf.SetXY(x+50, y)
			if r.Relation != nil {
				pdf.Cell(50, 10, *r.Relation)
			} else {
				pdf.Cell(50, 10, "None")
			}
			pdf.Ln(8)

			x, y = pdf.GetX(), pdf.GetY()
			pdf.SetFont("Arial", "B", 11)
			pdf.Cell(50, 10, "Details of Report: ")
			pdf.SetFont("Arial", "", 11)
			pdf.SetXY(x+50, y)
			pdf.MultiCell(100, 7, r.Details, "", "", false)
			pdf.Ln(8)

			_, y = pdf.GetX(), pdf.GetY()
			pdf.Line(10, y, 190, y)
		}

	} else {
		yStart := pdf.GetY()
		pdf.SetFont("Arial", "", 14)
		pdf.Cell(50, 10, "Policy Violated Milestone(s): ")
		pdf.Ln(10)
		pdf.Line(10, yStart+10, 190, yStart+10)

		pdf.SetFont("Arial", "", 11)
		for _, m := range item.Milestones {
			x, y = pdf.GetX(), pdf.GetY()
			pdf.SetFont("Arial", "B", 11)
			pdf.Cell(50, 10, "Milestone Title: ")
			pdf.SetFont("Arial", "", 11)
			pdf.SetXY(x+50, y)
			pdf.Cell(50, 10, m.Title)
			pdf.Ln(8)

			if m.Description != nil {
				x, y = pdf.GetX(), pdf.GetY()
				pdf.SetFont("Arial", "B", 11)
				pdf.Cell(50, 10, "Milestone Description: ")
				pdf.SetFont("Arial", "", 11)
				pdf.SetXY(x+50, y)
				pdf.Cell(50, 10, *m.Description)
				pdf.Ln(8)
			}

			x, y = pdf.GetX(), pdf.GetY()
			pdf.SetFont("Arial", "B", 11)
			pdf.Cell(50, 10, "Accumulated Fund: ")
			pdf.SetFont("Arial", "", 11)
			pdf.SetXY(x+50, y)
			pdf.Cell(50, 10, formatCurrency(m.CurrentFund))
			pdf.Ln(8)

			x, y = pdf.GetX(), pdf.GetY()
			pdf.SetFont("Arial", "B", 11)
			pdf.Cell(50, 10, "Total Released Fund: ")
			pdf.SetFont("Arial", "", 11)
			pdf.SetXY(x+50, y)
			if m.Completion != nil { // won't be nil, fix disputed to return only refuted milestones
				pdf.Cell(50, 10, formatCurrency(m.Completion.TransferAmount))
			} else {
				pdf.Cell(50, 10, formatCurrency(0))
			}
			pdf.Ln(8)

			_, y = pdf.GetX(), pdf.GetY()
			pdf.Line(10, y, 190, y)
		}
	}

	return pdf.Output(w)
}

func formatCurrency(amount int64) string {
	return fmt.Sprintf("%sVND", message.NewPrinter(language.English).Sprint(amount))
}
