package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Chirp struct {
	Body string `json:"body"`
}

type ChirpError struct {
	Error string `json:"error"`
}

type ChirpResp struct {
	CleanedBody string `json:"cleaned_body"`
}

func ValidateChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	chirp := Chirp{}

	w.Header().Set("Content-Type", "application/json")

	if err := decoder.Decode(&chirp); err != nil {
		log.Printf("Error decoding chirp: %s", err)
		_ = WriteError(w, 500, "Something went wrong")
		return
	}

	if len(chirp.Body) > 140 {
		if err := WriteError(w, 400, "Chirp is too long"); err != nil {
			panic(err)
		}
		return
	}

	validatedString := ValidateMessage(chirp.Body)

	if err := RespondWithJson(w, http.StatusOK, ChirpResp{CleanedBody: validatedString}); err != nil {
		panic(err)
	}
}

func WriteError(w http.ResponseWriter, status int, errMessage string) error {
	return RespondWithJson(w, status, ChirpError{Error: errMessage})
}

func RespondWithJson(w http.ResponseWriter, status int, payload interface{}) error {
	body, err := json.Marshal(&payload)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	w.Write(body)
	return nil
}

func ValidateMessage(message string) string {
	badWordsMap := make(map[string]bool)
	badWordsMap["kerfuffle"] = true
	badWordsMap["sharbert"] = true
	badWordsMap["fornax"] = true

	splited := strings.Split(message, " ")
	for i := range splited {
		if _, ok := badWordsMap[strings.ToLower(splited[i])]; ok {
			splited[i] = "****"
		}
	}
	return strings.Join(splited, " ")
}
