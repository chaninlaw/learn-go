package control

import (
	"errors"
	"fmt"
)

type Error struct {
	Msg string
}

type LoginError struct {
	Username string
	Msg      string
}

func (l *LoginError) Error() string {
	return fmt.Sprintf("Login error for user '%s': %s", l.Username, l.Msg)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func login(username, password string) error {
	if username != "admin" || password != "1234" {
		return &LoginError{Username: username, Msg: "Invalid username or password"}
	}
	return nil
}

func ErrorHandling() {
	fmt.Println("Handling error in Go!")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		// return
	}
	fmt.Println("Result:", result)

	// ------------------------------

	error := login("user", "pass")

	if error != nil {
		switch e := error.(type) {
		case *LoginError:
			fmt.Println("Custom error occurred:", e.Msg)
		default:
			fmt.Println("Unknown error occurred:", e)
		}
		return
	}
}
