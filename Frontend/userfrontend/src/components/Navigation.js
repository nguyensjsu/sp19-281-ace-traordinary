import React, { Component } from 'react';
import { Link } from "react-router-dom";

class Navigation extends Component {
    render() {
        return (
            <div className="Navigation">
              <img className={"brand"} src={require('./picasalogo.png')}/>
                <div className={"header-buttons"}>
                <ul><Link to={"/"}><li>Home</li></Link>
                    <Link to={"/myimages"}> <li>MyImages</li></Link>
                    <li>Logout</li>
                </ul>
                </div>
            </div>
        );
    }
}

export default Navigation;
