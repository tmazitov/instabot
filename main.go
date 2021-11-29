package main

import (
	"fmt"

	selenium "sourcegraph.com/sourcegraph/go-selenium"
)

func main() {
	ExampleFindElement()
}

func ExampleFindElement() {
	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	if webDriver, err = selenium.NewRemote(caps, "http://localhost:9515"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}

	err = webDriver.Get("https://www.instagram.com/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	// if title, err := webDriver.Title(); err == nil {
	// 	fmt.Printf("Page title: %s\n", title)
	// } else {
	// 	fmt.Printf("Failed to get page title: %s", err)
	// 	return
	// }

	// time.Sleep(time.Second * 3)

	// var elem selenium.WebElement
	// elem, err = webDriver.FindElement(selenium.ByCSSSelector, "._2hvTZ.pexuQ.zyHYP")
	// if err != nil {
	// 	fmt.Printf("Failed to find element: %s\n", err)
	// 	return
	// }

	// elem.SendKeys()

	// time.Sleep(time.Second * 3)

	// if text, err := elem.Text(); err == nil {
	// 	fmt.Printf("TEXT: %s\n", text)
	// } else {
	// 	fmt.Printf("Failed to get text of element: %s\n", err)
	// 	return
	// }

	// output:
	// Page title: go-selenium - Sourcegraph
	// Repository: go-selenium
}
