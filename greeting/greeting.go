package greeting

import "fmt"

// Message returns the application's greeting.
func Message() string {
	return "Hello from Go in Docker! 🐹🐳"
}

// Details formats runtime information for the command-line application.
func Details(goos, arch string) string {
	return fmt.Sprintf("OS: %s\nArch: %s", goos, arch)
}
