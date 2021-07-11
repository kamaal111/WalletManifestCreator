package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kamaal111/WalletManifestCreator/hasher"
)

func main() {
	assetsPath := os.Getenv("ASSET_PATH")

	hashedManifest, err := hasher.HashFiles(assetsPath, true)
	if err != nil {
		log.Fatal(err.Error())
	}

	manifest, err := json.MarshalIndent(hashedManifest, "", " ")
	if err != nil {
		log.Fatal(err.Error())
	}
	_ = ioutil.WriteFile("manifest.json", manifest, 0644)

	fmt.Println("Done!")
}
