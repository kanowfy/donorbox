package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /healthz", app.healthCheckHandler)

	router.HandleFunc("GET /users/{id}", app.getOneUserHandler)
	router.HandleFunc("POST /register", app.registerAccountHandler)
	router.HandleFunc("PATCH /users", app.updateAccountHandler)
	router.HandleFunc("PATCH /users/password", app.changePasswordHandler)

	router.HandleFunc("GET /projects", app.getAllProjectsHandler)
	router.HandleFunc("GET /projects/{id}", app.getOneProjectHandler)
	router.HandleFunc("POST /projects", app.createProjectHandler)
	router.HandleFunc("PATCH /projects/{id}", app.updateProjectHandler)
	router.HandleFunc("DELETE /projects/{id}", app.deleteProjectHandler)

	router.HandleFunc("POST /project_updates", app.createProjectUpdateHandler)

	router.HandleFunc("POST /project_comments", app.createProjectCommentHandler)

	router.HandleFunc("GET /categories", app.getAllCategoriesHandler)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	return app.requestLogging(v1)
}
