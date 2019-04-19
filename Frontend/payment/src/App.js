import React, {Component} from 'react';
import {Elements, StripeProvider} from 'react-stripe-elements';
import {render} from 'react-dom';
import CheckoutForm from './CheckoutForm';

class App extends Component {
  render() {
    return (
      <StripeProvider apiKey="pk_test_TYooMQauvdEDq54NiTphI7jx">
        <div className="chform">
          <h3>Payment page</h3>
          <Elements>
            <CheckoutForm />
          </Elements>
        </div>
      </StripeProvider>
    );
  }
}

export default App;
