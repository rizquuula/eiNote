import NotebookPreview from "../../models/notebook.preview"

const notebook1 = new NotebookPreview("1", "Personal Private")
const notebook2 = new NotebookPreview("2", "PT A")
const notebook3 = new NotebookPreview("3", "PT B")
notebook3.SetActive(true)
const notebook4 = new NotebookPreview("4", "Fiverr")

export const NotebookPreviewMock = [
  notebook1,
  notebook2,
  notebook3,
  notebook4,
]