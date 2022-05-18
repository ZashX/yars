package repository

import "fmt"

func NewNamer(projectName string) Namer {
	return Namer{
		projectName: projectName,
	}
}

type Namer struct {
	projectName string
}

func (n *Namer) actorsTable(recommendName string) string {
	return fmt.Sprintf("%s__%s__actors", n.projectName, recommendName)
}

func (n *Namer) objectsTable(recommendName string) string {
	return fmt.Sprintf("%s__%s__objects", n.projectName, recommendName)
}

func (n *Namer) squashActionsTable(recommendName string) string {
	return fmt.Sprintf("%s__%s__squash_actions", n.projectName, recommendName)
}

func (n *Namer) unitsTable(unitName string) string {
	return fmt.Sprintf("%s__%s__unit", n.projectName, unitName)
}

func (n *Namer) actionsTable() string {
	return fmt.Sprintf("%s__actions", n.projectName)
}
