import React, { Component } from 'react';
import Login from "./Login";
import {Redirect, Route} from 'react-router-dom';
import ImagesDashBoard from "./ImagesDashBoard";
import Registration from "./Registration";
import Navigation from "./Navigation"
import BuyImage from "./BuyImage";
import MyImages from "./MyImages"
import ViewImage from "./ViewImage"
class WebRouter extends Component {

constructor(props){
    super(props);
}

    render() {
        return (
            <div className="WebRouter">
                <div>
                <Navigation/>
                </div>
                <div className={"secondComponent"}> 
                <Route exact path="/" component={ImagesDashBoard}/>
                <Route path="/images/buy" component={BuyImage}/>
                <Route path="/images/comment" component={ViewImage}/>
                <Route exact path="/myimages" component={MyImages}/>
                </div>
            </div>
        );
    }
}

export default WebRouter;
