
package main

import (
	"fmt"
)

// this is a graph data structure where each location will have at least one connection
type Location struct {
    Name        string
    Connections []*Location
}

type GameMap struct {
    Locations map[string]*Location
}

func main() {
    gameMap := GameMap{
    Locations: make(map[string]*Location),
    }

    pallet_town := &Location{Name: "Pallet Town"}
    pallet_house := &Location{Name: "Your House"}
    pallet_lab := &Location{Name: "Professor Oak's Lab"}
    route_1 := &Location{Name: "Route 1"}
    route_21 := &Location{Name: "Route 21"}

    pallet_town.Connections = append(pallet_town.Connections,route_21,route_1,pallet_house,pallet_lab)
    pallet_lab.Connections = append(pallet_lab.Connections,pallet_town)
    pallet_house.Connections = append(pallet_house.Connections,pallet_town)

    viridian_city := &Location{Name: "Viridian City"}
    viridian_gym := &Location{Name: "Viridian Gym"}
    viridian_trainer_school := &Location{Name: "Viridian Trainer School"}
    viridian_pokemart := &Location{Name: "Viridian Pokemart"}

    viridian_city.Connections = append(viridian_city.Connections,route_1,viridian_gym,viridian_trainer_school,viridian_pokemart)
    viridian_gym.Connections = append(viridian_gym.Connections,viridian_city)
    viridian_trainer_school.Connections = append(viridian_trainer_school.Connections,viridian_city)
    viridian_pokemart.Connections = append(viridian_pokemart.Connections,viridian_city)


    gameMap.Locations["Pallet Town"] = pallet_town
    gameMap.Locations["Your House"] = pallet_house
    gameMap.Locations["Professor Oak's Lab"] = pallet_lab
    gameMap.Locations["Route 1"] = route_1
    gameMap.Locations["Route 21"] = route_21
    gameMap.Locations["Viridian City"] = viridian_city
    gameMap.Locations["Viridian Gym"] = viridian_gym
    gameMap.Locations["Viridian Trainer School"] = viridian_trainer_school
    gameMap.Locations["Viridian Pokemart"] = viridian_pokemart

    fmt.Println(gameMap, pallet_town, pallet_lab, pallet_house)

}


