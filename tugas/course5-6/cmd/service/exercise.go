package service

import (
	"course5-6/cmd/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ExerciseSvc struct {
	db *gorm.DB
}

func NewExerciseSvc(DB *gorm.DB) *ExerciseSvc {
	return &ExerciseSvc{DB}
}

func (exerciseSvc ExerciseSvc) CreateExercise(exercise model.Exercise) (int, *model.Exercise, error) {

	if exercise.Title == "" {
		return 417, nil, errors.New("title required")
	}
	if exercise.Description == "" {
		return 417, nil, errors.New("description required")
	}

	var count int64
	err := exerciseSvc.db.Model(&model.Exercise{}).Where("title = ?", exercise.Title).Count(&count).Error
	if err != nil {
		return 500, nil, err
	}
	if count > 0 {
		return 417, nil, errors.New("title already used")
	}

	newExercise := &model.Exercise{Title: exercise.Title, Description: exercise.Description, UserId: exercise.UserId}

	var createdExercise *model.Exercise
	err = exerciseSvc.db.Create(newExercise).Take(&createdExercise).Error
	if err != nil {
		return 500, nil, err
	}
	return 200, createdExercise, nil
}

func (exerciseSvc ExerciseSvc) GetExercise(id int) (int, *model.Exercise, error) {
	var exercise *model.Exercise
	err := exerciseSvc.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		errMsg := "exercise not found"
		return 404, nil, errors.New(errMsg)
		// return 500, nil, err
	}
	return 200, exercise, nil
}

func isValidCorrectAnswer(correctAnswer string) bool {
	switch correctAnswer {
	case
		"a",
		"b",
		"c",
		"d":
		return true
	}
	return false
}

func (exerciseSvc ExerciseSvc) CreateQuestion(question *model.Question) (int, error) {
	if question.Body == "" {
		return 417, errors.New("body required")
	}
	if question.OptionA == "" {
		return 417, errors.New("option a required")
	}
	if question.OptionB == "" {
		return 417, errors.New("option b required")
	}
	if question.OptionC == "" {
		return 417, errors.New("option c required")
	}
	if question.OptionD == "" {
		return 417, errors.New("option d required")
	}
	if question.Score < 1 {
		return 417, errors.New("score must greater than 1")
	}
	if !isValidCorrectAnswer(question.CorrectAnswer) {
		return 417, errors.New("invalid correct asnwer (must be a/b/c/d)")
	}

	var count int64
	err := exerciseSvc.db.Model(&model.Exercise{}).Where("id = ?", question.ExerciseId).Count(&count).Error
	if err != nil {
		return 500, err
	}
	if count == 0 {
		return 404, errors.New("exercise not found")
	}

	err = exerciseSvc.db.Create(question).Error
	if err != nil {
		return 500, err
	}
	return 201, nil
}

func (exerciseSvc ExerciseSvc) AnswerQuestion(answer model.Answer) (int, error) {
	if !isValidCorrectAnswer(answer.Answer) {
		return 417, errors.New("invalid answer (must be a/b/c/d)")
	}
	var count int64
	err := exerciseSvc.db.Model(&model.Exercise{}).Where("id = ?", answer.ExerciseId).Count(&count).Error
	if err != nil {
		return 500, err
	}
	if count == 0 {
		return 404, errors.New("exercise not found")
	}

	var existQuestion *model.Question
	err = exerciseSvc.db.Model(&model.Question{}).Where("id = ?", answer.QuestionId).Take(&existQuestion).Error
	if err != nil {
		return 500, err
	}
	if existQuestion == nil {
		return 404, errors.New("question not found")
	}

	if existQuestion.ExerciseId != answer.ExerciseId {
		return 417, errors.New("question not match with exercise")
	}

	if exerciseSvc.db.Model(&answer).Where("exercise_id = ? and question_id = ? and user_id = ?", answer.ExerciseId, answer.QuestionId, answer.UserId).Updates(&answer).RowsAffected == 0 {
		exerciseSvc.db.Create(&answer)
	}
	return 201, nil
}

func (exerciseSvc ExerciseSvc) GetScore(userId int, exerciseId int) (int, *model.GetScoreResponse, error) {
	var getScoreResponse *model.GetScoreResponse
	var listAnswer []model.Answer
	err := exerciseSvc.db.Where("user_id = ? and exercise_id = ?", userId, exerciseId).Find(&listAnswer).Error
	if err != nil {
		return 500, nil, err
	}

	var listQuestion []model.Question
	err = exerciseSvc.db.Where("exercise_id = ?", exerciseId).Find(&listQuestion).Error
	if err != nil {
		return 500, nil, err
	}

	var score int = 0
	for _, answer := range listAnswer {
		fmt.Printf("answer : %s", answer.Answer)
		for _, question := range listQuestion {
			if answer.QuestionId == question.ID {
				fmt.Printf("correct : %s", question.CorrectAnswer)
				if answer.Answer == question.CorrectAnswer {
					score += question.Score
				}
			}
		}
	}

	getScoreResponse = new(model.GetScoreResponse)
	getScoreResponse.Score = score
	return 200, getScoreResponse, nil
}
