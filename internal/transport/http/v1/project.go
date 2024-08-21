package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"project-management-service/internal/service"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

func (h *ProjectHandler) RegisterRoutes(router chi.Router) {
	router.Route("/projects", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.Create)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetByID)
			//	r.Put("/", h.Update)
			//	r.Delete("/", h.Delete)
			//	r.Get("/tasks", h.GetTasks)
		})
		//r.Get("/search", h.Search)
	})
}

func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input service.ProjectInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	projectID, err := h.projectService.Create(r.Context(), service.ProjectInput{})

	if err != nil {
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": projectID})
}

func (h *ProjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.projectService.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	project, err := h.projectService.GetByID(r.Context(), id)
	if err != nil {
		if err == service.ErrProjectNotFound {
			http.Error(w, "Project not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve project", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		http.Error(w, "Failed to encode project", http.StatusInternalServerError)
	}

}
