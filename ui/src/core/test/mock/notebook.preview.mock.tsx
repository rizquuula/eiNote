import NotebookPreview from "../../models/notebook.preview"

const notebook1 = new NotebookPreview("1", "Personal Private", new Date())
const notebook2 = new NotebookPreview("2", "PT A", new Date())
const notebook3 = new NotebookPreview("3", "PT B", new Date())
notebook3.SetActive(true)
const notebook4 = new NotebookPreview("4", "Fiverr", new Date())

export const NotebookPreviewMocks = [
  notebook1,
  notebook2,
  notebook3,
  notebook4,
]

export const NotebookPreviewMock = notebook1