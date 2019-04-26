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

	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/lithammer/shortuuid"
	"github.com/mgechev/revive/formatter"
	"github.com/unrolled/render"
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

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(router))
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	//mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	//mx.HandleFunc("/upload", UploadPictureHandler(formatter)).Methods("POST")
	//	mx.HandleFunc("/update/{pictureId}", UpdateHandler(formatter)).Methods("PUT")
	//	mx.HandleFunc("/delete/{pictureId}", deleteByPictureIdHandler(formatter)).Methods("DELETE")
	//	mx.HandleFunc("/delete/{userId}", deleteByUserIdHandler(formatter)).Methods("DELETE")
}

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
	log.Println(handler.Filename)
	var newpic Picture
	newpic.ImageId = shortuuid.New() + filepath.Ext(handler.Filename)
	newpic.Description = req.FormValue("description")
	newpic.OwnerId = req.FormValue("ownerid")
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

//UpdatePictureHandler API update i.e. change owner
func UpdatePictureHandler(w http.ResponseWriter, req *http.Request) {
	var updaterequest Payload
	_ = json.NewDecoder(req.Body).Decode(&paymentdetail)
	session, _ := mgo.Dial(mongodb_server)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	err := session.DB("admin").Login(mongo_user, mongo_pass)
	if err != nil {
		formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c := session.DB(mongodb_database).C(mongodb_collection)
	params := mux.Vars(req)
	var picid string = params["pictureId"]
	var newusrid string = params["userId"]
	fmt.Println(picid)
	fmt.Println(newusrid)
	var result PictureRequest
	err = c.Find(bson.M{"pictureId": picid}).One(&result)
	if err != nil {
		fmt.Println("Picture not found")
		formatter.JSON(w, http.StatusNotFound, "Picture Not Found")
		return
	}
	result.RequestStatus = "Updated"
	result.UserId = paymentdetail.UserId
	c.Update(bson.M{"pictureId": picid}, bson.M{"$set": bson.M{"requestStatus": result.RequestStatus, "userId": result.UserId, "ipaddress": getIp()}})
	fmt.Println("Request:", picid, usrid, "updated")
	json.NewEncoder(w).Encode(result)
}

//DeletePictureHandler Deletes pictureby ID
func DeletePictureHandler(w http.ResponseWriter, req *http.Request) {
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer session.Close()
	if err := session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
		formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	var orderdetail RequiredPayload
	_ = json.NewDecoder(req.Body).Decode(&orderdetail)
	params := mux.Vars(req)
	var picid string = params["pictureId"]
	var result BurgerOrder
	fmt.Println("User ID: ", usrid)
	err = c.Find(bson.M{"commandId": uuid}).One(&result)
	if err != nil {
		fmt.Println("Command not found")
		formatter.JSON(w, http.StatusNotFound, "Command Not Found")
		return
	}
	for i := 0; i < len(result.Cart); i++ {
		if result.Cart[i].ItemId == orderdetail.ItemId {
			result.TotalAmount = result.TotalAmount - result.Cart[i].Price
			result.Cart = append(result.Cart[0:i], result.Cart[i+1:]...)
			break
		}
	}
	c.Update(bson.M{"commandId": uuid}, bson.M{"$set": bson.M{"items": result.Cart, "totalAmount": result.TotalAmount, "ipaddress": getIp()}})
	fmt.Println("Delete Item: ", orderdetail.ItemId, "from order", uuid)
	formatter.JSON(w, http.StatusOK, result)
}

/**
// API Delete all pictures owned by user userId
func deleteByUserIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		defer session.Close()
		if err := session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var orderdetail RequiredPayload
		_ = json.NewDecoder(req.Body).Decode(&orderdetail)
		fmt.Println("order ID: ", orderdetail.OrderId)
		err = c.Remove(bson.M{"orderId": orderdetail.OrderId})
		if err != nil {
			fmt.Println("order not found")
			formatter.JSON(w, http.StatusNotFound, "Order Not Found")
			return
		}
		formatter.JSON(w, http.StatusOK, "Order: "+orderdetail.OrderId+" Deleted")
	}
}
**/
