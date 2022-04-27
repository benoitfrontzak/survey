package main

import "log"

func collectReport(hexid string) (f TemplateFeedback, err error) {
	// collect all evaluations of questionnaire hexid
	e, err := getAllEvaluations(hexid)
	if err != nil {
		log.Println(err)
		return
	}
	// collect questionnaire
	q, err := getQuestionnaire(hexid)
	if err != nil {
		log.Println(err)
	}

	f = TemplateFeedback{
		Evaluations:   e,
		Questionnaire: q,
	}
	return
}
