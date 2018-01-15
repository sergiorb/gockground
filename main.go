package main

import (
  "flag"
  "github.com/sergiorb/gockground/gockground"
  "os"
  "fmt"
)

var clientId  string
var folder    string
var galleryId string
var noDelete  bool

func init() {

  flag.StringVar(&clientId, "clientId", "", "Your own Imgur Client Id")
  flag.StringVar(&folder, "folder", "/tmp/gockground", "Folder to store downloaded images")
  flag.StringVar(&galleryId, "galleryId", "default", "Allows you to download and set an image from an Imgur gallery as your background")
  // flag.BoolVar(&noDelete, "noDelete", false, "Avoids last image deletion when downloading a new wallpaper.")
}

func main() {

  flag.Parse()

  if clientId == "" {

    fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", "clientId")
    os.Exit(2)
  }
  gg := gockground.NewGockground(clientId, folder)

  gallery := gg.GetGalleryInfo(galleryId)

  image := gockground.GetRandomImage(gallery.Images)

  gg.Apply(image)

}
