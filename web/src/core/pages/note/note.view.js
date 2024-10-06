import React from 'react'
import NoteList from '../../components/note/note.list';
import NotebookList from '../../components/note/notebook.list';
import NoteEditor from '../../components/note/editor';

class NoteView extends React.Component {
  render() {

    return (
      <div id='note' className='flex flex-row'>
        <div id='note.sidebar' className='w-2/12'>
          <div id='note'>
            <NoteList />
          </div>
          <div id='notebook'>
            <NotebookList />
          </div>
        </div>
        <div id='note.editor' className='w-10/12'>
          <NoteEditor />
        </div>
      </div>
    )
  }
}

export default NoteView;