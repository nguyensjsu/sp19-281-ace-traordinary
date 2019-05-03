import {ADDIMAGE,GETALLIMAGES,GETALLMYIMAGES,DELETEIMAGE,VIEWIMAGE} from "../actions/ImageAction";





let images={
    myimages:[],
    allimages:[],
    viewingimage:{}
}


export default function image(state = images ,action ) {
    switch (action.type) {
        case ADDIMAGE:
            state.myimages.push(action.resimage)
            return {...state}
        case GETALLIMAGES:
            return {...state,allimages:action.allimages};
        case GETALLMYIMAGES:
            return {...state,myimages:action.myimages};
        case VIEWIMAGE:
            return {...state,viewingimage:action.image};
        case DELETEIMAGE:
            let myimages=state.myimages.filter(image=>image.imageid!==action.imageid)
            let allimages=state.allimages.filter(image=>image.imageid!==action.imageid)
            return {allimages:allimages,myimages:myimages};
        default:
            return state;
    }
}

