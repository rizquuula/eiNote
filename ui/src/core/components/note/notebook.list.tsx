import NotebookPreview from '../../models/notebook.preview';
import Notebook from './notebook';

interface NotebookListProps {
  notebooks: NotebookPreview[]
  activeNotebook: NotebookPreview | null
  SetActiveNotebook(notebook: NotebookPreview): void
}

function NotebookList({ notebooks, activeNotebook, SetActiveNotebook }: NotebookListProps) {

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
        return <Notebook key={i} data={data} onClick={SetActiveNotebook} isActive={isActive} />
      })}
    </div>
  )
}

export default NotebookList;