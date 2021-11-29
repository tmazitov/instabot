package post

import (
	"fmt"
	"time"

	"github.com/tmazitov/instabot/active"
	selenium "sourcegraph.com/sourcegraph/go-selenium"
)

// uploadPhoto upload photo to post form
func uploadPhoto(webDriver selenium.WebDriver, photoPath string) error {

	// https://stackoverflow.com/questions/52864531/posting-uploading-an-image-to-instagram-using-selenium-not-using-an-api

	_, err := webDriver.ExecuteScript("HTMLInputElement.prototype.click = function() { "+
		"  if(this.type !== 'file') HTMLElement.prototype.click.call(this);  "+
		"};                                                                  ",
		nil)
	if err != nil {
		return err
	}

	elems, err := webDriver.FindElements(selenium.ByTagName, "input")
	if err != nil {
		fmt.Println("input")
		return err
	}

	elems[len(elems)-1].SendKeys(photoPath)
	_, err = webDriver.ExecuteScript("delete HTMLInputElement.prototype.click", nil)
	if err != nil {
		fmt.Println("execute2")
		return err
	}
	return nil
}

// CreateNewPost create new post
func CreateNewPost(webDriver selenium.WebDriver, gap time.Duration) error {

	// SELECTORS
	const (
		newPostSelector    = ".vZuFV"
		photoFieldSelector = ".tb_sK"
		nextDivSelector    = ".qF0y9.Igw0E.IwRSH.eGOV_._4EzTm.XfCBB.g6RW6"
		nextSelector       = ".sqdOP.yWX7d.y3zKF"
		textSelector       = ".lFzco"
	)

	// Example Data
	photoPath := "C:/Users/noobo/GoProjects/instabot/example.png"
	text := "Groku is so cute!"

	// Clict to "New Post"
	if err := active.ClickToButton("css", newPostSelector, webDriver); err != nil {
		return err
	}

	time.Sleep(gap)

	uploadPhoto(webDriver, photoPath)

	time.Sleep(gap * 2)

	// Select photo
	if err := active.ClickToButton("css", nextDivSelector, webDriver); err != nil {
		fmt.Println("photo")
		return err
	}

	time.Sleep(gap)

	// Select filter
	if err := active.ClickToButton("css", nextDivSelector, webDriver); err != nil {
		fmt.Println("filter")
		return err
	}
	time.Sleep(gap)

	// Upload simple text
	if err := active.SendDataToField("css", textSelector, text, webDriver); err != nil {
		return err
	}

	time.Sleep(gap)

	// Upload post
	if err := active.ClickToButton("css", nextDivSelector, webDriver); err != nil {
		return err
	}

	return nil
}
