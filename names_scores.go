// Names scores (Problem 22)

package main

import (
   "fmt"
   "net/http"
   "strings"
   "strconv"
   "sort"

)

const name_list_url = "https://projecteuler.net/project/resources/p022_names.txt"

func main() {
   name_list_str :=  getUrlBody(name_list_url)
   name_list_str = strings.Replace(name_list_str, "\"", "", -1)
   name_list := strings.Split(name_list_str, ",")
   fmt.Println("Num names: " + strconv.Itoa(len(name_list)))
   name_score := name_score_brute_force(name_list)
   fmt.Println("Name score is ", name_score)

}

func name_score_brute_force(name_list []string) (name_score int) {
   sort.Strings(name_list)   
   for i,v := range name_list {
      score := (i+1) * charValueSum(v)
      name_score += score
      fmt.Println(i+1, v, score)
   }
   return
}

func name_score_trie(name_list []string) (name_score int) {
   return
}

func charValueSum(str string) (sum int) {
   str = strings.ToLower(str)
   for _, r := range str {
      sum += int(r) - 96
   }
   return
}

// todo: move to another package
func getUrlBody(url string) (body string) {
   resp, err := http.Get(url)
   if err != nil {
      fmt.Println("Failed to open url " + url)
   }
   defer resp.Body.Close()

   buffer := make([]byte, 1000)
   for {
      bytes_read, read_err := resp.Body.Read(buffer)
      body += string(buffer[:bytes_read])
      if read_err != nil {
         break
      }
   }
   return
}

