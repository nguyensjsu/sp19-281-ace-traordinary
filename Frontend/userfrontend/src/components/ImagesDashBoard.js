import React, { Component } from 'react';
import '../css/imagedashboard.css';
import ImageCard from "./ImageCard";
import InfiniteScroll from 'react-infinite-scroller';
import {getallimages} from '../actions/ImageAction'
import {testimages} from "../resources/TestResourse"
import connect from "react-redux/es/connect/connect";

var imagecards=[]

class ImagesDashBoard extends Component {
    constructor(props) {
        super(props);

        this.state = {
            totalcount:1000,
            hasMoreItems: true,
            nextHref: null
        };
        this.loadimages=this.loadimages.bind(this);
    }
componentWillMount(){
        this.props.getallimages(1)
    this.setState({
        hasMoreItems:false
    })
}
    loadimages=(page)=> {
           this.props.getallimages(2)
    }
    render() {
        const loader = <div className="loader">Loading ...</div>;
        let allimages  = this.props.allimages;
        imagecards= allimages.map(image=> <ImageCard key={image.imageid}  likecount={10} commentcount={10} isliked={true} img={image}/>
         )
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

function mapStateToProps(state) {
    return{
        allimages:state.images.allimages,
    }
}
export default connect(mapStateToProps,{getallimages})(ImagesDashBoard);
