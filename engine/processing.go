package engine

import (
	"encoding/json"
	"errors"
	"log"
)
type JSONRequest struct {
	Rq_type string `json:"type"`
	Key     string `json:"key,omitempty"`
	Value   interface{} `json:"value,omitempty"`
}

// Démarshallise le JSON de la requête TCP
// et retourne son équivalent @EngineRequest avec la gestion des
// erreurs en cas de malformation de la requête.
func ProcessRequest(json_rq []byte) (error, EngineRequest) {
	//démarshallisation avec gestion des erreurs
	var decoded_json JSONRequest
	err := json.Unmarshal(json_rq, &decoded_json)
	if (err!=nil) {
		log.Println(err.Error())
		return errors.New("cannot decode JSON request"), EngineRequest{}
	}

	//traitement du décodé
	switch decoded_json.Rq_type {
	case "set":
		if decoded_json.Key=="" || decoded_json.Value=="" {
			return errors.New("malformed json"), EngineRequest{}
		}
		return nil, EngineRequest{
			Response: nil,
			Action:"set",
			Key: decoded_json.Key,
			Value: decoded_json.Value,
		}
	case "create":
		if decoded_json.Key=="" || decoded_json.Value=="" {
			return errors.New("malformed json"), EngineRequest{}
		}
		return nil, EngineRequest{
			Response: nil,
			Action:"create",
			Key: decoded_json.Key,
			Value: decoded_json.Value,
		}
	case "get":
		if decoded_json.Key=="" {
			return errors.New("malformed json"), EngineRequest{}
		}
		return nil, EngineRequest{
			Response: nil,
			Action:"get",
			Key: decoded_json.Key,
		}
	case "delete":
		if decoded_json.Key=="" {
			return errors.New("malformed json"), EngineRequest{}
		}
		return nil, EngineRequest{
			Response: nil,
			Action:"delete",
			Key: decoded_json.Key,
		}
	default:
		return errors.New("unknown json 'type' value"), EngineRequest{}
	}
}
