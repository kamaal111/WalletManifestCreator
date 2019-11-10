package hasher

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kamaal111/WalletManifestCreator/types"
)

func hashItem(filepath string) string {
	file, fileError := os.Open(filepath)
	if fileError != nil {
		panic(fileError.Error())
	}
	defer file.Close()

	hasher := sha1.New()
	if _, hasherError := io.Copy(hasher, file); hasherError != nil {
		panic(hasherError.Error())
	}

	fmt.Println("Hashed " + filepath)
	return hex.EncodeToString(hasher.Sum(nil))
}

// HashFiles returns a struct of all the necessary hashed values
func HashFiles() types.Manifest {
	data := types.Manifest{
		Pass: hashItem("pass.json"),
	}

	files, filesError := ioutil.ReadDir("./")
	if filesError != nil {
		panic(filesError.Error())
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

	return data
}
