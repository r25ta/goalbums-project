package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"br.com.goalbums/internal/service"
)

type AlbumHandlers struct {
	service *service.AlbumService
}

func NewAlbumHandler(service *service.AlbumService) *AlbumHandlers {
	return &AlbumHandlers{service: service}
}

func (h *AlbumHandlers) GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := h.service.GetAlbums()

	if err != nil {
		http.Error(w, "failed to get albums \n"+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(albums) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)

}

func (h *AlbumHandlers) GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}

	album, err := h.service.GetAlbumByID(id)

	if err != nil {
		http.Error(w, "Failed to get album", http.StatusInternalServerError)
	}

	if album == nil {
		http.Error(w, "Album not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

func (h *AlbumHandlers) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album service.Album

	err := json.NewDecoder(r.Body).Decode(&album)

	if err != nil {
		http.Error(w, "Invalid Request Payload!", http.StatusBadRequest)
		return

	}

	if err := h.service.CreateAlbum(&album); err != nil {
		http.Error(w, "Internal Server Error!", http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)
}

func (h *AlbumHandlers) UpdateAlbum(w http.ResponseWriter, r *http.Request) {

	strId := r.PathValue("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		http.Error(w, "Invalid Album ID", http.StatusBadGateway)
	}

	var album service.Album

	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, "Invalid Request Payload!", http.StatusBadRequest)
		return
	}

	album.ID = id

	err = h.service.UpdateAlbum(&album)

	if err != nil {
		http.Error(w, "Failed to Update Album", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(album)

}

func (h *AlbumHandlers) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		http.Error(w, "Invalid Album ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteAlbum(id); err != nil {
		http.Error(w, "Failed to Delete Album ID "+strId, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
