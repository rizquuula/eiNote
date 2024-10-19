import React from 'react'

function FullScreen({ children }: { children: React.ReactElement }) {
  return (
    <div className='min-h-screen flex flex-col justify-between md:mx-20 mx-4 md:py-8 py-2'>
      {children}
    </div>
  )
}

export default FullScreen;