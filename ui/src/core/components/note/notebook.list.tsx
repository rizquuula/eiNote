import NotebookPreview from '../../models/notebook.preview';
import Notebook from './notebook';

function NotebookList({ notebooks }: { notebooks: NotebookPreview[] }) {

  return (
    <div className='flex flex-col gap-2'>
      {notebooks.map((data, i) => {
        return <Notebook key={i} data={data} />
      })}
    </div>
  )
}

export default NotebookList;