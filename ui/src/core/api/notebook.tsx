import NotebookPreview from "../models/notebook.preview"
import { NotebookPreviewMock, NotebookPreviewMocks } from "../test/mock/notebook.preview.mock"

class NotebookAPI {
  Host: string
  UseMock: boolean

  constructor(host: string, mock: boolean = false) {
    this.Host = host
    this.UseMock = mock
  }

  async CreateNotebook(notebook: NotebookPreview): Promise<NotebookPreview> {
    if (this.UseMock) {
      return NotebookPreviewMock
    }

    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    const raw = JSON.stringify({
      "id": notebook.ID,
      "name": notebook.Name,
    });

    const requestOptions = {
      method: "POST",
      headers: myHeaders,
      body: raw,
    };

    const url = this.Host + "/v1/notebook"
    const response = await fetch(url, requestOptions);

    // Check if the response is OK (status in the range 200-299)
    if (!response.ok) {
      throw new Error(`HTTP error! url: ${url}, status: ${response.status}`);
    }

    const result = await response.json();
    const data = result["data"];
    return new NotebookPreview(data["id"], data["name"], new Date(result["updated_at"]))
  }

  async DeleteNotebook(notebook: NotebookPreview): Promise<void> {
    const requestOptions = {
      method: "DELETE",
    };

    const queryParams = {
      notebook_id: notebook ? notebook.ID : "",
    };

    const url = new URL(this.Host + "/v1/notebook");
    url.search = new URLSearchParams(queryParams).toString();

    const response = await fetch(url, requestOptions);

    // Check if the response is OK (status in the range 200-299)
    if (!response.ok) {
      throw new Error(`HTTP error! url: ${url}, status: ${response.status}`);
    }
  }

  async ReadNotebooks(): Promise<NotebookPreview[]> {
    if (this.UseMock) {
      return NotebookPreviewMocks
    }

    const requestOptions = {
      method: "GET",
    };

    const url = this.Host + "/v1/notebooks"
    const response = await fetch(url, requestOptions);

    // Check if the response is OK (status in the range 200-299)
    if (!response.ok) {
      throw new Error(`HTTP error! url: ${url}, status: ${response.status}`);
    }

    const result = await response.json();
    const notebooks: NotebookPreview[] = [];
    const data = result["data"]["notebooks"];

    for (let n of data) {
      notebooks.push(
        new NotebookPreview(
          n["id"],
          n["name"],
          new Date(n["updated_at"])
        )
      );
    }
    return notebooks;
  }
}

export default NotebookAPI