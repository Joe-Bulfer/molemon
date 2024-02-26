
package main

import (
	"fmt"
)

type attack struct {
    name string
    damage int
    atk_type string
}

type pokemon struct {
    name string
    health int
    attacks []attack
}

func choose_starter() string {
    var opt_starter string
    for {
        fmt.Println("Choose your pokemon\n1. Fire lizard\n2. Squirtle\n3. Bulbasaur")
        fmt.Scan(&opt_starter)
        switch opt_starter {
        case "1":
            fmt.Println("You chose Charmander!")
            return "Charmander"
        case "2":
            fmt.Println("You chose Squirtle!")
            return "Squirtle"
        case "3":
            fmt.Println("You chose Bulbasaur!")
            return "Bulbasaur"
        default:
            fmt.Println("Please choose a valid option.")
            // The loop continues, prompting the user again
        }
    }
}
func in_town_opt() string {
    		var opt_1 string
		for {
			fmt.Println("Choose an option.\n1. Go to Route 1\n2. Explore Pallet Town\n3. Go to Inventory") // Character representation
			fmt.Scan(&opt_1)
			switch opt_1 {
			case "1":
				retval := "Route 1"
				return retval
			case "2":
				retval := "Pallet Town"
				return retval
			case "3":
				retval := "Inventory"
				return retval
			default:
				fmt.Printf("Please choose a valid option.")
				continue
			}
		}
    }


//func (p pokemon) (name,health) {

func main() {
    fmt.Println("tittis")
    starter := choose_starter()
    fmt.Println(starter)
    opt := in_town_opt()
    fmt.Println(opt)
//    wild_pokemon_1 := [5]string{"Caterpie","Weedle","Pidgey","Rattatata"}
    scratch := attack{
	    name:   "Scratch",
	    damage: 10,
	    atk_type: "Normal",
    }

    weedle := pokemon{
	    name:   "Weedle",
	    health: 100,
	    attacks: []attack{scratch},
    }

    fmt.Printf("Weedle's Attack: %s, Damage: %d, Type: %s\n", weedle.attacks[0].name, weedle.attacks[0].damage, weedle.attacks[0].atk_type)
}
