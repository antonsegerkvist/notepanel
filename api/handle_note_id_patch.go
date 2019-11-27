package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/notepanel/model/note"
)

func handleNoteIDPatch(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type RequestBody struct {
		Note string `json:"note"`
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		io.WriteString(w, `{"status":415}`)
		return
	}

	noteIDString := mux.Vars(r)["id"]
	noteID, err := strconv.ParseUint(noteIDString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status":400}`)
		return
	}

	request := &RequestBody{}
	err = json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status":400}`)
		return
	}

	model := &note.ModelNote{
		ID:   noteID,
		Note: request.Note,
	}

	err = note.UpdateNote(model, 1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"reason":"MODEL_NOTE_UPDATE_NOTE"}`)
		return
	}

	io.WriteString(w, `{"status":202}`)

}
