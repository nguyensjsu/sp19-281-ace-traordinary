import React, { Component } from 'react';
import { Link } from "react-router-dom";
import connect from "react-redux/es/connect/connect";
import {userlogout} from '../actions/UserAction'

class Navigation extends Component {

constructor(props){
    super(props);
    this.logout=this.logout.bind(this);
}

    logout=()=>{
        localStorage.removeItem("state");
        this.props.userlogout()
    }

    render() {
        return (
            <div className="Navigation">
              <img className={"brand"} src={require('./picasalogo.png')}/>
                <div className={"header-buttons"}>
                <ul><Link to={"/"}><li>Home</li></Link>
                    <Link to={"/myimages"}> <li>MyImages</li></Link>
                    <Link to={"/"}> <li onClick={this.logout}>Logout</li></Link>
                </ul>
                </div>
            </div>
        );
    }
}
function mapStateToProps(state) {
    return{
        user:state.user,
    }
}
export default connect(mapStateToProps,{userlogout})(Navigation);
