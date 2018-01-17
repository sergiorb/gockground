package gockground

import (
  "fmt"
  "net/http"
  "encoding/json"
  "os"
)

const (
  API_URL = "https://api.imgur.com/3"
)

type Gockground struct {

  ApiUrl      string  `json:"api_url"`
  ClientId    string  `json:"client_id"`
  ImagesFolder string  `json:"image_folder"`
}

func NewGockground(clientId string, imagesFolder string) *Gockground {

  return &Gockground{
    API_URL,
    clientId,
    imagesFolder,
  }
}

func (g *Gockground) GetGalleryInfo(gelleryId string) *Gallery {

  url := fmt.Sprintf("%v/gallery/%v", g.ApiUrl, gelleryId)

  client := &http.Client{}

  r, _ := http.NewRequest("GET", url, nil)

  r.Header.Add("Authorization", fmt.Sprintf("Client-ID %v", g.ClientId))

  resp, _ := client.Do(r)

  defer resp.Body.Close()

  var apiResponseGallery ApiResponseGallery

  json.NewDecoder(resp.Body).Decode(&apiResponseGallery)

  return &apiResponseGallery.Data
}

func (g *Gockground) Apply(image Image) {

  imagePath := fmt.Sprintf("%v/%v.%v", g.ImagesFolder, image.Id, extractExtension(image))

  os.MkdirAll(g.ImagesFolder, os.ModePerm);

  downloadImage(image, imagePath)

  command := getSetWallPapperCommand(imagePath)

  _, stdErr := executeBashCommand(command)

  if stdErr != "" {

    fmt.Println(stdErr)
  }
}
