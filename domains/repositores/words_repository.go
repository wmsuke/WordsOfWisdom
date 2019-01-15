package repositores

import (
	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

type WordRepository interface {
	FindOne(ID int) (*models.Words, error)
}
