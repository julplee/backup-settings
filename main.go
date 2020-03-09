package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

const (
	root = "./settings-to-copy"
)

type config struct {
	UserPath        string   `json:"user_path"`
	BackupPath      string   `json:"backup_path"`
	FoldersToSave   []string `json:"folders_to_save"`
	FoldersToIgnore []string `json:"folders_to_ignore"`
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
		backupFollowingConfigFile(configFile)
	}

}

func backupFollowingConfigFile(configFile string) error {
	file, err := ioutil.ReadFile(configFile)

	if err != nil {
		return err
	}

	data := config{}

	_ = json.Unmarshal([]byte(file), &data)

	toSkip := make(map[string]struct{})
	for _, folderToSkip := range data.FoldersToIgnore {
		folderToSkip = filepath.Join(data.UserPath, folderToSkip)
		toSkip[folderToSkip] = struct{}{}
	}

	for _, folderToSave := range data.FoldersToSave {
		backupFolder := filepath.Join(data.BackupPath, folderToSave)
		folderToSave = filepath.Join(data.UserPath, folderToSave)

		err := copyFolderToBackupFolder(folderToSave, backupFolder, toSkip)

		if err != nil {
			return err
		}
	}

	return nil
}

func copyFolderToBackupFolder(folderToSave string, backupFolder string, toSkip map[string]struct{}) error {
	opt := copy.Options{Skip: func(src string) bool {
		if _, isToSkip := toSkip[src]; isToSkip {
			return true
		}

		return false
	}}

	err := copy.Copy(folderToSave, backupFolder, opt)

	if err != nil {
		return err
	}

	return nil
}
