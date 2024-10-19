import { NotePreviewMock } from "../../test/mock/note.preview.mock";
import Note from "./note";

function NoteList() {

  return (
    <div className='flex flex-col gap-2'>
      {NotePreviewMock.map((data, i) => {
        return <Note key={i} data={data} />
      })}
    </div>
  )
}


export default NoteList;