package gockground

import (
  "fmt"
  "net/http"
  "encoding/json"
  "os/exec"
  "os"
  "io"
  "strings"
  "bytes"
)

const (
  API_URL = "https://api.imgur.com/3"
  BACKGROUND_COMMAND = "gsettings set org.cinnamon.desktop.background picture-uri"
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

  command := fmt.Sprintf("%v file://%v", BACKGROUND_COMMAND, imagePath)

  cmd := exec.Command("bash", "-c", command)

  var stderr bytes.Buffer

  cmd.Stderr = &stderr

  err := cmd.Run()

  if err != nil {
    
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
  }
}

// https://stackoverflow.com/questions/22417283/save-an-image-from-url-to-file
func downloadImage(image Image, path string) {

  response, err := http.Get(image.Link)

  if err != nil {
      panic(err)
  }

  defer response.Body.Close()

  file, err := os.Create(path)

  if err != nil {
    panic(err)
  }

  // Use io.Copy to just dump the response body to the file.
  // This supports huge files
  _, err = io.Copy(file, response.Body)

  if err != nil {

      panic(err)
  }

  file.Close()
}

func extractExtension(image Image) string {

  return strings.Split(image.Itype, "/")[1]
}
