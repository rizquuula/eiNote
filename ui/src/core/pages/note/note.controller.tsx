import { useEffect, useState } from "react";
import NotebookAPI from "../../api/notebook"
import NoteView from "./note.view"
import NotebookPreview from "../../models/notebook.preview";
import NoteAPI from "../../api/note";
import NotePreview from "../../models/note.preview";


function NoteController({ notebookApi, noteApi }: { notebookApi: NotebookAPI, noteApi: NoteAPI }) {
  const [notebooks, setNotebooks] = useState<NotebookPreview[]>([]);
  const [activeNotebook, setActiveNotebook] = useState<NotebookPreview | null>(null);

  const [notes, setNotes] = useState<NotePreview[]>([]);
  const [activeNote, setActiveNote] = useState<NotePreview | null>(null);

  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const notebooksPreview = await notebookApi.ReadNotebooks();

        if (activeNotebook == null && notebooksPreview.length > 0) {
          setActiveNotebook(notebooksPreview[0])
        }

        setNotebooks(notebooksPreview);

        const notesPreview = await noteApi.ReadNotes(activeNotebook);

        if (activeNote == null && notesPreview.length > 0) {
          setActiveNote(notesPreview[0])
        }

        setNotes(notesPreview)
      } catch (err) {
        setError(`Failed to fetch data. ${err}`);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [notebookApi, noteApi, activeNotebook, activeNote]);

  const saveNote = async (note: NotePreview) => {
    await noteApi.CreateNote(note)
    setActiveNote(note)
  }

  const createNotebook = async (name: string) => {
    const newNotebook = await notebookApi.CreateNotebook(new NotebookPreview("", name, null))
    setNotebooks([...notebooks, newNotebook])
  }

  const deleteNotebook = async (notebook: NotebookPreview) => {
    await notebookApi.DeleteNotebook(notebook)
    let currNotebooks = []
    for (let n of notebooks) {
      if (n.ID !== notebook.ID) {
        currNotebooks.push(n)
      }
    }
    setNotebooks(currNotebooks)
  }

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return <NoteView
    notebooks={notebooks}
    activeNotebook={activeNotebook}
    setActiveNotebook={setActiveNotebook}
    createNotebook={createNotebook}
    notes={notes}
    activeNote={activeNote}
    setActiveNote={setActiveNote}
    saveNote={saveNote}
    deleteNotebook={deleteNotebook}
  />;
}

export default NoteController