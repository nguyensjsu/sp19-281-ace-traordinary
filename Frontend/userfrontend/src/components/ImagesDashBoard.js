import React, { Component } from 'react';
import '../css/imagedashboard.css';
import ImageCard from "./ImageCard";
import {testimages} from "../resources/TestResourse"


class ImagesDashBoard extends Component {

    render() {
        let images  = testimages;
         const image_cards = images.map(image=>{
             return (<ImageCard imagesrc={image.imageurl} likecount={image.likecount} commentcount={image.comment} isliked={image.isliked}/>)
         })
        return (

            <div className="ImagesDashBoard">
                {image_cards}
            </div>
        );
    }
}

export default ImagesDashBoard;
