import React from 'react';
import {injectStripe} from 'react-stripe-elements';

import CardSection from './CardSection';

class CheckoutForm extends React.Component {
  constructor(props) {
  super(props);
  this.state = {complete: false};
  this.submit = this.submit.bind(this);
}
  handleSubmit = (ev) => {
    // We don't want to let default form submission happen here, which would refresh the page.
    ev.preventDefault();

    // Within the context of `Elements`, this call to createToken knows which Element to
    // tokenize, since there's only one in this group.
    this.props.stripe.createToken({name: 'Jenny Rosen'}).then(({token}) => {
      console.log('Received Stripe token:', token);

    });



    // However, this line of code will do the same thing:
    //
    // this.props.stripe.createToken({type: 'card', name: 'Jenny Rosen'});

    // You can also use createSource to create Sources. See our Sources
    // documentation for more: https://stripe.com/docs/stripe-js/reference#stripe-create-source
    //
    // this.props.stripe.createSource({type: 'card', owner: {
    //   name: 'Jenny Rosen'
    // }});
    
  };

  async submit(ev) {
      let {token} = await this.props.stripe.createToken({name: "Jenny Rosen"});
      let response = await fetch("/charge", {
        method: "POST",
        headers: {"Content-Type": "text/plain"},
        body: token.id
      });

      if (response.ok) console.log("Purchase Complete!")
    }

  render() {
    return (
      <div id="ch">
        <form onSubmit={this.handleSubmit}>
        
          <CardSection />
          <button onClick={this.submit}>Confirm order</button>
        </form>
      </div>  
    );
  }
}

export default injectStripe(CheckoutForm);