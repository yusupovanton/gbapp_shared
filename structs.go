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
	User_id     uint           `json:"user_id" redis:"user_id"`
	Username    string         `json:"username,omitempty" redis:"username"`
	Price       string         `json:"price,omitempty" redis:"price"`
	Category    string         `json:"category,omitempty" redis:"category"`
	Title       string         `json:"title,omitempty" redis:"title"`
	Location    string         `json:"location,omitempty" redis:"location"`
	Contacts    string         `json:"contacts,omitempty" redis:"contacts"`
	Description string         `json:"description,omitempty" redis:"description"`
	Post        bool           `json:"post,omitempty" redis:"post"`
	Last_posted time.Time      `json:"last_posted,omitempty" redis:"last_posted"`
	Images      pq.StringArray `json:"images,omitempty" gorm:"type:string[]"`
	Image_count uint           `json:"image_count,omitempty" redis:"image_count"`
}
