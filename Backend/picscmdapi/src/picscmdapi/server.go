/*
	go-burger-order REST API (Version)
*/

package main

import (
	"strconv"
	"strings"

	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/lithammer/shortuuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB Config

//var mongodb_server = "10.0.0.117"
//var mongodb_server = "dockerhost"

var mongodb_server = "52.11.201.189" //os.Getenv("Server")
var mongodb_database = "cmpe281"     //os.Getenv("Database") // pics
var mongodb_collection = "Picture"   //os.Getenv("Collection") // picassa
var mongo_user = "masea"             //os.Getenv("User") // masea
var mongo_pass = "cmpe281"           //os.Getenv("Pass") // cmpe281

var s3_bucket_name_orig = "picassooriginal" //os.Getenv("BUCKET_NAME_ORIG")
var s3_bucket_name_tumb = "picassotumbnail" //os.Getenv("BUCKET_NAME_ORIG")

//S3 bucket credentials
//ACCESSKEYS3 ACCESSKEYS3
const ACCESSKEYS3 = "AKIAQKC4VVTZASCMOPBC"

//SECRETKEYS3 SECRETKEYS3
const SECRETKEYS3 = "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"

//BUCKETNAME BUCKETNAME
const BUCKETNAME = "cmpe281picassa"

//FOLDERS3 FOLDERS3
const FOLDERS3 = "pictures"

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//Create a function we use to display errors and exit.
func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func getIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	address := strings.Split(localAddr, ":")
	fmt.Println("address: ", address[0])
	return address[0]
}

// API Ping Handler
func pingHandler(w http.ResponseWriter, req *http.Request) {
	message := "Picture write command API Server Working on machine: " + getIp()
	json.NewEncoder(w).Encode(map[string]string{"result": message})
}

//UploadHandler upload new image to s3 and store data to MongoDB
func UploadPictureHandler(w http.ResponseWriter, req *http.Request) {
	session, err := mgo.Dial(mongodb_server)
	/**if err := session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
		formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}**/
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)

	file, handler, err := req.FormFile("myfile")
	if err != nil {
		log.Println("Error in uploading the file")
		return
	}
	defer file.Close()

	var newpic Picture
	newpic.ImageId = shortuuid.New() + filepath.Ext(handler.Filename)
	newpic.Description = req.FormValue("description")
	newpic.UserId = req.FormValue("userid")
	price, _ := strconv.ParseInt(req.FormValue("price"), 0, 64)
	newpic.Price = price
	newpic.Title = req.FormValue("title")
	//Inserting file to S3 Bucket
	res := InsertIntoS3(newpic.ImageId, file)
	if res == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT a Valid Request"))
	} else {
		newpic.OrigUrl = res
		err = c.Insert(newpic)
		if err != nil {
			log.Println("Exception inserting data to Database")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("NOT a Valid Request"))
		} else {
			json.NewEncoder(w).Encode(newpic)
		}
	}

}

//DeletePictureHandler API update i.e. change owner
func DeletePictureHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	imageid := params["imageid"]
	session, err := mgo.Dial(mongodb_server)
	/**if err := session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
		formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}**/
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	query := bson.M{"imageid": imageid}
	var result Picture
	err = c.Find(query).One(&result)
	if err != nil {
		log.Println("No Image Found")
	} else {
		err = DeleteFromS3(imageid)
		if err != nil {
			log.Println("Error while removing from S3Bucket")
		}
		err = c.Remove(query)
		if err != nil {
			log.Println("Error while removing from MongoDB")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("NOT a Valid Request"))
		}
		json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	}

}

//UpdatePictureHandler Deletes pictureby ID
func UpdatePictureHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	imageid := params["imageid"]
	log.Println("PicscmdApi In UpdatePictureHandler ImageID" + imageid)
	var picture Picture
	_ = json.NewDecoder(req.Body).Decode(&picture)
	session, err := mgo.Dial(mongodb_server)
	/**if err := session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
		formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}**/
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	query := bson.M{"imageid": picture.ImageId}
	err = c.Update(query, picture)
	if err != nil {
		log.Println("Error while Updating Document in UpdatePictureHandler Password")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT a Valid Request"))
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "Removed the image Successfully"})
}

//DeleteByUserIdHandler API Delete all pictures owned by user userId
func DeleteByUserIdHandler(w http.ResponseWriter, req *http.Request) {

}
