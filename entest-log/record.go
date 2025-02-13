// hai 13/02/2025
// implement log record
// https://github.com/inelpandzic/simpledb/blob/main/log/record.go

package entestlog

import "encoding/binary"

const intByteSize = 4

type Record struct {
	Length int
	Data   []byte
}

func NewRecord(data []byte) *Record {
	return &Record{
		Length: len(data),
		Data:   data,
	}
}

// bytes return whole record bytes, length 4-byte metadata field plus data
func (r *Record) bytes() []byte {
	lengthBytes := make([]byte, intByteSize)
	binary.NativeEndian.PutUint32(lengthBytes, uint32(r.Length))
	return append(lengthBytes, r.Data...)
}

// total length returns the total length of the record, including the length field
func (r *Record) TotalLength() int {
	return intByteSize + r.Length
}
