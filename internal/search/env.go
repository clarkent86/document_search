package search

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// init env is used by th environment version of the document search app
// to read in environment variables provided in the root of the application
// in the `.env` file
func (search *Search) InitEnv() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("couldn't read config")
	}

	term, ok := viper.Get("TERM").(string)
	if !ok {
		return errors.New("invalid type assertion")
	}

	search.Term = term

	methodInterface := viper.Get("METHOD")
	if !ok {
		return errors.New("invalid type assertion")
	}
	search.Method, err = strconv.Atoi(methodInterface.(string))
	if err != nil {
		return err
	}
	if search.Method < 1 || search.Method > 3 {
		return errors.New("invalid search method")
	}

	path, ok := viper.Get("PATH_TO_TEXTS").(string)
	if !ok {
		return errors.New("invalid type assertion")
	}

	texts_directory, err := os.Open(path)
	if err != nil {
		return err
	}
	defer texts_directory.Close()

	files, _ := texts_directory.Readdirnames(0) // 0 to read all files and folders

	var text Text

	for _, name := range files {
		text.path = path + "/" + name
		text.content, _ = readInFile(text.path)
		text.Name = name
		search.Texts = append(search.Texts, text)
	}

	return nil
}

// readInFile is a helper function to read in a text file
func readInFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content), err
}
