package model

import "time"

// TableName overrides the table name used by User to `profiles`
func (Exercise) TableName() string {
	return "EXERCISE"
}

type Exercise struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UserId      int        `json:"user_id"`
	Questions   []Question `json:"questions"`
}

type CreateExerciseResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (Question) TableName() string {
	return "QUESTION"
}

type Question struct {
	ID            int       `json:"id"`
	Body          string    `json:"body"`
	OptionA       string    `json:"option_a"`
	OptionB       string    `json:"option_b"`
	OptionC       string    `json:"option_c"`
	OptionD       string    `json:"option_d"`
	CorrectAnswer string    `json:"correct_answer"`
	ExerciseId    int       `json:"exercise_id"`
	Score         int       `json:"score"`
	UserId        int       `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateQuestionResponse struct {
	Body          string `json:"body"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	CorrectAnswer string `json:"correct_answer"`
}

func (Answer) TableName() string {
	return "ANSWER"
}

type Answer struct {
	ID         int       `json:"id"`
	ExerciseId int       `json:"exercise_id"`
	QuestionId int       `json:"question_id"`
	UserId     int       `json:"user_id"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateAnswerRequest struct {
	Answer string `json:"answer"`
}

type GetScoreResponse struct {
	Score int `json:"score"`
}
