package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func collectEvaluation(r *http.Request, sid string) (e Evaluation, err error) {
	t := time.Now()
	tos := r.FormValue("surveyTitle") // Title of the survey
	aea := allEvaluationAnswers(r)    // All evaluation answers

	// Evaluation
	e = Evaluation{
		ID:                 [12]byte{},
		RowID:              "",
		CreatedAt:          t.Format("2006-01-02 15:04:05"),
		QuestionnaireID:    sid,
		QuestionnaireTitle: tos,
		Answers:            aea,
	}

	return
}

func allEvaluationAnswers(r *http.Request) (answers []EvaluationAnswer) {
	// Number Of question of the survey
	noq, err := strconv.Atoi(r.FormValue("surveyQuestionsNumber"))
	if err != nil {
		log.Println(err)
		return
	}

	// Collect answer of each question
	for i := 0; i < noq; i++ {
		aoq := r.Form["answerLabel"+strconv.Itoa(i+1)]           // answer(s) of question [type=checkbox|radio]
		fa := r.FormValue("freeanswerLabel" + strconv.Itoa(i+1)) // answer of question [type=text] (freeanswer)
		fr := isSlice(fa, aoq)                                   // store form response as []string

		ea := EvaluationAnswer{
			Title:    "Question " + strconv.Itoa(i+1),
			Response: fr,
		}
		answers = append(answers, ea)
	}
	return
}
func isSlice(s string, sl []string) (slc []string) {
	if len(s) > 0 {
		// answer is text (freeanswer)
		slc = append(slc, s)
	} else {
		// answer is checkbox|radio
		slc = sl

	}
	return
}
