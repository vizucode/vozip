package services

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/visstars7/vozip/utils"
)

func Unzip(args []string) error {
	var fileName string
	if len(args) <= 0 {
		return errors.New("argument must filled")
	}

	argument := args[0]

	fileExist, err := utils.FileExists(argument)

	if err != nil {
		return err
	}

	if fileExist {
		fileName, err = filepath.Abs(argument)

		if err != nil {
			return err
		}
	} else {
		return errors.New("file doesn't exist, btw")
	}

	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	utils.Unzip(fileName, wd)

	return nil
}
