package controller

import (
	"course5-6/cmd/model"
	"course5-6/cmd/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExerciseController struct {
	exerciseSvc *service.ExerciseSvc
}

func NewExerciseController(DB *gorm.DB) *ExerciseController {
	exerciseSvc := service.NewExerciseSvc(DB)
	return &ExerciseController{exerciseSvc}
}

func (exerciseController ExerciseController) CreateExercise(c *gin.Context) {
	var createExerciseRequest model.Exercise
	if err := c.ShouldBind(&createExerciseRequest); err != nil {
		Handler(c, 400, err)
		return
	}
	ctxUserId, _ := c.Get("id")
	strUserId := fmt.Sprint(ctxUserId)
	intUserId, err := strconv.Atoi(strUserId)
	if err != nil {
		Handler(c, 400, err)
		return
	}

	createExerciseRequest.UserId = intUserId
	code, createdExercise, err := exerciseController.exerciseSvc.CreateExercise(createExerciseRequest)
	if err == nil {
		var createExerciseResponse model.CreateExerciseResponse = model.CreateExerciseResponse{ID: createdExercise.ID,
			Title: createdExercise.Title, Description: createdExercise.Description}
		c.JSON(code, createExerciseResponse)
	} else {
		Handler(c, code, err)
	}
}

func (exerciseController ExerciseController) GetExercise(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		Handler(c, 400, err)
		return
	}

	code, exercise, err := exerciseController.exerciseSvc.GetExercise(id)
	if err == nil {
		c.JSON(code, exercise)
	} else {
		Handler(c, code, err)
	}
}

func (exerciseController ExerciseController) CreateQuestion(c *gin.Context) {
	strExerciseId := c.Param("exerciseId")
	exerciseId, err := strconv.Atoi(strExerciseId)
	if err != nil {
		Handler(c, 400, err)
		return
	}

	var createQuestionRequest *model.Question
	if err := c.ShouldBind(&createQuestionRequest); err != nil {
		Handler(c, 400, err)
		return
	}

	ctxUserId, _ := c.Get("id")
	strUserId := fmt.Sprint(ctxUserId)
	intUserId, err := strconv.Atoi(strUserId)
	if err != nil {
		Handler(c, 400, err)
		return
	}

	createQuestionRequest.ExerciseId = exerciseId
	createQuestionRequest.UserId = intUserId

	code, err := exerciseController.exerciseSvc.CreateQuestion(createQuestionRequest)
	if err == nil {
		c.JSON(code, map[string]string{
			"message": "success to create question",
		})
	} else {
		Handler(c, code, err)
	}
}

func (exerciseController ExerciseController) AnswerQuestion(c *gin.Context) {
	var createAnswerRequest model.CreateAnswerRequest
	if err := c.ShouldBind(&createAnswerRequest); err != nil {
		Handler(c, 400, err)
		return
	}
	strExerciseId := c.Param("exerciseId")
	exerciseId, err := strconv.Atoi(strExerciseId)
	if err != nil {
		Handler(c, 400, err)
		return
	}

	strQuestionId := c.Param("questionId")
	questionId, err := strconv.Atoi(strQuestionId)
	if err != nil {
		Handler(c, 400, err)
		return
	}

	ctxUserId, _ := c.Get("id")
	strUserId := fmt.Sprint(ctxUserId)
	intUserId, err := strconv.Atoi(strUserId)
	if err != nil {
		Handler(c, 400, err)
		return
	}

	var answer model.Answer = model.Answer{UserId: intUserId, ExerciseId: exerciseId, QuestionId: questionId, Answer: createAnswerRequest.Answer}

	code, err := exerciseController.exerciseSvc.AnswerQuestion(answer)
	if err == nil {
		c.JSON(code, map[string]string{
			"message": "success to answer question",
		})
	} else {
		Handler(c, code, err)
	}
}

func (exerciseController ExerciseController) GetScore(c *gin.Context) {
	ctxUserId, _ := c.Get("id")
	strUserId := fmt.Sprint(ctxUserId)
	intUserId, err := strconv.Atoi(strUserId)
	if err != nil {
		Handler(c, 400, err)
		return
	}
	strExerciseId := c.Param("exerciseId")
	exerciseId, err := strconv.Atoi(strExerciseId)
	if err != nil {
		Handler(c, 400, err)
		return
	}

	code, getScoreResponse, err := exerciseController.exerciseSvc.GetScore(intUserId, exerciseId)
	if err == nil {
		c.JSON(code, getScoreResponse)
	} else {
		Handler(c, code, err)
	}
}
