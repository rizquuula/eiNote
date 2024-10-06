import React from 'react'
import Notebook from './notebook';

class NotebookList extends React.Component {
  render() {
    const array = Array.from({ length: 3 }, (v, i) => i + 1);
    return (
      array.map(() => {
        return <Notebook />
      })
    )
  }
}

export default NotebookList;