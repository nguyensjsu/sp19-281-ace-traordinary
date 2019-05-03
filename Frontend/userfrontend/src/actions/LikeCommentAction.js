import axios from "axios";
import {LIKECOMMENT_ROOTURL} from '../resources/constants'

export const LIKEIMAGE ="LIKEIMAGE"
export const COMMENTIMAGE ="COMMENTIMAGE"
export const DELETECOMMENT ="DELETECOMMENT"
export const UNLIKEIMAGE ="UNLIKEIMAGE"
export const GETREACTION ="GETREACTION"


export const likeimage= async (data)=>{
    const response = await axios.post(`${LIKECOMMENT_ROOTURL}/images`, data);
    let reaction =response.data;
    const action={
        type:LIKEIMAGE,
        reaction
    }
    return action;
}
export const comment= async (data)=>{
    const response = await axios.post(`${LIKECOMMENT_ROOTURL}/images`, data);
    let reaction =response.data;
    const action={
        type:COMMENTIMAGE,
        reaction
    }
    return action;
}

export const getreaction = async (imageid)=>{
    const response = await axios.get(`${LIKECOMMENT_ROOTURL}/images`,{
        params:{
            imageid:imageid
        }
    });

    let reaction=response.data;
    console.log(JSON.stringify(reaction))
    const action={
        type:GETREACTION,
        reaction
    }
    return action;
}

export const deletecomment = async (data)=>{
    const response = await axios.delete(`${LIKECOMMENT_ROOTURL}/images`,{
        params:{
            imageid:data.imageid,
            commentid:data.commentid,
            userid:data.userid
        }
    });

    let reaction=response.data;
    console.log(JSON.stringify(reaction))
    const action={
        type:DELETECOMMENT,
        reaction
    }
    return action;
}