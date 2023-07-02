package main

import (
	"net/http"
	"sort"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't retrieve chirps")
	}

	chirps := []Chirp{}

	for _, chirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:   chirp.ID,
			Body: chirp.Body,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	respondWithJson(w, http.StatusOK, chirps)
}
