import axios from "axios";
import {IMAGE_ROOTURL,IMAGE_CMDURL} from '../resources/constants'

export const ADDIMAGE="ADDIMAGE"
export const GETIMAGE="GETIMAGE"
export const GETALLIMAGES="GETALLIMAGES"
export const GETALLMYIMAGES="GETALLMYIMAGES"
export const DELETEIMAGE="DELETEIMAGE"
export const VIEWIMAGE="VIEWIMAGE"
export const UPDATEIMAGE="UPDATEIMAGE"

export const addimage = async (imagedata,inclass)=> {
    console.log("In ADD image")
    console.log(imagedata)
    let url=`${IMAGE_CMDURL}/images`
    const response = await axios({method: 'post',
        url: url,
        data: imagedata,
        config: { headers: {'Content-Type': 'multipart/form-data' }}
    }).catch(function (error) {
        if (error.response) {
            console.log(error.response.status);
            console.log(error.response.headers);
        }
    });
    let resimage;
    if (response == undefined) {
        resimage = undefined;
    }else{
        inclass.close();
        resimage=response.data
    }
    console.log(JSON.stringify(resimage))
    const action={
        type:ADDIMAGE,
        resimage
    }
    return action;
}
export const getallimages = async (pagenumber)=> {
    const response = await axios.get(`${IMAGE_ROOTURL}/pictures`, {
        params: {
            pagenumber: pagenumber
        }
    }).catch(function (error) {
        if (error.response) {
            console.log("Error in get all Images");
        }
    });
    let allimages;
    if (response === undefined) {
    allimages = [];
     }else{
        allimages=response.data
    }
    console.log(JSON.stringify(allimages))
    const action={
        type:GETALLIMAGES,
        allimages
    }
    return action;
}
export const getallmyimages = async (userid)=> {
    const response = await axios.get(`${IMAGE_ROOTURL}/users/${userid}/pictures`).catch(function (error) {
        if (error.response) {
            console.log("Error in get all My images");
        }
    });
    let myimages;
    if (response ===undefined) {
        myimages = [];
    }else{
        myimages=response.data
    }
    console.log(JSON.stringify(myimages))
    const action={
        type:GETALLMYIMAGES,
        myimages
    }
    return action;
}
export const deleteimage = async (imageid)=> {
    const response = await axios.delete(`${IMAGE_CMDURL}/images/${imageid}`).catch(function (error) {
        if (error.response) {
            console.log(error.response.status);
            console.log(error.response.headers);
        }
    });
    console.log(JSON.stringify(response))
    const action={
        type:DELETEIMAGE,
        imageid
    }
    return action;
}
export const updateimage = async (image)=> {
    let imageid =image.imageid
    const response = await axios.put(`${IMAGE_CMDURL}/images/${imageid}`,image).catch(function (error) {
        if (error.response) {
            console.log(error.response.status);
            console.log(error.response.headers);
        }
    });
    console.log(JSON.stringify(response))
    const action={
        type:DELETEIMAGE,
        imageid
    }
    return action;
}
export const viewimage=(image)=>{
    const action={
        type:VIEWIMAGE,
        image
    }
}