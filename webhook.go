
package main

import (
//"bytes"
//"encoding/json"
//"errors"
"fmt"
//"io"
//"io/ioutil"
//"log"
//"mime/multipart"
"net/http"
//"os"
//"path/filepath"
)
// This defines a bot
// Set Debug to true for debugging
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Say Hi")
}
func SetWebhook(w http.ResponseWriter, req *http.Request) {
switch req.Method {
case "GET":
if req.FormValue("hub.verify_token") == "Our_Bot" {
w.Write([]byte(req.FormValue("hub.challenge")))
return
}
w.WriteHeader(http.StatusUnauthorized)
return

default:
w.WriteHeader(http.StatusMethodNotAllowed)
return
}
}

//return callbackChan, mux


func main(){
  http.HandleFunc("/hi",handler)
  http.HandleFunc("/webhook",SetWebhook)
  http.ListenAndServe(":8000",nil)
}
