package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

type config struct {
	UserPath      string   `json:"user_path"`
	FoldersToSave []string `json:"folders_to_save"`
}

func main() {
	var configFiles []string

	root := "./settings-to-copy"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Base(path) == "backup-config.json" {
			configFiles = append(configFiles, path)
		} else {
			fmt.Println(filepath.Base(path))
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(configFiles)
	for _, configFile := range configFiles {
		fmt.Println(configFile)
		file, err := ioutil.ReadFile(configFile)

		if err != nil {
			panic(err)
		}

		data := config{}

		_ = json.Unmarshal([]byte(file), &data)

		for _, el := range data.FoldersToSave {
			path := filepath.Join(data.UserPath, el)

			err := copy.Copy(path, root+"/copy")
			if err != nil {
				panic(err)
			}
			fmt.Println("copy")
		}
	}

}
