package gockground

import (
  "math/rand"
  "time"
)

type Image struct {

  Id    string  `json:"id"`
  Title string  `json:"title"`
  Itype string  `json:"type"`
  Link  string  `json:"link"`
}

func GetRandomImage(images []Image) Image {

  rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

  return images[rand.Intn(len(images))]
}
