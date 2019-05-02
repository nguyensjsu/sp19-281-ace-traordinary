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
	"github.com/satori/go.uuid"
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
	mx.HandleFunc("/unlike/{ImageId}/{User}", imageUnlikeHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/comment", imageCommentHandler(formatter)).Methods("POST")
	mx.HandleFunc("/removecomment/{ImageId}", commentDeleteHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("like/image/{ImageId}/user/{User}", isImageLikedByUser(formatter)).Methods("GET")

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
		fmt.Println("Request: ", req)
		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)

		params := mux.Vars(req)
		var imageid string = params["ImageId"]
		fmt.Println("Image ID: ", imageid)

		var likes []Likes
		var comments []Comments

		likes, err = getLikesList(imageid)
		comments, err = getCommentsList(imageid)

		if comments == nil && likes == nil {
			if err != nil {
				_ = formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		_ = formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
		return
	}

		var reactions ImageReaction
		reactions.Imageid=imageid
		if len(likes) != 0 {
			reactions.Likes = likes
		}
		if len(comments) != 0 {
			reactions.Comments = comments
		}

		_ = json.NewDecoder(req.Body).Decode(&reactions)
		fmt.Println("Pictures: ", reactions)
		formatter.JSON(w, http.StatusOK, reactions)

	}
}

//API to update like
func imageLikeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("Request: ", req.URL.Query())
		fmt.Println("Request Param: ", req.URL.Query().Get("imageId"))
		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(COLLECTION)

		var imageid string = req.URL.Query().Get("imageId")
		var userid= req.URL.Query().Get("userId")
		var username= req.URL.Query().Get("userName")
		fmt.Println("Image ID: ", imageid)
		fmt.Println("User ID: ", userid)

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

		result, err = getLikesList(imageid)
		var reactions ImageReaction
		reactions.Imageid = imageid

		if result == nil {
			if err != nil {
				_ = formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			_ = formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}
		reactions.Likes = result
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
		var userid  = params["User"]

		fmt.Println( "Image ID: ", imageid )
		fmt.Println( "User ID: ", userid )

		var likes []Likes
		query := bson.M{"imageId":imageid,"userId": userid,"reactionType":"Like"}

		errin := c.Remove(query)
		if errin != nil {
			panic(err)
		}
		likes,err = getLikesList(imageid)
		if likes==nil {
			if err!=nil{
				_=formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			_=formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}
		var reactions ImageReaction
		reactions.Imageid=imageid
		if len(likes) != 0 {
			reactions.Likes = likes
		}
		_ = json.NewDecoder(req.Body).Decode(&reactions)
		fmt.Println("Reactions: ", reactions)
		formatter.JSON(w, http.StatusOK, reactions)

	}
}

func commentDeleteHandler(formatter *render.Render) http.HandlerFunc {
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
		var userid  = req.URL.Query().Get("userId")
		var commentId string = req.URL.Query().Get("commentId")
		fmt.Println( "Image ID: ", imageid )
		fmt.Println( "User ID: ", userid )

		var comments []Comments
		query := bson.M{"imageId":imageid,"userId": userid,"reactionType":"Comment", "commentId":commentId}
		fmt.Println("Debugging: 0")
		errin := c.Remove(query)
		if errin != nil {
			panic(err)
		}
		fmt.Println("Debugging: 1")
		comments,err = getCommentsList(imageid)
		fmt.Println("Debugging: 2")
		if comments==nil {

			if err!=nil{
				_=formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			_=formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}
		fmt.Println("Debugging: 3")
		fmt.Println("Reactions: ", comments)
		var reactions ImageReaction
		reactions.Imageid=imageid
		if len(comments) != 0 {
			reactions.Comments = comments
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

		uuid := uuid.NewV4()
		reaction.CommentId = uuid.String();

		errin := c.Insert(reaction)
		if errin != nil {
			panic(err)
		}
		comments,err = getCommentsList(imageid)

		if comments==nil {
			if err!=nil {
				_ = formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			_ = formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}

		var reactions ImageReaction
		reactions.Imageid=imageid
		if(len(comments) != 0) {
			reactions.Comments = comments
		}
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


func isImageLikedByUser(formatter *render.Render) http.HandlerFunc {
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
		var userid  = params["User"]

		fmt.Println( "Image ID: ", imageid )
		fmt.Println( "User ID: ", userid )

		type isLike struct {
			isLiked bool
		}
		result := isLike{
			isLiked: false,
		}

		query := bson.M{"imageId":imageid,"userId": userid,"reactionType":"Like"}

		n, errin := c.Find(query).Count()
		if errin != nil {
			panic(err)
		}

		if n>0 {

			result.isLiked= true

		}

		_ = json.NewDecoder(req.Body).Decode(&result)
		fmt.Println("Reactions: ", result)
		formatter.JSON(w, http.StatusOK, result)

	}
}

//Helper method for returning likes
func getLikesList(imageId string) ([]Likes,error) {
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
		fmt.Println( "Error: ", err )
		return nil,err
	}
	if len(likes)==0 {
		fmt.Println( "No such document: " )
		return nil,nil
	}
	return likes,nil
}

//Helper method to return comments
func getCommentsList(imageId string) ([]Comments,error ){
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
		"userName" : 1,
		"comment" : 1,
		"timeStamp": bson.M{
			"$dateToString": bson.M{"format": "%Y-%m-%d %H:%M:%S", "date": "$timeStamp"},

		},
		"commentId" : 1,
	}
	pipeline := []bson.M{
		{"$match": matchQuery},
		{"$project": projectQuery},
	}
	err = c.Pipe(pipeline).All(&comments)
	if err!=nil {
		fmt.Println( "Error: ", err )
		return nil,err
	}
	if len(comments)==0 {
		fmt.Println("No such document: ")
		return nil, nil
	}
	fmt.Println( "Comments: ", comments )
	return comments,nil

}


