package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"maker-checker/internal/service"
)

type MessageHandler struct {
	svc service.MessageService
}

func NewMessageHandler(svc service.MessageService) *MessageHandler {
	return &MessageHandler{svc: svc}
}

func (h *MessageHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Content   string `json:"content"`
		Recipient string `json:"recipient"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := r.Header.Get("X-User-Id")
	if user == "" {
		user = "default_maker"
	}

	msg, err := h.svc.CreateMessage(context.Background(), req.Content, req.Recipient, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func (h *MessageHandler) Approve(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	id := parts[2]

	checker := r.Header.Get("X-User-Id")
	if checker == "" {
		checker = "default_checker"
	}

	if err := h.svc.ApproveMessage(context.Background(), id, checker); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *MessageHandler) Reject(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	id := parts[2]

	checker := r.Header.Get("X-User-Id")
	if checker == "" {
		checker = "default_checker"
	}

	if err := h.svc.RejectMessage(context.Background(), id, checker); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *MessageHandler) ListPending(w http.ResponseWriter, r *http.Request) {
	msgs, err := h.svc.ListPending(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msgs)
}
