package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func writeFile() string {
  content := "Insert this into file"
	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err) // shut down execution
	}

	len, err :=  io.WriteString(file, content)
  if(err != nil) {
    panic(err)
  }

  fmt.Println("len of str: ", len)
  defer file.Close()

  return "./file.txt"
}

func readFile(path string) {
  data, err := ioutil.ReadFile(path)
  checkNilErr(err)

  // reading format is not string but byte
  // => Convert to string
  fmt.Println("Data: ", string(data))
}

func main() {
	fmt.Println("Handling Files in Go")

	path := writeFile()
  readFile(path)

}


// Separate fun for error
func checkNilErr(err error) {
  if(err != nil) {
    panic(err)
  }
}