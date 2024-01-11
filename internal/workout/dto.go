package workout

import (
	"github.com/gym-api/internal/forces"
	"time"
)

type Request struct {
	Date      time.Time         `json:"date"`
	Exercises []ExerciseRequest `json:"exercises"`
}

type ExerciseRequest struct {
	Name        string       `json:"name"`
	MuscleGroup string       `json:"muscle_group"`
	Sets        []SetRequest `json:"sets"`
}

type SetRequest struct {
	Reps   uint          `json:"reps"`
	Weight uint          `json:"weight"`
	Rest   time.Duration `json:"rest"`
}

func (r Request) ToModel() Workout {
	exercises := make([]Exercise, len(r.Exercises))

	for i, e := range r.Exercises {
		exercises[i].Name = e.Name
		exercises[i].MuscleGroup = e.MuscleGroup
		sets := make([]Set, len(e.Sets))
		for si, s := range e.Sets {
			sets[si].Reps = s.Reps
			sets[si].Weight = s.Weight
			sets[si].Rest = s.Rest
			sets[si].GravitationalForce = float64(s.Weight) * forces.EarthGravity

			s.Rest.Seconds()
		}
		exercises[i].Sets = sets
	}
	return Workout{
		Date:      r.Date,
		Exercises: exercises,
	}
}
