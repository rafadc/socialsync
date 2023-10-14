package socialsync

import (
	"github.com/charmbracelet/log"
	"os"
	"time"
)

var baseFolder string
var syncFileName string

func init() {
	baseFolder = os.Getenv("BASE_FOLDER")
	if baseFolder == "" {
		baseFolder = "/data"
	}

	syncFileName = baseFolder + "/last_sync.txt"
}

func LatestSyncDate() time.Time {
	if _, err := os.Stat(syncFileName); err == nil {
		content, err := os.ReadFile(syncFileName)
		if err != nil {
			log.Fatal(err)
		}
		format := "2006-01-02 15:04:05.999999999 -0700 MST"
		dateRead, err := time.Parse(format, string(content))
		if err != nil {
			log.Fatal(err)
		}
		return dateRead
	} else {
		return time.Now()
	}
}

func UpdateSyncDate() {
	f, err := os.Create(syncFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	val := time.Now().UTC().String()
	data := []byte(val)

	log.Debug("Writing new sync date: ", data)
	_, err = f.Write(data)

	if err != nil {
		log.Fatal(err)
	}
}
