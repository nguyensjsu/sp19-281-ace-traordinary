import React, { Component } from 'react';
import ImageBlockChain from "./ImageBlockChain";
import Payment from "./Payment";

import {Image } from "semantic-ui-react"

class BuyImage extends Component {
    
    
    constructor(props){
        super(props)
        this.state={
            imgurl:""
        }
    }
    componentDidMount(){
        const {imageurl} = this.props.location.state
        this.setState({
            imgurl:imageurl
        })
    }
    render() {
       // const { imageurl } = this.props.match.params
        console.log(this.state.imgurl)
        return (
            <div className="BuyImage">
                <Image classNmae='card-image' src={this.state.imgurl} size='medium' rounded />
                <ImageBlockChain/>
                <Payment/>
            </div>
        );
    }
}

export default BuyImage;
