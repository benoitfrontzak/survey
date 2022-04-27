package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func collectQuestionnaire(r *http.Request) (questionnaire Questionnaire, err error) {
	t := time.Now()

	// Collect questionnaire information from request
	st := r.FormValue("surveyTitle")                         // survey title
	sd := r.FormValue("surveyDescription")                   // survey description
	sa := r.FormValue("surveyAuthor")                        // survey author
	sc := r.FormValue("surveyCopyrights")                    // survey copyrights
	sac := r.FormValue("surveyAccessibility")                // survey accessibility
	sDate := string(t.Format("2006-01-02"))                  // survey date
	stl, err := strconv.Atoi(r.FormValue("surveyTimelimit")) // survey timelimits
	if err != nil {
		log.Println(err)
	}

	// total questions number
	tqn, err := strconv.Atoi(r.FormValue(("totalQuestions")))
	if err != nil {
		log.Println(err)
		return
	}

	// Store all question block
	aqb := collectQuestionBlock(tqn, r)

	// total rules number
	trn, err := strconv.Atoi(r.FormValue(("totalRules")))
	if err != nil {
		log.Println(err)
		return
	}
	//  Store all rules
	ar := collectRules(trn, r)

	questionnaire = Questionnaire{
		ID:            [12]byte{},
		RowID:         "",
		CreatedAt:     string(t.Format("2006-01-02 15:04:05")),
		UpdatedAt:     string(t.Format("2006-01-02 15:04:05")),
		DeletedAt:     "",
		Title:         st,
		Description:   sd,
		Author:        sa,
		Date:          sDate,
		Questionnaire: aqb,
		Blueprint:     ar,
		Status:        "",
		Accessibility: sac,
		Copyrights:    sc,
		Timelimit:     stl,
	}
	return
}

func collectQuestionBlock(n int, r *http.Request) (allQuestions []QuestionBlock) {

	// Collect all question block from request
	for i := 0; i < n; i++ {
		// Collect from request
		qs := r.FormValue("questionStory" + strconv.Itoa(i+1))                      // question story
		ql := r.FormValue("questionLabel" + strconv.Itoa(i+1))                      // question label
		qh := r.FormValue("questionHelper" + strconv.Itoa(i+1))                     // question helper
		qt := r.FormValue("questionType" + strconv.Itoa(i+1))                       // question type
		an, err := strconv.Atoi(r.FormValue(("answersNumber" + strconv.Itoa(i+1)))) // answers number
		if err != nil {
			log.Println(err)
			return
		}
		// Store all answers
		aa := make([]Answers, 0, an)
		// Collect all answers  with answers number
		for j := 0; j < an; j++ {
			// Collect from request
			al := r.FormValue("q" + strconv.Itoa(i+1) + "answer" + strconv.Itoa(j))                   // answer label
			ac := r.FormValue("q" + strconv.Itoa(i+1) + "comment" + strconv.Itoa(j))                  // answer comment
			as, err := strconv.Atoi(r.FormValue("q" + strconv.Itoa(i+1) + "score" + strconv.Itoa(j))) // answer score
			if err != nil {
				log.Println(err)
				return
			}
			// answer
			a := Answers{
				Label:   al,
				Score:   as,
				Comment: ac,
			}
			aa = append(aa, a)
		}
		questionBlock := QuestionBlock{
			Story:   qs,
			Label:   ql,
			Helper:  qh,
			Type:    qt,
			Answers: aa,
		}
		allQuestions = append(allQuestions, questionBlock)
	}
	return
}

func collectRules(n int, r *http.Request) (allRules []Report) {
	// Collect all rules from request
	for i := 0; i < n; i++ {
		// Collect from request
		ruleValue := r.FormValue("ruleValue" + strconv.Itoa(i+1))
		ruleOperator := r.FormValue("ruleOperator" + strconv.Itoa(i+1))
		ruleComment := r.FormValue("ruleComment" + strconv.Itoa(i+1))
		report := Report{
			TriggerValue:    ruleValue,
			TriggerOperator: ruleOperator,
			Comment:         ruleComment,
		}
		allRules = append(allRules, report)
	}
	return
}
