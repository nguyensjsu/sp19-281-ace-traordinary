import React, { Component } from 'react';
import { Button, Checkbox, Form ,Select,TextArea,Modal,Label,Menu,Image,Header} from 'semantic-ui-react'
import {addimage} from '../actions/ImageAction'
import connect from "react-redux/es/connect/connect";

const options = [
    { key: 'a', text: 'Art', value: 'Art' },
    { key: 'p', text: 'Photography', value: 'Photography' },
]
class NewImage extends Component {
    stat = {open: false}
    show = dimmer => () => this.setState({dimmer, open: true})
    close = () => this.setState({open: false})

    constructor(props){
        super(props);
        this.state={
            file:null,
            newimageurl:null,
            price:0,
            description:"",
            title:"Art"

        }
        this.handleChangeImage=this.handleChangeImage.bind(this)
        this.handleChage =this.handleChage.bind(this)
        this.handleSubmit =this.handleSubmit.bind(this)
    }
    handleChage(event){
        this.setState({[event.target.name]:event.target.value})
    }
    handleChangeImage(event){
        let file = event.target.files[0];
        console.log(file)
        this.setState({
            file: file,
            newimageurl:URL.createObjectURL(file)
        });
    }
    handleSubmit(){
        let formdata = new FormData()
        formdata.append('userid', this.props.user.userid);
        formdata.append('description', this.state.description);
        formdata.append('price', this.state.price);
        formdata.append('title', this.state.title);
        formdata.append('myfile', this.state.file);
        this.props.addimage(formdata)
    }
    render() {
        const { open, dimmer } = this.state
        let image =  <Image wrapped size='medium' src='https://s3-us-west-2.amazonaws.com/ravitejakommalapati.com/picasalogo.png' verticalAlign='top' />
if(this.state.newimageurl!=null){
    image=<Image wrapped size='medium' src={this.state.newimageurl} verticalAlign='top' />
}

let imageform=<div className="NewImage">
    <div className={"nemimage"}>
        <input type="file" accept={"image/png,image/jpg"} onChange={this.handleChangeImage}></input>
        {image}
    </div>
    <div className={"form"}>
        <Form>
            <Form.Field>
                <label>Enter Price</label>
                <input placeholder='Please Enter Price' name={"price"} onChange={this.handleChage}/>
            </Form.Field>
            <Form.Field control={Select} label='Category'  options={options} placeholder='Category' name={"title"}/>
            <TextArea placeholder='Descreption about Image' name="description" style={{ minHeight: 100 }}onChange={this.handleChage} />
        </Form>
    </div>
</div>
        return (
            <>
            <li onClick={this.show('blurring')}>NewImage</li>
            <Modal dimmer={dimmer} open={open} onClose={this.close}>
                <Modal.Header>Upload New Image</Modal.Header>
                <Modal.Content image>
                    {imageform}
                </Modal.Content>
                <Modal.Actions>
                    <Button color='red' onClick={this.close}>
                        Cancel
                    </Button>
                    <Button
                        color='green'
                        content="Upload"
                        onClick={this.handleSubmit}
                    />
                </Modal.Actions>
            </Modal>
                </>
        );
    }
}
function mapStateToProps(state) {
    return{
        myimages:state.images.myimages,
        user:state.user
    }
}
export default connect(mapStateToProps,{addimage})(NewImage);
