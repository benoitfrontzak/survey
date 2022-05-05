package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Evaluation struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	RowID              string             `bson:"rowID,omitempty"`
	CreatedAt          string             `bson:"createdAt,omitempty"`
	QuestionnaireID    string             `bson:"qid,omitempty"`
	QuestionnaireTitle string             `bson:"qtitle,omitempty"`
	Answers            []EvaluationAnswer `bson:"answers,omitempty"`
}
type EvaluationAnswer struct {
	Title    string   `bson:"title,omitempty"`
	Response []string `bson:"response,omitempty"`
}

type Questionnaire struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	RowID         string             `bson:"rowID,omitempty"`
	CreatedAt     string             `bson:"createdAt,omitempty"`
	UpdatedAt     string             `bson:"updatedAt,omitempty"`
	DeletedAt     string             `bson:"deletedAt,omitempty"`
	Title         string             `bson:"title,omitempty"`
	Description   string             `bson:"description,omitempty"`
	Author        string             `bson:"author,omitempty"`
	Date          string             `bson:"date,omitempty"`
	Questionnaire []QuestionBlock    `bson:"questionnaire,omitempty"`
	Blueprint     []Report           `bson:"blueprint,omitempty"`
	Status        string             `bson:"status,omitempty"`
	Accessibility string             `bson:"accessibility,omitempty"`
	Copyrights    string             `bson:"copyrights,omitempty"`
	Timelimit     int                `bson:"timelimit,omitempty"`
}

type QuestionBlock struct {
	Story   string    `bson:"story,omitempty"`
	Label   string    `bson:"label,omitempty"`
	Helper  string    `bson:"helper,omitempty"`
	Type    string    `bson:"type,omitempty"`
	Answers []Answers `bson:"answers,omitempty"`
}

type Answers struct {
	Label   string `bson:"label,omitempty"`
	Score   int    `bson:"score,omitempty"`
	Comment string `bson:"comment,omitempty"`
}

type Report struct {
	TriggerValue    string `bson:"triggervalue,omitempty"`
	TriggerOperator string `bson:"triggeroperator,omitempty"`
	Comment         string `bson:"comment,omitempty"`
}

type TemplateData struct {
	Login   string
	Avatar  string
	Surveys []TemplateSurvey
	Data    interface{}
}

type TemplateSurvey struct {
	ID          primitive.ObjectID
	RowID       string
	Title       string
	Description string
}

type TemplateFeedback struct {
	Evaluations   []Evaluation
	Questionnaire Questionnaire
}
