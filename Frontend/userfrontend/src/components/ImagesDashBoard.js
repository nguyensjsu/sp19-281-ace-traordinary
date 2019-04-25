import React, { Component } from 'react';
import '../css/imagedashboard.css';
import ImageCard from "./ImageCard";
import InfiniteScroll from 'react-infinite-scroller';
import {testimages} from "../resources/TestResourse"

var imagecards=[]

class ImagesDashBoard extends Component {
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
        let images  = testimages;
        this.state.tracks.map(image=>{
            imagecards.push(<ImageCard imagesrc={image.imageurl} likecount={image.likecount} commentcount={image.comment} isliked={image.isliked} buyoption={true}/>)
         })
        return (

            <div className="ImagesDashBoard">
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

export default ImagesDashBoard;
