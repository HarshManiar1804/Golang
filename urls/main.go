package main

import (
	"fmt"
	"net/url"
)

func main(){

	myurl := "https://example.com:8080/path/to/resourse?key1=value1&key2=value2"

	fmt.Printf("Type of the url is %T\n",myurl)

	parsedUrl,err:=url.Parse(myurl)
	if err != nil{
		fmt.Println("Error while converting string into url",err)
		return
	}
	fmt.Printf("Type of the parsedURL is %T\n",parsedUrl)

	fmt.Println("Schema:-",parsedUrl.Scheme)
	fmt.Println("Port:-",parsedUrl.Port())
	fmt.Println("Path:-",parsedUrl.Path)
	fmt.Println("Host:-",parsedUrl.Host)
	fmt.Println("Rawquery:-",parsedUrl.RawQuery)

	//update the url

	parsedUrl.Path="/newpath"
	parsedUrl.RawQuery="username=harsh"

	// constructing a URL string from a URL path
	newURL := parsedUrl.String()
	fmt.Println("new url:-", newURL)

}