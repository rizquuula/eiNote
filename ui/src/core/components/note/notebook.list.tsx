import Notebook from './notebook';

function NotebookList() {
  const array = Array.from({ length: 3 }, (v, i) => i + 1);

  return (
    <div>
      {array.map(() => {
        return <Notebook />
      })}
    </div>
  )
}

export default NotebookList;