import React, { Component } from 'react';
import Login from "./Login";
import {Redirect, Route} from 'react-router-dom';
import ImagesDashBoard from "./ImagesDashBoard";
import Registration from "./Registration";
import Navigation from "./Navigation"
class WebRouter extends Component {

constructor(props){
    super(props);
    this.state={
        isloggedin:true,
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
    landPage=()=>{
        let land_page=""
        if(this.state.isloggedin){
            land_page =  <><Navigation/><ImagesDashBoard/></>
        }else if(this.state.showlogin_page) {
           land_page = <Login showRegister={this.showRegister}/>
        }else if(this.state.showregister_page){
            land_page =<Registration showLogin={this.showLogin}/>
        }
        return land_page;
    }
    render() {
        return (
            <div className="WebRouter">
                <Route exact path="/" >
                    {this.landPage()}
                </Route>
            </div>
        );
    }
}

export default WebRouter;
