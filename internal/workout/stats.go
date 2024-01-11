package workout

import (
	"strings"
	"time"
)

type Stats struct {
	Date                  time.Time       `json:"date"`
	MuscleGroupCount      map[string]uint `json:"exercise_count"`
	MuscleGroupVolumeLoad map[string]uint `json:"muscle_group_volume_load"`
	ExerciseMaxWeight     map[string]uint `json:"exercise_maximum_weight"`
	ExerciseVolumeLoad    map[string]uint `json:"exercise_volume_load"`
}

func NewStats(w Workout) Stats {
	muscleGroupCount := make(map[string]uint)
	exerciseMaxWeight := make(map[string]uint)
	exerciseVolumeLoad := make(map[string]uint)
	muscleGroupVolumeLoad := make(map[string]uint)

	for _, e := range w.Exercises {
		for _, s := range e.Sets {
			exerciseVolumeLoad[e.Name] += s.Reps * s.Weight
			if s.Weight > exerciseMaxWeight[e.Name] {
				exerciseMaxWeight[e.Name] = s.Weight
			}
		}
		muscleGroupVolumeLoad[strings.ToLower(e.MuscleGroup)] += exerciseVolumeLoad[e.Name]
		muscleGroupCount[strings.ToLower(e.MuscleGroup)] += 1
	}
	return Stats{
		MuscleGroupCount:      muscleGroupCount,
		MuscleGroupVolumeLoad: muscleGroupVolumeLoad,
		ExerciseMaxWeight:     exerciseMaxWeight,
		ExerciseVolumeLoad:    exerciseVolumeLoad,
	}
}
