package api

import (
	"encoding/json"
	"github.com/notepanel/model/note"
	"io"
	"net/http"
)

func handleNoteDelete(w http.ResponseWriter, r *http.Request) {

	type RequestBody struct {
		Notes []uint64 `json:"ids"`
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

	err = note.DeleteNotes(request.Notes, 1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"reason":"MODEL_NOTE_DELETE_NOTES"}`)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	io.WriteString(w, `{"status":202}`)

}
