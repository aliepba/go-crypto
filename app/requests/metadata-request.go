package requests

import "github.com/aliepba/go-crypto/app/models"

type MetadataInput struct {
	CoinID       int    `json:"coin_id"`
	Website      string `json:"website"`
	TechnicalDoc string `json:"technical_doc"`
	Twitter      string `json:"twitter"`
	Reddit       string `json:"reddit"`
	MessageBoard string `json:"message_board"`
	SourceCode   string `json:"source_code"`
	User         models.User
}
