// test package includes common methods to run tests.
// It should not be included in release builds
package test

import (
	"testing"
)

const (
	logfile          = "testlog"
	blockfile        = "testfile"
	blockSize        = 400
	buffersAvaialble = 3
)

type Conf struct {
	DbFolder         string
	LogFile          string
	BlockFile        string
	BlockSize        int
	BuffersAvailable int
}

func DefaultConfig(t *testing.T) Conf {
	return Conf{
		// DbFolder:         t.TempDir(),
    DbFolder: "/ubuntu-data/home/ubuntu/workspace/go-workspace/entestdb/file",
		LogFile:          logfile,
		BlockFile:        blockfile,
		BlockSize:        blockSize,
		BuffersAvailable: buffersAvaialble,
	}
}
