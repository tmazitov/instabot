package auth

import (
	"errors"
	"os"
	"time"

	"github.com/tmazitov/instabot/active"
	selenium "sourcegraph.com/sourcegraph/go-selenium"
)

// SetAuth auth user
func SetAuth(webDriver selenium.WebDriver, gap time.Duration) error {

	// Get username
	username, exists := os.LookupEnv("LOGIN")
	if !exists {
		return errors.New("failed to get user login")
	}

	// Get password
	password, exists := os.LookupEnv("PASSWORD")
	if !exists {
		return errors.New("failed to get user password")
	}

	const (
		passwordSelector = "password"
		usernameSelector = "username"
		doAuth           = ".qF0y9.Igw0E.IwRSH.eGOV_._4EzTm"
		noSaveAccount    = ".sqdOP.yWX7d.y3zKF"
		noNotifications  = ".aOOlW.HoLwm "
	)

	// Set username to form
	if err := active.SendDataToField("name", usernameSelector, username, webDriver); err != nil {
		return err
	}

	time.Sleep(gap)

	// Set password to form
	if err := active.SendDataToField("name", passwordSelector, password, webDriver); err != nil {
		return err
	}

	time.Sleep(gap)

	// *Click* to auth
	if err := active.ClickToButton("css", doAuth, webDriver); err != nil {
		return err
	}

	time.Sleep(gap)

	// Not save account
	if err := active.ClickToButton("css", noSaveAccount, webDriver); err != nil {
		return err
	}

	time.Sleep(gap)

	// Not send notifications
	if err := active.ClickToButton("css", noNotifications, webDriver); err != nil {
		return err
	}
	return nil
}
