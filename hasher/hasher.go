package hasher

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/kamaal111/WalletManifestCreator/types"
	"github.com/kamaal111/kamaal-go-utils/files"
)

func hashItem(filepath string, logging bool) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha1.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}

	if logging {
		log.Printf("hashed %s\n", filepath)
	}

	return hex.EncodeToString(hasher.Sum(nil)), err
}

// HashFiles returns a struct of all the necessary hashed values or throws an error if something wrong
func HashFiles(assetsDirectory string, logging bool) (types.Manifest, error) {
	data := types.Manifest{}

	passHash, err := hashItem(files.AppendFileToPath(assetsDirectory, "pass.json"), logging)
	if err != nil {
		return data, err
	}

	data.Pass = passHash

	dirFiles, err := ioutil.ReadDir(assetsDirectory)
	if err != nil {
		return data, err
	}

	for _, file := range dirFiles {
		if file.IsDir() {
			continue
		}

		filePath := files.AppendFileToPath(assetsDirectory, file.Name())

		fileHash, err := hashItem(filePath, logging)
		if err != nil {
			return data, err
		}

		switch filePath {
		case files.AppendFileToPath(assetsDirectory, "icon.png"):
			data.Icon = fileHash
		case files.AppendFileToPath(assetsDirectory, "icon@2x.png"):
			data.Icon2x = fileHash
		case files.AppendFileToPath(assetsDirectory, "background.png"):
			data.Background = fileHash
		case files.AppendFileToPath(assetsDirectory, "background@2x.png"):
			data.Background2x = fileHash
		case files.AppendFileToPath(assetsDirectory, "logo.png"):
			data.Logo = fileHash
		case files.AppendFileToPath(assetsDirectory, "logo@2x.png"):
			data.Logo2x = fileHash
		case files.AppendFileToPath(assetsDirectory, "footer.png"):
			data.Footer = fileHash
		case files.AppendFileToPath(assetsDirectory, "footer@2x.png"):
			data.Footer2x = fileHash
		case files.AppendFileToPath(assetsDirectory, "strip.png"):
			data.Strip = fileHash
		case files.AppendFileToPath(assetsDirectory, "strip@2x.png"):
			data.Strip2x = fileHash
		case files.AppendFileToPath(assetsDirectory, "thumbnail.png"):
			data.Thumbnail = fileHash
		case files.AppendFileToPath(assetsDirectory, "thumbnail@2x.png"):
			data.Thumbnail2x = fileHash
		}
	}

	return data, nil
}
