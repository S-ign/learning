package maps

import (
	"fmt"
)

type ffxizone struct {
	name        string
	mapType     string
	connections []string
}

// FFXIZoneGetter .
type FFXIZoneGetter interface {
	String() string
	GetConnections() []string
	GetName() string
	GetMapType() string
	GetConnectionsObjList() []FFXIZoneGetter
	GetConnectionsObjMap() map[string]FFXIZoneGetter
}

func (f *ffxizone) String() string {
	return fmt.Sprintf("Name: %v\nMap Type: %v\nConnections: %v",
		f.name,
		f.mapType,
		f.connections)
}

func (f *ffxizone) GetConnections() []string {
	return f.connections
}

func (f *ffxizone) GetConnectionsObjList() []FFXIZoneGetter {
	objList := []FFXIZoneGetter{}
	for _, v := range f.connections {
		objList = append(objList, Vanadiel[v])
	}
	return objList
}

func (f *ffxizone) GetConnectionsObjMap() map[string]FFXIZoneGetter {
	objMap := map[string]FFXIZoneGetter{}
	for _, v := range f.connections {
		objMap[v] = Vanadiel[v]
	}
	return objMap
}

func (f *ffxizone) GetName() string {
	return f.name
}

func (f *ffxizone) GetMapType() string {
	return f.mapType
}

// NewMap Initilizer
func NewMap(name string, mapType string, connections []string) FFXIZoneGetter {
	return &ffxizone{
		name:        name,
		mapType:     mapType,
		connections: connections,
	}
}

var mapTypes = []string{"Dungeon", "Outdoor", "Town"}

// Vanadiel is the world map of FFXI
var Vanadiel = map[string]FFXIZoneGetter{
	"Horlais Peak":            NewMap("Horlais Peak", mapTypes[1], []string{"Yughott Grotto"}),
	"Yughott Grotto":          NewMap("Yughott Grotto", mapTypes[0], []string{"Ghelsea Outpost"}),
	"Ghelsea Outpost":         NewMap("Ghelsea Outpost", mapTypes[0], []string{"Ghelsea Fort", "Ronfaure West"}),
	"Ghelsea Fort":            NewMap("Ghelsea Fort", mapTypes[0], []string{"Ghelsea Outpost"}),
	"Ronfaure West":           NewMap("Ronfaure West", mapTypes[1], []string{"Ghelsea Outpost", "La Theine Plateau", "Ronfaure East", "South San D'Oria"}),
	"La Theine Plateau":       NewMap("La Theine Plateau", mapTypes[1], []string{"Ronfaure West", "Valkurm Dunes", "Jugner Forest"}),
	"Valkurm Dunes":           NewMap("Valkurm Dunes", mapTypes[1], []string{"Selbina", "La Theine Plateau", "Konschtat Highlands"}),
	"Selbina":                 NewMap("Selbina", mapTypes[2], []string{"Valkurm Dunes"}),
	"Konschtat Highlands":     NewMap("Konschtat Highlands", mapTypes[1], []string{"Valkurm Dunes", "Gusgen Mines", "Pashlow Marshlands", "Gustaberg North"}),
	"Pashlow Marshlands":      NewMap("Pashlow Marshlands", mapTypes[1], []string{"Rolanberry Fields", "Beaudeaux", "Konschtat Highlands"}),
	"Beaudeaux":               NewMap("Beaudeaux", mapTypes[0], []string{"Pashlow Marshlands"}),
	"Rolanberry Fields":       NewMap("Rolanberry Fields", mapTypes[1], []string{"Crawler's Nest", "Lower Jeuno", "Sauromugue Champaign"}),
	"Crawler's Nest":          NewMap("Crawler's Nest", mapTypes[0], []string{"Rolanberry Fields"}),
	"Gusgen Mines":            NewMap("Gusgen Mines", mapTypes[0], []string{"Konschtat Highlands"}),
	"Gustaberg North":         NewMap("Gustaberg North", mapTypes[1], []string{"Konschtat Highlands", "Palborough Mines", "Port Bastok", "Gustaberg South"}),
	"Gustaberg South":         NewMap("Gustaberg South", mapTypes[1], []string{"Gustaberg North", "Bastok Market", "Dangruf Wadi"}),
	"Dangruf Wadi":            NewMap("Dangruf Wadi", mapTypes[1], []string{"Gustaberg South"}),
	"Port Bastok":             NewMap("Port Bastok", mapTypes[2], []string{"Gustaberg North", "Gustaberg South", "Bastok Mines"}),
	"Bastok Market":           NewMap("Bastok Market", mapTypes[2], []string{"Gustaberg South", "Bastok Mines"}),
	"Bastok Mines":            NewMap("Bastok Mines", mapTypes[0], []string{"Bastok Market", "Zehrun Mines"}),
	"Zehrun Mines":            NewMap("Zehrun Mines", mapTypes[0], []string{"Palborough Mines", "Bastok Mines", "Korrolaka Tunnel"}),
	"Palborough Mines":        NewMap("Palborough Mines", mapTypes[0], []string{"Gustaberg North", "Zehrun Mines"}),
	"Korrolaka Tunnel":        NewMap("Korrolaka Tunnel", mapTypes[0], []string{"Zehrun Mines", "Altrepa Desert East", "Altrepa Desert West"}),
	"Altrepa Desert West":     NewMap("Altrepa Desert West", mapTypes[1], []string{"Altrepa Desert East", "Korrolaka Tunnel", "Rabao"}),
	"Altrepa Desert East":     NewMap("Altrepa Desert East", mapTypes[1], []string{"Altrepa Desert West", "Korrolaka Tunnel", "Rabao"}),
	"Rabao":                   NewMap("Rabao", mapTypes[2], []string{"Altrepa Desert West", "Altrepa Desert East"}),
	"South San D'Oria":        NewMap("South San D'Oria", mapTypes[2], []string{"Ronfaure West", "Ronfaure East", "Bostaunieux Oubliette"}),
	"Bostaunieux Oubliette":   NewMap("Bostaunieux Oubliette", mapTypes[0], []string{"Chateau D'Oraguille"}),
	"Chateau D'Oraguille":     NewMap("Chateau D'Oraguille", mapTypes[0], []string{"Bostaunieux Oubliette"}),
	"Ronfaure East":           NewMap("Ronfaure East", mapTypes[1], []string{"Ranguemont Pass", "Ranperre's Tomb", "Ronfaure West", "South San D'Oria"}),
	"Ranperre's Tomb":         NewMap("Ranperre's Tomb", mapTypes[0], []string{"Ronfaure East", "Jugner Forest"}),
	"Jugner Forest":           NewMap("Jugner Forest", mapTypes[1], []string{"Davoi", "Batallia Downs", "La Theine Plateau"}),
	"Davoi":                   NewMap("Davoi", mapTypes[0], []string{"Monastic Cavern"}),
	"Monastic Cavern":         NewMap("Monastic Cavern", mapTypes[0], []string{"Davoi"}),
	"Ranguemont Pass":         NewMap("Ranguemont Pass", mapTypes[1], []string{"Ronfaure East", "Beaucedine Glacier"}),
	"Beaucedine Glacier":      NewMap("Beaucedine Glacier", mapTypes[1], []string{"Xarcabard", "Batallia Downs", "Fei'Yin", "Ranguemont Pass"}),
	"Xarcabard":               NewMap("Xarcabard", mapTypes[1], []string{"Beaucedine Glacier"}),
	"Fei'Yin":                 NewMap("Fei'Yin", mapTypes[0], []string{"Beaucedine Glacier"}),
	"Batallia Downs":          NewMap("Batallia Downs", mapTypes[1], []string{"Upper Jeuno", "Eldieme Necropolis", "Jugner Forest", "Beaucedine Glacier"}),
	"Port Jeuno":              NewMap("Port Jeuno", mapTypes[2], []string{"Quifim Isle", "Sauromugue Champaign", "Rolanberry Fields", "Batallia Downs"}),
	"Lower Jeuno":             NewMap("Lower Jeuno", mapTypes[2], []string{"Rolanberry Fields"}),
	"Upper Jeuno":             NewMap("Upper Jeuno", mapTypes[2], []string{"Batallia Downs"}),
	"Eldieme Necropolis":      NewMap("Eldieme Necropolis", mapTypes[0], []string{"Batallia Downs"}),
	"Sauromugue Champaign":    NewMap("Sauromugue Champaign", mapTypes[1], []string{"Port Jeuno", "Meriphataud Mountains", "Rolanberry Fields"}),
	"Meriphataud Mountains":   NewMap("Meriphataud Mountains", mapTypes[1], []string{"Castle Oztroja", "Zi'Tah Sanctuary", "Sauromugue Champaign", "Tahrongi Canyon"}),
	"Castle Oztroja":          NewMap("Castle Oztroja", mapTypes[0], []string{"Meriphataud Mountains"}),
	"Tahrongi Canyon":         NewMap("Tahrongi Canyon", mapTypes[1], []string{"Meriphataud Mountains", "Sarutabaruta East", "Buburimu Peninsula", "Shakhrami Maze"}),
	"Shakhrami Maze":          NewMap("Shakhrami Maze", mapTypes[0], []string{"Buburimu Peninsula", "Tahrongi Canyon"}),
	"Sarutabaruta East":       NewMap("Sarutabaruta East", mapTypes[1], []string{"Tahrongi Canyon", "Horutoto Ruins Inner", "Windhurst Woods", "Sarutabaruta West"}),
	"Horutoto Ruins Inner":    NewMap("Horutoto Ruins Inner", mapTypes[0], []string{"Sarutabaruta East", "Sarutabaruta West", "Horutoto Ruins Outer"}),
	"Horutoto Ruins Outer":    NewMap("Horutoto Ruins Outer", mapTypes[0], []string{"Horutoto Ruins Inner", "Sarutabaruta West"}),
	"Sarutabaruta West":       NewMap("Sarutabaruta West", mapTypes[1], []string{"Horutoto Ruins Outer", "Port Windhurst"}),
	"Windhurst Woods":         NewMap("Windhurst Woods", mapTypes[2], []string{"Sarutabaruta East"}),
	"Port Windhurst":          NewMap("Port Windhurst", mapTypes[2], []string{"Sarutabaruta West"}),
	"Buburimu Peninsula":      NewMap("Buburimu Peninsula", mapTypes[1], []string{"Mhaura", "Shakhrami Maze", "Tahrongi Canyon"}),
	"Mhaura":                  NewMap("Mhaura", mapTypes[2], []string{"Buburimu Peninsula"}),
	"Zi'Tah Sanctuary":        NewMap("Zi'Tah Sanctuary", mapTypes[1], []string{"Ro'Maeve", "Meriphataud Mountains"}),
	"Ro'Maeve":                NewMap("Ro'Maeve", mapTypes[0], []string{"Zi'Tah Sanctuary", "Hall of the Gods"}),
	"Hall of the Gods":        NewMap("Hall of the Gods", mapTypes[0], []string{"Ro'Maeve"}),
	"Quifim Isle":             NewMap("Quifim Isle", mapTypes[1], []string{"Port Jeuno", "Behemoth's Dominion", "Delkfutt's Tower Lower"}),
	"Behemoth's Dominion":     NewMap("Behemoth's Dominion", mapTypes[1], []string{"Quifim Isle"}),
	"Delkfutt's Tower Lower":  NewMap("Delkfutt's Tower Lower", mapTypes[0], []string{"Quifim Isle", "Delkfutt's Tower Middle"}),
	"Delkfutt's Tower Middle": NewMap("Delkfutt's Tower Middle", mapTypes[0], []string{"Delkfutt's Tower Upper", "Delkfutt's Tower Lower"}),
	"Delkfutt's Tower Upper":  NewMap("Delkfutt's Tower Upper", mapTypes[0], []string{"Delkfutt's Tower Middle"}),
}
