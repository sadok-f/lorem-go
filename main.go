package main

import (
  "encoding/json"
  "net/http"
)

type Profile struct {
  Name    string
  SocialProfiles  []string
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Sadok", []string{"https://github.com/sadok-f", "https://twitter.com/sadok_f"}}

  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
