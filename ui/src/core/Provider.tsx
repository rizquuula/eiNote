import { BrowserRouter, Route, Routes } from "react-router-dom";
import { WebPage } from "./models/config";
import FullScreen from "./layouts/fullscreen.layout";
import HealthCheck from "./pages/health-check/healthcheck";
import NoteController from "./pages/note/note.controller";
import NotebookAPI from "./api/notebook";


function Provider() {
  const notebookApi = new NotebookAPI("http://localhost:2001", false)
  return (
    <BrowserRouter>
      <div className='bg-white dark:bg-black text-black dark:text-white'>
        <Routes>
          <Route path={WebPage.HEALTH_CHECK} element={<FullScreen><HealthCheck /></FullScreen>} />
          <Route path={WebPage.HOME} element={<FullScreen><NoteController notebookApi={notebookApi} /></FullScreen>} />
        </Routes>
      </div>
    </BrowserRouter>
  )
}

export default Provider;
