package main

import (
	"fmt"
	"os"
)

func main() {
	// (((( 0 ))))
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("filename:", fileInfo.Name())
	// fmt.Println("file or folder:", fileInfo.IsDir())
	// fmt.Println("size:", fileInfo.Size())
	// fmt.Println("permission:", fileInfo.Mode())
	// fmt.Println("modified at :", fileInfo.ModTime())

	// defer f.Close() //will close if any err
	// buffer := make([]byte, fileInfo.Size())

	// d, err := f.Read(buffer)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buffer); i++ {
	// 	println("data", d, string(buffer[i]))
	// }

	// (((( 1 ))))
	// not ideal if file size is bigger

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// (((( 2 ))))
	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// (((( 3 ))))
	// f, err := os.Create("2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// f.WriteString("hi feyhide go here \n")
	// f.WriteString("hi feyhide go here2 \n")

	// bytes := []byte("Hello feyhide")
	// f.Write(bytes)

	// (((( 4 ))))
	// DELETE A FILE
	err := os.Remove("2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File Deleted")
}
