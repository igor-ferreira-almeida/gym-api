package workout

import (
	"strings"
	"time"
)

type Workout struct {
	Date      time.Time  `json:"date"`
	Exercises []Exercise `json:"exercises"`
}

func (w Workout) CountExercises() map[string]uint {
	result := make(map[string]uint)
	for _, e := range w.Exercises {
		result[strings.ToLower(e.MuscleGroup)] += 1
	}
	return result
}

type Exercise struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"muscle_group"`
	Sets        []Set  `json:"sets"`
}

type Set struct {
	Reps               uint          `json:"reps"`
	Weight             uint          `json:"weight"`
	Rest               time.Duration `json:"rest"`
	GravitationalForce float64       `json:"gravitational_force"`
}
