package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handleRequests()  {
	http.HandleFunc("/",ServeHTTP)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := os.Getenv("EXTERNAL_SERVICE_KEY")
	url := "http://newsapi.org/v2/everything?q=covid19&language=en&sortBy=publishedAt&apiKey="
	log.Println("Fetch from URL:", url)
	response, err := http.Get(url+key)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)
}

func main()  {
	log.Print("Server is started!!!")
	handleRequests()
}