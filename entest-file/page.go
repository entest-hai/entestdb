// hai 06/02/2025
// implement page
// https://github.com/inelpandzic/simpledb/blob/main/file/page.go

package entestfile

import (
	"encoding/binary"
	"errors"
)

type Page struct {
	bytes []byte
}

// New page with the specified size
func NewPage(size int) *Page {
	return &Page{bytes: make([]byte, size)}
}

// Check size of the page 
func (p *Page) Size() int {
	return len(p.bytes)
}

// Write copies data from data slice to page at specified offset
func (p *Page) Write(offset int, data []byte) (int, error) {
	if offset+len(data) > p.Size() {
		return 0, errors.New("data exceeds page bounds")
	}

	n := copy(p.bytes[offset:], data)

	return n, nil
}

// WriteInt writes an integer value to the page at the specified offset
func (p *Page) WriteInt(offset int, value int) error {
	b := make([]byte, 4)
	binary.NativeEndian.PutUint32(b, uint32(value))

	_, error := p.Write(offset, b)
	return error
}


// Read copies data from the page at offset and write it to data slice 
func (p *Page) Read(offset int, data []byte) int {
  return copy(data, p.bytes[offset:])
}

// Bytes returns the byte of the page 
func (p *Page) Bytes() []byte {
	return p.bytes
} 