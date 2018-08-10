package Badwords

import (
	"encoding/json"
	"regexp"
	"strings"
	"os"
	"fmt"
	"io/ioutil"
	"errors"
)

type BadWordContent struct {
	Text string
	Lang string
	FileLocation string
}

type languages struct {
	Languages [][]struct {
		Initial string `json:"initial"`
		Name    string `json:"name"`
		Words   []words `json:"words"`
	} `json:"languages"`
}

type words struct {
	RelativeGood string `json:"relative_good"`
	BadWord      string `json:"bad_word"`
}

func openFile(filename string) languages{
	var l languages
	jsonFile, err := os.Open(filename)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &l)
	return l
}

func prepare(phrase string) (separate[]string) {
	var re = regexp.MustCompile(`[^A-Za-z\s]`)
	changed := re.ReplaceAllString(strings.ToLower(phrase), "")
	separate = strings.Split(changed," ")
	return
}

func (p BadWordContent) Search() ([]string, error) {

	words, err := matchInJSON(prepare(p.Text), p.Lang, p.FileLocation)

	var content []string
	for _, word := range words {
		content = append(content,word.BadWord)
	}

	return content , err
}

func matchInJSON(words []string, lang string, fileName string) (w []words, err error){
	langs := openFile(fileName)
	w, err = nil, nil

	for _, item := range langs.Languages[0] {
		if item.Initial == lang {
			for _,word := range words {
				for _, bw := range item.Words {
					if bw.BadWord == word {
						w = append(w, bw)
					}
				}
			}
		}
	}

	if w == nil {
		err = errors.New("language not exist in json file")
	}

	return
}

func (p BadWordContent) Clean() (string, error){
	return p.CleanWith("*")
}

func (p BadWordContent) ChangeToBetter() (phrase string, err error) {
	words, err := matchInJSON(prepare(p.Text), p.Lang, p.FileLocation)
	phrase = p.Text
	for _, word := range words {
		phrase = strings.Replace(phrase,word.BadWord,word.RelativeGood,-1)
	}
	return
}

func (p BadWordContent) CleanWith(value string) (phrase string, err error){
	badwords, err := p.Search()
	phrase = p.Text
	for _, badword := range badwords {
		phrase = strings.Replace(phrase,badword,strings.Repeat(value,len(badword)),-1)
	}
	return
}

func (p BadWordContent) ListLanguagesByName() (content []string){
	langs := openFile(p.FileLocation)

	for _, item := range langs.Languages[0] {
		content = append(content,item.Name)
	}

	return
}

func (p BadWordContent) ListLanguagesById() (content []string){
	langs := openFile(p.FileLocation)

	for _, item := range langs.Languages[0] {
		content = append(content,item.Initial)
	}

	return
}