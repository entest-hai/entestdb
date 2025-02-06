// hai 06/02/2025
// implement file
// https://github.com/inelpandzic/simpledb/blob/main/file/file.go

package entestfile

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var ErrBlockOutOfBound = fmt.Errorf("block number greater than file size")

type FileMgr struct { 
  BlockSize int 
  dataDir string 
  openedFiles map[string]*os.File 

  mu sync.Mutex

}

func NewFileMgr(dataDir string, blockSize int) *FileMgr {
  return &FileMgr{
    dataDir: dataDir,
    BlockSize: blockSize,
    openedFiles: make(map[string]*os.File),
  }
}

// Read reads the contents of specified block into provided page
func (fm *FileMgr) Read(blockID BlockID, p *Page) (int, error) {
  fm.mu.Lock()
  defer fm.mu.Unlock()

  f, error :=fm.getFile(blockID.Filename) 

  if error != nil {
    return 0, fmt.Errorf("Failed to get file; %w", error)
  }

  size, error := fm.FileSize(blockID.Filename)

  if error != nil {
    return 0, error
  }

  if blockID.Number >= size {
    return 0, ErrBlockOutOfBound
  }

  n, error := f.ReadAt(p.Bytes(), int64(blockID.Number * fm.BlockSize))

  if error != nil {
    return 0, fmt.Errorf("Failed to read block; %w", error)
  }

  return n, nil 
  
}

// Write writes the contents of the provided page into specified block
func (fm *FileMgr) Write(blockID BlockID, p *Page) (int, error) {
  fm.mu.Lock()
  defer fm.mu.Unlock()

  f, error := fm.getFile(blockID.Filename)

  if error != nil {
    return 0, error
  }

  // size, error := fm.FileSize(blockID.Filename)
  // if error != nil {
  //   return 0, error
  // }

  // if blockID.Number > size {
  //   return 0, ErrBlockOutOfBound
  // }

  n, error := f.WriteAt(p.Bytes(), int64(blockID.Number * fm.BlockSize))

  if error != nil {
    return 0, fmt.Errorf("Failed to write block; %w", error)
  }

  return n, nil
}

// Close closes all opened files
func (fm *FileMgr) Close() error {
  fm.mu.Lock()
  defer fm.mu.Unlock()

  for _, f := range fm.openedFiles {
    if err := f.Close(); err != nil {
      return fmt.Errorf("Failed to close file; %w", err)
    }
  }

  return nil
}

// fileLength returns the number of blocks in specified file
func (fm *FileMgr) FileSize(fileName string) (int, error) {
  file, error := fm.getFile(fileName)

  fileInfo, error := file.Stat()  

  if error != nil {
    return 0, fmt.Errorf("Failed to get file info; %w", error)
  }

  return int(fileInfo.Size() / int64(fm.BlockSize)), nil
} 

// getFile returns the file with specified file name, creating if it does not exit 
func (fm *FileMgr) getFile(fileName string) (*os.File, error) {
  if f, ok := fm.openedFiles[fileName]; ok {
    return f, nil
  }

  f, err := os.OpenFile(filepath.Join(fm.dataDir, fileName), os.O_RDWR|os.O_CREATE, 0666)
  if err != nil {
    return nil, err
  }

  fm.openedFiles[fileName] = f
  return f, nil
}