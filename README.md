# Bad Words Filter
###### Creating to help Go developers to filter bad words in phrases/comments/text context.

## For Developers
It would be nice if you want to contribute to this project. I really need your help, there are a lot of things to do.
Feel free to create a pull request or make an issue to report bugs/request new features. You can see the list of things to do at the end of file.
🙇

## Installation
Simple installs the package to your $GOPATH with the go tool from shell:
```
$ go get -u github.com/go-sql-driver/mysql
```
Make sure Git is installed on your machine and in your system's PATH.

## Features
* Search of Bad Words
* Clean bad words in a sentence
* Change bad words in a sentence with */or other specified
* Change bad words to relative and polite other words
* Evaluation of profanities level of sentence
* Badwords Dataset in JSON
* Created and maintained by the community

## Usage
To use the lib features is necessary create a variable using the BadWordContent struct, imported by the lib.
```
phrase := Badwords.BadWordContent{
Text: "...",
Lang: "en",
FileLocation:"./badwords",
}
```
###### The `Text` (string) attribute is responsible for getting the phrase to be evaluated. The `Lang` (string) attribute is responsible to search in the dataset thee data of bad words registered on the respective JSON file. The `FileLocation` is responsible to find the location of all main files of lib.

### Methods

#### Search()
##### return []string, error
Return all the bad words found in the sentence.
The `err`(error) return a possible error, like if the language passed by parameters in the struct has no file in the dataset.
```
search, _ := phrase.Search()
```
###### The `search`([]string) variable get the words found in the dataset.

#### Clean()
##### return string, error
Return the sentence with no bad words, changed the bad words found by default with `*`.
```
clean, _ := phrase.Clean()
```
###### The `clean`(string) variable receive the new string.

#### CleanWith(string, bool)
##### return string, error
Return the sentence with no bad words, changed the bad words found with the string passed by parameters.
The bool parameter means if the value used to replace the bad word is an unique char or don't.
```
cleanWith, _ := phrase.Clean()
```
###### The `cleanWith`(string) variable receive the new string.

#### ChangeToBetter()
##### return string, error
Return the sentence with the bad words changed by the relative meaning(but polite) of each word found in the dataset.
```
changedToBetter, _ = phrase.ChangeToBetter()
```
###### The `changedToBetter`(string) variable receive the new string.

#### ProfanityLevel()
##### return float64, error
Return the profanity level of the sentence. Each bad word has the own profanity level in the dataset. The result is the sum of each bad word level by the numbers of normal words.
```
level, _ := phrase.ProfanityLevel()
```
###### The `level`(float64) variable receive the level of profanity calculated.

See an example of use:
```
package main

import (
	"wiltonribeiro/badwords"
	"fmt"
	"log"
)

func checkFail(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func main(){
	phrase := Badwords.BadWordContent{
		Text: "no fucking way",
		Lang: "en",
		FileLocation:"../github.com/badwords",
	}

	search, err := phrase.Search()
	checkFail(err)
	fmt.Println(search)

	clean, err := phrase.CleanWith("inapropiate",true)
	checkFail(err)
	fmt.Println(clean)


	clean, err = phrase.Clean()
	checkFail(err)
	fmt.Println(clean)

	clean, err = phrase.ChangeToBetter()
	checkFail(err)
	fmt.Println(clean)

	level, err := phrase.ProfanityLevel()
	checkFail(err)
	fmt.Println(level)
}
```
Output:
```
no inapropiate way
no ******* way
no freaking way
1
```
## To Do
- [x] Basic Funcitons
- [ ] Complete the parameters `relative_good` and `profanity_level` of English and Portuguese Dataset
- [ ] Add new languages Dataset Support (German, Spanish, Italian, Chinese, Arabian, etc)
- [ ] Make it a RESTFul API
