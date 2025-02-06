// hai 06/02/2025
// go test -run TestBlock
// go test -rn TestPage
// test block

package entestfile

import (
  "fmt"
  "testing"
)

func TestBlock(t *testing.T){
  fmt.Println("test block")

  b := BlockID{"./data_base", 16}

  fmt.Printf("filename: %s, block number: %d\n", b.Filename, b.Number)
  
}


func TestPage(t *testing.T){
  fmt.Println("test page")

  // create a new page 
  p := NewPage(8)
  fmt.Printf("page size: %d\n", p.Size())
  fmt.Printf("page data: %v\n", p.Bytes())

  // insert data in page 
  data := []byte("hai")
  p.Write(0, data)
  fmt.Printf("%v\n", p.Bytes())
  fmt.Println(string(p.Bytes()[0:4]))

  // read data from page 
  buffer := make([]byte, 4)
  p.Read(0, buffer)
  fmt.Println(buffer)
  
}


func TestFileManager(t *testing.T){
  fmt.Println("test file")
  // create a file manager
  fm := NewFileMgr("./", 8)
  fmt.Println(fm.dataDir, fm.BlockSize)

  // getFile and create if it does not exit 
  f, error := fm.getFile("database")

  if error != nil {
    fmt.Println(error)
  }

  // get file information
  fileInfo, error := f.Stat()
  if error != nil {
    fmt.Println(error)
  }
  fmt.Println(fileInfo.Size(), fileInfo.Name(), fileInfo.IsDir())

  // write page to the file into block 
  page := NewPage(8)
  page.Write(0, []byte("hai"))
  n, error := fm.Write(BlockID{"database", 0}, page)

  if error != nil {
    fmt.Println("failed to write file", error)
  }

  fmt.Println(n)
  // read data from file to page 
  readPage := NewPage(8)
  fm.Read(BlockID{"database", 0}, readPage)
  fmt.Println(readPage.Bytes())
  fmt.Println(string(readPage.Bytes()[0:4]))

}