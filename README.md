# NotePanel

## Description

### Backend

There exists five different endpoints. These are as follows:

```http
GET    /api/note/id
PATCH  /api/note/id
GET    /api/note/list
DELETE /api/note
POST   /api/note
```

The first handles getting the content of a single note, the second updates the content of a note and updates the modified date, the third returns a list of all notes, the fourth handles deletions of notes and the last one adds a new note.

All endpoints use json for communications. Both the POST and PATCH endpoints expect an interface of the following structure:

```json
{
  "note": "Note text here"
}
```

And they return the following json structures respectively:

```json
{
  "status": 201,
  "id": 10
}

{
  "status": 202
}
```

### Frontend

The task is broken down into two main components. One toggle button and a panel.
More specifically the panel contains the view logic for the component, and the
subcomponents contains more specialized code for each specific task.

The button and the panel is split so that there will be no invisible div occupying
rest of the view during the closed state. It also allows for customized toggle
buttons.

The panel's subcomponents are a list view and a editor view. The list view contains a list of notes which is fetched from the backend using JSON.

The note list just prints the text of the note content as text and under it it prints when the note was created and the date of last modification. The modification date is only printed if the modification date is defferent from the create date.