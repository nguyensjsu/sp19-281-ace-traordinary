package src

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/sp19-281-ace-traordinary/Backend/Reactions"

)

//MongoDB connection details
var mongodbServer = "52.11.201.189"
var mongodbDatabase = "cmpe281"
var COLLECTION = "User"

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
	mx.HandleFunc("/like/{ImageId}", imageLikeHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/dislike/{ImageId}", imageUnlikeHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/comment/{ImageId}", imageCommentHandler(formatter)).Methods("PUT")

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

		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)


		//c := session.DB(mongodbDatabase).C(COLLECTION)

		params := mux.Vars(req)
		var image_id string = params["ImageId"]
		fmt.Println( "Image ID: ", image_id )

		var likes []Likes
		var comments []Comments


		likes = getLikesList(image_id)
		comments = getCommentsList(image_id)

		if (comments==nil && likes ==nil) {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if (len(comments) == 0 && len(likes)==0) {
			formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}

		var reactions ImageReaction
		reactions.Imageid=image_id
		if(len(likes) != 0) {
			reactions.likes = likes
		}
		if(len(comments) != 0) {
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

		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(COLLECTION)

		params := mux.Vars(req)
		var image_id string = params["ImageId"]
		var user_id string = params["UserId"]
		var user_name string = params["UserName"]
		fmt.Println( "Image ID: ", image_id )
		fmt.Println( "User ID: ", user_id )

		var result []Likes
		var reaction model.Reaction
		reaction.image_id = image_id
		reaction.Reaction_type = "Like"
		reaction.UserId = user_id
		reaction.Username = user_name
		reaction.Timestamp = time.Now()

		errin := c.Insert(reaction)
		if errin != nil {
			panic(err)
		}

		result = getLikesList(image_id)

		if result==nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if len(result) == 0 {
			formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}
		_ = json.NewDecoder(req.Body).Decode(&result)
		fmt.Println("Pictures: ", result)
		formatter.JSON(w, http.StatusOK, result)



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
		var image_id string = params["ImageId"]
		var user_id string = params["UserId"]

		fmt.Println( "Image ID: ", image_id )
		fmt.Println( "User ID: ", user_id )

		var likes []Likes
	//	var reaction model.Reaction
		query := bson.M{"image_id":image_id,"userid": user_id,"Reaction_type":"Like"}


		errin := c.Remove(query)
		if errin != nil {
			panic(err)
		}

		likes = getLikesList(image_id)

		if likes==nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if len(likes) == 0 {
			formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}

		var reactions ImageReaction
		reactions.Imageid=image_id
		if(len(likes) != 0) {
			reactions.likes = likes
		}
		_ = json.NewDecoder(req.Body).Decode(&reactions)
		fmt.Println("Pictures: ", reactions)
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

		params := mux.Vars(req)
		var image_id string = params["ImageId"]
		var user_id string = params["UserId"]
		var user_name string = params["UserName"]
		var comment string = params["Comment"]
		fmt.Println( "Image ID: ", image_id )
		fmt.Println( "User ID: ", user_id )

		var comments []Comments
		var reaction model.Reaction
		reaction.image_id = image_id
		reaction.Reaction_type = "Comment"
		reaction.UserId = user_id
		reaction.Username = user_name
		reaction.Timestamp = time.Now()
		reaction.Comment = comment

		errin := c.Insert(reaction)
		if errin != nil {
			panic(err)
		}

		comments = getCommentsList(image_id)

		if comments==nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if len(comments) == 0 {
			formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist anymore.")
			return
		}
		var reactions ImageReaction
		reactions.Imageid=image_id

		if(len(comments) != 0) {
			reactions.comments = comments
		}

		_ = json.NewDecoder(req.Body).Decode(&reactions)
		fmt.Println("Pictures: ", reactions)
		formatter.JSON(w, http.StatusOK, reactions)

	}
}


func getReactionList(imageId string) []model.Reaction {
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
	err = c.Find(bson.M{"image_id":image_id}).All(&result)
	if err!=nil {
		//formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return nil
	}

return result

}

//Helper method for returning likes
func getLikesList(imageId string) []model.Likes {
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
	err = c.Find(bson.M{"image_id":image_id,"Reaction_type":"Like"}).Select(bson.M({"Username":1,"TimeStamp":1})).All(&likes)
	if err!=nil {
		//formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return nil
	}

	return likes

}

//Helper method to return comments
func getCommentsList(imageId string) []model.Comments {
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
	err = c.Find(bson.M{"image_id":image_id,"Reaction_type":"Comment"}).Select(bson.M({"Username":1,"Comment":1,"TimeStamp":1})).All(&comments)
	if err!=nil {
		//formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return nil
	}

	return comments

}