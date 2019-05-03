import React, { Component } from 'react';
import Comment from "./Comment";
import {Image,Grid,TextArea,Form,Button,Icon } from "semantic-ui-react"
import {testcomments} from '../resources/TestResourse'
import {Link} from "react-router-dom";
import ImageBlockChain from "./ImageBlockChain";
import {deleteimage,a} from '../actions/ImageAction'
class ViewImage extends Component {


    constructor(props){
        super(props)
        this.state={
            imgurl:"https://i.pinimg.com/236x/4d/f8/58/4df85823d89a34522dabf8dd49cdfbd8.jpg",
            img:{},
            newcomment:""
        }
        this.comment=this.comment.bind(this);
    }

    componentDidMount(){
        const {imageurl,img} = this.props.location.state
        this.setState({
            imgurl:imageurl,
            img:img
        })
    }
    comment(){
        if(this.state.newcomment.length>5){

        }
    }
    render() {

        let comments=testcomments;
        
        let commentslist =comments.map(comment=>{
           return <><Comment name ={comment.name} timestamp={comment.timestamp}comment={comment.comment}/><hr/></>
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
                                <TextArea placeholder='Your Comment goes here' style={{ minHeight: 100 }} />
                            </Form>
                            <Button primary>Comment</Button>
                        </Grid.Column>


                    </Grid.Row>

                </Grid>
                </div>
            </div>
        );
    }
}

export default ViewImage;
