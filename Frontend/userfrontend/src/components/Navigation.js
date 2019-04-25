import React, {Component} from 'react';
import {Link} from "react-router-dom";
import connect from "react-redux/es/connect/connect";
import {userlogout} from '../actions/UserAction'
import {Button, Header, Image, Modal,Dropdown,Menu,Label} from 'semantic-ui-react'
import NewImage from './NewImage'
import firebase from 'firebase'

const options = [
    { key: 1, text: 'Choice 1', value: 1 },
    { key: 2, text: 'Choice 2', value: 2 },
    { key: 3, text: 'Choice 3', value: 3 },
]
class Navigation extends Component {
    state = {open: false}
    show = dimmer => () => this.setState({dimmer, open: true})
    close = () => this.setState({open: false})

    constructor(props) {
        super(props);
        this.logout = this.logout.bind(this);
    }
    logout = () => {
        localStorage.removeItem("state");
        firebase.auth().signOut();
        this.props.userlogout()
    }
    
    render() {
        const { open, dimmer } = this.state
        const {user} =this.props
        const trigger = (<span><Image avatar src={user.profilepic} /> {user.firstname}</span>
        )
        return (
            <div>
            <div className="Navigation">
                <div>
                <img className={"brand"} src={require('./picasalogo.png')}/>
                </div>
                <div className={"header-buttons"}>
                    <ul><Link to={"/"} className={"link"}>
                        <li>Home</li>
                    </Link>
                        <Link to={"/myimages"} className={"link"}>
                            <li>MyImages</li>
                        </Link>
                        <li onClick={this.show('blurring')}>NewImage</li>


                        <Dropdown   trigger={trigger} pointing='top left' icon={null}  >
                            <Dropdown.Menu>
                                <Dropdown.Item><Link to={"/myprofile"} className={"link"}>
                                    My Profile
                                </Link></Dropdown.Item>
                                <Dropdown.Item>Remove Account</Dropdown.Item>
                                <Dropdown.Divider />
                                <Dropdown.Item><Link to={"/"} className={"link"}>
                                    <span onClick={this.logout}>Log Out</span>
                                </Link></Dropdown.Item>
                            </Dropdown.Menu>
                        </Dropdown>

                    </ul>

                </div>
                <Modal dimmer={dimmer} open={open} onClose={this.close}>
                    <Modal.Header>Upload New Image</Modal.Header>
                    <Modal.Content image>
                        <NewImage/>
                    </Modal.Content>
                    <Modal.Actions>
                        <Button color='red' onClick={this.close}>
                            Cancel
                        </Button>
                        <Button
                            color='green'
                            content="Upload"
                            onClick={this.close}
                        />
                    </Modal.Actions>
                </Modal>
            </div>
            </div>
        );
    }
}

function mapStateToProps(state) {
    return {
        user: state.user,
    }
}

export default connect(mapStateToProps, {userlogout})(Navigation);
