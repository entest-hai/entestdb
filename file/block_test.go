// hai tran 01/01/2025
// book: database design and implementation, chapter 3
// implement the logical file block
package file

import (
	"fmt"
	"testing"
)

func TestBlockName(t *testing.T) {
	block := NewBlock("demo", 10)
	fmt.Printf("block name: %s, id: %s, number: %d \n", block.Filename(), block.ID(), block.Number())
}
