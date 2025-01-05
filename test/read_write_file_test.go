// hai tran 05/01/2025
// test read and write file
// check throughput and IOPS
// go test -run TestWriteFile
// stat -c %s data

package test

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)

const NumUser = 100
const StringSize = 1024 * 1024
const SegmentSize = 1024 * 1024


func TestWriteFile(t *testing.T) {

	// open file
	file, error := os.OpenFile("data", os.O_WRONLY|os.O_CREATE, 0666)

	if error != nil {
		t.Error(error)
	}

	defer file.Close()

	// create a wait group
	var wg sync.WaitGroup

	// generate a long string
	data := []byte(strings.Repeat("a", StringSize))

	// concurrently write to file by thread
	for i := 0; i < NumUser; i++ {

		wg.Add(1)

		go func(user int) {

			for {
				fmt.Printf("user %d write fo file \n", user)
				file.WriteAt(data, int64(user*SegmentSize))
				// time.Sleep(1 * time.Second)
			}

			// defer wg.Done()

		}(i)
	}

	// wait until wait group counter is zero
	wg.Wait()

}
