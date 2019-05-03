import React, { Component } from 'react';
import { Header, Segment } from 'semantic-ui-react'

class ImageBlockChain extends Component {
    render() {
        return (
            <div className="ImageBlockChain">
                <div>
                    <Header as='h2' attached='top' color='blue'>
                        About Image
                    </Header>
<div>
                        {this.props.description}
</div>
                </div>

            </div>
        );
    }
}

export default ImageBlockChain;
