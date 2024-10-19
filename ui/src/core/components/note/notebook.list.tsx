import { NotebookPreviewMock } from '../../test/mock/notebook.preview.mock';
import Notebook from './notebook';

function NotebookList() {

  return (
    <div className='flex flex-col gap-2'>
      {NotebookPreviewMock.map((data, i) => {
        return <Notebook key={i} data={data} />
      })}
    </div>
  )
}

export default NotebookList;