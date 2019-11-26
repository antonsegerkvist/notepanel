package note

import (
	"database/sql"

	"github.com/notepanel/mysql"
)

//
// GetAllNotes returns a list of all notes accessable to a user.
//
func GetAllNotes(userID uint64) (*[]ModelNote, error) {

	db, err := mysql.Open()
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(`
		SELECT c_id, c_note, c_create_date, c_modified_date
		FROM t_note
		WHERE c_user_id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err == sql.ErrNoRows {
		return &[]ModelNote{}, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret := &[]ModelNote{}
	for rows.Next() {
		buffer := ModelNote{}
		err = rows.Scan(
			&buffer.ID,
			&buffer.Note,
			&buffer.CreateDate,
			&buffer.ModifiedDate,
		)
		if err != nil {
			return nil, err
		}
		*ret = append(*ret, buffer)
	}

	return ret, nil

}
