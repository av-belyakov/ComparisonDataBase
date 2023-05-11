package modulemongodb

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/av-belyakov/comparisondatabase/datamodels"
	methodstixobjects "github.com/av-belyakov/methodstixobjects"
)

type definingTypeSTIXObject struct {
	datamodels.CommonPropertiesObjectSTIX
}

// GetListElementSTIXObject возвращает, из БД, список STIX объектов
func GetListElementSTIXObject(cur *mongo.Cursor) (datamodels.ElementSTIXObject, error) {
	element := datamodels.ElementSTIXObject{}
	var modelType definingTypeSTIXObject

	if err := cur.Decode(&modelType); err != nil {
		return element, err
	}

	switch modelType.Type {
	//* *** Domain Objects STIX ***
	case "attack-pattern":
		tmpObj := methodstixobjects.AttackPatternDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil
	case "campaign":
		tmpObj := methodstixobjects.CampaignDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "course-of-action":
		tmpObj := methodstixobjects.CourseOfActionDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "grouping":
		tmpObj := methodstixobjects.GroupingDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "identity":
		tmpObj := methodstixobjects.IdentityDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "indicator":
		tmpObj := methodstixobjects.IndicatorDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "infrastructure":
		tmpObj := methodstixobjects.InfrastructureDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "intrusion-set":
		tmpObj := methodstixobjects.IntrusionSetDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "location":
		tmpObj := methodstixobjects.LocationDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "malware":
		tmpObj := methodstixobjects.MalwareDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "malware-analysis":
		tmpObj := methodstixobjects.MalwareAnalysisDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "note":
		tmpObj := methodstixobjects.NoteDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "observed-data":
		tmpObj := methodstixobjects.ObservedDataDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "opinion":
		tmpObj := methodstixobjects.OpinionDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "report":
		tmpObj := methodstixobjects.ReportDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "threat-actor":
		tmpObj := methodstixobjects.ThreatActorDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "tool":
		tmpObj := methodstixobjects.ToolDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "vulnerability":
		tmpObj := methodstixobjects.VulnerabilityDomainObjectsSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	//* *** Relationship Objects ***
	case "relationship":
		tmpObj := methodstixobjects.RelationshipObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "sighting":
		tmpObj := methodstixobjects.SightingObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	//* *** Cyber-observable Objects STIX ***
	case "artifact":
		tmpObj := methodstixobjects.ArtifactCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "autonomous-system":
		tmpObj := methodstixobjects.AutonomousSystemCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "directory":
		tmpObj := methodstixobjects.DirectoryCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "domain-name":
		tmpObj := methodstixobjects.DomainNameCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "email-addr":
		tmpObj := methodstixobjects.EmailAddressCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "email-message":
		tmpObj := methodstixobjects.EmailMessageCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "file":
		tmpObj := methodstixobjects.FileCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "ipv4-addr":
		tmpObj := methodstixobjects.IPv4AddressCyberObservableSimilarObjectSTIX{}
		//tmpObj := datamodels.IPv4AddressCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data: methodstixobjects.IPv4AddressCyberObservableObjectSTIX{
				CommonPropertiesObjectSTIX:                        tmpObj.CommonPropertiesObjectSTIX,
				OptionalCommonPropertiesCyberObservableObjectSTIX: tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX,
				Value:          tmpObj.Value,
				ResolvesToRefs: tmpObj.ResolvesToRefs,
				BelongsToRefs:  tmpObj.BelongsToRefs,
			},
		}, nil

	case "ipv6-addr":
		tmpObj := methodstixobjects.IPv6AddressCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "mac-addr":
		tmpObj := methodstixobjects.MACAddressCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "mutex":
		tmpObj := methodstixobjects.MutexCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "network-traffic":
		tmpObj := methodstixobjects.NetworkTrafficCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "process":
		tmpObj := methodstixobjects.ProcessCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "software":
		tmpObj := methodstixobjects.SoftwareCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "url":
		tmpObj := methodstixobjects.URLCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "user-account":
		tmpObj := methodstixobjects.UserAccountCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "windows-registry-key":
		tmpObj := methodstixobjects.WindowsRegistryKeyCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil

	case "x509-certificate":
		tmpObj := methodstixobjects.X509CertificateCyberObservableObjectSTIX{}
		err := cur.Decode(&tmpObj)
		if err != nil {
			return element, err
		}

		return datamodels.ElementSTIXObject{
			DataType: modelType.Type,
			Data:     tmpObj,
		}, nil
	}

	return element, errors.New("the stix object type was not found")
}
