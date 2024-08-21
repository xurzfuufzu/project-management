package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"project-management-service/internal/domain"
	"project-management-service/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterRoutes(router chi.Router) {
	router.Route("/users", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Get("/", h.GetAll)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetByID)
			r.Delete("/", h.Delete)
			r.Put("/", h.Update)
			//r.Get("/projects", h.GetUserProjects)
		})
		r.Get("/search", h.Search)
	})
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input service.UserInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := h.userService.Create(context.Background(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAll(r.Context())
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

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.userService.GetByID(context.TODO(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", service.ErrUserNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.userService.Delete(context.TODO(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var input service.UserInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.userService.Update(r.Context(), id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) Search(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	if (name == "" && email == "") || (name != "" && email != "") {
		http.Error(w, "no parameter", http.StatusBadRequest)
		return
	}
	var err error

	if name != "" {
		var users []domain.User
		users, err := h.userService.SearchByName(r.Context(), name)
		if err != nil {
			http.Error(w, "Users not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)

	} else if email != "" {
		var user *domain.User
		user, err = h.userService.SearchByEmail(r.Context(), email)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)

	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

//func (h *UserHandler) GetUserProjects(w http.ResponseWriter, r *http.Request){
//	id := chi.URLParam(r, "id")
//
//	projects, err := h.userService.GetProjectsByUserID(context.TODO(), id)
//	if err != nil{
//		http.Error(w, , http.StatusNotFound)
//	}
//}
