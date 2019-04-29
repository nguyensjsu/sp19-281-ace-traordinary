import React, { Component } from 'react';
import { Link } from "react-router-dom";
import {Image } from "semantic-ui-react"
import { Segment, Icon } from 'semantic-ui-react'
import connect from "react-redux/es/connect/connect";
import {deleteimage} from '../actions/ImageAction'

import "../css/imagecard.css"


class ImageCard extends Component {
    render() {
        let img=this.props.img;
        let buylink;
        let deleteimage;
        let imageID ="IR77bjSuubjdk9jduHHg"
        let like =<Icon className="heart outline icon likebutton"></Icon>
        if(this.props.isliked) {
            like = <Icon className="heart  icon inverted likebutton" color='red'></Icon>
        }

        if(img.userid==this.props.user.userid){
            deleteimage =<div className={"delete-icon"} onClick={()=>this.props.deleteimage(img.imageid)}> <Icon className = "trash icon " size='large' color='red'></Icon></div>
        }else{
            buylink =<Link to={{pathname:'/images/buy',
                state:{
                    imageurl:this.props.imagesrc
                }
            }}className={"link"}> <div className={"buy-button"}>Buy</div></Link>
        }
        return (
            <div className="Imagecard">
                <div>
                <Image classNmae='card-image' src={img.origurl} size='medium' rounded />

                    {deleteimage}
                    {buylink}
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
function mapStateToProps(state) {
    return{
        user:state.user
    }
}
export default connect(mapStateToProps,{deleteimage})(ImageCard);
