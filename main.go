package main

import (
  "encoding/json"
  "sync/atomic"
  "log"
  "time"
  "math/rand"
  "net/http"
)

func main() {

  isReady := &atomic.Value{}
  isReady.Store(true)

  http.HandleFunc("/", home())
  http.HandleFunc("/healthz", healthz)
  http.HandleFunc("/readyz", readyz(isReady))

  http.ListenAndServe(":3000", nil)
}

type Rand struct {
  Title string `json:"title"`
  Description string `json:"description"`
  Date time.Time `json:"date"`
  Version string `json:"version"`
}

var wordsList = []string{
  "ipsum", "semper", "habeo", "duo", "ut", "vis", "aliquyam", "eu", "splendide", "Ut", "mei", "eteu", "nec", "antiopam", "corpora", "kasd", "pretium", "cetero", "qui", "arcu", "assentior", "ei", "his", "usu", "invidunt", "kasd", "justo", "ne", "eleifend", "per", "ut", "eam", "graeci", "tincidunt", "impedit", "temporibus", "duo", "et", "facilisis", "insolens", "consequat", "cursus", "partiendo", "ullamcorper", "Vulputate", "facilisi", "donec", "aliquam", "labore", "inimicus", "voluptua", "penatibus", "sea", "vel", "amet", "his", "ius", "audire", "in", "mea", "repudiandae", "nullam", "sed", "assentior", "takimata", "eos", "at", "odio", "consequat", "iusto", "imperdiet", "dicunt", "abhorreant", "adipisci", "officiis", "rhoncus", "leo", "dicta", "vitae", "clita", "elementum", "mauris", "definiebas", "uonsetetur", "te", "inimicus", "nec", "mus", "usu", "duo", "aenean", "corrumpit", "aliquyam", "est", "eum",
}

func getRandomWord() string {
  return wordsList[rand.Intn(len(wordsList))]
}

func randomWords(length int) string {
  result := "Lorem "
  for i := 0; i < length-1; i++ {
    result += getRandomWord() + " "
  }
  return result

}

func home() http.HandlerFunc {
  return func(w http.ResponseWriter, _ *http.Request) {

    var title = randomWords(10)
    var desc = randomWords(rand.Intn(100))
    rand := Rand{
      title,
      desc, 
      time.Now(),
      "0.0.2",
    }

    body, err := json.Marshal(rand)
    if err != nil {
      log.Printf("Could not encode rand data: %v", err)
      http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(body)
  }
}

func healthz(w http.ResponseWriter, _ *http.Request) {
  w.WriteHeader(http.StatusOK)
}

func readyz(isReady *atomic.Value) http.HandlerFunc {
  return func(w http.ResponseWriter, _ *http.Request) {
    if isReady == nil || !isReady.Load().(bool) {
      http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
      return
    }
    w.WriteHeader(http.StatusOK)
  }
}