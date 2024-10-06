import React from 'react'
import { SwitchTheme } from '../../utils/switchTheme';

class HealthCheck extends React.Component {
  render() {

    return (
      <div className='flex flex-col justify-center items-center flex-grow'>
        <p className=' text-3xl'>Web is healthy!</p>
        <button className='bg-black dark:bg-white px-4 py-2 text-white dark:text-black' onClick={() => SwitchTheme()}>Switch Theme</button>
      </div>
    )
  }
}

export default HealthCheck;