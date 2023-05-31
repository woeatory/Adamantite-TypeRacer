package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/woeatory/Adamantite-TypeRacer/cmd/server"
	"io"
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

func LogIn() {
	username, err := usernamePrompt.Run()
	if err != nil {
		return
	}

	password, err := passwordPrompt.Run()
	if err != nil {
		return
	}
	payload := []byte(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password))
	const PATH = "http://" + server.ADDRESS + "/" + server.AuthGroupPath + server.AuthLogin

	resp, err := http.Post(PATH, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(body))
}

func SignUp() {
	username, err := usernamePrompt.Run()
	if err != nil {
		return
	}

	password, err := passwordPrompt.Run()
	if err != nil {
		return
	}
	payload := []byte(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password))
	const PATH = "http://" + server.ADDRESS + "/" + server.AuthGroupPath + server.AuthSignUp

	resp, err := http.Post(PATH, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(body))
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
