package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirpIDString := chi.URLParam(r, "chirpID")
	chirpID, err := strconv.Atoi(chirpIDString)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}

	dbChirp, err := cfg.DB.GetChirp(chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not get chirp")
		return
	}
	respondWithJson(w, http.StatusOK, Chirp{
		ID:   dbChirp.ID,
		Body: dbChirp.Body,
	})

}
