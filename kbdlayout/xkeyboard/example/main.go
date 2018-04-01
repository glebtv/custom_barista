package main

import (
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/davecgh/go-spew/spew"
	"github.com/glebtv/custom_barista/kbdlayout/xkeyboard"
)

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	err = xkeyboard.Init(X)
	if err != nil {
		log.Fatal(err)
	}
	// this is really UseExtension message
	xkeyboard.GetVersion(X, 1, 0)
	//spew.Dump(resp.Reply())

	// GetAtomName for atom=0x1f2 (symbolsName atom)
	anresp := xproto.GetAtomName(X, xproto.Atom(0x1f2))
	anreply, err := anresp.Reply()
	log.Println("symbol names reply:")
	log.Println(anreply.Name)
	//spew.Dump(anreply, err)

	log.Println("get state start")
	sresp := xkeyboard.GetState(X, 3)
	//spew.Dump(sresp)
	sreply, err := sresp.Reply()
	log.Println("getstate reply:")
	spew.Dump(sreply, err)

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
}
