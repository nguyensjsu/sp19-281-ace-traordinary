package main


import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"os"
)

var pool = newPool()
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
	mx.HandleFunc("/ws/{Id}", socketHandler(formatter)).Methods("GET")
}


func socketHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("entered...",req.URL);
		serveWs(pool, w, req)



	}
}
func main(){

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	go pool.run()
	//http.HandleFunc("/", serveHome)
	fmt.Println("http server started on :8000")
	server := NewServer()
	server.Run(":" + port)
	//mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("entered...",r.URL);
	//	serveWs(pool, w, r)
	//})
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
//	panic( err)
	//}




	//http.HandleFunc("/ws", handleConnections)
	//go handleMessages()
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
	//	fmt.Errorf("ListenAndServe: ", err)
//	}

}

