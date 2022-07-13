// Hi, I am Misitebao.
//
// Usage:
//
// Install the package using the `cargo install` command:
//
// 		go install github.com/misitebao/misitebao/cmd/misitebao@latest
//
// Once that's done, you can run this command inside the terminal...
//
// 		misitebao
package main

import (
	"fmt"

	"github.com/pkg/browser"
)

func main() {
	var path = "https://misitebao.com"

	err := browser.OpenURL(path)
	if err != nil {
		panic(fmt.Sprintf("An error occurred when opening '%s': %s", path, err))
	}

	fmt.Printf("Opened '%s' successfully.", path)
}
