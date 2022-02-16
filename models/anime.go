package models

import "gorm.io/gorm"

type Anime struct {
  gorm.Model
  Name string
  Description string
  CoverImage string
  Status string  `gorm:"comment:'on-going, finished, running, dropped'`
  PublishedYear int
  Episodes []Episode
}