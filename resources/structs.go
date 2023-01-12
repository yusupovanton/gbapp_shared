// STRUCTS
package resources

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Ads           pq.Int64Array `json:"ads" gorm:"type:integer[]"`
	CurrentAd     uint          `json:"currentad,omitempty"`
	Username      string        `json:"username"`
	StateName     string        `json:"statename,omitempty"`
	StateOverride bool          `json:"stateoverride,omitempty"`
	Password      string        `json:"password"`
}

// type State struct {
// 	gorm.Model
// 	Id        int
// 	StateName string
// 	Override  bool
// 	Last_changed int64
// }

type Ad struct {
	gorm.Model
	User_id     uint
	Username    string
	Price       string
	Category    string
	Title       string
	Location    string
	Contacts    string
	Description string
	Post        bool
	Last_posted time.Time
	Images      pq.StringArray `gorm:"type:string[]"`
}
