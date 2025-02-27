package filecopy

import (
	"fmt"
	"os"
	"io"
)

func Copyfile(){
	sourcefile := "D:\\go-lang\\sourcefile.txt"
	destinationfile := "D:\\go-lang\\destination.txt"

	src, err := os.Open(sourcefile)
	if err != nil {
		fmt.Println("Source file having some issue", err)
		return
	}
	defer src.Close()

	// dest, err := os.Create(destinationfile)
	dest, err := os.OpenFile(destinationfile, os.O_APPEND| os.O_CREATE |os.O_WRONLY , 0664)
	if err != nil {
		fmt.Println("destination file having some issue",err)
		return
	}
	defer dest.Close()

	_ , err = io.Copy(dest,src)
	if err !=nil {
		fmt.Println("Copying is unsuccessful",err)
		return
	}
	fmt.Println("file copied successfylly!")
}