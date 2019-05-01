import {combineReducers} from 'redux';
import user from './UserReducer'
import images from './ImageReducer'
import {USER_LOGOUT} from "../actions/UserAction";

const appReducer =combineReducers(
    {user,images}
);

const rootReducer = (state, action) => {

    if (action.type === 'USER_LOGOUT') {
        state = undefined;
    }

    return appReducer(state, action)
}





export  default  rootReducer;