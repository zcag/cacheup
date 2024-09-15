package main

import (
	"github.com/zcag/cacheup/util"
)

func Write(path, content string) error {
		err := util.SetContent("", path, content)
		if err != nil { return err }

		return nil
}

func Read(path, max_age, command string) (string, error) {
	if command != "" {
		valid, err := IsValid(path, max_age)
		if err != nil { return "", err }

		if !valid {
			stdout, err := util.Exec(command)
			if err != nil { return "", err }
			err = Write(path, stdout)
			if err != nil { return "", err }

			return stdout, nil
		}
	}

	content, err := util.GetContent("", path)
	if err != nil { return "", err }

	return content, nil
}

func IsValid(path, max_age string) (bool, error) {
	return util.IsCacheValid("", path, max_age)
}
