package engine

import (
	"gitlab.com/yars-ai/yars/interfaces"
	"gitlab.com/yars-ai/yars/services/action"
	"gitlab.com/yars-ai/yars/services/recommend"
	"gitlab.com/yars-ai/yars/services/recommend/mf"
	"gitlab.com/yars-ai/yars/services/unit"
)

type Fabric interface {
	Unit(unitName string, entity interface{}) (unit.Store, interfaces.Migration)
	Action(actionName string) (action.Store, interfaces.Migration)
	Actor(recommendName string) (recommend.ActorStore, interfaces.Migration)
	ActorMF(recommendName string) (mf.ActorStore, interfaces.Migration)
	Object(recommendName string) (recommend.ObjectStore, interfaces.Migration)
	ObjectMF(recommendName string) (mf.ObjectStore, interfaces.Migration)
	SquashAction(recommendName string) (recommend.SquashActionStore, interfaces.Migration)
}
