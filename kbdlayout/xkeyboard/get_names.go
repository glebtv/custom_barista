package xkeyboard

import (
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/davecgh/go-spew/spew"
)

// GetVersionCookie is a cookie used only for GetVersion requests.
type GetNamesCookie struct {
	*xgb.Cookie
}

const XkbSymbolsNameMask = (1<<2)
const XkbGroupNamesMask = (1<<12)

// GetVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetVersionCookie.Reply()
func GetNames(c *xgb.Conn, which uint32) GetNamesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XKEYBOARD"]; !ok {
		panic("Cannot issue request 'GetVersion' using the uninitialized extension 'XTEST'. xtest.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getNamesRequest(c, which), cookie)
	return GetNamesCookie{cookie}
}

//BYTE	type;
//BYTE	deviceID;
//CARD16	sequenceNumber B16;
//CARD32	length B32;
//CARD32	which B32;
//KeyCode	minKeyCode;
//KeyCode	maxKeyCode;
//CARD8	nTypes;
//CARD8	groupNames;
//CARD16	virtualMods B16;
//KeyCode	firstKey;
//CARD8	nKeys;
//CARD32	indicators B32;
//CARD8	nRadioGroups;
//CARD8	nKeyAliases;
//CARD16	nKTLevels B16;
//CARD32	pad3 B32;

// GetNamesReply represents the data returned from a GetNames request.
type GetNamesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	//MajorVersion byte
	//MinorVersion uint16
}

// Reply blocks and returns the reply data for a GetNames request.
func (cook GetNamesCookie) Reply() (*GetNamesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getNamesReply(buf), nil
}

// getNamesReply reads a byte slice into a GetNamesReply value.
func getNamesReply(buf []byte) *GetNamesReply {
	log.Println("parsing reply:")
	spew.Dump(buf)
	v := new(GetNamesReply)
	b := 1 // skip reply determinant

	//v.MajorVersion = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	//v.MinorVersion = xgb.Get16(buf[b:])
	//b += 2

	return v
}

// Write request to wire for GetNames
// getNamesRequest writes a GetNames request to a byte slice.
func getNamesRequest(c *xgb.Conn, which uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XKEYBOARD"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 17 // request opcode X_kbGetNames
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], uint16(0)) // deviceSpec ?
	b += 2

	xgb.Put16(buf[b:], uint16(0)) // pad ?
	b += 2

	xgb.Put32(buf[b:], which)
	b += 4
	log.Println("request:")
	spew.Dump(buf)
	return buf
}
