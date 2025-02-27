package filecopy

import (
	"fmt"
	"os"
	"io"
)

func Copyfile(){
	sourcefile := "D:\\go-lang\\filecopy\\sourcefile.txt" // source file path
	destinationfile := "D:\\go-lang\\filecopy\\destination.txt" //destination file path

	src, err := os.Open(sourcefile) //using os.Open() from os package to open the source file 
	if err != nil {
		fmt.Println("Source file having some issue", err)
		return
	}
	defer src.Close() //using defer to close the file properly even if any error occurs defer closes file properly

	// dest, err := os.Create(destinationfile) //os.create used to create the file if destination file is not created and whenever you copy it overwrites the data
	dest, err := os.OpenFile(destinationfile, os.O_APPEND| os.O_CREATE |os.O_WRONLY , 0664) //os.OpenFile is used to copy the content on already existed file
	if err != nil {
		fmt.Println("destination file having some issue",err)
		return
	}
	defer dest.Close()

	_ , err = io.Copy(dest,src) //io.Copy used to copy the content from one file to another file.(dest,src)--> copying data from src to dest,(src,dest) --> copying data from dest to src
	if err !=nil {
		fmt.Println("Copying is unsuccessful",err)
		return
	}
	fmt.Println("file copied successfylly!")
}