import React from 'react'

import { BrowserRouter, Routes, Route } from 'react-router-dom'
import HealthCheck from './pages/health-check/healthcheck'


class Provider extends React.Component {
    constructor(props) {
      super(props)
    }
  
    render() {

      return (
        <BrowserRouter>
          <div className='min-h-screen flex flex-col justify-between bg-white dark:bg-black text-black dark:text-white'>
            <Routes>
              <Route path={"/"} element={<HealthCheck />} />
            </Routes>
          </div>
        </BrowserRouter>
      )
    }
  }
  
  export default Provider