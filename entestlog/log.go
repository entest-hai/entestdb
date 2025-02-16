// hai 16/02/2025
// implement log
// https://github.com/inelpandzic/simpledb/blob/main/log/log.go

package entestlog

import (
	"entestdb/entestfile"
	"sync"
)

type LogMgr struct {
	logFile          string
	fm               *entestfile.FileMgr
	logPage          *entestfile.Page
	currentBlock     *entestfile.BlockID
	latestLSN        int
	latestDurableLSN int

	mu sync.Mutex
}

func NewLogMgr(fm *entestfile.FileMgr, logFile string) *LogMgr {

	currentBlock := entestfile.BlockID{
		Filename: logFile,
		Number:   0,
	}

	logPage := entestfile.NewPage(fm.BlockSize)

	logSize, error := fm.FileSize(logFile)

	if error != nil {
		panic(error)
	}

	if logSize == 0 {
		//  initial offset is the block size
		error := logPage.WriteInt(0, fm.BlockSize)
		if error != nil {
			panic(error)
		}

		_, error = fm.Write(currentBlock, logPage)
		if error != nil {
			panic(error)
		}
	} else {
		currentBlock.Number = logSize - 1

		_, error = fm.Read(currentBlock, logPage)

		if error != nil {
			panic(error)
		}
	}

	return &LogMgr{
		logFile:          logFile,
		fm:               fm,
		logPage:          logPage,
		currentBlock:     &currentBlock,
		latestLSN:        0,
		latestDurableLSN: 0,
		mu:               sync.Mutex{},
	}

}
