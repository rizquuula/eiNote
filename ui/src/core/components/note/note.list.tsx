import NotePreview from "../../models/note.preview";
import Note from "./note";

interface NoteListProps {
  notes: NotePreview[]
  activeNote: NotePreview | null
  SetActiveNote(note: NotePreview): void
}

function NoteList({ notes, activeNote, SetActiveNote }: NoteListProps) {

  return (
    <div className='flex flex-col gap-2'>
      {notes.map((data, i) => {
        const isActive = data.ID === (
          activeNote
            ? activeNote.ID
            : notes.length > 0
              ? notes[0].ID
              : null
        )
        return <Note key={i} data={data} onClick={SetActiveNote} isActive={isActive} />
      })}

    </div>
  )
}


export default NoteList;