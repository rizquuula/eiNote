import NotebookPreview from '../../models/notebook.preview';
import Notebook from './notebook';

interface NotebookListProps {
  notebooks: NotebookPreview[]
  activeNotebook: NotebookPreview | null
  setActiveNotebook(notebook: NotebookPreview): void
  deleteNotebook(notebook: NotebookPreview): void
}

function NotebookList({
  notebooks,
  activeNotebook,
  setActiveNotebook,
  deleteNotebook,
}: NotebookListProps) {

  return (
    <div className='flex flex-col gap-2'>
      {notebooks.map((data, i) => {
        const isActive = data.ID === (
          activeNotebook
            ? activeNotebook.ID
            : notebooks.length > 0
              ? notebooks[0].ID
              : null
        )
        return <Notebook
          key={i}
          data={data}
          setActiveNotebook={setActiveNotebook}
          isActive={isActive}
          deleteNotebook={deleteNotebook}
        />
      })}
    </div>
  )
}

export default NotebookList;