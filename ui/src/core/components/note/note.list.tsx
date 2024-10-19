import Note from "./note";

function NoteList() {
  const array = Array.from({ length: 3 }, (v, i) => i + 1);

  return (
    <div className='flex flex-col gap-2'>
      {array.map(() => {
        return <Note />
      })}
    </div>
  )
}


export default NoteList;