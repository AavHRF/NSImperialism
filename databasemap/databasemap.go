package databasemap

import "errors"

type DatabaseCell struct {
	ID       string
	Resident string
}

type DatabaseWar struct {
	Attacker      string
	Defender      string
	Score         int
	ID            string
	TerritoryName string
	IsOngoing     bool
}

func NewWar(attacker string, defender string, id string, territoryName string) DatabaseWar {
	return DatabaseWar{Attacker: attacker, Defender: defender, Score: 0, ID: id, TerritoryName: territoryName, IsOngoing: true}
}

type DatabaseMap struct {
	ID    string
	Year  int
	Cells map[string]DatabaseCell
	Wars  map[string]DatabaseWar
}

func NewBlankDatabaseMap() DatabaseMap {
	return DatabaseMap{
		Cells: make(map[string]DatabaseCell),
		Wars:  make(map[string]DatabaseWar),
	}
}

func NewDatabaseMapWithTerritories(territoryNames []string) DatabaseMap {
	databaseMap := NewBlankDatabaseMap()
	for _, territoryName := range territoryNames {
		databaseMap.Cells[territoryName] = DatabaseCell{territoryName, ""}
	}
	return databaseMap
}

func (databaseMap *DatabaseMap) SetResident(territoryName string, nationID string) error {

	_, doesCellExist := databaseMap.Cells[territoryName]
	if !doesCellExist {
		return errors.New("Territory doesn't exist")
	}

	databaseMap.Cells[territoryName] = DatabaseCell{territoryName, nationID}

	return nil
}

func (databaseMap DatabaseMap) GetResident(territoryName string) (string, error) {

	territory, doesCellExist := databaseMap.Cells[territoryName]
	if !doesCellExist {
		return "", errors.New("Territory doesn't exist")
	}

	return territory.Resident, nil
}

func (databaseMap DatabaseMap) HasResident(territoryName string) (bool, error) {

	resident, err := databaseMap.GetResident(territoryName)
	if err != nil {
		return false, err
	}

	return len(resident) != 0, nil
}

func (databaseMap DatabaseMap) GetWars() []DatabaseWar {

	warsToReturn := make([]DatabaseWar, 0, len(databaseMap.Wars))
	for _, war := range databaseMap.Wars {
		warsToReturn = append(warsToReturn, war)
	}

	return warsToReturn
}

func (databaseMap DatabaseMap) PutWars(warsToAdd []DatabaseWar) {

	for _, warToAdd := range warsToAdd {

		databaseMap.Wars[warToAdd.ID] = warToAdd
	}
}
