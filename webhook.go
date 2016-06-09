
package main

import (

"fmt"

"net/http"

)

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




func main(){
  http.HandleFunc("/hi",handler)
  http.HandleFunc("/webhook",SetWebhook)
  http.ListenAndServe(":8000",nil)
}
