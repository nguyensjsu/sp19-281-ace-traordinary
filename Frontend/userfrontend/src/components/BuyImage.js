import React, { Component } from 'react';
import ImageBlockChain from "./ImageBlockChain";
import Payment from "./Payment";
import {Image ,Grid,Icon} from "semantic-ui-react"
import {Link} from "react-router-dom";
class BuyImage extends Component {
    
    
    constructor(props){
        super(props)
        this.state={
            imgurl:""
        }
    }
    componentDidMount(){
        const {imageurl} = this.props.location.state
        this.setState({
            imgurl:imageurl
        })
    }
    render() {
       // const { imageurl } = this.props.match.params
        console.log(this.state.imgurl)
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
                <ImageBlockChain/>
                        </Grid.Column>
                            <Grid.Column>
                <Payment/>
                            </Grid.Column>
                    </Grid.Row>
                </Grid>
                </div>
            </div>
        );
    }
}

export default BuyImage;
