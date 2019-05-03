package main

import (
	//"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(mx))
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/reaction/{ImageId}", imageReactionHandler(formatter)).Methods("GET")
	mx.HandleFunc("/like", imageLikeHandler(formatter)).Methods("POST")
	mx.HandleFunc("/unlike/{imageid}/{userid}", imageUnlikeHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/comment", imageCommentHandler(formatter)).Methods("POST")
	mx.HandleFunc("/removecomment/{ImageId}", commentDeleteHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/images/{imageid}/user/{userid}", isImageLikedByUser(formatter)).Methods("GET")
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
		glog.Info("Entered imageReactionHandler")
		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)

		params := mux.Vars(req)
		var imageid string = params["ImageId"]
		glog.Info("Incoming data", imageid)
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
		reactions.Imageid = imageid
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
		glog.Info("Entered imageLikeHandler")
		var reaction Reaction
		_ = json.NewDecoder(req.Body).Decode(&reaction)
		glog.Info("Incoming data", reaction)
		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(COLLECTION)
		var result []Likes
		reaction.Reaction_type = "Like"
		reaction.Timestamp = time.Now()

		errin := c.Insert(reaction)
		if errin != nil {
			panic(err)
		}

		result, err = getLikesList(reaction.Image_id)
		var reactions ImageReaction
		reactions.Imageid = reaction.Image_id

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
		glog.Info("Reactions: ", reactions)
		_ = formatter.JSON(w, http.StatusOK, reactions)
	}
}

//Image like
func isImageLikedByUser(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		glog.Info("Inside isImageLikedByUser")
		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(COLLECTION)

		params := mux.Vars(req)
		var imageid = params["imageid"]
		var userid = params["userid"]
		glog.Info("Incoming Data", imageid, userid)

		type isLike struct {
			isliked bool
		}
		result := isLike{
			isliked: false,
		}

		query := bson.M{"imageId": imageid, "userId": userid, "reactionType": "Like"}

		n, errin := c.Find(query).Count()
		if errin != nil {
			panic(err)
		}

		if n > 0 {

			result.isliked = true

		}

		_ = json.NewDecoder(req.Body).Decode(&result)
		glog.Info("Response sending ", result)
		formatter.JSON(w, http.StatusOK, result.isliked)
	}
}

//API to update the unlike
func imageUnlikeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		glog.Info("Inside imageUnlikeHandler")
		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(COLLECTION)

		params := mux.Vars(req)
		var imageid = params["imageid"]
		var userid = params["userid"]
		glog.Info("Incomingdata", imageid, userid)

		var likes []Likes
		query := bson.M{"imageId": imageid, "userId": userid, "reactionType": "Like"}

		errin := c.Remove(query)
		if errin != nil {
			panic(err)
		}
		likes, err = getLikesList(imageid)
		if likes == nil {
			if err != nil {
				_ = formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			_ = formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}
		var reactions ImageReaction
		reactions.Imageid = imageid
		if len(likes) != 0 {
			reactions.Likes = likes
		}
		_ = json.NewDecoder(req.Body).Decode(&reactions)
		glog.Info("Reactions: ", reactions)
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
		var imageid = params["ImageId"]
		var userid = req.URL.Query().Get("userId")
		var commentId string = req.URL.Query().Get("commentId")
		fmt.Println("Image ID: ", imageid)
		fmt.Println("User ID: ", userid)

		var comments []Comments
		query := bson.M{"imageId": imageid, "userId": userid, "reactionType": "Comment", "commentId": commentId}
		fmt.Println("Debugging: 0")
		errin := c.Remove(query)
		if errin != nil {
			panic(err)
		}
		fmt.Println("Debugging: 1")
		comments, err = getCommentsList(imageid)
		fmt.Println("Debugging: 2")
		if comments == nil {

			if err != nil {
				_ = formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			_ = formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}
		fmt.Println("Debugging: 3")
		fmt.Println("Reactions: ", comments)
		var reactions ImageReaction
		reactions.Imageid = imageid
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

		values := req.URL.Query()
		var imageid string = values.Get("imageId")
		var userid = values.Get("userId")
		var username = values.Get("userName")
		var comment = values.Get("comment")

		fmt.Println("Image ID: ", imageid)
		fmt.Println("User ID: ", userid)

		var comments []Comments
		var reaction Reaction
		reaction.Image_id = imageid
		reaction.Reaction_type = "Comment"
		reaction.UserId = userid
		reaction.Username = username
		reaction.Timestamp = time.Now()
		reaction.Comment = comment

		uuid := uuid.NewV4()
		reaction.CommentId = uuid.String()

		errin := c.Insert(reaction)
		if errin != nil {
			panic(err)
		}
		comments, err = getCommentsList(imageid)

		if comments == nil {
			if err != nil {
				_ = formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			_ = formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}

		var reactions ImageReaction
		reactions.Imageid = imageid
		if len(comments) != 0 {
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

	fmt.Println("Image ID: ", imageId)

	var result []Reaction
	err = c.Find(bson.M{"imageId": imageId}).All(&result)
	if err != nil {
		//formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("Error: ", err)
		return nil
	}
	return result
}

//Helper method for returning likes
func getLikesList(imageId string) ([]Likes, error) {
	glog.Info("Entered getLikesList function ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(COLLECTION)
	var likes []Likes
	err = c.Find(bson.M{"imageId": imageId, "reactionType": "Like"}).Select(bson.M{"userName": 1, "timeStamp": 1}).All(&likes)
	if err != nil {
		glog.Error("Error finding the data", err)
		return nil, err
	}
	if len(likes) == 0 {
		glog.Info("No such document")
		return nil, nil
	}
	return likes, nil
}

//Helper method to return comments
func getCommentsList(imageId string) ([]Comments, error) {
	glog.Info("Entered getCommentsList function  ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(COLLECTION)
	var comments []Comments
	matchQuery := bson.M{"imageId": imageId, "reactionType": "Comment"}
	projectQuery := bson.M{
		"userName": 1,
		"comment":  1,
		"timeStamp": bson.M{
			"$dateToString": bson.M{"format": "%Y-%m-%d %H:%M:%S", "date": "$timeStamp"},
		},
		"commentId": 1,
	}
	pipeline := []bson.M{
		{"$match": matchQuery},
		{"$project": projectQuery},
	}
	err = c.Pipe(pipeline).All(&comments)
	if err != nil {
		glog.Error("Error: ", err)
		return nil, err
	}
	if len(comments) == 0 {
		glog.Info("No such document")
		return nil, nil
	}
	return comments, nil

}
