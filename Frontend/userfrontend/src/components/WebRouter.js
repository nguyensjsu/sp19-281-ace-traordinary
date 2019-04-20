import React, { Component } from 'react';
import Login from "./Login";
import {Redirect, Route} from 'react-router-dom';
import ImagesDashBoard from "./ImagesDashBoard";
import Registration from "./Registration";
import Navigation from "./Navigation"
import BuyImage from "./BuyImage";
import MyImages from "./MyImages"
class WebRouter extends Component {

constructor(props){
    super(props);
    this.state={
        isAuthenticated:false,
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
           land_page = (<><video autoPlay muted loop className={"videobaground"}>
               <source src={require('./LandPageBaground.mp4')} type="video/mp4"/>
           </video><Login showRegister={this.showRegister}/></>)
        }else if(this.state.showregister_page){
            land_page =(<><video autoPlay muted loop className={"videobaground"}>
                <source src={require('./LandPageBaground.mp4')} type="video/mp4"/>
            </video><Registration showLogin={this.showLogin}/></>)
        }
        return land_page;
    }

    buyImageRoute =()=>{
        return (<>
            <Route exact path="/images/buy" component={Navigation}/>
            <Route exact path="/images/buy" component={BuyImage}/>
        </>)
    }
    myImagesRoute =()=>{
        return (<>
            <Route exact path="/myimages" component={Navigation}/>
            <Route exact path="/myimages" component={MyImages}/>
        </>)
    }
    render() {
        return (
            <div className="WebRouter">
                {this.landPageRoute()}
                {this.buyImageRoute()}
                {this.myImagesRoute()}
            </div>
        );
    }
}

export default WebRouter;
