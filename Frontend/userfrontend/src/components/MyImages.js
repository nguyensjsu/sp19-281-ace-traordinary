import React, { Component } from 'react';
import InfiniteScroll from 'react-infinite-scroller';
import {testimages} from "../resources/TestResourse"
import ViewImage from "./ViewImage"
import ImageCard from "./ImageCard";
var imagecards=[]

class MyImages extends Component {

    constructor(props) {
        super(props);

        this.state = {
            tracks: testimages,
            totalcount:50,
            hasMoreItems: true,
            nextHref: null
        };
        this.loadimages=this.loadimages.bind(this);
    }

    loadimages=(page)=> {
        this.setState({
            tracks:testimages
        })
    }
    render() {

        const loader = <div className="loader">Loading ...</div>;
        console.log("IN mY INag"+JSON.stringify(this.state.tracks));

        this.state.tracks.map(image=>{
            imagecards.push(<ImageCard imagesrc={image.imageurl} likecount={image.likecount} commentcount={image.comment} isliked={image.isliked}/>)
        })
        return (
            <div className="MyImages">
                oooovfdsjnjkk
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

export default MyImages;
