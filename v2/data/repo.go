package data

import (
	"github.com/dimitrisCBR/shameboardAPI/v2/model"
	"fmt"
)

var currentId int;

var Shames model.Shames;

func init() {
	RepoCreateShame(model.Shame{ID : 1 ,Name: "Fuckup"})
	RepoCreateShame(model.Shame{ID : 1 ,Name: "Useless"})
}

func RepoCreateShame(shame model.Shame) model.Shame{
	currentId += 1;
	shame.ID = currentId;
	Shames = append(Shames,shame)
	return shame;
}

func RepoFindShame(id int) model.Shame {
	for _, shame := range Shames {
		if shame.ID == id {
			return shame
		}
	}

	return model.Shame{}
}

func RepoDeleteShame(id int) error {
	for i, shame := range Shames {
		if shame.ID == id {
			Shames = append(Shames[:i], Shames[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find shame with id %d to delete", id)
}