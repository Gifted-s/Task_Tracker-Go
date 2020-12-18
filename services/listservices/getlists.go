package listservices

import (
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
)

func GetListsService() ([]models.List, error) {
	result, err := tododb.GetAllList()
	if err != nil {
		return nil, err
	}
	return result, nil
}
