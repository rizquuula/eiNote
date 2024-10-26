import NoteEditor from "../../components/note/editor"
import NoteList from "../../components/note/note.list"
import NotebookList from "../../components/note/notebook.list"
import NotebookPreview from "../../models/notebook.preview"

function NoteView({ notebooks }: { notebooks: NotebookPreview[] }) {
  return <div id='note' className='flex flex-row gap-4'>
    <div id='note.sidebar' className='w-2/12 flex flex-col gap-8'>
      <div id='note' className="flex flex-col gap-2">
        <p className="font-semibold text-main-normal">My Notes</p>
        <NoteList />
      </div>
      <div id='notebook' className="flex flex-col gap-2">
        <p className="font-semibold text-main-normal">My Notebooks</p>
        <NotebookList notebooks={notebooks} />
      </div>
    </div>
    <div id='note.editor' className='flex flex-col gap-2 w-10/12'>
      <p className="font-semibold text-main-normal">Note Editor</p>
      <NoteEditor />
    </div>
  </div>
}

export default NoteView