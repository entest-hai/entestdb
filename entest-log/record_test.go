// hai 13/02/2025
// test record
// go test -run TestBinaryEndianBytes

package entestlog

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBinaryEndianBytes(t *testing.T) {
	fmt.Println("BinaryEndianBytesTest")
	lengthBytes := make([]byte, 4)

	binary.NativeEndian.PutUint32(lengthBytes, uint32(255))

	fmt.Println(lengthBytes)
	fmt.Printf("%v", lengthBytes)
	fmt.Printf("%08b", lengthBytes)
}

func TestLogRecord(t *testing.T) {

	data := []byte("hello")
	fmt.Printf("%s\n", data)
	fmt.Printf("%08b\n", data)

	expected := []byte{5, 0, 0, 0, 'h', 'e', 'l', 'l', 'o'}
	record := NewRecord(data)

	got := record.bytes()

	fmt.Printf("%08b\n", got)
	fmt.Printf("%08b\n", expected)

	for i, v := range got {
		if v != expected[i] {
			t.Errorf("got %v, expected %v", v, expected[i])
		}
	}
}

