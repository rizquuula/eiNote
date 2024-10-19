import Notepreview from '../../models/note.preview';

function Note() {
  const data = new Notepreview(
    "Title Example",
    "Note description this can be expanded up to 3 lines, after that begin blank blank blank, like this",
    new Date(),
  )

  return (
    <div className='flex flex-row flex-grow'>
      <div className='w-1 bg-green-500'></div>
      <div className='flex flex-col bg-zinc-900 p-2 flex-grow'>
        <p className=''>{data.Title}</p>
        <p className='text-sm'>{data.Description}</p>
        <div className='flex flex-row justify-between'>
          <div className='flex flex-row gap-2'>
            <p className='text-xs'>Saturday</p>
          </div>
          <div className='flex flex-row gap-1'>
            <p className='text-xs'>{data.GetDate()}</p>
            <p className='text-xs'>{data.GetTime()}</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Note;