// hai tran 05/01/2025
// implement file manager

package file

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

const (
	tmpTablePrefix = "__tmp__"
	TmpTablePrefix = tmpTablePrefix + "%d"

	// Max size of logged table file name
	MaxLoggedTableFileNameSize = 255
)

// Implements methods that read and write pages to disk blocks.
// It always read and writes a block-sized number of bytes from a file, always at a block boundary.
// This ensure that each call to read, write or append will occur exactly one disk access.

type FileManager struct {
	folder    string
	blockSize int
	isNew     bool
	// Maps a file name to an open file
	// Files are opened in RWS mode
	openFiles map[string]*os.File
	sync.Mutex
}

func NewFileManager(path string, blockSize int) *FileManager {
	_, err := os.Stat(path)

	isNew := os.IsNotExist(err)
	// If the folder does not exist, create one
	if isNew {
		os.Mkdir(path, os.ModeSticky|os.ModePerm)
	}

	if !isNew && err != nil {
		panic(err)
	}

	// Clear all tmp files in the folder
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, v := range entries {
		if strings.HasPrefix(v.Name(), tmpTablePrefix) {
			fn := filepath.Join(path, v.Name())
			if err := os.Remove(fn); err != nil {
				panic(err)
			}
		}
	}

	return &FileManager{
		folder:    path,
		blockSize: blockSize,
		isNew:     isNew,
		openFiles: make(map[string]*os.File),
	}
}

func (manager *FileManager) getFile(fname string) *os.File {
	f, ok := manager.openFiles[fname]

	if !ok {
		p := path.Join(manager.folder, fname)
		table, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR|os.O_SYNC, 0o755)
		if err != nil {
			panic(err)
		}
		manager.openFiles[fname] = table
		return table
	}

	return f
}

func (manager *FileManager) IsNew() bool {
	return manager.isNew
}

func (manager *FileManager) BlockSize() int {
	return manager.blockSize
}

// Read reads the content of a block id blk into Page page
func (manager *FileManager) Read(blk Block, p *Page) {
	manager.Lock()

	defer manager.Unlock()

	f := manager.getFile(blk.Filename())

	//
	if _, err := f.ReadAt(p.contents(), int64(blk.Number()*manager.blockSize)); err != io.EOF && err != nil {
		panic(err)
	}
}

// Write writes Page p to BlockID block, persisted to a file
func (manager *FileManager) Write(blk Block, p *Page) {
	manager.Lock()

	defer manager.Unlock()

	f := manager.getFile(blk.Filename())

	f.WriteAt(p.contents(), int64(blk.Number()*manager.blockSize))
}

// Size returns the size, in blocks of the given file
func (manager *FileManager) Size(filename string) int {
	f := manager.getFile(filename)
	finfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	return int(finfo.Size() / int64(manager.blockSize))
}

// Append seeks to the end of the file and writes an empty array of bytes to the file
func (manager *FileManager) Append(filename string) Block {
	newBlkNum := manager.Size(filename)
	block := NewBlock(filename, newBlkNum)
	buf := make([]byte, manager.blockSize)

	f := manager.getFile(filename)
	f.WriteAt(buf, int64(newBlkNum*manager.blockSize))
	return block
}
