package main

import "log"

func collectFeedback(hexid string) (f TemplateFeedback, err error) {
	// collect evaluation
	e, err := getEvaluation(hexid)
	if err != nil {
		log.Println(err)
		return
	}
	// add evaluation to []Evaluation
	var ae = []Evaluation{}
	ae = append(ae, e)
	// collect questionnaire
	q, err := getQuestionnaire(e.QuestionnaireID)
	if err != nil {
		log.Println(err)
	}

	f = TemplateFeedback{
		Evaluations:   ae,
		Questionnaire: q,
	}
	return
}
