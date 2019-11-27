package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/notepanel/model/note"
)

func handleNoteIDGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	noteIDString := mux.Vars(r)["id"]
	noteID, err := strconv.ParseUint(noteIDString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status":400}`)
		return
	}

	model, err := note.GetNote(noteID, 1)
	if err == note.ErrNoteNotFound {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `{"status":404}`)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"reason":"MODEL_NOTE_GET_NOTE"}`)
		return
	}

	result := map[string]interface{}{
		"status": 200,
		"note": map[string]interface{}{
			"id":           model.ID,
			"note":         model.Note,
			"createDate":   model.CreateDate,
			"modifiedDate": model.ModifiedDate,
		},
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"reason":"JSON_MARSHAL"}`)
		return
	}

	io.WriteString(w, string(jsonBytes))

}
