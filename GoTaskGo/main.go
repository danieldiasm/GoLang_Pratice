package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	assembleur "./assembleur"
)

//
func main() {
	splashScreen()
}

//Func that clear the screen ONLY ON MS WINDOWS
func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func splashScreen() {
	cls()
	err := assembleur.Overture(70, 30, "-= GO TASK GO =-", "/")
	if err != "" {
		fmt.Println(err)
	}
	time.Sleep(2 * time.Second)
	cls()
}

func mainMenu() {

}
