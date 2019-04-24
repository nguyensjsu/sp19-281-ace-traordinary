/*
	go-burger-order REST API (Version)
*/

package main

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/handlers"
)

// MongoDB Config

//var mongodb_server = "10.0.0.117"
//var mongodb_server = "dockerhost"

var mongodb_server = "mongodb"//os.Getenv("Server") 
var mongodb_database = "pics"//os.Getenv("Database") // pics
var mongodb_collection = "picasso"//os.Getenv("Collection") // picassa
var mongo_user = "masea"//os.Getenv("User") // masea
var mongo_pass = "cmpe281"//os.Getenv("Pass") // cmpe281

var s3_bucket_name_orig = "picassooriginal"//os.Getenv("BUCKET_NAME_ORIG") 
var s3_bucket_name_tumb = "picassotumbnail"//os.Getenv("BUCKET_NAME_ORIG") 

// RabbitMQ Config
// var rabbitmq_server = "rabbitmq"
// var rabbitmq_port = "5672"
// var rabbitmq_queue = "gumball"

// var rabbitmq_user = "guest"
// var rabbitmq_pass = "guest"

// NewServer configures and returns a Server.
/*func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}*/

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
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/upload", uploadHandler(formatter)).Methods("POST")
	mx.HandleFunc("/update/{pictureId}", updateHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/delete/{pictureId}", deleteByPictureIdHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/delete/{userId}", deleteByUserIdHandler(formatter)).Methods("DELETE")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// log s3 access requests
func LogAccess() {
	svc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
}

// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html
// a function to display errors and exit.
func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}

// Creates a S3 Bucket in the region configured in the shared config
// or AWS_REGION environment variable.
//
// Usage:
//    go run s3_upload_object.go BUCKET_NAME FILENAME
func insertIntoS3(filename string, bucketname string) {
	object, err := os.Open("my-testfile")
	if err != nil {
		log.Fatalln(err)
	}
	defer object.Close()
	objectStat, err := object.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	n, err := s3Client.PutObject("my-bucketname", "my-objectname", object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Uploaded", "my-objectname", " of size: ", n, "Successfully.")

    // if len(os.Args) != 3 {
    //     exitErrorf("bucket and file name required\nUsage: %s bucket_name filename",
    //         os.Args[0])
    // }

    // bucket := os.Args[1]
    // filename := os.Args[2]

    file, err := os.Open(filename)
    if err != nil {
        exitErrorf("Unable to open file %q, %v", err)
    }

    defer file.Close()
    fileStat, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2")},
    )

	n, err := s3Client.PutObject("my-bucketname", "my-objectname", object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

    // Setup the S3 Upload Manager. Also see the SDK doc for the Upload Manager
    // for more information on configuring part size, and concurrency.
    //
    // http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
    uploader := s3manager.NewUploader(sess)

    // Upload the file's body to S3 bucket as an object with the key being the
    // same as the filename.
    _, err = uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String(bucket),

        // Can also use the `filepath` standard library package to modify the
        // filename as need for an S3 object key. Such as turning absolute path
        // to a relative path.
        Key: aws.String(filename),

        // The file to be uploaded. io.ReadSeeker is preferred as the Uploader
        // will be able to optimize memory when uploading large content. io.Reader
        // is supported, but will require buffering of the reader's bytes for
        // each part.
        Body: file,
    })
    if err != nil {
        // Print the error and exit.
        exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
    }

    fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}

func deleteFromS3(filename string) {
	    if len(os.Args) != 3 {
        exitErrorf("Bucket and object name required\nUsage: %s bucket_name object_name",
            os.Args[0])
    }

    bucket := os.Args[1]
    obj := os.Args[2]

    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2")},
    )

    // Create S3 service client
    svc := s3.New(sess)

    // Delete the item
    _, err = svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(obj)})
    if err != nil {
        exitErrorf("Unable to delete object %q from bucket %q, %v", obj, bucket, err)
    }

    err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(obj),
    })
    if err != nil {
        exitErrorf("Error occurred while waiting for object %q to be deleted, %v", obj)
    }

    fmt.Printf("Object %q successfully deleted\n", obj)
}


// Deletes all of the objects in the specified S3 Bucket in the region configured in the shared config
// or AWS_REGION environment variable.
//
// Usage:
//    go run s3_delete_objects BUCKET
func deleteMultipleObjectsFromS3() {
    if len(os.Args) != 2 {
        exitErrorf("Bucket name required\nUsage: %s BUCKET", os.Args[0])
    }

    bucket := os.Args[1]

    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2")},
    )

    // Create S3 service client
    svc := s3.New(sess)

    // Setup BatchDeleteIterator to iterate through a list of objects.
    iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
        Bucket: aws.String(bucket),
    })

    // Traverse iterator deleting each object
    if err := s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter); err != nil {
        exitErrorf("Unable to delete objects from bucket %q, %v", bucket, err)
    }

    fmt.Printf("Deleted object(s) from bucket: %s", bucket)
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
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "Picture write command API Server Working on machine: " + getSystemIp()
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}


// API upload new picture
func uploadFileHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
       fmt.Println("method:", req.Method)
       if req.Method == "GET" {
           crutime := time.Now().Unix()
           h := md5.New()
           io.WriteString(h, strconv.FormatInt(crutime, 10))
           token := fmt.Sprintf("%x", h.Sum(nil))

           t, _ := template.ParseFiles("upload.gtpl")
           t.Execute(w, token)
       } else {
           req.ParseMultipartForm(32 << 20)
           file, handler, err := req.FormFile("uploadfile")
           if err != nil {
               fmt.Println(err)
               return
           }
           defer file.Close()
           fmt.Fprintf(w, "%v", handler.Header)
           f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
           if err != nil {
               fmt.Println(err)
               return
           }
           defer f.Close()
           io.Copy(f, file)
       }
   }
}


// API upload new picture
func uploadHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Open MongoDB Session
		var requestdetail UploadRequest
		_ = json.NewDecoder(req.Body).Decode(&requestdetail)
		session, err := mgo.Dial(mongodb_server)
		if err:= session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)		
		c := session.DB(mongodb_database).C(mongodb_collection)

		var result UploadRequest
		var newpic Picture

		params := mux.Vars(req)
		var file []byte] = params["file"]

		// upload file
		req.ParseMultipartForm(32 << 20)
		file, handler, err := req.FormFile("uploadfile")
		if err != nil {
		   fmt.Println(err)
		   return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
		   fmt.Println(err)
		   return
		}
		defer f.Close()
		io.Copy(f, file)
		// upload file

		// Find pictureId if it exist
		err = c.Find(bson.M{"pictureId" : requestdetail.PictureId}).One(&result)

		if err == nil {
			formatter.JSON(w, http.StatusNotFound, "This pictureid already exists in the database.")
			return
		} else {
			newpic = Picture {		
			    PictureId: requestdetail.PictureId,
				UserId: requestdetail.UserId,
				Title: requestdetail.Title,  
				Price: requestdetail.Price, 	
				Description: requestdetail.Description, 
				TumbnailUrl: "", 
				OrigUrl: "",
			}
			fmt.Println("Picture not found, inserting a new record into database")	
			pic = Picture {
				PictureId:     requestdetail.PictureId,
				UserId:        requestdetail.UserId,
				TotalAmount:   newpic.Price,
				IpAddress:	  getIp(), 
			}
			_ = json.NewDecoder(req.Body).Decode(&pic)
			err = c.Insert(pic)
			if err != nil {
				formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}
		fmt.Println("Orders: ", orders)
		formatter.JSON(w, http.StatusOK, order)
	}
}

// API update i.e. change owner
func updateHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var updaterequest Payload
		_ = json.NewDecoder(req.Body).Decode(&paymentdetail)
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
		c.Update(bson.M{"pictureId": picid}, bson.M{"$set": bson.M{"requestStatus" : result.RequestStatus, "userId" : result.UserId, "ipaddress" : getIp()}})
        fmt.Println("Request:", picid, usrid, "updated" )
		formatter.JSON(w, http.StatusOK, result)
	} 
} 

// API Delete picture by pictureId
func deleteByPictureIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()
		if err:= session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
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
		err = c.Find(bson.M{"commandId":uuid}).One(&result)
		if err!=nil{
			fmt.Println("Command not found")
			formatter.JSON(w, http.StatusNotFound, "Command Not Found")
			return
		}
		for i := 0; i < len(result.Cart); i++ {
			if result.Cart[i].ItemId == orderdetail.ItemId {
				result.TotalAmount = result.TotalAmount - result.Cart[i].Price
				result.Cart = append(result.Cart[0:i],result.Cart[i+1:]...)
				break
			}
		}
		c.Update(bson.M{"commandId": uuid}, bson.M{"$set": bson.M{"items": result.Cart, "totalAmount": result.TotalAmount, "ipaddress": getIp()}})
		fmt.Println("Delete Item: ", orderdetail.ItemId, "from order", uuid)
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Delete all pictures owned by user userId
func deleteByUserIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		defer session.Close()
		if err:= session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		  }
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var orderdetail RequiredPayload
		_ = json.NewDecoder(req.Body).Decode(&orderdetail)
		fmt.Println("order ID: ", orderdetail.OrderId)
		err = c.Remove(bson.M{"orderId": orderdetail.OrderId})
		if err!=nil{
			fmt.Println("order not found")
			formatter.JSON(w, http.StatusNotFound, "Order Not Found")
			return
		}
		formatter.JSON(w, http.StatusOK, "Order: " + orderdetail.OrderId + " Deleted")
	} 
} 