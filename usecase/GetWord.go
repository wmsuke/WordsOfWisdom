package usecase

import "github.com/wmsuke/WordsOfWisdom/domains/models"

func GetWord() models.Words {
	var words models.Words
	words.Word = "人生において、人から与えられたものは何もなかった。いつも自分で取りにいかなければならない。"
	words.Author = "ケビン・デュラント"
	return words

}
