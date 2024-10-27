import NotebookPreview from "../../models/notebook.preview";

interface NotebookProps {
  data: NotebookPreview
  onClick(notebook: NotebookPreview): void
  isActive: boolean
}

function Notebook({ data, onClick, isActive }: NotebookProps) {
  const indicatorBgColor = isActive ? 'bg-main-normal' : ''

  return (
    <div
      className="border-b border-r border-black hover:border-background-accent-hover px-2 pb-2 duration-500 cursor-pointer"
      onClick={() => onClick(data)}
    >
      <div className="flex flex-row items-center gap-2">
        <div className={"h-1 w-1 " + indicatorBgColor}></div>
        <p className="text-sm">
          {data.Title}
        </p>
      </div>
    </div>
  )

}

export default Notebook;