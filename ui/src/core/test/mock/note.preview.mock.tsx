import NotePreview from "../../models/note.preview"

const notePreview1 = new NotePreview(
  "41cb2657-7745-45d2-ab6b-1cab38e6440b",
  "31cb2657-7745-45d2-ab6b-1cab38e6440b",
  "Title Example 1",
  "Note description this can be expanded up to 3 lines, after that begin blank blank blank, like this",
  new Date(),
)

const notePreview2 = new NotePreview(
  "d7165c14-8eba-495b-9309-0a62fd3f16d1",
  "31cb2657-7745-45d2-ab6b-1cab38e6440b",
  "Title Example 2",
  "Note description this can be expanded up to 3 lines, after that begin blank blank blank, like this",
  new Date(),
)

const notePreview3 = new NotePreview(
  "d7165c14-8eba-495b-9309-0a62fd3f16d2",
  "31cb2657-7745-45d2-ab6b-1cab38e6440b",
  "Title Example 3",
  "Note description this can be expanded up to 3 lines, after that begin blank blank blank, like this",
  new Date(),
)

export const NotePreviewMock = [
  notePreview1,
  notePreview2,
  notePreview3,
]