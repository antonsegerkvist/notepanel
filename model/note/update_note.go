package note

import "github.com/notepanel/mysql"

//
// UpdateNote updates the note content.
//
func UpdateNote(note *ModelNote, userID uint64) error {

	db, err := mysql.Open()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
		UPDATE t_note SET c_note = ?, c_modified_date = NOW()
		WHERE c_id = ? AND c_user_id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(note.Note, note.ID, userID)
	if err != nil {
		return err
	}

	return nil

}
