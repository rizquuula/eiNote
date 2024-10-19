import { useEffect, useRef } from "react"

import EasyMDE from "easymde";
import "easymde/dist/easymde.min.css";
import "./editor.darktheme.css";
import "./editor.preview.css";


export default function NoteEditor() {

  const textareaRef = useRef<HTMLTextAreaElement>(null);
  const easyMdeRef = useRef<EasyMDE | null>(null);

  useEffect(() => {

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
          delay: 1000,
          submit_delay: 0,
          timeFormat: {
            locale: 'en-US',
            format: {
              year: 'numeric',
              month: 'long',
              day: '2-digit',
              hour: '2-digit',
              minute: '2-digit',
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
    }

  }, [])

  return <textarea title="editor" className="note-editor" ref={textareaRef} />
}