package img2dcm

import (
	"fmt"
	"log"
	"os"
)

func Abc() {
	fmt.Println("abc")

	f, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		log.Println("[Error] : ", err.Error())
	}

	// l, err := f.Write([]byte("abc"))
	// if err != nil {
	// 	log.Println("[Error] : ", err.Error())
	// }

	// fmt.Println(l)

	// f.Sync()
}

func Def() {
	Abc()
}
