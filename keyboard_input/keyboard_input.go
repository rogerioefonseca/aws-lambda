package keyboard

import (
	"bufio"
	"fmt"
	"os"
)

func GetCity() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name of the city")
	city, _ := reader.ReadString('\n')
	fmt.Print("You live in " + city)
}
