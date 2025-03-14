package main 
import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"context"
	"os"
	"os/signal"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware" 
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	//mgo is a library that helps us to talk to the database
	"gopkg.in/mgov2/bson"
)

var rend *renderer.Render
var db *mgo.Database
// the constants to be used throughout the application
const (
	hostName   string = "localhost:27017"
	dbName   string = "todo"
	collectionName string ="todos"
	port  string = "9001" 
)

type(
	todomodel struct {
		ID bson.ObjectId `bson:"_id,omitempty"`
		Title string `bson:"title"`
		Completed bool `bson:"completed"`
		CreatedAt time.Time `json:"created_at"`
		//bon to interacte with the database
	}
	todo struct {
		ID string `json:"id"`
		Title string `json:"title"`
		Completed bool `json:"completed"`
		Created time.Time `json:"created_at"`
		//json to interact with the frontend
	}
) 

func init(){
	rend = renderer.New()
	sess, err := mgo.Dial(hostName)
	if err != nil {
		log.Fatal(err)
	}
	db = sess.DB(dbName)
}
func homeHandler(w http.ResponseWriter,r*http.Request){
	err := rend.Template(w,http.StatusOK,[]string{"static/home.tpl"},nil)
	//.Template comes together
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func main(){
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Mount("/todo", todoHandlers())
	//Mount groups related routes together ... in this case it makes any rute that stats with /todo call todoHandlers function 
srv := &http.Server{
	Addr:port,
	Handler:r,
	ReadTiemout: 60 * time.Second,
	WriteTimeout: 60 * time.Second,
	IdleTimeout: 60 * time.Second,
}
go func(){
	log.Println("Listening on port " + port)
	if err:=srv.ListenandServer(); err != nil {
		log.Printf("listen:%s\n",err)
	}
}()
}


func todoHandlers() http.Handler {
	rg:= chi.NewRouter()
	rg.Group(func(r chi.Router){
		r.Post("/", createTodoHandler)
		r.Get("/", getAllTodosHandler)
	
		r.Put("/{id}", updateTodoHandler)
		r.Delete("/", deleteTodoHandler)
	})
}