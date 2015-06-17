package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//  url := "http://" + os.Getenv("ISORENDER_PORT_3000_TCP_ADDR") + ":" + os.Getenv("ISORENDER_PORT_3000_TCP_PORT") + "/render"
	url := "http://127.0.0.1:3000/render"
	var query = []byte(`{"name":"mike"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(body))
}

func main() {

	//    fmt.Println("response Status:", resp.Status)
	//    fmt.Println("response Headers:", resp.Header)
	//    fmt.Println("response Body:", string(body))

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)

}
