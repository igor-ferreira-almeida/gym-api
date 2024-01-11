package workout

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStats(t *testing.T) {
	as := assert.New(t)
	workout := GetChestDayWorkout()
	stats := NewStats(workout)

	as.Equal(uint(4), stats.MuscleGroupCount["chest"])
	as.Equal(uint(3), stats.MuscleGroupCount["triceps"])
	as.Equal(uint(1), stats.MuscleGroupCount["abs"])

	as.Equal(uint(80), stats.ExerciseMaxWeight["Barbell Bench Press"])
}

func GetChestDayWorkout() Workout {
	e1 := Exercise{
		Name:        "Barbell Bench Press",
		MuscleGroup: "chest",
		Sets: []Set{
			{Reps: 10, Weight: 40},
			{Reps: 5, Weight: 50},
			{Reps: 4, Weight: 60},
			{Reps: 4, Weight: 70},
			{Reps: 2, Weight: 80},
			{Reps: 20, Weight: 40},
		},
	}
	e2 := Exercise{
		Name:        "Incline Dumbbell Press",
		MuscleGroup: "chest",
		Sets: []Set{
			{Reps: 12, Weight: 64},
			{Reps: 8, Weight: 64},
			{Reps: 6, Weight: 64},
		},
	}
	e3 := Exercise{
		Name:        "Incline Dumbbell Fly",
		MuscleGroup: "chest",
		Sets: []Set{
			{Reps: 8, Weight: 40},
			{Reps: 7, Weight: 40},
			{Reps: 9, Weight: 40},
		},
	}
	e4 := Exercise{
		Name:        "Fly Machine",
		MuscleGroup: "chest",
		Sets: []Set{
			{Reps: 10, Weight: 40},
			{Reps: 11, Weight: 40},
			{Reps: 10, Weight: 40},
			{Reps: 10, Weight: 40},
			{Reps: 10, Weight: 40},
		},
	}

	e5 := Exercise{
		Name:        "Cable Barbell Push-Down",
		MuscleGroup: "triceps",
		Sets: []Set{
			{Reps: 15, Weight: 40},
			{Reps: 15, Weight: 50},
			{Reps: 11, Weight: 60},
			{Reps: 8, Weight: 70},
		},
	}
	e6 := Exercise{
		Name:        "Unilateral Overhead Triceps Extension",
		MuscleGroup: "triceps",
		Sets: []Set{
			{Reps: 10, Weight: 24},
			{Reps: 8, Weight: 24},
			{Reps: 11, Weight: 24},
			{Reps: 10, Weight: 24},
		},
	}
	e7 := Exercise{
		Name:        "Cable Rope High Pulley Triceps Extension",
		MuscleGroup: "triceps",
		Sets: []Set{
			{Reps: 12, Weight: 25},
			{Reps: 10, Weight: 25},
			{Reps: 10, Weight: 25},
		},
	}

	e8 := Exercise{
		Name:        "Incline Leg Raise",
		MuscleGroup: "abs",
		Sets: []Set{
			{Reps: 11, Weight: 1},
			{Reps: 10, Weight: 1},
			{Reps: 10, Weight: 1},
		},
	}
	return Workout{Exercises: []Exercise{e1, e2, e3, e4, e5, e6, e7, e8}}
}
