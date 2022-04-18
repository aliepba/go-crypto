package helpers

import "github.com/aliepba/go-crypto/app/models"

type MetadataFormatter struct {
	Website      string `json:"website"`
	TechnicalDoc string `json:"technical_doc"`
	Twitter      string `json:"twitter"`
	MessageBoard string `json:"message_board"`
	SourceCode   string `json:"source_code"`
	CoinID       int    `json:"coin_id"`
}

func FormatMetadata(metadata models.MetadataCoin) MetadataFormatter {
	formatter := MetadataFormatter{}

	formatter.Website = metadata.Website
	formatter.TechnicalDoc = metadata.TechnicalDoc
	formatter.Twitter = metadata.Twitter
	formatter.MessageBoard = metadata.MessageBoard
	formatter.SourceCode = metadata.SourceCode
	formatter.CoinID = metadata.CoinID

	return formatter
}

func FormatMetadataCoins(metadataCoins []models.MetadataCoin) []MetadataFormatter {
	metadatasFormatter := []MetadataFormatter{}

	for _, metadata := range metadataCoins {
		metadataFormatter := FormatMetadata(metadata)
		metadatasFormatter = append(metadatasFormatter, metadataFormatter)
	}

	return metadatasFormatter
}
