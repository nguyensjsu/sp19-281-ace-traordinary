import React, { Component } from 'react';
import {Image } from "semantic-ui-react"
import { Button, Checkbox, Form ,Select} from 'semantic-ui-react'

const options = [
    { key: 'a', text: 'Art', value: 'art' },
    { key: 'p', text: 'Photography', value: 'photography' },
]
class NewImage extends Component {


    constructor(props){
        super(props);
        this.state={
            
        }
    }
    render() {
        return (
            <div className="NewImage">
                <Image wrapped size='medium' src='https://s3-us-west-2.amazonaws.com/ravitejakommalapati.com/picasalogo.png' verticalAlign='top' />
                <Form>
                    <Form.Field>
                        <label>Enter Price</label>
                        <input placeholder='Please Enter Price In $' />
                    </Form.Field>
                    <Form.Field control={Select} label='Category' options={options} placeholder='Category' />
                    <Button type='submit'>Submit</Button>
                </Form>
            </div>
        );
    }
}

export default NewImage;
