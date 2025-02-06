package router

import (
	"net/http"

	"github.com/kanowfy/donorbox/internal/config"
	"github.com/kanowfy/donorbox/internal/handler"
	"github.com/kanowfy/donorbox/internal/middleware"
)

// Setup configure the endpoints for the handlers and returns a router.
func Setup(handlers handler.Handlers, authMiddleware middleware.Auth, cfg config.Config) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /healthz", handler.Healthcheck)

	router.HandleFunc("GET /users/authenticated", authMiddleware.RequireUserAuthentication(handlers.User.GetAuthenticatedUser))
	router.HandleFunc("GET /users/{id}", handlers.User.GetUserByID)
	router.HandleFunc("POST /users/register", handlers.Auth.Register)
	router.HandleFunc("POST /users/verify", handlers.Auth.ActivateUser) //TODO: change this to /users/activate
	router.HandleFunc("POST /users/login", handlers.Auth.Login)
	router.HandleFunc("PATCH /users", authMiddleware.RequireUserAuthentication(handlers.User.UpdateAccount))
	router.HandleFunc("PATCH /users/password", authMiddleware.RequireUserAuthentication(handlers.User.ChangePassword))
	//	-- router.HandleFunc("GET /users/backings", authMiddleware.RequireUserAuthentication(app.getBackingsForUserHandler))
	router.HandleFunc("GET /users/auth/{provider}", handlers.Auth.StartOAuth)
	router.HandleFunc("GET /users/auth/{provider}/callback", handlers.Auth.GoogleAuthCallback)
	router.HandleFunc("GET /users/auth/{provider}/token", handlers.Auth.GetAuthenticationToken)
	router.HandleFunc("POST /users/password/forgot", handlers.Auth.ForgotPassword)
	router.HandleFunc("POST /users/password/reset", handlers.Auth.ResetPassword)
	router.HandleFunc("POST /users/uploadDocument", authMiddleware.RequireUserAuthentication(handlers.User.UploadDocument))
	router.HandleFunc("GET /users/pendingVerification", authMiddleware.RequireEscrowAuthentication(handlers.User.GetPendingVerificationUsers))
	router.HandleFunc("GET /users/backings", authMiddleware.RequireUserAuthentication(handlers.Backing.GetBackingsForUser))

	router.HandleFunc("GET /escrow/authenticated", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.GetAuthenticatedEscrow))
	router.HandleFunc("POST /escrow/login", handlers.Escrow.Login)
	router.HandleFunc("POST /escrow/register", handlers.Auth.RegisterEscrow)
	router.HandleFunc("POST /escrow/approve/project", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.ApproveOfProject))
	router.HandleFunc("POST /escrow/approve/verification", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.ApproveUserVerification))
	router.HandleFunc("POST /escrow/approve/proof", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.ApproveSpendingProof))
	router.HandleFunc("POST /escrow/resolve/{id}", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.ResolveMilestone))
	router.HandleFunc("POST /escrow/review/report", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.ReviewReport))
	router.HandleFunc("GET /escrow/statistics", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.GetStatistics))
	router.HandleFunc("POST /escrow/resolve/dispute", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.ResolveDispute))

	router.HandleFunc("GET /projects", handlers.Project.GetAllProjects)
	router.HandleFunc("POST /projects/search", handlers.Project.SearchProjects)
	router.HandleFunc("GET /projects/ended", handlers.Project.GetEndedProjects)
	router.HandleFunc("GET /projects/pending", authMiddleware.RequireEscrowAuthentication(handlers.Project.GetPendingProjects))
	router.HandleFunc("GET /projects/{id}", handlers.Project.GetProjectDetails)
	router.HandleFunc("POST /projects", authMiddleware.RequireUserAuthentication(handlers.Project.CreateProject))
	router.HandleFunc("GET /projects/authenticated", authMiddleware.RequireUserAuthentication(handlers.Project.GetProjectsForUser))
	router.HandleFunc("PATCH /projects/{id}", authMiddleware.RequireUserAuthentication(handlers.Project.UpdateProject))
	router.HandleFunc("DELETE /projects/{id}", authMiddleware.RequireUserAuthentication(handlers.Project.DeleteProject))
	router.HandleFunc("GET /projects/disputed", authMiddleware.RequireEscrowAuthentication(handlers.Project.GetDisputedProjects))

	router.HandleFunc("POST /projects/{id}/reports", handlers.Project.CreateProjectReport)
	router.HandleFunc("GET /projects/reports", authMiddleware.RequireEscrowAuthentication(handlers.Project.GetProjectReports))

	router.HandleFunc("GET /milestones/funded", authMiddleware.RequireEscrowAuthentication(handlers.Project.GetFundedMilestones))
	router.HandleFunc("POST /milestones/proofs", authMiddleware.RequireUserAuthentication(handlers.Project.CreateMilestoneProof))

	router.HandleFunc("POST /upload/image", handlers.ImageUploader.UploadImage)

	router.HandleFunc("POST /projects/paymentIntent", handlers.Backing.CreatePaymentIntent)
	router.HandleFunc("GET /projects/{id}/backings", handlers.Backing.GetBackingsForProject)
	router.HandleFunc("POST /projects/{id}/backings", handlers.Backing.CreateProjectBacking)
	router.HandleFunc("GET /projects/{id}/backings/stats", handlers.Backing.GetProjectBackingStats)

	router.HandleFunc("GET /categories", handlers.Project.GetAllCategories)
	router.HandleFunc("GET /categories/{name}", handlers.Project.GetCategoryByName)

	router.HandleFunc("GET /notifications/{id}", handlers.Notification.GetNotificationsForUser)
	router.HandleFunc("POST /notifications/{id}/read", handlers.Notification.UpdateReadNotification)
	router.HandleFunc("GET /notifications/events", handlers.Notification.NotificationStreamHandler)

	router.HandleFunc("POST /rag/documents", handlers.Rag.AddDocuments)
	router.HandleFunc("POST /rag/query", handlers.Rag.Query)

	router.HandleFunc("GET /audits", authMiddleware.RequireEscrowAuthentication(handlers.AuditTrail.GetAuditHistory))

	router.HandleFunc("POST /reports/generate", handlers.Escrow.GenerateReport)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	return middleware.EnableCors([]string{cfg.ClientUrl, cfg.DashboardUrl}, middleware.LogRequest(v1))
}
