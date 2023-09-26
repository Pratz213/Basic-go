package main

import (
	"fmt"
	"net/url"
)

func main() {
	myurl := "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

	rawUrl, _ := url.Parse(myurl)

	fmt.Println(rawUrl.Host+rawUrl.RequestURI(), rawUrl.User)

}
