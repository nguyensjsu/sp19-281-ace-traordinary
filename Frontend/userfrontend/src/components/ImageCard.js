import React, { Component } from 'react';
import {LIKECOMMENT_ROOTURL}from "../resources/constants"
import { Link } from "react-router-dom";
import {Image } from "semantic-ui-react"
import { Segment, Icon } from 'semantic-ui-react'
import connect from "react-redux/es/connect/connect";
import {deleteimage,viewimage} from '../actions/ImageAction'
import axios from "axios";

import "../css/imagecard.css"


class ImageCard extends Component {

constructor(props){
    super(props)
    this.state={
        likes:[],
        comments:[],
        likecount:0,
        commentcount:0,
        isliked:false
    }
    this.like =this.like.bind(this);
    this.unlike =this.unlike.bind(this);
}
    componentDidMount(){
        let ImageId =this.props.img.imageid;
        let userid=this.props.user.userid;
        let url=`${LIKECOMMENT_ROOTURL}/reaction/${ImageId}`
        axios.get(url)
            .then(res => {
                const likes = res.data.Likes;
                const comments = res.data.Comments
                let likeslength=0;
                if(likes!==undefined){
                    likeslength=likes.length;
                }
                let commentslength=0;
                if(comments!==undefined)
                    commentslength=comments.length
                console.log("this are comments",res.data)
                this.setState({ likes:likes,likecount:likeslength,comments:comments,commentcount:commentslength });
            }).catch(err=>{
            console.log("No data Available")
        });
        let likeurl=`${LIKECOMMENT_ROOTURL}/images/${ImageId}/user/${userid}`
        axios.get(likeurl)
            .then(res => {
                const isliked = res.data;
                this.setState({ isliked:isliked });
            });
    }
    async unlike(){

           let image_id=this.props.img.imageid
               let userid=this.props.user.userid;

        let url=`${LIKECOMMENT_ROOTURL}/unlike/${image_id}/${userid}`
        axios.delete(url).then(res=>{
            const likes = res.data.Likes;
            this.setState({ likes:likes,likecount:likes.length,isliked:false });
        });

    }
    async like(){
        let reqdata={
            image_id:this.props.img.imageid,
            userid:this.props.user.userid,
            username:this.props.user.firstname}
        let url=`${LIKECOMMENT_ROOTURL}/like`
        let res= await axios.post(url, reqdata);
        const likes = res.data.Likes;
        this.setState({ likes:likes,likecount:likes.length,isliked:true });
    }
    render() {
        let img=this.props.img;
        let buylink;
        let deleteimage;
        let like =<Icon className="heart outline icon likebutton" onClick={this.like}></Icon>
        if(this.state.isliked) {
            like = <Icon className="heart  icon inverted likebutton" color='red' onClick={this.unlike}></Icon>
        }

        if(img.userid==this.props.user.userid){
            deleteimage =<div className={"delete-icon"} onClick={()=>this.props.deleteimage(img.imageid)}> <Icon className = "trash icon " size='large' color='red'></Icon></div>
        }else{
            buylink =<Link to={{pathname:'/images/buy',
                state:{
                    imageurl:img.origurl,
                    img:img
                }
            }}className={"link"}> <div className={"buy-button"} >Buy{img.price}$</div></Link>
        }
        return (
            <div className="Imagecard">
                <div>
                <Image classNmae='card-image' src={img.origurl} size='medium' rounded />

                    {deleteimage}
                    {buylink}
                </div>
                <div className={"lc-container"}>
                <div className={"like-container"}><span>{like} {this.state.likecount}</span></div>
                    <Link to={{pathname:'/images/comment',
                        state:{
                            imageurl:img.origurl,
                            img:img,
                            comments:this.state.comments
                        }
                    }} className={"link"}>
                        <div className={"comment-container"}><span><Icon className="comment outline" ></Icon> {this.state.commentcount}</span></div></Link>
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
export default connect(mapStateToProps,{deleteimage,viewimage})(ImageCard);
