// hai tran 05/01/2025
// file manager test

package file

import (
	"entestdb/test"
	"testing"
)

func TestFile(t *testing.T) {

	confg := test.DefaultConfig(t)
	fman := NewFileManager(confg.DbFolder, confg.BlockSize)

	block := NewBlock(confg.BlockFile, 2)
	page := NewPage()

	pos := 88

	const val = "abcdefghilmno"
	const invt = 352
  page.SetString(pos, val)

	pos2 := pos + StrLength(len(val))

	page.SetInt(pos2, invt)

	fman.Write(block, page)

	p2 := NewPage()
	fman.Read(block, p2)

	if got := p2.Int(pos2); got != invt {

		t.Fatalf("expected %d at offset %d. Got %d", invt, pos2, got)
	}
	t.Logf("offset %d contains %d", pos2, p2.Int(pos2))

	if got := p2.String(pos); got != val {
		t.Fatalf("expected %q at offset %d. Got %q", val, pos, got)
	}

	t.Logf("offset %d contains %s", pos, p2.String(pos))

}
