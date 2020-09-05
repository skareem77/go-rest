package app

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/src/app/handler"
	"rest-api/src/app/model"
	"rest-api/src/app/util"
	"rest-api/src/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// App struct
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// RequestHandlerFunction type to handle request
type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

// Init initializer
func (a *App) Init(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.DBName,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)

	if err != nil {
		log.Fatalln(err.Error())
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/api/projects", a.handleRequest(handler.GetAllProjects)).Methods("GET")
	a.Router.HandleFunc("/api/projects", a.handleRequest(handler.SaveProject)).Methods("POST")
	a.Router.HandleFunc("/api/login", a.handleRequest(handler.Login)).Methods("POST")
	a.Router.Use(util.JwtAuthentication)
}

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

//Run application
func (a *App) Run(host string) {
	err := http.ListenAndServe(host, a.Router)
	log.Fatalln(err.Error())
}
