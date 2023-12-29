package logs

import (
	"log"
	"os"
	"path/filepath"
)

var GenLog *log.Logger
var ErrLog *log.Logger

func InitLogger(filePath string) error {

	dir := filepath.Dir(filePath)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	generalLog, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	GenLog = log.New(generalLog, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrLog = log.New(generalLog, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)

	GenLog.Println("\n\n\nApplication restarted ---------")

	return nil

}
