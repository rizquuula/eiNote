import NotePreview from "../../models/note.preview"

const notePreview1 = new NotePreview(
  "Title Example",
  "Note description this can be expanded up to 3 lines, after that begin blank blank blank, like this",
  new Date(),
)

const notePreview2 = new NotePreview(
  "Title Example",
  "Note description this can be expanded up to 3 lines, after that begin blank blank blank, like this",
  new Date(),
)
notePreview2.SetActive(true)

const notePreview3 = notePreview1
const notePreview4 = notePreview1

export const NotePreviewMock = [
  notePreview1,
  notePreview2,
  notePreview3,
  notePreview4,
]