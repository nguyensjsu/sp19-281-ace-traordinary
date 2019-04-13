import React, { Component } from 'react';

class Navigation extends Component {
    render() {
        return (
            <div className="Navigation">
                <ul><li>Home</li>
                 <li>MyImages</li>
                 <li>Buys</li>
                  <li>Logout</li>
                </ul>
            </div>
        );
    }
}

export default Navigation;
