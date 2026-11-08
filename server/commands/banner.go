package commands

import (
	"fmt"

	"github.com/fatih/color"
)

func printBanner() {
	const banner = `
╔══════════════════════════════════════╗
║                                      ║
║    TAP — A shared-world adventure    ║
║                                      ║
╚══════════════════════════════════════╝
`
	colorBanner := color.New(color.FgGreen)
	fmt.Print(colorBanner.Sprint(banner) + "\n")
}
