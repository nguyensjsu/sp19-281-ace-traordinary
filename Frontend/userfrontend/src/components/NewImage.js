import React, { Component } from 'react';
import {Image } from "semantic-ui-react"
import { Button, Checkbox, Form ,Select,TextArea} from 'semantic-ui-react'


const options = [
    { key: 'a', text: 'Art', value: 'art' },
    { key: 'p', text: 'Photography', value: 'photography' },
]
class NewImage extends Component {


    constructor(props){
        super(props);
        this.state={
            file:null,
            newimageurl:null
        }
        this.handleChangeImage=this.handleChangeImage.bind(this)
    }
    handleChangeImage(event){
        let file = event.target.files[0];
        console.log(file)
        this.setState({
            file: file,
            newimageurl:URL.createObjectURL(file)
        });
    }
    render() {
        let image =  <Image wrapped size='medium' src='https://s3-us-west-2.amazonaws.com/ravitejakommalapati.com/picasalogo.png' verticalAlign='top' />
if(this.state.newimageurl!=null){
    image=<Image wrapped size='medium' src={this.state.newimageurl} verticalAlign='top' />
}
        return (
            <div className="NewImage">
                <div className={"nemimage"}>
                    <input type="file" accept={"image/png,image/jpg"} onChange={this.handleChangeImage}></input>
                    {image}
                </div>
                <div className={"form"}>
                <Form>
                    <Form.Field>
                        <label>Enter Price</label>
                        <input placeholder='Please Enter Price' />
                    </Form.Field>
                    <Form.Field control={Select} label='Category' options={options} placeholder='Category' />
                    <TextArea placeholder='Descreption about Image' style={{ minHeight: 100 }} />
                </Form>
                </div>
            </div>
        );
    }
}

export default NewImage;
