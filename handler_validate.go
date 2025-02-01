package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func cleanString(s string) string {
	words := strings.Fields(s)
	var cleaned []string
	for _, word := range words {
		switch strings.ToLower(word) {
		case "kerfuffle":
			word = "****"
		case "sharbert":
			word = "****"
		case "fornax":
			word = "****"
		}
		cleaned = append(cleaned, word)
	}
	return strings.Join(cleaned, " ")
}

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleanString(params.Body),
	})
}
