package tokenizer

import (
	"unsafe"
)

func b2s(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func s2b(s string) (b []byte) {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func isNumberByte(b byte) bool {
	return '0' <= b && b <= '9'
}

func bytesStarts(prefix []byte, b []byte) bool {
	if len(prefix) > len(b) {
		return false
	}
	return b2s(prefix) == b2s(b[0:len(prefix)])
}

func bytesEnds(suffix []byte, b []byte) bool {
	if len(suffix) > len(b) {
		return false
	}
	return b2s(suffix) == b2s(b[len(b)-len(suffix):])
}
