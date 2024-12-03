package conv

import "unsafe"

func ToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}

	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func ToStr(b []byte) string {
	if len(b) == 0 {
		return ""
	}

	return unsafe.String(unsafe.SliceData(b), len(b))
}
