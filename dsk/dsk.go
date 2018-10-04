package dsk

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"barista.run/bar"
	"barista.run/modules/diskio"
	"barista.run/modules/diskspace"
	"barista.run/outputs"
	"barista.run/pango"
	"github.com/glebtv/custom_barista/utils"
)

func AddTo(modules []bar.Module) []bar.Module {
	modules = append(modules, diskspace.New("/"))

	path, err := exec.LookPath("findmnt")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(path, "/")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		parts := strings.Split(string(out), "\n")
		info := strings.Fields(parts[1])
		dn := strings.Split(info[1], "/")
		name := dn[len(dn)-1]

		diskio.RefreshInterval(2 * time.Second)

		sda := diskio.New(name).
			Output(func(io diskio.IO) bar.Output {
				//spew.Dump(io)
				return outputs.Pango(
					pango.Textf("io "),
					pango.Textf("%9s", outputs.IByterate(io.Input)),
					utils.Spacer,
					pango.Textf("%9s", outputs.IByterate(io.Output)),
				)
			})

		modules = append(modules, sda)
	}

	return modules
}
