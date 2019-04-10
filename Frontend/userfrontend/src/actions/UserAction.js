import axios from "axios";

export const USER_REGESTRATION ="USER_REGESTRATION"
export const USER_LOGIN="USER_LOGIN"
export const USER_LOGOUT="USER_LOGOUT"
export const USER_FORGOT_PASSWORD="USER_FORGOT_PASSWORD"


export const userregestration= async (user)=>{

    const action={
        type:USER_REGESTRATION,
        user
    }
    return action;
}
export const userlogin= async (user)=>{

    const action={
        type:USER_LOGIN,
        user
    }
    return action;
}
export const userlogout= async (user)=>{

    const action={
        type:USER_LOGOUT,
        user
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