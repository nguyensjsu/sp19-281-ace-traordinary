import React, { Component } from 'react';
import {Card, Button, Checkbox, Form } from 'semantic-ui-react'
import connect from "react-redux/es/connect/connect";
import '../css/registration.css';
import {userregestration} from '../actions/UserAction'


class Registration extends Component {
    constructor(props){
        super(props);
        this.state={
            userid:"",
            password:"",
            phonenumber:"",
            firstname:"",
            lastname:""
        }
        this.handlechange=this.handlechange.bind(this);
        this.registeruser=this.registeruser.bind(this);

    }

    handlechange=(event)=>{
        this.setState([event.target.name]=event.target.value)
    }
    registeruser=()=>{

        this.props.userregestration(this.state)
    }


    render(){
        return(
            <div className={"Registration"}>
                <Card className={"Registration-Card"}>
                <Form>
                    <Form.Field>
                        <label>Email</label>
                        <input placeholder='Email' name="email" onChange={this.handlechange} required={true} maxLength={40} />
                    </Form.Field>
                    <Form.Field>
                        <label>Password</label>
                        <input placeholder='Please enter password' name="password" onChange={this.handlechange} required={true} maxLength={20}/>
                    </Form.Field>
                    <Form.Field>
                        <label>First Name</label>
                        <input placeholder='First Name'name="firstname"onChange={this.handlechange} required={true} maxLength={20}/>
                    </Form.Field>
                    <Form.Field>
                        <label>Last Name</label>
                        <input placeholder='Last Name' name="lastname"  onChange={this.handlechange} required={true} maxLength={20}/>
                    </Form.Field>
                    <Button type='submit' onClick={this.registeruser} negative>Register</Button>
                </Form></Card>
            </div>
        );
    }
}


function mapStateToProps(state) {
    return{
        user:state.user,
    }
}

export default connect(mapStateToProps,{userregestration})(Registration);