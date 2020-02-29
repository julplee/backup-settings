package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

const (
	var root = "./settings-to-copy"
)

type config struct {
	UserPath      string   `json:"user_path"`
	FoldersToSave []string `json:"folders_to_save"`
}

func main() {
	var configFiles []string

	// look for backup config files
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Base(path) == "backup-config.json" {
			configFiles = append(configFiles, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, configFile := range configFiles {
		copyFollowingConfigFile(configFile)
	}

}

func copyFollowingConfigFile(configFile string) error {
	file, err := ioutil.ReadFile(configFile)

	if err != nil {
		return err
	}

	data := config{}

	_ = json.Unmarshal([]byte(file), &data)

	for _, folderToSave := range data.FoldersToSave {
		folderToSave = filepath.Join(data.UserPath, folderToSave)
		backupFolder := root + "/backup"

		err := copyFolderToBackupFolder(folderToSave, backupFolder)

		if err != nil {
			return err
		}
	}

	return nil
}

func copyFolderToBackupFolder(folderToSave string, backupFolder string) error {
	err := copy.Copy(folderToSave, backupFolder)

	if err != nil {
		return err
	}

	return nil
}
