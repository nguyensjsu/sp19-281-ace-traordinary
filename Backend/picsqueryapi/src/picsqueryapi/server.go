/*
	picsqueryapi REST API (Version)
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net"
	"os"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/handlers"

	//"github.com/aws/aws-sdk-go/service/s3"
)

// MongoDB Config

//var mongodb_server = "10.0.0.117"
//var mongodb_server = "dockerhost"

var mongodb_server = os.Getenv("Server")
var mongodb_database = os.Getenv("Database")
var mongodb_collection = os.Getenv("Collection")
var mongo_user = os.Getenv("User")
var mongo_pass = os.Getenv("Pass")   

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD",  "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	n.UseHandler(handlers.CORS(allowedHeaders,allowedMethods , allowedOrigins)(router))
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/picsquery/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/picsquery", pictureQueryByPageNumberAndCount(formatter)).Methods("GET")
	mx.HandleFunc("/picsquery/{pictureId}", pictureQueryByPictureId(formatter)).Methods("GET")
	mx.HandleFunc("/picsquery/{userId}", pictureQueryByUserId(formatter)).Methods("GET")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func getSystemIp() string {
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
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "Picture Query API Server Working on machine: " + getSystemIp()
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}

// API Get Pictures by Page Number and Count Handler
func pictureQueryByPageNumberAndCount(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		err:= session.DB("admin").Login(mongo_user, mongo_pass)
		if err!=nil{
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		//var qryid string = params["queryId"]
		//var usrid string = params["userId"]	
		//var howMany string = params["count"]
		var whichPage string = params["pageNumber"]// == "" ? 1 : params["pageNumber"]);
		if  whichPage == "" {
		    whichPage = "1"
		}
		//if whichpage == "" {
			var pics_array []Payload
			
			// mongodb
			err = c.Find(bson.M{}).All(&pics_array)
			
            // riak
   //          rowsPerPage := uint32(howmany)
			// page := uint32(whichPage) // 2
			// start := rowsPerPage * (page - uint32(1))

			// cmd, err := riak.NewSearchCommandBuilder().
			//     WithIndexName("tumbnail").
			//     WithQuery("*:*").
			//     WithStart(start).
			//     WithNumRows(rowsPerPage).
			//     Build();
			// if err != nil {
			//     return err
			// }

			// if err := cluster.Execute(cmd); err != nil {
			//     return err
			// }
			// riak



			fmt.Println("Pictures:", pics_array)
			formatter.JSON(w, http.StatusOK, pics_array)
		// } else {
		// 	fmt.Println("queryId: ", qryid)
		// 	var result Payload
		// 	err = c.Find(bson.M{"queryId":qryid}).One(&result)
		// 	if err!=nil {
		// 		formatter.JSON(w, http.StatusNotFound, "Order Not Found")
		// 		return
		// 	}
		// 	_ = json.NewDecoder(req.Body).Decode(&result)
		// 	fmt.Println("Pictures: ", result)
		// 	formatter.JSON(w, http.StatusOK, result)
		// }
	}
}

// API Get Picture by Picture Id Handler
func pictureQueryByPictureId(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		err:= session.DB("admin").Login(mongo_user, mongo_pass)
		if err!=nil{
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		var qryid string = params["queryId"]
		var usrid string = params["userId"]
		var picid string = params["pictureId"]
		fmt.Println("userId: ", usrid)
		fmt.Println("pictureId: ", qryid)
		var result []Payload
		err = c.Find(bson.M{"pictureId":picid}).All(&result)
		if err!=nil {
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

// API Get Pictures by User Id
func pictureQueryByUserId(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		err:= session.DB("admin").Login(mongo_user, mongo_pass)
		if err!=nil{
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		//var qryid string = params["queryId"]
		var usrid string = params["userId"]
		fmt.Println("userId: ", usrid)
		var result []Payload
		err = c.Find(bson.M{"userId":usrid}).All(&result)
		if err!=nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if len(result) == 0 {
			formatter.JSON(w, http.StatusNotFound, "No results found for this user.")
			return
		}
		_ = json.NewDecoder(req.Body).Decode(&result)
		fmt.Println("Pictures: ", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}
