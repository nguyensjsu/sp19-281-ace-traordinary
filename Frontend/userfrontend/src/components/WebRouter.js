import React, { Component } from 'react';
import Login from "./Login";
import {Redirect, Route} from 'react-router-dom';
import ImagesDashBoard from "./ImagesDashBoard";
import Registration from "./Registration";
import Navigation from "./Navigation"
import BuyImage from "./BuyImage";
class WebRouter extends Component {

constructor(props){
    super(props);
    this.state={
        isAuthenticated:true,
        showlogin_page:true,
        showregister_page:false
    }
    this.showLogin = this.showLogin.bind(this);
    this.showRegister = this.showRegister.bind(this);
    this.landPageRoute =this.landPageRoute.bind(this);
    this.buyImageRoute=this.buyImageRoute.bind(this);
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
    landPageRoute=()=>{
        let land_page=(<div></div>)
        if(this.state.isAuthenticated){
            land_page = ( <>
                <Route exact path="/" component={Navigation}/>
                <Route exact path="/" component={ImagesDashBoard}/>
               </>)
        }else if(this.state.showlogin_page) {
           land_page = (<Login showRegister={this.showRegister}/>)
        }else if(this.state.showregister_page){
            land_page =(<Registration showLogin={this.showLogin}/>)
        }
        return land_page;
    }

    buyImageRoute =()=>{
        return (<>
            <Route exact path="/images/buy" component={Navigation}/>
            <Route exact path="/images/buy" component={BuyImage}/>
        </>)
    }
    render() {
        return (
            <div className="WebRouter">
                {this.landPageRoute()}
                {this.buyImageRoute()}
            </div>
        );
    }
}

export default WebRouter;
