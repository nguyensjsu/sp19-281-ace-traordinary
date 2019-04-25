package main

import (
	//"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	//"github.com/sp19-281-ace-traordinary/Backend/Reactions"

)

//MongoDB connection details
var mongodbServer = "54.241.227.172"
var mongodbDatabase = "cmpe281"
var COLLECTION = "Reaction"

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/reaction/{ImageId}", imageReactionHandler(formatter)).Methods("GET")
	mx.HandleFunc("/like", imageLikeHandler(formatter)).Methods("POST")
	mx.HandleFunc("/unlike/{ImageId}", imageUnlikeHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/comment", imageCommentHandler(formatter)).Methods("POST")

}

//Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "Pong"
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}

//API handler for GET reactions for Image..
func imageReactionHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println( "Request: ", req )
		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)


		//c := session.DB(mongodbDatabase).C(COLLECTION)

		params := mux.Vars(req)
		var imageid string = params["ImageId"]
		fmt.Println( "Image ID: ", imageid )

		var likes []Likes
		var comments []Comments


		likes = getLikesList(imageid)
		comments = getCommentsList(imageid)

		if (comments==nil && likes ==nil) {
			_=formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if (len(comments) == 0 && len(likes)==0) {
			_=formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}

		var reactions ImageReaction
		reactions.Imageid=imageid
		if len(likes) != 0 {
			reactions.likes = likes
		}
		if len(comments) != 0 {
			reactions.comments = comments
		}

		_ = json.NewDecoder(req.Body).Decode(&reactions)
		fmt.Println("Pictures: ", reactions)
		formatter.JSON(w, http.StatusOK, reactions)

	}
}

//API to update like
func imageLikeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Println( "Request: ", req.URL.Query() )
		fmt.Println( "Request Param: ", 	req.URL.Query().Get("imageId"))
		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(COLLECTION)

		//params := mux.Vars(req)
		var imageid string  = req.URL.Query().Get("imageId")
		var userid  = req.URL.Query().Get("userId")
		var username  = req.URL.Query().Get("userName")
		fmt.Println( "Image ID: ", imageid )
		fmt.Println( "User ID: ", userid )

		//vals:= req.URL.Query().Get("imageId")


		var result []Likes
		var reaction Reaction
		reaction.Image_id = imageid
		reaction.Reaction_type = "Like"
		reaction.UserId = userid
		reaction.Username = username
		reaction.Timestamp = time.Now()

		errin := c.Insert(reaction)
		if errin != nil {
			panic(err)
		}

		result = getLikesList(imageid)
var reactions ImageReaction
		reactions.Imageid = imageid

		if result==nil {
			_ = formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if len(result) == 0 {
		_ =	formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}
		reactions.likes = result
		_ = json.NewDecoder(req.Body).Decode(&reactions)
		fmt.Println("Reactions: ", reactions)
		_ = formatter.JSON(w, http.StatusOK, reactions)



	}
}

//API to update the unlike
func imageUnlikeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(COLLECTION)

		params := mux.Vars(req)
		var imageid  = params["ImageId"]
		var userid  = params["userId"]

		fmt.Println( "Image ID: ", imageid )
		fmt.Println( "User ID: ", userid )

		var likes []Likes
	//	var reaction model.Reaction
		query := bson.M{"imageId":imageid,"userId": userid,"reactionType":"Like"}


		errin := c.Remove(query)
		if errin != nil {
			panic(err)
		}

		likes = getLikesList(imageid)

		if likes==nil {
			_=formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if len(likes) == 0 {
			_=formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}

		var reactions ImageReaction
		reactions.Imageid=imageid
		if len(likes) != 0 {
			reactions.likes = likes
		}
		_ = json.NewDecoder(req.Body).Decode(&reactions)
		fmt.Println("Reactions: ", reactions)
		formatter.JSON(w, http.StatusOK, reactions)




	}
}
//API to insert comments
func imageCommentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(COLLECTION)

		//params := mux.Vars(req)
		//var imageid  = params["ImageId"]
	//	var userid   = params["userId"]
		//var username = params["userName"]
	//	var comment  = params["comment"]
		//fmt.Println( "Image ID: ", imageid )
		//fmt.Println( "User ID: ", userid )

		values :=req.URL.Query()
		var imageid string  = values.Get("imageId")
		var userid  = values.Get("userId")
		var username  = values.Get("userName")
		var comment  = values.Get("comment")

		fmt.Println( "Image ID: ", imageid )
		fmt.Println( "User ID: ", userid )


		var comments []Comments
		var reaction Reaction
		reaction.Image_id = imageid
		reaction.Reaction_type = "Comment"
		reaction.UserId = userid
		reaction.Username = username
		reaction.Timestamp = time.Now()
		reaction.Comment = comment

		errin := c.Insert(reaction)
		if errin != nil {
			panic(err)
		}

		comments = getCommentsList(imageid)

		if comments==nil {
			_ = formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if len(comments) == 0 {
			_ = formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}
		var reactions ImageReaction
		reactions.Imageid=imageid

		if(len(comments) != 0) {
			reactions.comments = comments
		}
		fmt.Println("Reactions:before marshal ", reactions)
		//reactions.likes = &[]Likes{}
		data,_   := json.Marshal(reactions)
		fmt.Println("Reactions: ", string(data))
		  json.NewDecoder(req.Body).Decode(&reactions)
		fmt.Println("Reactions: ", reactions)
		 formatter.JSON(w, http.StatusOK, reactions)

	}
}


func getReactionList(imageId string) []Reaction {
	fmt.Println("Entered LoginDao function  ")

	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(COLLECTION)

	fmt.Println( "Image ID: ", imageId )

	var result []Reaction
	err = c.Find(bson.M{"imageId":imageId}).All(&result)
	if err!=nil {
		//formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println( "Error: ", err )
		return nil
	}

return result

}

//Helper method for returning likes
func getLikesList(imageId string) []Likes {
	fmt.Println("Entered LoginDao function  ")

	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(COLLECTION)

	fmt.Println( "Image ID: ", imageId )

	var likes []Likes
	err = c.Find(bson.M{"imageId": imageId,"reactionType":"Like"}).Select(bson.M{"userName": 1,"timeStamp": 1 }).All(&likes)

	if err!=nil {
		//formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println( "Error: ", err )
		return nil
	}
	if len(likes)==0 {
		//formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println( "No such document: " )
		return nil
	}
	return likes

}

//Helper method to return comments
func getCommentsList(imageId string) []Comments {
	fmt.Println("Entered LoginDao function  ")

	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(COLLECTION)

	fmt.Println( "Image ID: ", imageId )

	var comments []Comments

	matchQuery := bson.M{"imageId":imageId,"reactionType":"Comment"}
	projectQuery := bson.M{
		"userName" : 1 ,
		"comment" : 1,
		"timeStamp": bson.M{
			"$dateToString": bson.M{"format": "%Y-%m-%d %H:%M:%S", "date": "$timeStamp"},

		},
	}
	pipeline := []bson.M{
		{"$match": matchQuery},
		{"$project": projectQuery},
	}

	//err = c.Find(bson.M{"imageId":imageId,"reactionType":"Comment"}).Select(bson.M{"userName":1,"comment":1,"timeStamp":1}).All(&comments)
	err = c.Pipe(pipeline).All(&comments)
	if err!=nil {
		//formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println( "Error: ", err )
		return nil
	}

	return comments

}


