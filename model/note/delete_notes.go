package note

import (
	"strconv"
	"strings"

	"github.com/notepanel/mysql"
)

//
// DeleteNotes removes a list of notes.
//
func DeleteNotes(notes []uint64, userID uint64) error {

	if len(notes) == 0 {
		return nil
	}

	stringNotes := []string{}
	for _, v := range notes {
		stringNotes = append(stringNotes, strconv.FormatUint(v, 10))
	}

	db, err := mysql.Open()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
		DELETE FROM t_note
		WHERE
			c_id IN (` + strings.Join(stringNotes, ",") + `) AND
			c_user_id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}

	return nil

}
