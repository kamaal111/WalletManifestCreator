package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kamaal111/WalletManifestCreator/hasher"
)

func main() {
	manifest, manifestError := json.MarshalIndent(hasher.HashFiles(), "", " ")
	if manifestError != nil {
		log.Fatal(manifestError.Error())
	}
	_ = ioutil.WriteFile("manifest.json", manifest, 0644)

	fmt.Println("Done!")
}
