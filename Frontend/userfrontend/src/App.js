import React, { Component } from 'react';
import Registration from "./components/Registration";
import "./scss/Picasso.scss"
import './App.css';
import BrowserRouter from "react-router-dom/es/BrowserRouter";
import WebRouter from "./components/WebRouter"
class App extends Component {
  render() {
    return (
        <BrowserRouter>
      <div className="App">
          <WebRouter/>
      </div>
        </BrowserRouter>
    );
  }
}

export default App;
