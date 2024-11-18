package models

import (
	"time"
)

type Project struct {
	ID                  int       `db:"id" json:"id"`
	CreatedAt           time.Time `db:"created_at" json:"created_at"`
	Name                string    `db:"name" json:"name"`
	Type                string    `db:"type" json:"type"`
	GithubLink          *string   `db:"github_link" json:"github_link"`
	ProjectURL          *string   `db:"project_url" json:"project_url"`
	Year                *int      `db:"year" json:"year"`
	DevTime             *string   `db:"dev_time" json:"dev_time"`
	Technologies        *string   `db:"technologies" json:"technologies"`
	MainImage           *string   `db:"main_image" json:"main_image"`
	SecondaryImage      *string   `db:"secondary_image" json:"secondary_image"`
	Overview            *string   `db:"overview" json:"overview"`
	DetailedDescription *string   `db:"detailed_description" json:"detailed_description"`
	InProgress          *bool     `db:"in_progress" json:"in_progress"`
	BGColor             *string   `db:"bg_color" json:"bg_color"`
	TextColor           *string   `db:"text_color" json:"text_color"`
}

type ProjectList struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	MainImage  string `db:"main_image" json:"main_image"`
	InProgress *bool  `db:"in_progress" json:"in_progress"`
	BGColor    string `db:"bg_color" json:"bg_color"`
	TextColor  string `db:"text_color" json:"text_color"`
}
