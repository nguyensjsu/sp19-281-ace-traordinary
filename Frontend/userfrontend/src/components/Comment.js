import React, { Component } from 'react';
import {Card } from "semantic-ui-react"


class Comment extends Component {
    render() {
        return (
            <div className="Comment">
                <div className={"header"}>
                    <span><b> {this.props.name}</b></span>
                    <span>{this.props.timestamp}</span>
                </div>
                <div className={"comment"}> 
                    <p>{this.props.comment}</p>
                </div>
            </div>
        );
    }
}

export default Comment;
