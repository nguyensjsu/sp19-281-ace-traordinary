import React, { Component } from 'react';
import '../css/imagedashboard.css';
import ImageCard from "./ImageCard";
//https://i.pinimg.com/236x/5f/67/86/5f6786f7e998ed17f059155561378ff2.jpg?b=t
//https://i.pinimg.com/236x/4d/f8/58/4df85823d89a34522dabf8dd49cdfbd8.jpg

class ImagesDashBoard extends Component {
    render() {
        return (
            <div className="ImagesDashBoard">
                <ImageCard imagesrc={"https://i.pinimg.com/236x/4d/f8/58/4df85823d89a34522dabf8dd49cdfbd8.jpg"}/>
                <ImageCard imagesrc={"https://i.pinimg.com/236x/5f/67/86/5f6786f7e998ed17f059155561378ff2.jpg?b=t"}/>


                <ImageCard imagesrc={"https://i.pinimg.com/236x/4d/f8/58/4df85823d89a34522dabf8dd49cdfbd8.jpg"}/>
                <ImageCard imagesrc={"https://i.pinimg.com/236x/4d/f8/58/4df85823d89a34522dabf8dd49cdfbd8.jpg"}/>
                <ImageCard imagesrc={"https://i.pinimg.com/236x/5f/67/86/5f6786f7e998ed17f059155561378ff2.jpg?b=t"}/>
                <ImageCard imagesrc={"https://i.pinimg.com/236x/5f/67/86/5f6786f7e998ed17f059155561378ff2.jpg?b=t"}/>
                <ImageCard imagesrc={"https://i.pinimg.com/236x/4d/f8/58/4df85823d89a34522dabf8dd49cdfbd8.jpg"}/>
                <ImageCard imagesrc={"https://i.pinimg.com/236x/4d/f8/58/4df85823d89a34522dabf8dd49cdfbd8.jpg"}/>
                <ImageCard imagesrc={"https://i.pinimg.com/236x/5f/67/86/5f6786f7e998ed17f059155561378ff2.jpg?b=t"}/>
                <ImageCard imagesrc={"https://i.pinimg.com/236x/5f/67/86/5f6786f7e998ed17f059155561378ff2.jpg?b=t"}/>

            </div>
        );
    }
}

export default ImagesDashBoard;
