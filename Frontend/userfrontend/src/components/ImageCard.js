import React, { Component } from 'react';
import { Link } from "react-router-dom";
import {Image } from "semantic-ui-react"
import "../css/imagecard.css"



class ImageCard extends Component {
    render() {
        let imageID ="IR77bjSuubjdk9jduHHg"
        return (
            <div className="Imagecard">
                <div>
                <Image classNmae='card-image' src={this.props.imagesrc} size='medium' rounded />
                <Link to={{pathname:'/images/buy',
                state:{
                    imageurl:this.props.imagesrc
                }
                }}> <div className={"buy-button"}>Buy</div></Link>
                </div>
                <div className={"lc-container"}>
                <div className={"like-container"}><span>Like {this.props.likecount}</span></div>
                    <div className={"comment-container"}><span>Comment {this.props.commentcount}</span></div>
                </div>
            </div>
        );
    }
}

export default ImageCard;
