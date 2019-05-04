import React, { Component } from 'react';
import ImageBlockChain from "./ImageBlockChain";
import Payment from "./Payment";
import {Image, Grid, Icon, Segment} from "semantic-ui-react"
import {Link} from "react-router-dom";
import connect from "react-redux/es/connect/connect";
import {viewimage} from "../actions/ImageAction";
class BuyImage extends Component {
    
    
    constructor(props){
        super(props)
        this.state={
            imgurl:"",
            img:{}
        }
    }
    componentWillMount(){
        const {imageurl,img} = this.props.location.state
        this.setState({
            imgurl:imageurl,
            img:img
        })
    }
    render() {
        return (
            <div className="BuyImage">
                <div className={"BuyImageSecond"}>
                    <Link to={"/"} className={"link"}><Icon  name='angle left' size='huge' color={"blue"} className={"back"}/></Link>

                    <Grid columns={3} >
                    <Grid.Row>
                    <Grid.Column>
                <Image classNmae='card-image' src={this.state.imgurl} size='medium' rounded />
                    </Grid.Column>
                        <Grid.Column>
                <ImageBlockChain description={this.state.img.description}/>
                        </Grid.Column>
                            <Grid.Column>
                                <Segment inverted color={"blue"} raised><b>Payment</b></Segment>
                            <Payment img={this.state.img} user={this.props.user}/>
                            </Grid.Column>
                    </Grid.Row>
                </Grid>
                </div>
            </div>
        );
    }
}
function mapStateToProps(state) {
    return {
        image: state.images.viewingimage,
        user:state.user
    }
}
export default connect(mapStateToProps, {})(BuyImage);
