package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://googlr.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "http://linkedin.com"

	delete(websites, "Google")
}
