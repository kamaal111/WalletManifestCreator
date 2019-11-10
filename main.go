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
	Icon   string `json:"icon.png"`
	Icon2x string `json:"icon@2x.png"`
	Logo   string `json:"logo.png"`
	Logo2x string `json:"logo@2x.png"`
	Pass   string `json:"pass.json"`
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
		Icon:   hashItem("icon.png"),
		Icon2x: hashItem("icon@2x.png"),
		Logo:   hashItem("logo.png"),
		Logo2x: hashItem("logo@2x.png"),
		Pass:   hashItem("pass.json"),
	}

	manifest, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("manifest.json", manifest, 0644)

	fmt.Println("Done!")
}
