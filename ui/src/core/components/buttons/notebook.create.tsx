import { useState } from "react";

interface CreateNotebookBtnProps {
  CreateNotebook(name: string): void
}

function CreateNotebookBtn({ CreateNotebook }: CreateNotebookBtnProps) {
  const [newNotebook, setNewNotebook] = useState<string>("");

  return <div className="flex flex-row justify-between">
    <input id="new-notebook" className="text-black bg-light-normal flex-grow outline-none px-2" value={newNotebook} onInput={(e) => setNewNotebook(e.currentTarget.value)} type="text" />
    <button type="button" className="px-2 py-1 bg-main-normal hover:bg-main-hover duration-500" onClick={() => CreateNotebook(newNotebook)}>
      add
    </button>
  </div>
}

export default CreateNotebookBtn