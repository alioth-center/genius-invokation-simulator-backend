package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Run() {
	for {
		if input, _, err := bufio.NewReader(os.Stdin).ReadLine(); err != nil {
			fmt.Printf("[cli.log] error occurred in cli.Run() %v\n", err)
			time.Sleep(time.Second)
		} else {
			if strings.Split(string(input), " ")[0] == "exit" {
				break
			}
			fmt.Printf("%+v\n", string(input))
		}
	}
}

func Quit() {}
