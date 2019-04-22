import React, { Component } from 'react';
import {Card } from "semantic-ui-react"


class Comment extends Component {
    render() {
        return (
            <div className="Comment">
                <Card>
                    <span><b> {this.props.name}</b></span>
                    <span>{this.props.timestamp}</span>
                    <p>{this.props.comment}</p>
                </Card>
            </div>
        );
    }
}

export default Comment;
