package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Manifest represents the structure of the manifest.json file
type Manifest struct {
	Icon         string `json:"icon.png,omitempty"`
	Icon2x       string `json:"icon@2x.png,omitempty"`
	Logo         string `json:"logo.png,omitempty"`
	Logo2x       string `json:"logo@2x.png,omitempty"`
	Background   string `json:"background.png,omitempty"`
	Background2x string `json:"background@2x.png,omitempty"`
	Footer       string `json:"footer.png,omitempty"`
	Footer2x     string `json:"footer@2x.png,omitempty"`
	Strip        string `json:"strip.png,omitempty"`
	Strip2x      string `json:"strip@2x.png,omitempty"`
	Thumbnail    string `json:"thumbnail.png,omitempty"`
	Thumbnail2x  string `json:"thumbnail@2x.png,omitempty"`
	Pass         string `json:"pass.json,omitempty"`
}

func hashItem(filepath string) string {
	file, fileError := os.Open(filepath)

	if fileError != nil {
		panic(fileError.Error())
	}

	defer file.Close()

	hasher := sha1.New()

	if _, hasherError := io.Copy(hasher, file); hasherError != nil {
		log.Fatal(hasherError)
	}

	fmt.Println("Hashed " + filepath)
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	data := Manifest{
		Pass: hashItem("pass.json"),
	}

	files, filesError := ioutil.ReadDir("./")
	if filesError != nil {
		log.Fatal(filesError)
	}

	for _, file := range files {
		switch file.Name() {
		case "icon.png":
			data.Icon = hashItem(file.Name())
		case "icon@2x.png":
			data.Icon2x = hashItem(file.Name())
		case "background.png":
			data.Background = hashItem(file.Name())
		case "background@2x.png":
			data.Background2x = hashItem(file.Name())
		case "logo.png":
			data.Logo = hashItem(file.Name())
		case "logo@2x.png":
			data.Logo2x = hashItem(file.Name())
		case "footer.png":
			data.Footer = hashItem(file.Name())
		case "footer@2x.png":
			data.Footer2x = hashItem(file.Name())
		case "strip.png":
			data.Strip = hashItem(file.Name())
		case "strip@2x.png":
			data.Strip2x = hashItem(file.Name())
		case "thumbnail.png":
			data.Thumbnail = hashItem(file.Name())
		case "thumbnail@2x.png":
			data.Thumbnail2x = hashItem(file.Name())
		}
	}

	manifest, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("manifest.json", manifest, 0644)

	fmt.Println("Done!")
}
