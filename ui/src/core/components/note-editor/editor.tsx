import { useEffect, useRef } from "react"

import EasyMDE from "easymde";
import "easymde/dist/easymde.min.css";
import "./editor.darktheme.css";
import "./editor.preview.css";
import NotePreview from "../../models/note.preview";

interface NoteEditorProps {
  notebookId: string | null
  note: NotePreview | null
  saveNote(note: NotePreview): void
}

export default function NoteEditor({ notebookId, note, saveNote }: NoteEditorProps) {

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

    } else {
      // never update the textarea based on active note, it caused loop or textarea refreshed
    }

    // Set up the interval
    const intervalId = setInterval(() => {

      if (easyMdeRef.current !== null) {
        const newContent = easyMdeRef.current.value();
        if (note === null) {
          if (notebookId) {
            saveNote(new NotePreview("", notebookId, "", newContent, new Date()));
          }
        } else {
          note.UpdateContent(newContent);
          saveNote(note);
        }
      }
    }, saveDelay);

    // clear interval on unmount
    return () => clearInterval(intervalId);

  }, [note, notebookId, saveNote])

  return <textarea title="editor" className="note-editor" ref={textareaRef} />
}