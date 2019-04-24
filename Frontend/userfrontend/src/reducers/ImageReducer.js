import {ADDIMAGE} from "../actions/ImageAction";




export default function image(state = [] ,action ) {
    switch (action.type) {
        case ADDIMAGE:
            return action.image
        case GETIMAGE:
            return action.image;
        default:
            return state;
    }
}

