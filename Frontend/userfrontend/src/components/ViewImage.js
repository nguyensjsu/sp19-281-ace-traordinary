import React, { Component } from 'react';
import Comment from "./Comment";
import {Image,Grid,TextArea,Form,Button } from "semantic-ui-react"
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
           return <><Comment name ={comment.name} timestamp={comment.timestamp}comment={comment.comment}/><hr/></>
        })
        return (
            <div className="ViewImage">
                <div className={"ViewImageSecond"}>
                <Grid columns={3} >
                    <Grid.Row>
                        <Grid.Column>
                            <Image  src={this.state.imgurl} size='medium' rounded />
                        </Grid.Column>
                        <Grid.Column></Grid.Column>
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
