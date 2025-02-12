package note

import "testing"

func TestSort(t *testing.T) {
	SortCase01()
	SortCase02()
	SortCase03()
	SortCase04()
}

func TestSlices(t *testing.T) {
	SlicesCase01()
	SlicesCase02()
	SlicesCase03()
}

func TestStrconv(t *testing.T) {
	StrconvCase()
}

func TestStrings(t *testing.T) {
	StringsCase01()
	StringsCase02()
}

func TestBytesCase(t *testing.T) {
	BytesCase01()
}

func TestUnicodeCase(t *testing.T) {
	UnicodeCase01()
}

func TestUTF8Case(t *testing.T) {
	UTF8Case01()
}
