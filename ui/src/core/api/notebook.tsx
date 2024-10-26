import NotebookPreview from "../models/notebook.preview"
import { NotebookPreviewMock } from "../test/mock/notebook.preview.mock"

class NotebookAPI {
  Host: string
  UseMock: boolean

  constructor(host: string, mock: boolean = false) {
    this.Host = host
    this.UseMock = mock
  }

  async ReadNotebook(): Promise<NotebookPreview[]> {
    if (this.UseMock) {
      return NotebookPreviewMock
    }

    const requestOptions = {
      method: "GET",
    };

    const response = await fetch(this.Host + "/v1/notebooks", requestOptions);

    // Check if the response is OK (status in the range 200-299)
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const result = await response.json();
    const notebooks: NotebookPreview[] = [];
    const data = result["data"]["notebooks"];

    for (let n of data) {
      notebooks.push(
        new NotebookPreview(
          n["id"],
          n["name"]
        )
      );
    }
    return notebooks;
  }
}

export default NotebookAPI