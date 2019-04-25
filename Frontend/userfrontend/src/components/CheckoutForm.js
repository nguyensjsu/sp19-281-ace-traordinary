import React, {Component} from 'react';
import {CardElement, injectStripe} from 'react-stripe-elements';
import {Image,Grid,TextArea,Form,Button,Icon } from "semantic-ui-react"

class CheckoutForm extends Component {
    constructor(props) {
        super(props);
        this.submit = this.submit.bind(this);
    }

    async submit(ev) {
        // User clicked submit
    }

    render() {
        return (
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
        );
    }
}

export default injectStripe(CheckoutForm);