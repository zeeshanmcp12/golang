package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Find empty files!")

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Provide directory name!")
		return
	}

	files, err := ioutil.ReadDir(args[0])

	checkNilErr(err)

	newFile, err := os.Create("./filenames.txt")

	for _, file := range files {
		fmt.Println(file.Name())
		name := file.Name()

		io.WriteString(newFile, name)

	}

	defer newFile.Close()

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
