package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /healthz", app.healthCheckHandler)

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

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"status":  "up",
		"version": "0.0.1",
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
