package main

import (
	"fmt"
	"github.com/mabetle/mcore"
)

func demoJoinPath() {
	fmt.Printf("%s\n", mcore.JoinPath("c.txt"))                     // a/b/c.txt
	fmt.Printf("%s\n", mcore.JoinPath("a", "b", "c.txt"))           // a/b/c.txt
	fmt.Printf("%s\n", mcore.JoinPath("$HOME", "b", "c.txt"))       // a/b/c.txt
	fmt.Printf("%s\n", mcore.JoinPath("$HOME", "$GOPATH", "c.txt")) // a/b/c.txt
}

func demoExpand() {
	fmt.Printf("Home: %s\n", mcore.Home())
	fmt.Printf(".vim: %s\n", mcore.ExpandPath("$HOME/.vim"))
	fmt.Printf(".vim: %s\n", mcore.ExpandPath("~/.vim"))
	fmt.Printf("MISC: %s\n", mcore.ExpandPath("$HOME$GOPATH"))
	readme := mcore.ExpandPath("$GOROOT/README.md")
	fmt.Printf("README: %s\n", readme)
	//cat GOROOT README.md
	//mcore.PrintFile(readme)
}

func main() {
	demoJoinPath()
}
