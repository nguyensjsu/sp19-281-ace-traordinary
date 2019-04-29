import React, { Component } from 'react';
import InfiniteScroll from 'react-infinite-scroller';
import {testimages} from "../resources/TestResourse"
import '../css/myImages.css';
import {getallmyimages} from '../actions/ImageAction'
import ViewImage from "./ViewImage"
import ImageCard from "./ImageCard";
import connect from "react-redux/es/connect/connect";
var imagecards=[]

class MyImages extends Component {

    constructor(props) {
        super(props);

        this.state = {
            tracks: testimages,
            totalcount:50,
            hasMoreItems: false,
            nextHref: null
        };
        this.loadimages=this.loadimages.bind(this);
    }
componentWillMount(){
        this.props.getallmyimages(this.props.user.userid)
}
    loadimages=(page)=> {

    }
    render() {
        const loader = <div className="loader">Loading ...</div>;
        imagecards=this.props.myimages.map(image=> <ImageCard  key={image.imageid} imagesrc={image.origurl} img={image}likecount={image.likecount} commentcount={image.comment} isliked={image.isliked} buyoption={false} />)
        return (
            <div className="MyImages">
                <InfiniteScroll
                    pageStart={0}
                    loadMore={this.loadimages}
                    hasMore={this.state.hasMoreItems}
                    loader={loader}>

                    <div className="tracks">
                        {imagecards}
                    </div>
                </InfiniteScroll>
            </div>
        );
    }
}
function mapStateToProps(state) {
    return{
        myimages:state.images.myimages,
        user:state.user
    }
}
export default connect(mapStateToProps,{getallmyimages})(MyImages);
