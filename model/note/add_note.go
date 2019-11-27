package note

import "github.com/notepanel/mysql"

//
// AddNote adds a note to the database.
//
func AddNote(note string, userID uint64) (uint64, error) {

	db, err := mysql.Open()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(`
		INSERT INTO t_note (c_user_id, c_note)
		VALUES (?, ?)
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(userID, note)
	if err != nil {
		return 0, err
	}

	ret, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ret), nil

}
