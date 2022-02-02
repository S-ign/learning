package pathfinding

import (
	"fmt"

	"github.com/S-ign/ffxigps/maps"
)

// FindPath .
func FindPath(startingZone string, destination string) []maps.FFXIZoneGetter {
	// add starting zone to list
	gpsPath := []maps.FFXIZoneGetter{maps.Vanadiel[startingZone]}

	// store destination connections in a list
	destinationConnectionsObj := [][]maps.FFXIZoneGetter{maps.Vanadiel[destination].GetConnectionsObjList()}

	// append connections of destination connections until starting zone is found
	lookingForStatingZone := true
	for lookingForStatingZone {
		for _, connections := range destinationConnectionsObj {
			for _, zone := range connections {
				if zone.GetName() == startingZone {
					lookingForStatingZone = false
				}
			}
		}
		destinationConnectionsObj = getConnections(destinationConnectionsObj)
	}

	mergedConnects := removeDuplicates(destinationConnectionsObj)
	badPath := []maps.FFXIZoneGetter{}

	for gpsPath[len(gpsPath)-1].GetName() != destination {
		gpsPath, badPath = findCorrectPath(mergedConnects, badPath, destination)
	}

	return gpsPath
}

func getConnections(destinationConnectionsObj [][]maps.FFXIZoneGetter) [][]maps.FFXIZoneGetter {
	for _, connections := range destinationConnectionsObj {
		for _, zone := range connections {
			destinationConnectionsObj = append(destinationConnectionsObj, maps.Vanadiel[zone.GetName()].GetConnectionsObjList())
		}
	}
	return destinationConnectionsObj
}

func removeDuplicates(destinationConnectionsObj [][]maps.FFXIZoneGetter) []maps.FFXIZoneGetter {
	mergedList := []maps.FFXIZoneGetter{}
	for _, connections := range destinationConnectionsObj {
		for _, zone := range connections {
			if !objInList(mergedList, zone) {
				mergedList = append(mergedList, zone)
			}
		}
	}
	return mergedList
}

func objInList(gpsPath []maps.FFXIZoneGetter, obj maps.FFXIZoneGetter) bool {
	for _, v := range gpsPath {
		if v.GetName() == obj.GetName() {
			return true
		}
	}
	return false
}

func findCorrectPath(gpsPath, badPath []maps.FFXIZoneGetter, destination string) ([]maps.FFXIZoneGetter, []maps.FFXIZoneGetter) {
	// iterate stored connections for what connection is attached to starting
	// zone prepend and it to list
	for _, connections := range gpsPath {
		//fmt.Println("Currently looking at:")
		//fmt.Println(connections)
		//fmt.Println()
		for _, concon := range connections.GetConnections() {
			//	fmt.Println("concon is:")
			//	fmt.Println(concon, path.GetName())
			//	fmt.Println()
			if concon == gpsPath[len(gpsPath)-1].GetName() && !objInList(gpsPath, connections) && !objInList(badPath, connections) {
				fmt.Println("!!!!!!!***********************HIT**********************!!!!!!!")
				fmt.Println(concon)
				gpsPath = append(gpsPath, connections)
			}
		}
	}
	//fmt.Println(gpsPath[:len(gpsPath)-1])
	return gpsPath, badPath
}
