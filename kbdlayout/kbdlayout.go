package kbdlayout

import (
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/davecgh/go-spew/spew"
	"github.com/glebtv/custom_barista/kbdlayout/xkeyboard"
)

func GetLayout() (string, error) {
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	err = xkeyboard.Init(X)
	if err != nil {
		log.Fatal(err)
	}
	//newKeymap, keyErr := xproto.GetKeyboardMapping(xu.Conn(), min, byte(max-min+1)).Reply()
	//spew.Dump(newKeymap, keyErr)

	//resp := xkeyboard.GetVersion(X, 0, 0)
	//spew.Dump(resp.Reply())

	nresp := xkeyboard.GetNames(X, xkeyboard.XkbSymbolsNameMask)
	log.Println("reply:")
	//spew.Dump(nresp)
	repl, err := nresp.Reply()
	spew.Dump(repl, err)

	nresp = xkeyboard.GetNames(X, xkeyboard.XkbGroupNamesMask)
	log.Println("reply:")
	//spew.Dump(nresp)
	repl, err = nresp.Reply()
	spew.Dump(repl, err)

	// Get the window id of the root window.
	//setup := xproto.Setup(X)
	//root := setup.DefaultScreen(X).Root

	return "test", nil
}
