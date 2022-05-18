package recommend

type Action struct {
	ActorID  string  `json:"actor_id" gorm:"size:256;primaryKey"`
	ObjectID string  `json:"object_id" gorm:"size:256;primaryKey"`
	Rate     float32 `json:"rate"`
}

type Object struct {
	ID string `gorm:"size:256;primaryKey"`
}

type Actor struct {
	ID string `gorm:"size:256;primaryKey"`
}
