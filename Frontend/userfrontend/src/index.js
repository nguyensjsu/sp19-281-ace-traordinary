import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import {Provider} from 'react-redux'
import rootReducer from './reducers'
import promise from "redux-promise";
import {applyMiddleware, compose, createStore} from "redux";


function savetoLocalStorage(state) {

    try{
        const resstate = JSON.stringify(state);
        localStorage.setItem('state',resstate)

    }catch (e) {
        console.log(e);

    }
}

function loadfromLocalStorage() {
    try{
        const state = localStorage.getItem('state');
        if(state===null)return undefined;
        return JSON.parse(state);
    }catch (e) {
        console.log(e);
        return undefined;
    }
}

const persistedState =loadfromLocalStorage();
const composePlugin = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const createStoreWithMiddleware = applyMiddleware(promise)(createStore);
//For Redux
const store = createStore(rootReducer, persistedState,composePlugin(applyMiddleware(promise)));
store.subscribe(()=>savetoLocalStorage(store.getState()))

ReactDOM.render( <Provider store={store}><App /></Provider>, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
