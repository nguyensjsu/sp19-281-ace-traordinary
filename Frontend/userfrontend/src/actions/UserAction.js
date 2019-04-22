import axios from "axios";

export const USER_REGESTRATION ="USER_REGESTRATION"
export const USER_LOGIN="USER_LOGIN"
export const USER_LOGOUT="USER_LOGOUT"
export const USER_FORGOT_PASSWORD="USER_FORGOT_PASSWORD"



const USERURL  = "http://localhost:8000";

export const userregestration= async (inuser)=>{
    const response = await axios.post(`${USERURL}/users`, inuser);
let user =response.data;
    const action={
        type:USER_REGESTRATION,
        user
    }
    return action;
}
export const userlogin= async (inuser)=>{
    const response = await axios.post(`${USERURL}/user`, inuser);
    let user =response.data;
    const action={
        type:USER_LOGIN,
        user
    }
    return action;
}
export const userlogout= async ()=>{

    const action={
        type:USER_LOGOUT,
    }
    return action;
}
export const userforgotpassword= async (user)=>{

    const action={
        type:USER_FORGOT_PASSWORD,
        user
    }
    return action;
}