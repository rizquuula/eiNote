import { BrowserRouter, Route, Routes } from "react-router-dom";
import { WebPage } from "./config/config";
import FullScreen from "./layouts/fullscreen.layout";
import HealthCheck from "./pages/health-check/healthcheck";
import NoteController from "./pages/note/note.controller";
import NotebookAPI from "./api/notebook";
import NoteAPI from "./api/note";


function Provider() {
  const apiHost = process.env.REACT_APP_BACKEND_API ? process.env.REACT_APP_BACKEND_API : ""
  const notebookApi = new NotebookAPI(apiHost, false)
  const noteApi = new NoteAPI(apiHost, false)
  return (
    <BrowserRouter>
      <div className='bg-white dark:bg-black text-black dark:text-white'>
        <Routes>
          <Route path={WebPage.HEALTH_CHECK} element={<FullScreen><HealthCheck /></FullScreen>} />
          <Route path={WebPage.HOME} element={<FullScreen><NoteController notebookApi={notebookApi} noteApi={noteApi} /></FullScreen>} />
        </Routes>
      </div>
    </BrowserRouter>
  )
}

export default Provider;
