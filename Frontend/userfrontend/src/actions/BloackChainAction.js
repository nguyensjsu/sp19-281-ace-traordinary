import {BLOCKCHAIN_ROOTURL} from '../resources/constants'

export const ADDTOCHAIN="ADDTOCHAIN"
export const GETBLOCKCHANIN="GETBLOCKCHANIN"

export const addblockchain= async (data)=>{
    const response = await axios.post(`${BLOCKCHAIN_ROOTURL}/images`, inuser);
    let chain =response.data;
    const action={
        type:ADDTOCHAIN,
        chain
    }
    return action;
}
export const getblockchain= async (imageid)=>{
const response = await axios.get(`${BLOCKCHAIN_ROOTURL}/images`,{
    params:{
        imageid:imageid
    }
});

let chain=response.data;
console.log(JSON.stringify(chain))
const action={
    type:GETBLOCKCHANIN,
    chain
}
return action;
}
