package active

import (
	"errors"
	"fmt"

	selenium "sourcegraph.com/sourcegraph/go-selenium"
)

func ClickToButton(selectorType string, selectorValue string, webDriver selenium.WebDriver) error {

	var elem selenium.WebElement
	var err error

	switch selectorType {
	case "css":
		elem, err = webDriver.FindElement(selenium.ByCSSSelector, selectorValue)
	case "name":
		elem, err = webDriver.FindElement(selenium.ByName, selectorValue)
	default:
		return errors.New("this type of selectors is not supported")
	}

	if err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
		return err
	}

	err = elem.Click()
	if err != nil {
		fmt.Printf("Failed to click: %s\n", err)
		return err
	}

	return err
}

func SendDataToField(selectorType string, selectorValue string, data string, webDriver selenium.WebDriver) error {

	var elem selenium.WebElement
	var err error

	switch selectorType {
	case "css":
		elem, err = webDriver.FindElement(selenium.ByCSSSelector, selectorValue)
	case "name":
		elem, err = webDriver.FindElement(selenium.ByName, selectorValue)
	default:
		return errors.New("this type of selectors is not supported")
	}

	if err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
		return err
	}

	err = elem.SendKeys(data)

	if err != nil {
		fmt.Printf("Failed to send data: %s\n", err)
		return err
	}

	return nil

}
