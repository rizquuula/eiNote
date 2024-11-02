import NotePreview from "../models/note.preview"
import NotebookPreview from "../models/notebook.preview"
import { NotePreviewMock } from "../test/mock/note.preview.mock"

class NoteAPI {
  Host: string
  UseMock: boolean

  constructor(host: string, mock: boolean = false) {
    this.Host = host
    this.UseMock = mock
  }

  async ReadNotes(notebook: NotebookPreview | null): Promise<NotePreview[]> {
    if (this.UseMock) {
      return NotePreviewMock
    }

    const requestOptions = {
      method: "GET",
    };

    const queryParams = {
      notebook_id: notebook ? notebook.ID : "",
    };

    const url = new URL(this.Host + "/v1/notes");
    url.search = new URLSearchParams(queryParams).toString();

    const response = await fetch(url, requestOptions);

    // Check if the response is OK (status in the range 200-299)
    if (!response.ok) {
      throw new Error(`HTTP error! url: ${url}, status: ${response.status}`);
    }

    const result = await response.json();
    const notebooks: NotePreview[] = [];
    const data = result["data"]["notes"];

    for (let n of data) {
      const parsedDate = new Date(n["updated_at"].replace(" ", "T"));

      notebooks.push(
        new NotePreview(
          n["id"],
          n["notebook_id"],
          n["title"],
          n["content"],
          parsedDate,
        )
      );
    }
    return notebooks;
  }

  async SaveNote(note: NotePreview): Promise<NotePreview> {

    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    const raw = JSON.stringify({
      "id": note.ID,
      "notebook_id": note.NotebookID,
      "content": note.Content,
    });

    const requestOptions = {
      method: "POST",
      headers: myHeaders,
      body: raw,
    };

    const url = new URL(this.Host + "/v1/note");

    const response = await fetch(url, requestOptions);

    // Check if the response is OK (status in the range 200-299)
    if (!response.ok) {
      throw new Error(`HTTP error! url: ${url}, status: ${response.status}`);
    }

    const result = await response.json();
    const data = result["data"];

    return new NotePreview(
      data["id"],
      data["notebook_id"],
      data["title"],
      data["content"],
      new Date(data['updated_at']),
    )
  }

  async DeleteNot(note: NotePreview): Promise<void> {
    const requestOptions = {
      method: "DELETE",
    };

    const queryParams = {
      note_id: note ? note.ID : "",
    };

    const url = new URL(this.Host + "/v1/note");
    url.search = new URLSearchParams(queryParams).toString();

    const response = await fetch(url, requestOptions);

    // Check if the response is OK (status in the range 200-299)
    if (!response.ok) {
      throw new Error(`HTTP error! url: ${url}, status: ${response.status}`);
    }
  }
}


export default NoteAPI