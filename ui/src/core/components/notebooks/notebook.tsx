import NotebookPreview from "../../models/notebook.preview";

interface NotebookProps {
  data: NotebookPreview
  setActiveNotebook(notebook: NotebookPreview): void
  isActive: boolean
  deleteNotebook(notebook: NotebookPreview): void
}

function Notebook({
  data,
  setActiveNotebook,
  isActive,
  deleteNotebook
}: NotebookProps) {
  const indicatorBgColor = isActive ? 'bg-main-normal' : ''

  return (
    <div
      className="border-b border-r border-black hover:border-background-accent-hover px-2 pb-2 duration-500 cursor-pointer"
      onClick={() => setActiveNotebook(data)}
    >
      <div className="flex justify-between">
        <div className="flex flex-row items-center gap-2">
          <div className={"h-1 w-1 " + indicatorBgColor}></div>
          <p className="text-sm">
            {data.Name}
          </p>
        </div>
        <div>
          <svg
            className="stroke-accent-hover hover:stroke-red-500 hover:bg-background-accent-hover rounded-full duration-500 p-1"
            onClick={() => deleteNotebook(data)}
            xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M3 6h18" />
            <path d="M8 6V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2" />
            <path d="M10 11v6" />
            <path d="M14 11v6" />
            <rect x="5" y="6" width="14" height="14" rx="2" ry="2" />
          </svg>
        </div>
      </div>
    </div>
  )

}

export default Notebook;