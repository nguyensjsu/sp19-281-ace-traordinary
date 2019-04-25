import React, { Component } from 'react';
import {Card, Button, Checkbox, Form } from 'semantic-ui-react'
import connect from "react-redux/es/connect/connect";
import '../css/login.css';
import {userregestration,userlogin,firbaselogin} from '../actions/UserAction'
import firebase from 'firebase'
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth'

firebase.initializeApp({
    apiKey:"AIzaSyBlv6VOEesnKK9vCL2pc5LSIBoaLbyZCw8",
    authDomain:"cmpe281-238705.firebaseapp.com"
})
class Login extends Component {
    constructor(props){
        super(props);
        this.state={
            userid:"",
            password:"",
            isSignedIn:false,
        }
        this.handlechange=this.handlechange.bind(this);
        this.loginuser=this.loginuser.bind(this);
        this.uiConfig = {
            signInFlow: "popup",
            signInOptions: [
                firebase.auth.GoogleAuthProvider.PROVIDER_ID,
                firebase.auth.FacebookAuthProvider.PROVIDER_ID,
            ],
            callbacks: {
                signInSuccess: () => false
            }
        }
    }
    componentDidMount = () => {
        firebase.auth().onAuthStateChanged(user => {
            this.setState({ isSignedIn: !!user })
            if(!!user){
                let firbaseuser={
                    userid: user.email,
                    profilepic:user.photoURL,
                    firstname:user.displayName,
                    phonenumber:user.phoneNumber
                }
                this.props.firbaselogin(firbaseuser)
            }
            console.log("user", user)
            console.log("user",JSON.stringify(user))
        })
    }
    handlechange=(event)=>{
        this.setState({[event.target.name]:event.target.value})
    }
    loginuser=()=>{

        this.props.userlogin(this.state)
    }


    render(){
        return(
            <div>

            <div className={"Login"}>

                <div className={"Login-Card"}>
                    <h3>Login</h3>
                    <hr/>
                    <Form>
                        <Form.Field>
                            <label>Email</label>
                            <input placeholder='Email' name="userid" onChange={this.handlechange} required={true} maxLength={40} />
                        </Form.Field>
                        <Form.Field>
                            <label>Password</label>
                            <input placeholder='Please enter password' type={"password"} name="password" onChange={this.handlechange} required={true} maxLength={20}/>
                        </Form.Field>
                        <Button type='submit' onClick={this.loginuser} negative>Login</Button>
                        <a><b className={"rb"}>Forgot password</b></a>
                        <br/>
                        <p><span>Not a Member?</span><b className={"rb"}onClick={this.props.showRegister} >SignUP</b></p>
                    </Form>
                    <StyledFirebaseAuth
                        uiConfig={this.uiConfig}
                        firebaseAuth={firebase.auth()}
                    />
                </div>
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

export default connect(mapStateToProps,{userlogin,firbaselogin})(Login);