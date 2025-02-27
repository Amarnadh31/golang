//////////////////////////////LOOPS USING FUNCTIONS////////////////////
package loops

import (
	"fmt"
)

func LoopsFunc( ){

	var start ,end int
	m := "Left"
	n := "Right"

	for {
		fmt.Println("Enter the start number: ")
		_, err1 := fmt.Scan(&start)
		if err1 != nil {
			fmt.Println("Please enter vaild input(0-9)")

			var discord string
			fmt.Scanln(&discord)
			continue	
		} 
		break
	}
	
	for{
		fmt.Println("Enter the end number: ")
		_,err2 := fmt.Scan(&end)
		if err2 != nil {
			fmt.Println("Please enter vaild input(0-9)")

			var discord string
			fmt.Scanln(&discord)
			continue
		}
		break
	}

	for i:= start; i< end; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(m+n)
		}else if i%5 == 0 {
			fmt.Println(n)
		}else if i%3 == 0{
			fmt.Println(m)
		}else {
			fmt.Println(i)
		}

	}
}