package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("File Handling!")
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close() 
	// info, finfoErr := f.Stat()
	// if finfoErr != nil {
	// 	panic(finfoErr)
	// }
	// fmt.Printf("File Name: %s\nSize: %d bytes\nIsDir: %v\n", info.Name(), info.Size(), info.IsDir())


	// buffer:=make([]byte, 1024)
	// data,err :=f.Read(buffer)
	// if err!=nil{
	// 	panic(err)
	// }
	// fmt.Println(string(buffer))
	// fmt.Println(data)


	// read Folders:
	// dir, err:=os.Open("../")

	// if err!=nil{
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err:=dir.ReadDir(-1)
	// if err!=nil{
	// 	panic(err)
	// }
	// for _, fi:= range fileInfo{
	// 	fmt.Println(fi.Name())
	// }


	// f, err:=os.Create("example2.txt")
	// if err!=nil{
	// 	panic(err)
	// }
	// defer f.Close()
	// f.WriteString("Hello Go from JS")
	// f.WriteString(" Hello Go from JS")
	// fmt.Println("File Created Successfully: ",f.Name())


	// delete the file:
	err := os.Remove("example2.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist, nothing to delete.")
		} else {
			panic(err)
		}
	} else {
		fmt.Println("File deleted Successfully!")
	}
}