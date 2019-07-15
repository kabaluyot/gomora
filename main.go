/*
	Googol API Framework
	A minimalist web API framework using the service-go-pattern

	@Author: Karl Anthony B. Baluyot
*/

package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
)

//* entry point of the googol framework
func main() {

	//get the API_URL_PORT
	port := os.Getenv("API_URL_PORT")
	if len(port) == 0 {
		port = "8000" //! default port 8000
	}

	//listen for requests and serve
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), ChiRouter().InitRouter())
	if err != nil {
		log.Println("Listen and Serve:", err)
	}
}
