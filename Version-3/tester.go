// package main

// import (
//   "fmt"
//   "strings"
//   "net/http"
//   "io/ioutil"
// )

// func main() {

//   url := "localhost:8080/videos"
//   method := "POST"

//   payload := strings.NewReader(`{
//     "title":"learning video",
//     "description": "Go is one of the fastest programming languages out there",
//     "url": "https://www.youtube.com/watch?v=qR0WnWL2o1Q&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w"
// }`)

//   client := &http.Client {
//   }
//   req, err := http.NewRequest(method, url, payload)

//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   req.Header.Add("Authorization", "Basic YWpheTpyYWo=")
//   req.Header.Add("Content-Type", "application/json")

//   res, err := client.Do(req)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   defer res.Body.Close()

//   body, err := ioutil.ReadAll(res.Body)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   fmt.Println(string(body))
// }