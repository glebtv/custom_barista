package utils

import (
	"os/user"
	"path/filepath"

	"barista.run/pango"
)

var Spacer = pango.Text(" ").XXSmall()

func Home(path string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, path)
}
