package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"reflect"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

type owner struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}
type item struct {
	ID          int    `json:"id"`
	Owner       *owner `json:"owner"`
	Description string `json:"description"`
}

type all struct {
	TotalCount int    `json:"total_count"`
	Items      []item `json:"items"`
}

func main() {
	//var qs = "https://api.github.com/search/repositories?q=porn"
	//var allDate all
	//resp,err := http.Get(qs)
	//if err != nil{
	//}
	//defer resp.Body.Close()
	//json.NewDecoder(resp.Body).Decode(&allDate)
	//fmt.Println(allDate)

	var buf bytes.Buffer
	command := exec.Command("/bin/bash", "-c", "git")
	command.Stdout = &buf
	command.Run()
	fmt.Println(buf.String())
	s, err := fmt.Fprint(command.Stdout)

	if err != nil {

	}
	fmt.Println(s)

}

func testStruct() {
	s, err := json.Marshal(movies)
	s2, err := json.MarshalIndent(movies, "|", "    ")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	fmt.Printf("%T,%[1]v\n", reflect.TypeOf(movies))
	fmt.Printf("%T,%[1]v\n", reflect.ValueOf(movies))
	fmt.Printf("%T,%T", s, s2)
	//fmt.Printf("%s,%T",s2,s2)
}
