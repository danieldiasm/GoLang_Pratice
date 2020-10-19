package assembleur

import (
	"fmt"
)

//Assemble and print a splash screen
func Overture(largeur, hauteur int, titre, entonnoir string) (err string) {

	if verifTaille(entonnoir) {
		for i := 0; i < hauteur; i++ {
			//Stops execution in the middle to put the title
			if i == hauteur/2 {
				//Calculates each side filler size
				cote := ((largeur - len(titre)) / 2)
				//Fill up one size with characters
				for x := 0; x < cote; x++ {
					fmt.Print(entonnoir)
				}
				fmt.Print(titre)
				//Now fill up other side
				for x := 0; x < cote; x++ {
					fmt.Print(entonnoir)
				}
				fmt.Print("\n")
				//If not in the middle, just fill up
			} else {
				for j := 0; j < largeur; j++ {
					fmt.Print(entonnoir)
				}
				fmt.Print("\n")
			}
		}
		err = ""
		return err
	} else {
		err = "Filler is greater than 1 character"
		return err
	}
}

//Assemble and print a menu with given specs
func Menu(largeur, hauteur int, titre string, options []string) (err string) {
	printSeparator(largeur)
	printSeparator(largeur)
	err = ""
	return err
}

//Check filler characters qty
func verifTaille(entonnoir string) (result bool) {
	if len(entonnoir) > 1 {
		return false
	} else {
		return true
	}
}

//Print menu separator, also for top and bottom
func printSeparator(largeur int) {
	for i := 0; i < largeur; i++ {
		if i == 0 || i == largeur {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
		fmt.Print("\n")
	}
}

//Print options line
func printOptLine(largeur int, option string) {
	for i := 0; i < largeur; i++ {
		if i == 0 || i == largeur {
			fmt.Print("|")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}
