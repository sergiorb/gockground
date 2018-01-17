package gockground

import (
	"strings"
	"os"
  "os/exec"
  "net/http"
  "io"
  "bytes"
  "fmt"
)

const (

  CINNAMON  = "cinnamon"
  UNITY     = "unity"
  GNOME     = "gnome"
  UNKNOWN   = "unknown"
  DEBIAN    = "debian"
  UBUNTU    = "ubuntu"
  GNOME_BACKGROUND_COMMAND = "gsettings set org.gnome.desktop.background picture-uri"
)

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

func executeBashCommand(command string) (string, string) {


  cmd := exec.Command("bash", "-c", command)

  var stdout bytes.Buffer
  var stderr bytes.Buffer

  cmd.Stdout = &stdout
  cmd.Stderr = &stderr

  cmd.Run()

  return stdout.String(), stderr.String()
}

func getDesktop() string {

  stdOut, stdErr := executeBashCommand("uname -srv")

  if stdErr != "" {

    return UNKNOWN;
  }

  stdOutLowercase := strings.ToLower(stdOut)

  if (strings.Contains(stdOutLowercase, DEBIAN) || strings.Contains(stdOutLowercase, UBUNTU)) {

      return GNOME
  }

  return UNKNOWN
}

func getSetWallPapperCommand(imagePath string) string {

  var command string

  desktopType := getDesktop()

  if desktopType == GNOME {

    command = fmt.Sprintf("%v file://%v", GNOME_BACKGROUND_COMMAND, imagePath)
  }

  if desktopType == UNKNOWN {

    panic("Can't determinate desktop type.")
  }

  return command
}
