package handlers

import (
	"github.com/JGugino/survival-game-go/utils"
	"github.com/google/uuid"
)

type Entity interface {
	Update()
	Move()
	Draw()
	HandleInput()
}

type ActiveEntity struct {
	Id         uuid.UUID
	EntityData *utils.EntityData
	Entity     *Entity
}

type Entities struct {
	ActiveEntities map[uuid.UUID]*ActiveEntity
}

func (e *Entities) AddNewEntity() {

}
