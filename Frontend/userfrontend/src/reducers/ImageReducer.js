import {ADDIMAGE,GETALLIMAGES,GETALLMYIMAGES,DELETEIMAGE} from "../actions/ImageAction";





let images={
    myimages:[],
    allimages:[]
}


export default function image(state = images ,action ) {
    switch (action.type) {
        case ADDIMAGE:
            return {...state,myimages:action.myimages}
        case GETALLIMAGES:
            return {...state,allimages:action.allimages};
        case GETALLMYIMAGES:
            return {...state,myimages:action.myimages};
        case DELETEIMAGE:
            let myimages=state.myimages.filter(image=>image.imageid!=action.imageid)
            let allimages=state.allimages.filter(image=>image.imageid!=action.imageid)
            return {allimages:allimages,myimages:myimages};
        default:
            return state;
    }
}

