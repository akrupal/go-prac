package todolist

import (
	"encoding/json"
	"net/http"
	"strconv"

	"todolist/pkg/structs"

	"github.com/go-chi/chi/v5"
)

const (
	MediaTypeJSON = "application/json"
)

type ItemsHandlers struct {
	ItemsService ItemsService
}

func (h *ItemsHandlers) ConfigureRoutes(r chi.Router) {
	r.Route("/todolist", func(r chi.Router) {
		r.Post("/", h.createItem)
		r.Get("/", h.listItems)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.getItem)
			r.Put("/", h.updateItem)
			r.Delete("/", h.deleteItem)

			r.Route("/{newOrder}", func(r chi.Router) {
				r.Put("/", h.reorderItem)
			})
		})
	})
}

func requestAs(r *http.Request, v interface{}) error {
	if r.ContentLength == 0 {
		return nil
	} else { // assume JSON by default
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(v); err != nil {
			return err
		}
	}
	return nil
}

func (h *ItemsHandlers) createItem(w http.ResponseWriter, r *http.Request) {
	var item structs.TodoItem
	err := requestAs(r, &item)
	if err != nil {
		http.Error(w, "Failed", http.StatusBadRequest)
		return
	}

	err = h.ItemsService.AddItem(r.Context(), &item)
	if err != nil {
		http.Error(w, "Failed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *ItemsHandlers) listItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.ItemsService.ListItems(r.Context())
	if err != nil {
		http.Error(w, "Failed", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(items)
}

func (h *ItemsHandlers) deleteItem(w http.ResponseWriter, r *http.Request) {
	deploymentId := chi.URLParam(r, "id")
	err := h.ItemsService.DeleteItem(r.Context(), deploymentId)
	if err != nil {
		http.Error(w, "Failed", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ItemsHandlers) updateItem(w http.ResponseWriter, r *http.Request) {
	deploymentId := chi.URLParam(r, "id")

	var item structs.TodoItem
	err := requestAs(r, &item)
	if err != nil {
		http.Error(w, "Failed", http.StatusBadRequest)
		return
	}

	item.Id = deploymentId

	err = h.ItemsService.UpdateItem(r.Context(), &item)
	if err != nil {
		http.Error(w, "Failed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *ItemsHandlers) getItem(w http.ResponseWriter, r *http.Request) {
	deploymentId := chi.URLParam(r, "id")

	deployment, err := h.ItemsService.GetItem(r.Context(), deploymentId)
	if err != nil {
		http.Error(w, "Failed", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(deployment)
}

func (h *ItemsHandlers) reorderItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	newOrder := chi.URLParam(r, "newOrder")
	inewOrder, err := strconv.Atoi(newOrder)
	if err != nil {
		http.Error(w, "Invalid order", http.StatusBadRequest)
	}
	err = h.ItemsService.ReorderItem(r.Context(), id, inewOrder)
	if err != nil {
		http.Error(w, "Failed to reorder", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
