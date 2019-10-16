package model
 
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
 
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