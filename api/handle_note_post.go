package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/notepanel/model/note"
)

func handleNotePost(w http.ResponseWriter, r *http.Request) {

	type RequestBody struct {
		Note string `json:"note"`
	}

	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		io.WriteString(w, `{"status":415}`)
		return
	}

	request := &RequestBody{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status":400}`)
		return
	}

	noteID, err := note.AddNote(request.Note, 1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"reason":"MODEL_NOTE_ADD_NOTE"}`)
		return
	}

	response := map[string]interface{}{
		"status": http.StatusCreated,
		"id":     noteID,
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"reason":"JSON_MARSHAL"}`)
		return
	}

	io.WriteString(w, string(jsonBytes))

}
