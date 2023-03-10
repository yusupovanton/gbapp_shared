package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	telegraph "github.com/StarkBotsIndustries/telegraph"
	tgbotapi "gitlab.com/kingofsystem/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func PostAd(user *User, ad *Ad, db *gorm.DB) (uint, bool, error) {

	success := true
	var err error

	result := db.Table("gbapp_ads").Create(&ad)

	if result.Error != nil {
		success = false
		log.Printf("There was an error posting an ad: %v", result.Error)
	} else {
		user.Ads = append(user.Ads, int64(ad.ID))
		db.Table("gbapp_users").Save(&user)
	}

	return ad.ID, success, err
}

func CheckUser(user_id uint, db *gorm.DB) (bool, error) {

	var new = false

	result := db.Table("gbapp_users")

	if result.Error == gorm.ErrRecordNotFound {
		new = true
	} else {
		new = false
	}

	return new, nil
}

func InitNewUser(user *User, db *gorm.DB) (bool, error) {
	//1. write the user to users
	//2. assign the state to main menu
	//3. tie it all together in the session table

	var success = true

	user.StateName = "main_menu"
	user.StateOverride = false

	result := db.Table("gbapp_users").Create(&user)

	errorCheck(result.Error)
	log.Printf("A user has been created with user id: %v", user.ID)

	return success, nil
}

func UpdateUser(user *User, db *gorm.DB) (bool, error) {

	result := db.Table("gbapp_users").Save(&user)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func UpdateAd(ad *Ad, db *gorm.DB) (bool, error) {

	result := db.Table("gbapp_ads").Save(&ad)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func GetUser(user_id uint, db *gorm.DB) (*User) {

	//1. gets state and ad ids by using the user id in the session table
	//2. gets the state and current ad details by their ids

	var err error
	var user *User

	result := db.Table("gbapp_users").Where("id = ?", user_id).Find(&user)

	if result.Error == gorm.ErrRecordNotFound {
		log.Printf("There is no user with the user_id %v", user_id)
	} else if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("An unexpected error while getting session from the database, %v", err)
	}

	return user
}

func GetAd(ad_id uint, db *gorm.DB) (*Ad) {

	var err error
	var ad *Ad

	result := db.Table("gbapp_ads").Where("ad_id = ?", ad_id).Last(&ad)

	//Ad can be empty
	if result.Error == gorm.ErrRecordNotFound {
		err = nil
		log.Printf("There is no ad for the ad_id %d", ad_id)
	} else if err != nil && err != sql.ErrNoRows {
		log.Printf("An unexpected error while getting ad from the database")
	}

	return ad
}

func DeleteUser(user_id uint, db *gorm.DB) (bool, error) {

	// Deletes rows by a user_id
	var success = true
	var user *User

	result := db.Table("gbapp_users").Where("user_id = ?", user_id).Delete(&user)

	if result.Error != nil {
		success = false
		log.Printf("There was an error deleting the user from the db! %v", result.Error)
	}

	return success, nil
}

func CreateTelegraphPage(ad *Ad) *telegraph.Page {

	var images string
	var url string

	for _, id := range ad.Images {
		url = fmt.Sprintf("https://storage.googleapis.com/gotgbot_bucket/%v", id)
		images = images + fmt.Sprintf(`<img src="%v" width="500" height="600">`, url)
	}
	log.Printf("%v", images)
	htmlData := fmt.Sprintf(`%v<h1>%v</h1><h2>%v</h2><b>%v</b><p>%v</p><u>%v</u>`,
		images, ad.Title, ad.Location, ad.Contacts, ad.Description, ad.CreatedAt)

	page, err := telegraph.CreatePage(&telegraph.CreatePageOpts{
		Title:       ad.Title + " | " + ad.Location,
		Content:     htmlData,
		AccessToken: os.Getenv("TELEGRAPH_KEY")})
	errorCheck(err)

	return page
}

func CreateAdMessage(ad *Ad, modifyOpts bool) tgbotapi.MessageConfig {

	// Don't forget to enter chat id!

	var msgRet tgbotapi.MessageConfig

	page := CreateTelegraphPage(ad)

	markdownTemplate := fmt.Sprintf("**%v**\n---\n%v\n_%v_\n---\n[%v](%v)", ad.Title, ad.Description, ad.Price, page.Title, page.URL)

	// Adding commangs for change
	if modifyOpts {
		changeLink := fmt.Sprintf(`???????????????? ????????????????????: */changeAd_%v*`, ad.ID)
		markdownTemplate += fmt.Sprintf("\n---\n%v", changeLink)
	}

	msgRet.Text = markdownTemplate
	msgRet.ParseMode = "markdown"

	// TODO: photo media group
	// var medias []interface{}
	// media1 := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL("http://example.com/test1.jpg"))

	// mediaArr := append(medias, media1)
	// var mediaGr = tgbotapi.MediaGroupConfig{ChatID: -1001515307140, Media: mediaArr}

	return msgRet
}
