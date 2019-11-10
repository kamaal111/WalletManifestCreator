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
}

func main() {
	icon, iconError := os.Open("icon.png")
	icon2x, icon2xError := os.Open("icon@2x.png")

	if iconError != nil {
		panic(iconError.Error())
	}
	if icon2xError != nil {
		panic(icon2xError.Error())
	}

	defer icon.Close()
	defer icon2x.Close()

	hashIcon := sha1.New()
	hashIcon2x := sha1.New()

	if _, hashIconError := io.Copy(hashIcon, icon2x); hashIconError != nil {
		log.Fatal(hashIconError)
	}
	if _, hashIcon2xError := io.Copy(hashIcon2x, icon2x); hashIcon2xError != nil {
		log.Fatal(hashIcon2xError)
	}

	finalIconHash := hex.EncodeToString(hashIcon2x.Sum(nil))
	fmt.Println("Hashed icon.png")
	finalIcon2xHash := hex.EncodeToString(hashIcon2x.Sum(nil))
	fmt.Println("Hashed icon@2x.png")

	data := Manifest{
		Icon:   finalIconHash,
		Icon2x: finalIcon2xHash,
	}

	manifest, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("test.json", manifest, 0644)

	fmt.Println("Done!")
}
