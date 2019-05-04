import React, { Component } from 'react';
import CheckoutForm from './CheckoutForm'
import {Elements, StripeProvider} from 'react-stripe-elements';
class Payment extends Component {

    render() {
        return (
            <StripeProvider apiKey="pk_test_TYooMQauvdEDq54NiTphI7jx">
            <div className="Payment">
                <Elements>
              <CheckoutForm img={this.props.img} user={this.props.user}/>
                </Elements>
            </div>
            </StripeProvider>
        );
    }
}

export default Payment;
