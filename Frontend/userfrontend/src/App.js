import React, { Component } from 'react';
import Registration from "./components/Registration";

import './App.css';
import BrowserRouter from "react-router-dom/es/BrowserRouter";
import Login from "./components/Login";

class App extends Component {
  render() {
    return (
        <BrowserRouter>
      <div className="App">
<Registration/>
          <Login/>
      </div>
        </BrowserRouter>
    );
  }
}

export default App;
