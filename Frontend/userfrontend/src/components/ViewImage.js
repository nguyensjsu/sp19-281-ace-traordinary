import React, { Component } from 'react';
import Comment from "./Comment";
import {Image,Grid,TextArea,Form,Button,Icon } from "semantic-ui-react"
import {testcomments} from '../resources/TestResourse'
import {Link} from "react-router-dom";
import ImageBlockChain from "./ImageBlockChain";
import {deleteimage, a, viewimage} from '../actions/ImageAction'
import {LIKECOMMENT_ROOTURL}from "../resources/constants"
import axios from "axios";
import connect from "react-redux/es/connect/connect";

class ViewImage extends Component {


    constructor(props){
        super(props)
        this.state={
            imgurl:"https://i.pinimg.com/236x/4d/f8/58/4df85823d89a34522dabf8dd49cdfbd8.jpg",
            img:{},
            newcomment:"",
            comments:[]
        }
        this.comment=this.comment.bind(this);
        this.deletecomment=this.deletecomment.bind(this);
        this.handleChange =this.handleChange.bind(this);
    }

    componentWillMount(){
        const {imageurl,img,comments} = this.props.location.state
        this.setState({
            imgurl:imageurl,
            img:img,
            comments:comments
        })
    }
    handleChange(event){
        this.setState({[event.target.name]:event.target.value})
    }
    deletecomment(comments){
        this.setState({
            comments:comments
        })
    }
    comment(){
        if(this.state.newcomment.length>5){
            let url=`${LIKECOMMENT_ROOTURL}/comment`;
            let reqdata={
                image_id:this.state.img.imageid,
                userid:this.props.user.userid,
                username:this.props.user.firstname,
                comment:this.state.newcomment}
            axios.post(url,reqdata).then(res=>{
                let comments=res.data.Comments;
                this.setState({comments:comments})
                console.log("New Comment Respose"+res.data)
            })
        }
    }
    render() {
        let commentslist;
        let comments=this.state.comments;
        if(comments!==undefined)
         commentslist =comments.map(comment=>{
           return <><Comment key={comment.CommentId}comment ={comment} userid={this.props.user.userid}imageid={this.state.img.imageid} deletecomment={this.deletecomment}/><hr/></>
        })
        return (
            <div className="ViewImage">
                <div className={"ViewImageSecond"}>
                    <Link to={"/"} className={"link"}><Icon  name='angle left' size='huge' color={"blue"} className={"back"}/></Link>
                <Grid columns={3} >
                    <Grid.Row>
                        <Grid.Column>
                            <Image  src={this.state.imgurl} size='medium' rounded />
                        </Grid.Column>
                        <Grid.Column>
                            <ImageBlockChain description={this.state.img.description}/>
                        </Grid.Column>
                        <Grid.Column>
                            <div className={"commentsdisplay"}>
                            {commentslist}
                            </div>
                            <Form>
                                <TextArea placeholder='Your Comment goes here' style={{ minHeight: 100 }} name={"newcomment"} onChange={this.handleChange}/>
                            </Form>
                            <Button primary onClick={this.comment}>Comment</Button>
                        </Grid.Column>


                    </Grid.Row>

                </Grid>
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
export default connect(mapStateToProps,{})(ViewImage);
