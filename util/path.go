package util

import (
	"fmt"
	"path/filepath"
)

func GetTeamplte(baseTemplate string) (paths []string) {
	files, err := filepath.Glob("./layouts/*.gohtml")
	if err != nil {
		panic(err)
	}
	paths = append([]string{fmt.Sprintf("./template/%s.gohtml", baseTemplate)}, files...)
	return paths
}
