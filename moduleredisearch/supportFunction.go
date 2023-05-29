package moduleredisearch

import (
	"github.com/av-belyakov/comparisondatabase/datamodels"
)

// GetIndex формирует тип с индексами
func GetIndex(elem datamodels.ElementSTIXObject) datamodels.IndexObject {
	mt := elem.Data.GeneratingDataForIndexing()
	indexObject := datamodels.IndexObject{
		Type: elem.DataType,
	}

	for k, v := range mt {
		if k == "id" {
			indexObject.ID = v
		}

		if k == "name" {
			indexObject.Name = v
		}

		if k == "description" {
			indexObject.Description = v
		}

		if k == "street_address" {
			indexObject.StreetAddress = v
		}

		if k == "abstract" {
			indexObject.Abstract = v
		}

		if k == "aliases" {
			indexObject.Aliases = v
		}

		if k == "content" {
			indexObject.Content = v
		}

		if k == "value" {
			indexObject.Value = v
		}
	}

	return indexObject
}
