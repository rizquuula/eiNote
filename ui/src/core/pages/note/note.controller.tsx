import { useEffect, useState } from "react";
import NotebookAPI from "../../api/notebook"
import NoteView from "./note.view"
import NotebookPreview from "../../models/notebook.preview";
import { NOTEBOOK_KEY } from "../../models/config";

function checkActiveNotebook(notebooks: NotebookPreview[]): NotebookPreview[] {
  const key = localStorage.getItem(NOTEBOOK_KEY)
  if (key) {
    for (let n of notebooks) {
      if (n.ID === key) {
        n.SetActive(true)
        return notebooks
      }
    }
  }

  // if note with key in localStorage not found, mark the first item in notebooks as active
  if (notebooks.length > 0) {
    notebooks[0].SetActive(true)
    localStorage.setItem(NOTEBOOK_KEY, notebooks[0].ID)
  }
  return notebooks
}

function NoteController({ notebookApi }: { notebookApi: NotebookAPI }) {
  const [notebooks, setNotebooks] = useState<NotebookPreview[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchNotebooks = async () => {
      try {
        let data = await notebookApi.ReadNotebook();
        data = checkActiveNotebook(data)
        setNotebooks(data);
      } catch (err) {
        setError("Failed to fetch notebooks.");
      } finally {
        setLoading(false);
      }
    };

    fetchNotebooks();
  }, [notebookApi]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return <NoteView notebooks={notebooks} />;
}

export default NoteController