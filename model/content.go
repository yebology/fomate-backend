package model

type Course struct {
	CourseName string `json:course_name`
	CourseDuration uint64 `json:course_duration`
	CourseTitle string `json:course_title`
	CourseDescription string `json:course_description`
	CoursePrice uint64 `json:course_price`
	CourseInsights []string `json:course_insights`
}