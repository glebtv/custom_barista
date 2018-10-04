package netm

import (
	"net"
	"strings"
	"time"

	"github.com/glebtv/custom_barista/utils"
	"barista.run/bar"
	"barista.run/colors"
	"barista.run/modules/netspeed"
	"barista.run/modules/wlan"
	"barista.run/outputs"
	"barista.run/pango"
)

func AddTo(modules []bar.Module) []bar.Module {
	ifs, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, ifc := range ifs {
		//spew.Dump(ifc)
		if strings.HasPrefix(ifc.Name, "en") || strings.HasPrefix(ifc.Name, "wl") {
			//log.Println("add interface", ifc.Name)
			ift := ifc
			net := netspeed.New(ift.Name).
				RefreshInterval(2 * time.Second).
				Output(func(s netspeed.Speeds) bar.Output {
					// to update flags
					ift, err := net.InterfaceByName(ift.Name)
					if err != nil {
						return outputs.Error(err)
					}
					addrs, err := ift.Addrs()
					if err != nil {
						return outputs.Error(err)
					}

					up := ift.Flags&net.FlagUp != 0
					var up_text string
					if up {
						up_text = " UP "
					} else {
						up_text = "DOWN"
					}
					var upSeg *pango.Node
					if up {
						upSeg = pango.Text(up_text).Color(colors.Scheme("good"))
					} else {
						upSeg = pango.Text(up_text).Color(colors.Scheme("bad"))
					}

					ips := make([]string, 0)
					if err == nil {
						// handle err
						for _, addr := range addrs {
							var ip net.IP
							switch v := addr.(type) {
							case *net.IPNet:
								ip = v.IP
							case *net.IPAddr:
								ip = v.IP
							}
							//spew.Dump(addr, ip)
							if ip.To4() != nil {
								ips = append(ips, ip.String())
							}
						}
						//log.Println("addrs:", ift.Name, ips)
					}

					things := []interface{}{
						pango.Textf("%s", ift.Name),
						utils.Spacer,
						upSeg,
					}

					if up {
						upthings := []interface{}{
							pango.Icon("material-file-upload"),
							utils.Spacer,
							pango.Textf("%8s", outputs.Byterate(s.Tx)),
							pango.Text(" ").Small(),
							pango.Icon("material-file-download"),
							utils.Spacer,
							pango.Textf("%8s", outputs.Byterate(s.Rx)),
							utils.Spacer, pango.Textf("%s", strings.Join(ips, "|")),
						}
						things = append(things, upthings...)
					}

					return outputs.Pango(things...)
				})
			modules = append(modules, net)
		}

		if strings.HasPrefix(ifc.Name, "wl") {
			wlan := wlan.Named(ifc.Name)
			modules = append(modules, wlan)
		}
	}
	return modules
}
