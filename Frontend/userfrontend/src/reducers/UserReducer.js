import {USER_REGESTRATION,USER_LOGIN,USER_FORGOT_PASSWORD,FIREBASE_LOGIN} from "../actions/UserAction";

export default function user(state = {} ,action ) {
    switch (action.type) {
        case USER_REGESTRATION:
            return action.user
        case USER_LOGIN:
            return action.user;
        case USER_FORGOT_PASSWORD:
            return action.user;
        case FIREBASE_LOGIN:
            return action.user;

        default:
            return state;
    }
}

