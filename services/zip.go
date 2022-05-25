package services

import (
	"errors"
	"path/filepath"

	"github.com/visstars7/vozip/utils"
)

func Zip(args []string) error {

	var source string

	if len(args) <= 0 {
		return errors.New("argument must filled")
	}

	file := args[0]

	fileExist, err := utils.FileExists(file)

	if err != nil {
		return err
	}

	if fileExist {
		source, err = filepath.Abs(file)

		if err != nil {
			return err
		}

	} else {
		return errors.New("file doesn't exist, btw")
	}

	target := utils.LastDirName(source) + ".zip"
	err = utils.Zip(source, target)

	if err != nil {
		return err
	}

	return nil
}
