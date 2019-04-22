import React, {Component} from 'react';
import {Link} from "react-router-dom";
import connect from "react-redux/es/connect/connect";
import {userlogout} from '../actions/UserAction'
import {Button, Header, Image, Modal} from 'semantic-ui-react'

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
        this.props.userlogout()
    }
    
    render() {
        const { open, dimmer } = this.state
        return (
            <div className="Navigation">
                <img className={"brand"} src={require('./picasalogo.png')}/>
                <div className={"header-buttons"}>
                    <ul><Link to={"/"} className={"link"}>
                        <li>Home</li>
                    </Link>
                        <Link to={"/myimages"} className={"link"}>
                            <li>MyImages</li>
                        </Link>
                        <Link to={"/"} className={"link"}>
                            <li onClick={this.logout}>Logout</li>
                        </Link>
                        <li onClick={this.show('blurring')}>NewImage</li>
                        <li>RemoveAccount</li>
                    </ul>
                </div>
                <Modal dimmer={dimmer} open={open} onClose={this.close}>
                    <Modal.Header>Upload New Image</Modal.Header>
                    <Modal.Content image>
                        <Image wrapped size='medium' src='https://react.semantic-ui.com/images/avatar/large/rachel.png' />
                        <Modal.Description>
                            <Header>Default Image</Header>
                            <p>We've found the following gravatar image associated with your e-mail address.</p>
                            <p>Is it okay to use this photo?</p>
                        </Modal.Description>
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
        );
    }
}

function mapStateToProps(state) {
    return {
        user: state.user,
    }
}

export default connect(mapStateToProps, {userlogout})(Navigation);
