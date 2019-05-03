import React, { Component } from 'react';
import {Card } from "semantic-ui-react"
import { Icon } from 'semantic-ui-react'
import {LIKECOMMENT_ROOTURL} from "../resources/constants";
import axios from "axios";


class Comment extends Component {

    constructor(props){
        super(props);
        this.deletecomment=this.deletecomment.bind(this)
    }
deletecomment(){
    let image_id=this.props.imageid
    let userid=this.props.userid;
    let commentid=this.props.comment.CommentId
    let url=`${LIKECOMMENT_ROOTURL}/removecomment/${image_id}/${userid}/${commentid}`
    axios.delete(url).then(res=>{
        const comments = res.data.Comments;
        this.props.deletecomment(comments)
    });
}
    render() {
        return (
            <div className="Comment">
                <div className={"header"}>
                    <span><b> {this.props.comment.Username}</b></span>
                    <span>{this.props.comment.TimeStamp}</span>
                    <span><Icon disabled name='close' color={"red"} onClick={this.deletecomment}/></span>
                </div>
                <div className={"comment"}> 
                    <p>{this.props.comment.Comment}</p>
                </div>
            </div>
        );
    }
}

export default Comment;
