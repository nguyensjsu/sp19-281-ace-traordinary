import React, { Component } from 'react';
import {Image } from "semantic-ui-react"
import "../css/imagecard.css"


//https://i.pinimg.com/236x/4d/f8/58/4df85823d89a34522dabf8dd49cdfbd8.jpg
class ImageCard extends Component {
    render() {
        return (
            <div className="Imagecard">
                <Image classNmae='card-image' src={this.props.imagesrc} size='medium' rounded />
                <div className={"buy-button"}>Buy</div>
            </div>
        );
    }
}

export default ImageCard;
