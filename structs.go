package lib

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Ads           pq.Int64Array `json:"ads" gorm:"type:integer[]"`
	CurrentAd     uint          `json:"currentad,omitempty"`
	Username      string        `json:"username" redis:"username"`
	StateName     string        `json:"statename,omitempty" redis:"statename"`
	StateOverride bool          `json:"stateoverride,omitempty" redis:"stateoverride"`
	Password      string        `json:"password"`
}


type Ad struct {
	gorm.Model
	User_id     uint           `redis:"user_id"`
	Username    string         `redis:"username"`
	Price       string         `redis:"price"`
	Category    string         `redis:"category"`
	Title       string         `redis:"title"`
	Location    string         `redis:"location"`
	Contacts    string         `redis:"contacts"`
	Description string         `redis:"description"`
	Post        bool           `redis:"post"`
	Last_posted time.Time      `redis:"last_posted"`
	Images      pq.StringArray `gorm:"type:string[]"`
	Image_count uint           `redis:"image_count"`
}
