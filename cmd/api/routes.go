package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /healthz", app.healthCheckHandler)

	router.HandleFunc("GET /users/authenticated", app.requireUserAuthentication(app.getAuthenticatedUserHandler))
	router.HandleFunc("GET /users/{id}", app.getUserByIDHandler)
	router.HandleFunc("POST /users/register", app.registerAccountHandler)
	router.HandleFunc("POST /verify", app.activateUserHandler) // change when incorporate frontend
	router.HandleFunc("POST /users/login", app.loginHandler)
	router.HandleFunc("PATCH /users", app.requireUserAuthentication(app.updateAccountHandler))
	router.HandleFunc("PATCH /users/password", app.requireUserAuthentication(app.changePasswordHandler))
	router.HandleFunc("GET /users/backings", app.requireUserAuthentication(app.getBackingsForUserHandler))
	router.HandleFunc("GET /users/auth/google", app.startGoogleAuthHandler)
	router.HandleFunc("GET /users/auth/google/callback", app.googleAuthCallbackHandler)
	router.HandleFunc("GET /users/logout/google", app.googleAuthLogoutHandler)

	router.HandleFunc("GET /projects", app.getAllProjectsHandler)
	router.HandleFunc("POST /projects/search", app.searchProjectsHandler)
	router.HandleFunc("GET /projects/{id}", app.getOneProjectHandler)
	router.HandleFunc("POST /projects", app.requireUserAuthentication(app.createProjectHandler))
	router.HandleFunc("PATCH /projects/{id}", app.requireUserAuthentication(app.updateProjectHandler))
	router.HandleFunc("DELETE /projects/{id}", app.requireUserAuthentication(app.deleteProjectHandler))
	router.HandleFunc("POST /projects/updates", app.requireUserAuthentication(app.createProjectUpdateHandler))
	router.HandleFunc("POST /projects/comments", app.requireUserAuthentication(app.createProjectCommentHandler))

	router.HandleFunc("GET /projects/{id}/backings", app.getBackingsForProjectHandler)
	router.HandleFunc("POST /projects/{id}/backings", app.requireUserAuthentication(app.createProjectBackingHandler))
	router.HandleFunc("GET /projects/{id}/backings/stats", app.getProjectBackingStats)

	router.HandleFunc("GET /transactions", app.requireEscrowAuthentication(app.getAllTransactionsHandler))
	router.HandleFunc("GET /transactions/{id}", app.requireEscrowAuthentication(app.getOneTransactionHandler))
	router.HandleFunc("GET /transactions/audit/{id}", app.requireEscrowAuthentication(app.getTransactionAuditHandler))

	router.HandleFunc("GET /categories", app.getAllCategoriesHandler)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	return app.enableCors(app.requestLogging(v1))
}
