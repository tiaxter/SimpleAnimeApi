package models

import "gorm.io/gorm"

type Episode struct {
  gorm.Model
  EpisodeNumber int
  Title string
  AnimeID int
  Anime Anime
}