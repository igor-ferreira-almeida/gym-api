package workout

import "time"

type Workout struct {
	Date      time.Time  `json:"date"`
	Exercises []Exercise `json:"exercises"`
}

type Exercise struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"muscle_group"`
}
