import React, { Component } from 'react';
import Registration from "./components/Registration";
import "./scss/Picasso.scss"
import './App.css';
import BrowserRouter from "react-router-dom/es/BrowserRouter";
import WebRouter from "./components/WebRouter"
import Login from "./components/Login";

class App extends Component {


    constructor(props){
        super(props);
        this.state={
            isAuthenticated:true,
            showlogin_page:true,
            showregister_page:false
        }
        this.showLogin = this.showLogin.bind(this);
        this.showRegister = this.showRegister.bind(this);
    }
    showLogin=()=>{
        this.setState({
            showlogin_page:true,
            showregister_page:false
        })
    }
    showRegister=()=>{
        console.log("HI I got called");
        this.setState({
            showlogin_page:false,
            showregister_page:true
        })
    }
    

  render() {
      console.log(this.props)
    if(this.state.isAuthenticated){
        return (
            <BrowserRouter>
                <div className="App">
                    <WebRouter/>
                </div>
            </BrowserRouter>
        );
    }else if(this.state.showlogin_page) {
        return (<><video autoPlay muted loop className={"videobaground"}>
            <source src={require('./resources/LandPageBaground.mp4')} type="video/mp4"/>
        </video><Login showRegister={this.showRegister}/></>)
    }else if(this.state.showregister_page){
        return (<><video autoPlay muted loop className={"videobaground"}>
            <source src={require('./resources/LandPageBaground.mp4')} type="video/mp4"/>
        </video><Registration showLogin={this.showLogin}/></>)
    }
  }
}

export default App;
