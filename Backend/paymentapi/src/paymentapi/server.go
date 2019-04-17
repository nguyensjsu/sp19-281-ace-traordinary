package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	//"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
	"net/http"
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
			id        int
			userid    string
			imageid   int
			paymentid int
			amount    float64
		)
		rows, err := db.Query("select id, userid, imageid, paymentid, amount from orders where id = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &userid, &imageid, &paymentid, &amount)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, userid, imageid, paymentid, amount)
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
	mx.HandleFunc("/order", allOrdersHandler(formatter)).Methods("POST")

}

func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Payment API is alive"})
	}
}

// API get all orders
func allOrdersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var orders []order

		//var orderRow order
		var (
			id         int
			userid     string
			imageid    int
			paymentid  int
			amount     float64
			created_on string
		)
		db, err := sql.Open("mysql", mysql_connect)
		defer db.Close()
		if err != nil {
			log.Fatal(err)
		} else {
			rows, err := db.Query("select id, userid, imageid, paymentid, amount, created_on from orders")

			fmt.Println(rows)
			if err != nil {
				log.Fatal(err)
			}

			defer rows.Close()
			for rows.Next() {

				err := rows.Scan(&id, &userid, &imageid, &paymentid, &amount, &created_on)

				if err != nil {
					log.Fatal(err)
				}
				orders = append(orders, order {
					
						Id:        id,
						userid:    userid,
						imageid:   imageid,
						paymentid: paymentid,
						amount:    amount,
					},
				)

				//log.Println(id, userid, imageid, paymentid, amount)
			}

		}

		
		//ordersJson, err := json.Marshal(orders)
    	//if err != nil {
       	 //log.Fatal("Cannot encode to JSON ", err)
   		 //}
         //fmt.Println(ordersJson)

		fmt.Println("All Orders:", orders)
		formatter.JSON(w, http.StatusOK, orders)
		
	}
}


