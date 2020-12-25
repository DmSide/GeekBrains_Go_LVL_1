package lesson7_cfgReader

import (
	"flag"
	"log"
)
import "github.com/go-playground/validator/v10"

var (
	lesson = flag.Int("lesson", 6, "Number of the lesson")
)

type Lesson struct {
	lesson int `validate:"required"`
}

var validate *validator.Validate

func ConfigInitByCmdParams() int {
	flag.Parse()

	validate = validator.New()

	lesson := &Lesson{
		lesson: *lesson,
	}
	err := validate.Struct(lesson)
	if err != nil {
		log.Fatalf("Validation error: %+v", err)
	}

	return lesson.lesson
}
