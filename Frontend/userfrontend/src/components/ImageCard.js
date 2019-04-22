import React, { Component } from 'react';
import { Link } from "react-router-dom";
import {Image } from "semantic-ui-react"
import "../css/imagecard.css"
import { Segment, Icon } from 'semantic-ui-react'




class ImageCard extends Component {
    render() {
        let imageID ="IR77bjSuubjdk9jduHHg"
        
        let like =<Icon className="heart outline icon likebutton"></Icon>
        if(this.props.isliked) {
            like = <Icon className="heart  icon inverted likebutton" color='red'></Icon>
        }
        return (
            <div className="Imagecard">
                <div>
                <Image classNmae='card-image' src={this.props.imagesrc} size='medium' rounded />
                <Link to={{pathname:'/images/buy',
                state:{
                    imageurl:this.props.imagesrc
                }
                }} className={"link"}> <div className={"buy-button"}>Buy</div></Link>
                </div>
                <div className={"lc-container"}>
                <div className={"like-container"}><span>{like} {this.props.likecount}</span></div>
                    <Link to={{pathname:'/images/comment',
                        state:{
                            imageurl:this.props.imagesrc
                        }
                    }} className={"link"}>
                        <div className={"comment-container"}><span><Icon className="comment outline" ></Icon> {this.props.commentcount}</span></div></Link>
                </div>
            </div>
        );
    }
}

export default ImageCard;
