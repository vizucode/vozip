package services

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/visstars7/vozip/utils"
)

func Unzip(args []string) error {

	var (
		fileName string
		wd       string
	)

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

	if len(args) > 1 {
		dirPath := args[1]

		fileExist, err := utils.FileExists(dirPath)

		if err != nil {
			return nil
		}

		if fileExist {
			wd = dirPath
		} else {
			return errors.New("destination path not found")
		}
	} else {
		wd, err = os.Getwd()
	}

	if err != nil {
		return err
	}

	utils.Unzip(fileName, wd)

	return nil
}
