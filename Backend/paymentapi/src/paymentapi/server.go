package main

import (
	//"fmt"
	"github.com/codegangsta/negroni"

	"net/http"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
   "github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"

)

var mysql_connect = "root:cmpe281@tcp(localhost:3306)/cmpe281"


// NewServer configures and returns a Server.
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

func init() {

	db, err := sql.Open("mysql", mysql_connect)
	if err != nil {
		log.Fatal(err)
	} else {
		var (
			id int
			userid string
			imageid int
			paymentid int

		)
		rows, err := db.Query("select id, userid, imageid, paymentid, amount, created_on from orders where id = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &userid, &imageid, &paymentid)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, userid, imageid, paymentid)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	defer db.Close()

}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/orders", allOrdersHandler(formatter)).Methods("GET")

}

func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Payment API is alive"})
	}
}

// API get all orders
func allOrdersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println( "Orders:", orders )
			var orders_array [] order
			for key, value := range orders {
    			fmt.Println("Key:", key, "Value:", value)
    			orders_array = append(orders_array, value)
			}
			formatter.JSON(w, http.StatusOK, orders_array)

	}
}

