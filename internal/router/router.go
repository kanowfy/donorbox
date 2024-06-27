package router

import (
	"net/http"

	"github.com/kanowfy/donorbox/internal/config"
	"github.com/kanowfy/donorbox/internal/handler"
	"github.com/kanowfy/donorbox/internal/middleware"
)

func Setup(handlers handler.Handlers, authMiddleware middleware.Auth, cfg config.Config) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /healthz", handler.Healthcheck)

	router.HandleFunc("GET /users/authenticated", authMiddleware.RequireUserAuthentication(handlers.User.GetAuthenticatedUser))
	router.HandleFunc("GET /users/{id}", handlers.User.GetUserByID)
	router.HandleFunc("POST /users/register", handlers.Auth.Register)
	router.HandleFunc("POST /users/verify", handlers.Auth.ActivateUser)
	router.HandleFunc("POST /users/login", handlers.Auth.Login)
	router.HandleFunc("PATCH /users", authMiddleware.RequireUserAuthentication(handlers.User.UpdateAccount))
	router.HandleFunc("PATCH /users/password", authMiddleware.RequireUserAuthentication(handlers.User.ChangePassword))
	//	-- router.HandleFunc("GET /users/backings", authMiddleware.RequireUserAuthentication(app.getBackingsForUserHandler))
	router.HandleFunc("GET /users/auth/{provider}", handlers.Auth.StartOAuth)
	router.HandleFunc("GET /users/auth/{provider}/callback", handlers.Auth.GoogleAuthCallback)
	router.HandleFunc("GET /users/auth/{provider}/token", handlers.Auth.GetAuthenticationToken)
	router.HandleFunc("POST /users/password/forgot", handlers.Auth.ForgotPassword)
	router.HandleFunc("POST /users/password/reset", handlers.Auth.ResetPassword)

	router.HandleFunc("GET /escrow/authenticated", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.GetAuthenticatedEscrow))
	router.HandleFunc("POST /escrow/login", handlers.Escrow.Login)
	router.HandleFunc("POST /escrow/{id}/payout", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.Payout))
	router.HandleFunc("POST /escrow/{id}/refund", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.Refund))
	router.HandleFunc("GET /escrow/statistics", authMiddleware.RequireEscrowAuthentication(handlers.Escrow.GetStatistics))

	router.HandleFunc("GET /projects", handlers.Project.GetAllProjects)
	router.HandleFunc("POST /projects/search", handlers.Project.SearchProjects)
	router.HandleFunc("GET /projects/ended", handlers.Project.GetEndedProjects)
	router.HandleFunc("GET /projects/{id}", handlers.Project.GetProjectDetails)
	router.HandleFunc("POST /projects", authMiddleware.RequireUserAuthentication(handlers.Project.CreateProject))
	router.HandleFunc("GET /projects/authenticated", authMiddleware.RequireUserAuthentication(handlers.Project.GetProjectsForUser))
	router.HandleFunc("PATCH /projects/{id}", authMiddleware.RequireUserAuthentication(handlers.Project.UpdateProject))
	router.HandleFunc("DELETE /projects/{id}", authMiddleware.RequireUserAuthentication(handlers.Project.DeleteProject))
	router.HandleFunc("POST /projects/updates", authMiddleware.RequireUserAuthentication(handlers.Project.CreateProjectUpdate))
	router.HandleFunc("GET /projects/{id}/updates", handlers.Project.GetProjectUpdates)

	router.HandleFunc("POST /upload/image", handlers.ImageUploader.UploadImage)

	router.HandleFunc("GET /cards/{id}/project", authMiddleware.RequireUserAuthentication(handlers.Card.GetCard)) //TODO: fix this sh
	router.HandleFunc("POST /escrow/transfer", authMiddleware.RequireEscrowAuthentication(handlers.Card.SetupEscrowCard))
	router.HandleFunc("POST /projects/{id}/transfer", authMiddleware.RequireUserAuthentication(handlers.Card.SetupProjectCard))

	router.HandleFunc("GET /projects/{id}/backings", handlers.Backing.GetBackingsForProject)
	router.HandleFunc("POST /projects/{id}/backings", handlers.Backing.CreateProjectBacking)
	router.HandleFunc("GET /projects/{id}/backings/stats", handlers.Backing.GetProjectBackingStats)

	router.HandleFunc("GET /transactions", authMiddleware.RequireEscrowAuthentication(handlers.Transaction.GetAllTransactions))
	router.HandleFunc("GET /transactions/{id}", authMiddleware.RequireEscrowAuthentication(handlers.Transaction.GetOneTransaction))
	router.HandleFunc("GET /transactions/audit/{id}", authMiddleware.RequireEscrowAuthentication(handlers.Transaction.GetTransactionAudit))

	router.HandleFunc("GET /categories", handlers.Project.GetAllCategories)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	return middleware.EnableCors([]string{cfg.ClientUrl, cfg.DashboardUrl}, middleware.LogRequest(v1))
}
