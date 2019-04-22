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
}

    render() {
        return (
            <div className="WebRouter">
                <Navigation/>
                <Route exact path="/" component={ImagesDashBoard}/>
                <Route path="/images/buy" component={BuyImage}/>
                <Route exact path="/myimages" component={MyImages}/>
            </div>
        );
    }
}

export default WebRouter;
