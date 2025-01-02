// hai tran 01/01/2025
// book: database design and implementation, chapter 3
// implement the logical file block
package file

import "fmt"

type BlockID string

type Block struct {
	id       BlockID
	filename string
	number   int
}

func NewBlock(filename string, number int) Block {
	return Block{
		id:       BlockID(fmt.Sprintf("f:%sb:%d", filename, number)),
		filename: filename,
		number:   number,
	}
}

func (bid Block) ID() BlockID {
	return bid.id
}

func (bid Block) Filename() string {
	return bid.filename
}

func (bid Block) Number() int {
	return bid.number
}

func (bid Block) Equals(other Block) bool {
	return bid.filename == other.filename && bid.number == other.number
}

func (bid Block) String() string {
	return fmt.Sprintf("file %q block %d", bid.filename, bid.number)
}
