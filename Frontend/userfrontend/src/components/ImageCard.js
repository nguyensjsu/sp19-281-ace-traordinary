import React, { Component } from 'react';
import { Link } from "react-router-dom";
import {Image } from "semantic-ui-react"
import "../css/imagecard.css"



class ImageCard extends Component {
    render() {
        let imageID ="IR77bjSuubjdk9jduHHg"
        return (
            <div className="Imagecard">
                <Image classNmae='card-image' src={this.props.imagesrc} size='medium' rounded />
                <Link to={"/images/buy"}> <div className={"buy-button"}>Buy</div></Link>

            </div>
        );
    }
}

export default ImageCard;
