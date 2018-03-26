// Dapatkan jumlah Star, Watch dari sebuah repository di github.com
package main

import (
	"fmt"
	"os"
)

func writeHelp() {
	fmt.Println("repovalue.exe golang/go golang/proposal")
}

func main() {
	repos := os.Args
	if len(repos) < 2 {
		writeHelp()
		os.Exit(1)
	}
	for i := 1; i < len(repos); i++ {
		fmt.Println("Menganalisa ", repos[i])
	}

}
