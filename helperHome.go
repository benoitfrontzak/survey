package main

import "log"

// Get all questionnaires title
func getTemplateSurveys() (templateSurveys []TemplateSurvey) {

	aq, err := getQuestionnaires() // all questionnaires
	if err != nil {
		log.Println(err)
		return
	}

	for _, q := range aq {

		s := TemplateSurvey{
			ID:          q.ID,
			RowID:       q.ID.String(),
			Title:       q.Title,
			Description: q.Description,
		}
		templateSurveys = append(templateSurveys, s)
	}

	return
}
