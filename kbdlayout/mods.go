package kbdlayout

const NUM_LOCK = 16
const CAPS_LOCK = 2

func GetMods(mods uint8) []string {
	ret := make([]string, 0)
	if mods&NUM_LOCK == NUM_LOCK {
		ret = append(ret, "NUM")
	}
	if mods&CAPS_LOCK == CAPS_LOCK {
		ret = append(ret, "CAPS")
	}
	//log.Println("getmods", mods, ret)
	return ret
}
