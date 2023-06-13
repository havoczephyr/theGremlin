package models

import "gorm.io/gorm"

type Questions struct {
	gorm.Model
	Body   string
	Author string
}

//status will have a number of states:

//Active indicates that the question is pending to be used
//Used indicates that the question has been used.
//PendingCuration indicates that the question is awaiting curation from the mod team.
//CurationFailed indicates the question was never used and was instead, denied.
