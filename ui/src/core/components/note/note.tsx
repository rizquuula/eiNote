import NotePreview from '../../models/note.preview';


function Note({ data }: { data: NotePreview }) {

  return (
    <div className='flex flex-row flex-grow'>
      <div className={data.IsActive ? 'w-1 bg-main-normal' : 'w-1 bg-background-accent-hover'}></div>
      <div className='flex flex-col bg-background-accent-normal hover:bg-background-accent-hover p-2 flex-grow duration-500 cursor-pointer'>
        <p className=''>{data.Title}</p>
        <p className='text-xs font-thin'>{data.Description}</p>
        <div className='flex flex-row justify-between mt-2'>
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