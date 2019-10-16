package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "encoding/json"
    "time"
    "github.com/bijakas/sf-gamify-backend/config/routes"
	"github.com/bijakas/sf-gamify-backend/config/config"
)
var db *gorm.DB
var err error

type User struct{
    Id          int    `json:"id"`
    Username    string `json:"username"`
    Password    string    `json:"password"`
    Status      int `json:"status"`
    CreatedAt   time.Time `gorm:"type:timestamp" json:"createdAt,omitempty"`
    UpdatedAt   time.Time `gorm:"type:timestamp" json:"updatedAt,omitempty"`
    CreatedBy   int `json:"createdBy"`
    UpdatedBy   int `json:"updatedBy"`
   }

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "delete called"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")

    userID := -1
    var err error
    if val, ok := pathParams["userID"]; ok {
        userID, err = strconv.Atoi(val)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "need a number"}`))
            return
        }
    }

    commentID := -1
    if val, ok := pathParams["commentID"]; ok {
        commentID, err = strconv.Atoi(val)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "need a number"}`))
            return
        }
    }

    query := r.URL.Query()
    location := query.Get("location")

    w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

func mainxx() {
    r := mux.NewRouter()
    db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/sf-gamify?charset=utf8&parseTime=True&loc=Local")

    api := r.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("", get).Methods(http.MethodGet)
    api.HandleFunc("/users", returnAllUsers).Methods(http.MethodGet)
    api.HandleFunc("", post).Methods(http.MethodPost)
    api.HandleFunc("", put).Methods(http.MethodPut)
    api.HandleFunc("", delete).Methods(http.MethodDelete)

    api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

    log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
    config := config.GetConfig()
	app := &routes.App{}
	app.Initialize(config)
	app.Run(":3000")
}


func returnAllUsers(w http.ResponseWriter, r *http.Request){
    users := []User{}
    db.Find(&users)
    fmt.Println("Endpoint Hit: returnAllUsers")
    json.NewEncoder(w).Encode(users)
   }

//https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj
//http://www.gorillatoolkit.org/pkg/mux/
//https://tutorialedge.net/golang/creating-restful-api-with-golang/
//https://gist.github.com/phnah/1750482
//https://social.msdn.microsoft.com/Forums/sqlserver/en-US/02a3ba7f-6e8d-453c-a7ee-4aa2dfffba92/designing-security-model-around-bitwise-operations?forum=transactsql
//https://codeburst.io/using-javascript-bitwise-operators-in-real-life-f551a731ff5
//https://stackoverflow.com/questions/9385341/how-to-use-mysql-bitwise-operations-in-php
//https://github.com/TutorialEdge/create-rest-api-in-go-tutorial
//https://www.golangprograms.com/golang-restful-api-using-grom-and-gorilla-mux.html