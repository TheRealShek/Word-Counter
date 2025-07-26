package filescanner

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindTxtFiles(root string) ([]string, error) {
	var txtFiles []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Preventing walk: %v \n", err)
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			txtFiles = append(txtFiles, path)
		}
		return nil
	})

	return txtFiles, err
}
