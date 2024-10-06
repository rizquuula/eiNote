import React from 'react'

class Note extends React.Component {
  render() {

    return (
      <div className='flex flex-row'>
        <div className='w-1 bg-green-500'></div>
        <div className='flex flex-col bg-zinc-900 p-2'>
          <p className=''>Note Title</p>
          <p className='text-sm'>Note description</p>
          <p className='text-xs'>Saturday, 01 August 2024</p>
        </div>
      </div>
    )
  }
}

export default Note;