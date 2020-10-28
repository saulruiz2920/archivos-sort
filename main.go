package main

import "fmt"
import "sort"
import "os"


func writeFile(filename string, slice []string) {
	file, err := os.Create(filename + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < len(slice); i++ {
		file.WriteString(slice[i] + "\n")
	} 
}

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	var slice []string
	for i := 0; i < n ; i++ {
	  var input string
	  fmt.Scan(&input)
	  slice = append(slice, input)
	}
	sort.Strings(slice)
	writeFile("asecendente", slice)
	sort.Sort(sort.Reverse(sort.StringSlice(slice)))
	writeFile("descendente", slice)
}