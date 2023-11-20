package exercise

import (
	"encoding/json"
	"time"
)

type Set struct {
	Reps   uint64        `json:"reps"`
	Rest   time.Duration `json:"rest"`
	Weight uint64        `json:"weight"`
	Volume uint64        `json:"volume"`
}

type Execution struct {
	Method string `json:"method"`
	Sets   []Set  `json:"sets"`
	Volume uint64 `json:"volume"`
}

type Exercise struct {
	Name        string      `json:"name"`
	MuscleGroup string      `json:"muscle_group"`
	Executions  []Execution `json:"executions"`
}

func (e Exercise) GetByte() ([]byte, error) {
	return json.Marshal(e)
}

type Training struct {
	MuscleGroups         []string          `json:"muscle_groups"`
	Exercises            []Exercise        `json:"exercise"`
	Date                 time.Time         `json:"date"`
	Volume               uint64            `json:"volume"`
	VolumePerMuscleGroup map[string]uint64 `json:"volume_per_muscle_group"`
}
