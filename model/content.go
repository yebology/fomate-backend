package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ContentId primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CourseName string `json:course_name`
	CourseDuration uint64 `json:course_duration`
	CourseTitle string `json:course_title`
	CourseDescription string `json:course_description`
	CoursePrice uint64 `json:course_price`
	CourseInsights []string `json:course_insights`
}