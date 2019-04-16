import React, { Component } from 'react';
import ImageBlockChain from "./ImageBlockChain";
import Payment from "./Payment";

class BuyImage extends Component {
    render() {
        return (
            <div className="BuyImage">
            <div >Hi I got called</div>
                <ImageBlockChain/>
                <Payment/>
            </div>
        );
    }
}

export default BuyImage;
