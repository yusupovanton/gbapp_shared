package lib_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	. "github.com/yusupovanton/gbapp_shared"
)

func errorCheck(err error) {
	if err != nil {
		log.Printf("An error occured in the lib_test module: %v", err)
	}
}

func init() {
	os.Setenv("ENV_FILE_LOCATION", "/Users/yustas/go/gbapp/gbapp_shared/lib-local.env")
	err := godotenv.Load(MustGetEnv("ENV_FILE_LOCATION"))
	errorCheck(err)
}

func TestDbActions(t *testing.T) {

	t.Parallel()

	ad := new(Ad)
	ad.Images = pq.StringArray{"AgACAgIAAxkBAAPbY7wgHkDCBftv0a7NswLxtDLfJ68AArfFMRsYheFJNa-EUaXpBE8BAAMCAAN5AAMtBA,AgACAgIAAxkBAAPcY7wgHtYUI7r64EUKF5bsxib8fSoAArnFMRsYheFJSzbtPu3z17ABAAMCAAN5AAMtBA", "AgACAgIAAxkBAAPdY7wgHgJG_NAkCn6Ygk99vK9m-BoAArrFMRsYheFJMgZNJF0bgwcBAAMCAAN5AAMtBA"}
	ad.Title = "Test Title"
	ad.Price = "Test Price"
	ad.Contacts = "Test Contacts"
	ad.Category = "Test Category"
	msgPost := CreateAdMessage(ad, false)

	log.Printf(msgPost.Text)
	require.NotNil(t, msgPost)

}
