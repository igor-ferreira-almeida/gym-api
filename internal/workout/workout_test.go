package workout

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkout_CountExercises(t *testing.T) {
	as := assert.New(t)

	e1 := Exercise{MuscleGroup: "chest"}
	e2 := Exercise{MuscleGroup: "chest"}
	e3 := Exercise{MuscleGroup: "chest"}
	e4 := Exercise{MuscleGroup: "chest"}

	e5 := Exercise{MuscleGroup: "triceps"}
	e6 := Exercise{MuscleGroup: "triceps"}
	e7 := Exercise{MuscleGroup: "triceps"}

	e8 := Exercise{MuscleGroup: "abs"}
	e9 := Exercise{MuscleGroup: "abs"}

	workout := Workout{Exercises: []Exercise{e1, e2, e3, e4, e5, e6, e7, e8, e9}}

	result := workout.CountExercises()

	for key, value := range result {
		fmt.Println(fmt.Sprintf("muscle_group: %s - exercises: %d", key, value))
	}

	as.Equal(uint(4), result["chest"])
	as.Equal(uint(3), result["triceps"])
	as.Equal(uint(2), result["abs"])
}
