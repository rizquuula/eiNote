import './App.css';
import Provider from './core/Provider';

function App() {
  const webTheme = localStorage.getItem("theme")
  return (
    <html className={ webTheme }>
      <Provider />
    </html>
  );
}

export default App;
