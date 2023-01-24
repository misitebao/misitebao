// Hi, I am Misite Bao.
//
// Usage:
//
// Install the package using the `go install` command:
//
//	go install github.com/misitebao/misitebao/cmd/misitebao@latest
//
// Once that's done, you can run this command inside the terminal...
//
//	misitebao
package main

import (
	"fmt"

	"github.com/pkg/browser"
)

const URL = "https://misitebao.com"

func main() {

	err := browser.OpenURL(URL)

	if err != nil {
		panic(fmt.Sprintf("An error occurred when opening '%s': %s", URL, err))
	}

	fmt.Printf("Opened '%s' successfully.", URL)
}
