import React from 'react'

class FullScreen extends React.Component {
  render() {

    return (
      <div className='min-h-screen flex flex-col justify-between md:mx-20 mx-4 md:py-8 py-2'>
        {this.props.children}
      </div>
    )
  }
}

export default FullScreen;