package comments

import (
	"comments-api/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	comments, err := h.service.GetAllComments()
	if err != nil {
		utils.SendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input NewComment
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.SendJSONError(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	comment, err := h.service.CreateComment(input)
	if err != nil {
		utils.SendJSONError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendJSONError(w, "El ID debe ser un número entero", http.StatusBadRequest)
		return
	}

	comment, err := h.service.GetCommentById(id)
	if err != nil {
		utils.SendJSONError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendJSONError(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.service.RemoveComment(id); err != nil {
		utils.SendJSONError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Comentario eliminado"})
}