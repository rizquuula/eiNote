import { useEffect, useRef } from "react"

import EasyMDE from "easymde";
import "easymde/dist/easymde.min.css";
import "./editor.darktheme.css";
import "./editor.preview.css";
import NotePreview from "../../models/note.preview";

interface NoteEditorProps {
  note: NotePreview | null
  SaveNote(note: NotePreview): void
}

export default function NoteEditor({ note, SaveNote }: NoteEditorProps) {

  const textareaRef = useRef<HTMLTextAreaElement>(null);
  const easyMdeRef = useRef<EasyMDE | null>(null);

  useEffect(() => {
    const saveDelay = 5000
    if (!textareaRef.current) {
      throw new Error("Textarea ref not found.")
    }

    // We only ever want EasyMDE to instantiate itself once.
    // ie. We're doing this to avoid double render problems that show themselves in React 18. 
    if (!easyMdeRef.current) {
      easyMdeRef.current = new EasyMDE({
        autosave: {
          enabled: true,
          uniqueId: "tmpNote",
          delay: saveDelay,
          timeFormat: {
            locale: 'en-US',
            format: {
              year: 'numeric',
              month: 'long',
              day: '2-digit',
              hour: '2-digit',
              minute: '2-digit',
              second: '2-digit',
            },
          },
          text: "Autosaved: "
        },
        element: textareaRef.current,
        hideIcons: ["side-by-side"],
        shortcuts: {
          "toggleSideBySide": null,
        },
        renderingConfig: {
          codeSyntaxHighlighting: true
        },
        minHeight: window.innerHeight * 0.75 + "px",
        maxHeight: window.innerHeight * 0.75 + "px",
        scrollbarStyle: "native",
      });

      setInterval(() => {
        if (note !== null && easyMdeRef.current !== null) {
          note.UpdateContent(easyMdeRef.current.value())
          SaveNote(note)
        }
      }, saveDelay)

    } else {
      const content = note ? note.Content : ""
      easyMdeRef.current.value(content);
    }

  }, [note, SaveNote])

  return <textarea title="editor" className="note-editor" ref={textareaRef} />
}