import React, { Component } from 'react';
import Comment from "./Comment";
import {Image } from "semantic-ui-react"
import {testcomments} from '../resources/TestResourse'
class ViewImage extends Component {


    constructor(props){
        super(props)
        this.state={
            imgurl:"https://i.pinimg.com/236x/4d/f8/58/4df85823d89a34522dabf8dd49cdfbd8.jpg"
        }
    }
    componentDidMount(){
        const {imageurl} = this.props.location.state
        this.setState({
            imgurl:imageurl
        })
    }
    render() {

        let comments=testcomments;
        
        let commentslist =comments.map(comment=>{
           return <Comment name ={comment.name} timestamp={comment.timestamp}comment={comment.comment}/>
        })
        return (
            <div className="ViewImage">
                <Image classNmae='card-image' src={this.state.imgurl} size='medium' rounded />
                <div>
                    {commentslist}
                </div>
<div>
    
</div>
            </div>
        );
    }
}

export default ViewImage;
