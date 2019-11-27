package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/notepanel/model/note"
)

func handleNoteListGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	notes, err := note.GetAllNotes(1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"reason":"MODEL_NOTE_GET_ALL_NOTES"}`)
		return
	}

	notesInterface := []map[string]interface{}{}
	for _, v := range *notes {
		notesInterface = append(notesInterface, map[string]interface{}{
			"id":           v.ID,
			"note":         v.Note,
			"createDate":   v.CreateDate,
			"modifiedDate": v.ModifiedDate,
		})
	}

	response := map[string]interface{}{
		"status": 200,
		"notes":  notesInterface,
	}

	jsonBody, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"reason":"JSON_MARSHAL"}`)
		return
	}

	io.WriteString(w, string(jsonBody))

}
