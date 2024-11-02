import CreateNotebookBtn from "../../components/buttons/notebook.create"
import NoteEditor from "../../components/note-editor/editor"
import NoteList from "../../components/note/note.list"
import NotebookList from "../../components/notebooks/notebook.list"
import NotePreview from "../../models/note.preview"
import NotebookPreview from "../../models/notebook.preview"

interface NoteViewProps {
  notebooks: NotebookPreview[]
  activeNotebook: NotebookPreview | null
  setActiveNotebook(notebook: NotebookPreview): void
  createNotebook(name: string): void
  deleteNotebook(notebook: NotebookPreview): void

  notes: NotePreview[]
  activeNote: NotePreview | null
  setActiveNote(note: NotePreview): void
  saveNote(note: NotePreview): void
}

function NoteView({
  notebooks,
  activeNotebook,
  setActiveNotebook,
  createNotebook,
  deleteNotebook,

  notes,
  activeNote,
  setActiveNote: SetActiveNote,
  saveNote,
}: NoteViewProps) {

  return <div className='flex flex-row gap-4'>
    <div id='note.sidebar' className='w-2/12 flex flex-col gap-8'>
      <div id='note' className="flex flex-col gap-2">
        <p className="font-semibold text-main-normal">My Notes</p>
        <NoteList notes={notes} activeNote={activeNote} SetActiveNote={SetActiveNote} />
      </div>
      <div id='notebook' className="flex flex-col gap-2">
        <p className="font-semibold text-main-normal">My Notebooks</p>
        <NotebookList
          notebooks={notebooks}
          activeNotebook={activeNotebook}
          setActiveNotebook={setActiveNotebook}
          deleteNotebook={deleteNotebook}
        />
        <CreateNotebookBtn CreateNotebook={createNotebook} />
      </div>
    </div>
    <div id='note.editor' className='flex flex-col gap-2 w-10/12'>
      <p className="font-semibold text-main-normal">Note Editor</p>
      <NoteEditor note={activeNote} notebookId={activeNotebook ? activeNotebook.ID : null} saveNote={saveNote} />
    </div>
  </div>
}

export default NoteView