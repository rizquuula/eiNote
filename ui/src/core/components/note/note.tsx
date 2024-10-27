import NotePreview from '../../models/note.preview';

interface NoteProps {
  data: NotePreview
  onClick(note: NotePreview): void
  isActive: boolean
}

function Note({ data, onClick, isActive }: NoteProps) {

  return (
    <div className='flex flex-row flex-grow'>
      <div className={isActive ? 'w-1 bg-main-normal' : 'w-1 bg-background-accent-hover'}></div>

      <div
        className='flex flex-col bg-background-accent-normal hover:bg-background-accent-hover p-2 flex-grow duration-500 cursor-pointer'
        onClick={() => onClick(data)}>
        <p className=''>{data.Title}</p>
        <p className='text-xs font-thin'>{data.ContentPreview}</p>
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