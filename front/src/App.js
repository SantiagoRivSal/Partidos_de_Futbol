import React from 'react';
import {Header} from './componentes/Header';
import 'boxicons';
import {BrowserRouter as Router} from "react-router-dom";
import {Paginas} from './componentes/Paginas';
function App() {
  return (
    
    <div className="App">
      <Router>
        <Header/>
        <Paginas/>
      </Router>
    </div>
 
  );
}

export default App;