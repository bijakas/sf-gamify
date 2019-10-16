package config
 
import (
	"fmt"
	"log"
	"net/http"
 
	"sf-gamify-backend/handler/users"
	"sf-gamify-backend/model/model"
	"sf-gamify-backend/config/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)
 
// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}
 
// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)
 
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/sf-gamify?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Could not connect database")
	}
 
	a.Router = mux.NewRouter()
	a.setRouters()
}
 
// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/users", a.GetAllEmployees)
}
 
// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
 
// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
 
// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}
 
// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
 
// Handlers to manage Employee Data
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users.GetAllUsers(a.DB, w, r)
}
 
// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}