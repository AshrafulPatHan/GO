package main
import (
   "fmt"
   "io/ioutil"
   "strings"
)

func main() {
   data, err := ioutil.ReadFile("file1.txt")  //read the file
   if err != nil {
      fmt.Println("Error reading file:", err) //print the error if occurs while reading file
      return
   }
   read_lines := strings.Split(string(data), "\n")
   for _, line := range read_lines {
      fmt.Println(line) //print the content line by line
   }
}
