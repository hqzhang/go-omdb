package main

import (
        "io/ioutil"
        "fmt"
        "log"
        "net/http"
        "encoding/json"
        "net/url"
        "os"
        "strings"
)
type Respbody struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}
var serverID string

func main() {
        
        query := os.Args[1]        
        //query := "Long Shot"
        mystr:=url.QueryEscape(query)
	fmt.Println(mystr)

        resp, err := http.Get("http://www.omdbapi.com/?apikey=d8528afb&t="+mystr)
        if err != nil {
            log.Fatalln(err)
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
           log.Fatalln(err)
        }
        //log.Println(string(body))
        
        var myresp *Respbody
        err4 := json.Unmarshal(body, &myresp)
        if err4 != nil {
                panic(err4)
        }
        for _, tmp := range myresp.Ratings {
          if strings.Contains(tmp.Source, "Rotten Tomatoes") {
               log.Println(tmp.Source,"rating is:",tmp.Value)
               break
          }
       }
}

func Middle(l *log.Logger, f http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
                l.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
                f(w, r)
        }
}

func Root(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%s: you are great!\n", serverID)
}
