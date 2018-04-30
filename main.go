//chapter new
package main

import "fmt"

func main()  {
	sum:= 1
	for ; sum<6;{
		sum += sum
	}
	fmt.Println(sum)

}
