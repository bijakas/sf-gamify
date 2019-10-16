package handler
 
import (
	"encoding/json"
	"net/http"
 
	"github.com/gorilla/model/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)
 
func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}
