import NoteEditor from "../../components/note/editor"
import NoteList from "../../components/note/note.list"
import NotebookList from "../../components/note/notebook.list"
import NotePreview from "../../models/note.preview"
import NotebookPreview from "../../models/notebook.preview"

interface NoteViewProps {
  notebooks: NotebookPreview[]
  activeNotebook: NotebookPreview | null
  SetActiveNotebook(notebook: NotePreview): void

  notes: NotePreview[]
  activeNote: NotePreview | null
  SetActiveNote(note: NotePreview): void
  SaveNote(note: NotePreview): void
}

function NoteView({
  notebooks,
  activeNotebook,
  SetActiveNotebook,

  notes,
  activeNote,
  SetActiveNote,
  SaveNote,
}: NoteViewProps) {
  return <div className='flex flex-row gap-4'>
    <div id='note.sidebar' className='w-2/12 flex flex-col gap-8'>
      <div id='note' className="flex flex-col gap-2">
        <p className="font-semibold text-main-normal">My Notes</p>
        <NoteList notes={notes} activeNote={activeNote} SetActiveNote={SetActiveNote} />
      </div>
      <div id='notebook' className="flex flex-col gap-2">
        <p className="font-semibold text-main-normal">My Notebooks</p>
        <NotebookList notebooks={notebooks} activeNotebook={activeNotebook} SetActiveNotebook={SetActiveNotebook} />
      </div>
    </div>
    <div id='note.editor' className='flex flex-col gap-2 w-10/12'>
      <p className="font-semibold text-main-normal">Note Editor</p>
      <NoteEditor note={activeNote} SaveNote={SaveNote} />
    </div>
  </div>
}

export default NoteView