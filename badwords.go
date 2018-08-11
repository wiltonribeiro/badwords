package Badwords

import (
	"encoding/json"
	"regexp"
	"strings"
	"os"
	"io/ioutil"
	"errors"
)

type BadWordContent struct {
	Text string
	Lang string
	FileLocation string
}

type language struct {
	Initial string `json:"initial"`
	Name    string `json:"name"`
	Words   []struct {
		RelativeGood string `json:"relative_good"`
		BadWord      string `json:"bad_word"`
	} `json:"words"`
}

type words struct {
	RelativeGood string `json:"relative_good"`
	BadWord      string `json:"bad_word"`
}

func (p BadWordContent) listLanguagesData() (content []string, err error){
	files, err := ioutil.ReadDir(p.FileLocation+"/dataset")
	for _, f := range files {
		content = append(content,strings.Replace(f.Name(),".json","",-1))
	}
	return
}

func prepare(phrase string) (separate[]string) {
	var re = regexp.MustCompile(`[^A-Za-z\s]`)
	changed := re.ReplaceAllString(strings.ToLower(phrase), "")
	separate = strings.Split(changed," ")
	return
}

func (p BadWordContent) CheckLanguageExits(lang string) (err error){
	languages, err := p.listLanguagesData()
	for _, item := range languages {
		if item == lang {
			return
		}
	}
	return errors.New("language not exist in dataset")
}

func (p BadWordContent) openFile(filename string, lang string) (data language, err error) {

	err = p.CheckLanguageExits(lang)
	if err != nil {
		return
	}

	jsonFile, _ := os.Open(filename+"/dataset/"+lang+".json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &data)
	return
}

func (p BadWordContent) getWordsData(wordsPhrase []string, lang string, fileName string) (w []words, err error) {
	langData, err := p.openFile(fileName, lang)
	w = nil

	for _,word := range wordsPhrase {
		for _, bw := range langData.Words {
			if bw.BadWord == word {
				w = append(w, bw)
			}
		}
	}

	return
}

func (p BadWordContent) Search() ([]string, error) {

	words, err := p.getWordsData(prepare(p.Text), p.Lang, p.FileLocation)

	var content []string
	for _, word := range words {
		content = append(content,word.BadWord)
	}

	return content , err
}


func (p BadWordContent) Clean() (string, error){
	return p.CleanWith("*")
}

func (p BadWordContent) ChangeToBetter() (phrase string, err error) {
	words, err := p.getWordsData(prepare(p.Text), p.Lang, p.FileLocation)
	phrase = p.Text
	for _, word := range words {
		phrase = strings.Replace(phrase,word.BadWord,word.RelativeGood,-1)
	}
	return
}

func (p BadWordContent) CleanWith(value string) (phrase string, err error){
	badWords, err := p.Search()
	phrase = p.Text
	for _, badWord := range badWords {
		phrase = strings.Replace(phrase,badWord,strings.Repeat(value,len(badWord)),-1)
	}
	return
}