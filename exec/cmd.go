package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func ExecLine() {
	fmt.Print("blDownloader>  ")
	ch := make(chan string)
	go getStdin(ch)
	for cmd := range ch {
		fmt.Printf("%s", "blDownloader>  ")
		switch cmd {
		case "quit", "exit":
			fmt.Printf("%s\n", "bye")
			os.Exit(0)
		default:
			fmt.Printf("processCmd: %s", cmd)
			fmt.Printf("\n%s", "blDownloader>  ")
		}

	}

}
func getStdin(ch chan string) {
	defer close(ch)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ch <- input.Text()
	}
}
