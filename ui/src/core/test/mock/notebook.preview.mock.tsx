import NotebookPreview from "../../models/notebook.preview"

const notebook1 = new NotebookPreview("Personal Private")
const notebook2 = new NotebookPreview("SALT")
const notebook3 = new NotebookPreview("PSN")
notebook3.SetActive(true)
const notebook4 = new NotebookPreview("Fiverr")

export const NotebookPreviewMock = [
  notebook1,
  notebook2,
  notebook3,
  notebook4,
]