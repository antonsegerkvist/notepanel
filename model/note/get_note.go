package note

import (
	"database/sql"

	"github.com/notepanel/mysql"
)

//
// GetNote returns a single note belonging to a user.
//
func GetNote(noteID, userID uint64) (*ModelNote, error) {

	db, err := mysql.Open()
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(`
		SELECT c_id, c_note, c_create_date, c_modified_date
		FROM t_note
		WHERE c_id = ? AND c_user_id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ret := &ModelNote{}
	err = stmt.QueryRow(noteID, userID).Scan(
		&ret.ID,
		&ret.Note,
		&ret.CreateDate,
		&ret.ModifiedDate,
	)
	if err == sql.ErrNoRows {
		return nil, ErrNoteNotFound
	} else if err != nil {
		return nil, err
	}

	return ret, nil

}
