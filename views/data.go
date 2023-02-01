package views

import (
	"fmt"
	"github.com/Ayoub-Moulahi/MyYouTube/models"
)

// Alert  hold an alert type ie:success,error ...
// and message to display to the user
type Alert struct {
	AlertType string
	Msg       string
}

// Data define the data type accepted by the Render method
// it has 2 fields to distinguish where the data should be executed
// an Alert field for the alert template and a Content field for
// content template
type Data struct {
	Alert   *Alert
	Content interface{}
}

type PublicError interface {
	error
	ToDisplay() string
}

const (
	AlertError   = "alert-error"
	AlertWarning = "alert-warning"
	AlertInfo    = "alert-info"
	AlertSuccess = "alert-success"
)

// CreateAlert create an error Alert to be displayed to the end user
func CreateAlert(err error) Data {
	var dt Data
	if pubErr, ok := err.(PublicError); ok {
		dt.Alert = &Alert{
			AlertType: AlertError,
			Msg:       pubErr.Error(),
		}

	} else {
		fmt.Println(ok)
		dt.Alert = &Alert{
			AlertType: AlertError,
			Msg:       models.ErrApp.Error(),
		}
	}
	return dt
}

// used to create other Alert type such as success,
func CreateOtherAlert(alertType string, msg string) Data {
	var dt Data
	dt.Alert = &Alert{
		AlertType: alertType,
		Msg:       msg,
	}
	return dt
}
