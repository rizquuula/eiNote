import React from 'react'

import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { WebPage } from './models/config'
import HealthCheck from './pages/health-check/healthcheck'
import NoteController from './pages/note/note.controller'
import FullScreen from './layouts/fullscreen.layout'


class Provider extends React.Component {
  render() {
    return (
      <BrowserRouter>
        <div className='bg-white dark:bg-black text-black dark:text-white'>
          <Routes>
            <Route path={WebPage.HEALTH_CHECK} element={<FullScreen><HealthCheck /></FullScreen>} />
            <Route path={WebPage.HOME} element={<FullScreen><NoteController /></FullScreen>} />
          </Routes>
        </div>
      </BrowserRouter>
    )
  }
}

export default Provider