import React, {Component} from 'react';
import {CardElement, injectStripe} from 'react-stripe-elements';
import {Image,Grid,TextArea,Form,Button,Icon ,Modal,Header} from "semantic-ui-react"
import axios from "axios";
import {IMAGE_CMDURL} from "../resources/constants";
import  { Redirect } from 'react-router-dom'
class CheckoutForm extends Component {


    handleOpen = () => this.setState({ modalOpen: true })

    handleClose = () => this.setState({ modalOpen: false })
    constructor(props) {
        super(props);
        this.submit = this.submit.bind(this);
        this.state = {complete: false,modalOpen: false};
        this.handleOpen=this.handleOpen.bind(this);
        this.handleClose=this.handleClose.bind(this);
    }

    async submit(ev) {
        let {token} = await this.props.stripe.createToken({name: "Random"});
        let img=this.props.img
        img.userid=this.props.user.userid
        const response = await axios.put(`${IMAGE_CMDURL}/images/${img.imageid}`,img).catch(function (error) {
            if (error.response) {
                console.log(error.response.status);
                console.log(error.response.headers);
            }
        });
        if (response.status===200) {
            this.setState({modalOpen: true })
        }

    }

    render() {
        if(this.state.complete){
            this.handleOpen()
           return <Redirect to='/myimages' />
        }
        return (
            <>
            <div className={"CheckoutForm"}>
            <div className="checkout">
                <p>Please enter card details to complete Purchase</p>
                <div className={"element"}>
                <CardElement />
                </div>
                <div className={"checkoutbtn"}>
                <Button icon color={'blue'} labelPosition='right' onClick={this.submit}>
                    Complete Purchase
                    <Icon name='right arrow' />
                </Button>
                </div>
            </div>
            </div>
                <Modal
                    open={this.state.modalOpen}
                    onClose={this.handleClose}
                    basic
                    size='small'
                >
                    <Header icon='browser' content='Payment Status' />
                    <Modal.Content>
                        <h3>Your Payment is success and sent confirmation email</h3>
                    </Modal.Content>
                    <Modal.Actions>
                        <Button color='green' onClick={this.handleClose} inverted>
                            <Icon name='checkmark' onClick={()=>this.setState({complete:true})}/> Got it
                        </Button>
                    </Modal.Actions>
                </Modal>
                </>
        );
    }
}
export default injectStripe(CheckoutForm);
