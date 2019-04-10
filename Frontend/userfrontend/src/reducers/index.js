import {combineReducers} from 'redux';

const appReducer =combineReducers(
    {}
);

const rootReducer = (state, action) => {
    if (action.type === 'USER_LOGOUT') {
        state = undefined;
        localStorage.clear();
       // cookie.remove('user');
    }

    return appReducer(state, action)
}





export  default  rootReducer;