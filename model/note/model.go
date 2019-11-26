package note

//
// ModelNote contains the data model for a single note.
//
type ModelNote struct {
	ID           uint64
	Note         string
	CreateDate   string
	ModifiedDate string
}
