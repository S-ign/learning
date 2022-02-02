package main

import (
	"fmt"

	"github.com/S-ign/ffxigps/maps"
	"github.com/S-ign/ffxigps/pathfinding"
)

func main() {
	path := pathfinding.FindPath("Dangruf Wadi", "Valkurm Dunes")
	for _, zone := range path {
		fmt.Println(zone.GetName())
	}
}

func listAllMaps() {
	for _, v := range maps.Vanadiel {
		fmt.Println(v)
	}
}
