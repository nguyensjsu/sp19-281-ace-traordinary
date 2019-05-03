package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var mysql_connect = "root:cmpe281@tcp(localhost:3306)/cmpe281"

//var mysql_connect = "cmpe281:cmpe281@tcp(10.0.2.230:3306)/cmpe281"

// NewServer configures and returns a Server.
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

func init() {
	db, err := sql.Open("mysql", mysql_connect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/orders", allOrdersHandler(formatter)).Methods("GET")
	mx.HandleFunc("/placeorder", placeorderHandler(formatter)).Methods("POST")

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
			imageid    string
			paymentid  string
			amount     float64
			created_on string
		)
		db, err := sql.Open("mysql", mysql_connect)
		defer db.Close()
		if err != nil {
			log.Fatal(err)
		} else {
			rows, err := db.Query("select id, userid, imageid, paymentid, amount, created_on from orders")

			//fmt.Println(rows)
			if err != nil {
				log.Fatal(err)
			}

			defer rows.Close()
			for rows.Next() {

				err := rows.Scan(&id, &userid, &imageid, &paymentid, &amount, &created_on)

				if err != nil {
					log.Fatal(err)
				}
				orders = append(orders, order{

					Id:        id,
					Userid:    userid,
					Imageid:   imageid,
					Paymentid: paymentid,
					Amount:    amount,
				},
				)

			}

		}

		fmt.Println("All Orders:", orders)
		formatter.JSON(w, http.StatusOK, orders)

	}

}

func placeorderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Am I here?")
		data, _ := ioutil.ReadAll(req.Body)

		db, err := sql.Open("mysql", mysql_connect)
		if err != nil {
			panic(err)
		}

		defer db.Close()

		var order order
		err = json.Unmarshal(data, &order)
		fmt.Println(order)

		//err := decoder.Decode(&data)

		if err != nil {
			//panic(err)
			fmt.Println(err)
		}

		if order.Imageid == "" || order.Userid == "" || order.Amount == 0 || order.Paymentid == "" {
			formatter.JSON(w, http.StatusBadRequest, "Order Information Incomplete. Please send Imageid, userid,  amount, Paymentid")
		} else {
			insertStmt := "insert into orders (userid, imageid, paymentid, amount, payment_method_id) VALUES(?,?,?,?,1);"
			ret, err := db.Exec(insertStmt, order.Userid, order.Imageid, order.Paymentid, order.Amount)
			if err != nil {
				formatter.JSON(w, http.StatusInternalServerError, "Payment Failed")
				panic(err)

			} else {
				formatter.JSON(w, http.StatusOK, "Payment Successful")
				id, _ := ret.LastInsertId()

				subject := "Payment Successful"
				body := fmt.Sprintf(`Hi, <br>

						The payment for order id: %d is successful.  Your purchase is complete.
						Please save your order details:<br>
						Order ID: %d
						Image code: %s
						Payment ID: %s`,
					id, id, order.Imageid, order.Paymentid)

				success := sendEmail(order.Userid, body, subject)
				fmt.Println(success)
			}
		}

	}
}

func sendEmail(to string, body string, subject string) int {
	fmt.Println(to)
	fmt.Println(subject)
	fmt.Println(body)
	return 1
}
