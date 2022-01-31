package main

import "fmt"

func main() {
	// Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
	var website = "\thttps://www.Google.com\t\n"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var siteDescription = `\t\tGoole is a search engine which is widely used by millions of users around the World.\t\n`

	fmt.Println(website, siteDescription)
}
