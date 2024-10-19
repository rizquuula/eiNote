import React from 'react';
import Provider from './core/Provider';
import './App.css'

function App() {
  const webTheme = localStorage.getItem("theme")

  return (
    <div className={webTheme || ''}>
      <Provider />
    </div>
  );
}

export default App;
