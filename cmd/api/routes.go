package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /healthz", app.healthCheckHandler)

	router.HandleFunc("GET /users/{id}", app.getOneUserHandler)
	router.HandleFunc("POST /users/register", app.registerAccountHandler)
	router.HandleFunc("POST /users/login", app.loginHandler)
	router.HandleFunc("PATCH /users", app.requireUserAuthentication(app.updateAccountHandler))
	router.HandleFunc("PATCH /users/password", app.requireUserAuthentication(app.changePasswordHandler))

	router.HandleFunc("GET /projects", app.getAllProjectsHandler)
	router.HandleFunc("GET /projects/{id}", app.getOneProjectHandler)
	router.HandleFunc("POST /projects", app.requireUserAuthentication(app.createProjectHandler))
	router.HandleFunc("PATCH /projects/{id}", app.requireUserAuthentication(app.updateProjectHandler))
	router.HandleFunc("DELETE /projects/{id}", app.requireUserAuthentication(app.deleteProjectHandler))

	router.HandleFunc("POST /project-updates", app.requireUserAuthentication(app.createProjectUpdateHandler))

	router.HandleFunc("POST /project-comments", app.requireUserAuthentication(app.createProjectCommentHandler))

	router.HandleFunc("GET /categories", app.getAllCategoriesHandler)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	return app.requestLogging(v1)
}
