package main
import (
   "fmt"
   "os"
)
func main() {
   file, err := os.Open("filename.txt")

   if err != nil {
      panic(err)
   }
   defer file.Close()

   var word string

   for {
      _, err := fmt.Fscanf(file, "%s", &word)
      if err != nil {
         break
      }
      fmt.Println(word)
   }
}
