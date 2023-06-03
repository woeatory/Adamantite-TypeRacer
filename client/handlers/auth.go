package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/woeatory/Adamantite-TypeRacer/client/http_client"
	"github.com/woeatory/Adamantite-TypeRacer/server/server"
	"io/ioutil"
	"net/http"
	"strings"
)

// todo there is a lot of copy-past. find better solution
var usernamePrompt = promptui.Prompt{
	Label:       "username",
	Default:     "",
	AllowEdit:   true,
	Validate:    usernameValidator,
	Mask:        0,
	HideEntered: true,
	Templates:   nil,
	IsConfirm:   false,
	IsVimMode:   false,
	Pointer:     nil,
	Stdin:       nil,
	Stdout:      nil,
}

var passwordPrompt = promptui.Prompt{
	Label:       "password",
	Default:     "",
	AllowEdit:   true,
	Validate:    passwordValidator,
	Mask:        '*',
	HideEntered: true,
	Templates:   nil,
	IsConfirm:   false,
	IsVimMode:   false,
	Pointer:     nil,
	Stdin:       nil,
	Stdout:      nil,
}

func LogIn() error {
	username, err := usernamePrompt.Run()
	if err != nil {
		return err
	}
	password, err := passwordPrompt.Run()
	if err != nil {
		return err
	}
	payload := []byte(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password))
	var PATH = "http://" + server.ADDRESS + "/" + server.AuthGroupPath + server.AuthLogin
	req, err := http.NewRequest(http.MethodPost, PATH, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http_client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func SignUp() error {
	username, err := usernamePrompt.Run()
	if err != nil {
		return err
	}

	password, err := passwordPrompt.Run()
	if err != nil {
		return err
	}
	payload := []byte(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password))
	var PATH = "http://" + server.ADDRESS + "/" + server.AuthGroupPath + server.AuthSignUp
	req, err := http.NewRequest(http.MethodPost, PATH, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http_client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func usernameValidator(s string) error {
	if len(s) < 6 || len(s) > 16 || strings.Contains(s, " ") {
		return errors.New("bad username")
	}
	return nil
}

func passwordValidator(s string) error {
	if len(s) < 6 || len(s) > 16 || strings.Contains(s, " ") {
		return errors.New("bad password")
	}
	return nil
}
