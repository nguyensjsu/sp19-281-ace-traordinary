import React, { Component } from 'react';
import { Link } from "react-router-dom";

class Navigation extends Component {
    render() {
        return (
            <div className="Navigation">
                <p></p>
                <div className={"brand"}>
              <img src={require('./picasalogo.png')}/>
                </div>
                <div className={"header-buttons"}>
                <ul><Link to={"/"}><li>Home</li></Link>
                    <Link to={"/"}> <li>MyImages</li></Link>
                    <li>Logout</li>
                </ul>
                </div>
            </div>
        );
    }
}

export default Navigation;
