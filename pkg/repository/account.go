package repository

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/google/uuid"
)

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

func (a *AccountRepo) FindWorkoutByDate(
	ctx context.Context,
	accountId uuid.UUID,
	dayOfWeek int,
	date time.Time) (*model.Workout, error) {
	query := `SELECT w.id, w.name, w.description, w.workout_type, w.difficulty, w.calories_estimate, exe.id AS exercise_id, exe.name AS exercise_name, exe.muscle_group, exe.video_url, we.reps, we.sets, we.weight, we.rest, we.notes
    FROM workout_exercise AS we
    JOIN exercise AS exe ON exe.id = we.exercise
    JOIN workout AS w ON w.id = we.workout
    JOIN plan_workouts AS pw ON pw.workout = w.id
    JOIN training_plan AS pl ON pl.id = pw.training_plan
    JOIN account_training_plan AS atp ON atp.training_plan = pl.id
    WHERE atp.account = $1 AND atp.status = 'Active' AND pw.day_of_week = $2 AND pw.week = (($3  - atp.start_date) / 7) + 1
    ORDER BY we.order_number ASC`

	rows, err := a.db.QueryContext(ctx, query, accountId, dayOfWeek, date)

	if err != nil {
		slog.Error("Cannot get the workout for the day", "Error", err.Error())
		return nil, err
	}

	defer rows.Close()

	workout := new(model.Workout)

	for rows.Next() {
		newExercise := new(model.Exercise)

		if err := rows.Scan(
			&workout.Id,
			&workout.Name,
			&workout.Description,
			&workout.WorkoutType,
			&workout.Difficulty,
			&workout.CaloriesEstimate,
			&newExercise.Id,
			&newExercise.Name,
			&newExercise.MuscleGroup,
			&newExercise.VideoUrl,
			&newExercise.Reps,
			&newExercise.Sets,
			&newExercise.Weight,
			&newExercise.Rest,
			&newExercise.Notes); err != nil {
			return nil, err
		}

		workout.Exercises = append(workout.Exercises, newExercise)
	}

	return workout, nil
}
