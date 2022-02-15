package engine

import (
	"log"
)

type EngineRequest struct {
	Response chan EngineResponse
	Action string //get set delete create
	Key string
	Value interface{}
}
type EngineResponse struct {
	Err   error
	Value interface{}
}

var c_engine chan EngineRequest //implémente une channel pour les requests au moteur

//initialisation
func init() {
	c_engine = make(chan EngineRequest)
	go Engine()
	log.Println("[ENGINE] Just started.")
}

//ajoute à la queu
func AddToQueu(e EngineRequest) error{
	log.Println("[ENGINE] element added to queu.")
	c_engine<-e
	return nil
}

func Engine() { //implémente une routine
	request := <-c_engine//attends une requete

	switch request.Action {
	case "get":
		err, v := getValue(request.Key)
		if err != nil {
			request.Response<-EngineResponse{
				Err:   err,
				Value: v,
			}
			break
		}
		request.Response<-EngineResponse{Err:nil, Value: v}
	case "set":
		err := setValue(request.Key, request.Value)
		if err != nil { 
			request.Response<-EngineResponse{
				Err:   err,
				Value: nil,
			}
			break
		}
		request.Response<-EngineResponse{
			Err:   nil,
			Value: nil,
		}
	case "create":
		err := createValue(request.Key, request.Value)
		if err != nil { 
			request.Response<-EngineResponse{
				Err:   err,
				Value: nil,
			}
			break
		}
		request.Response<-EngineResponse{
			Err:   nil,
			Value: nil,
		}
	case "delete":
		err := deleteValue(request.Key)
		if err != nil { 
			request.Response<-EngineResponse{
				Err:   err,
				Value: nil,
			}
			break
		}
		request.Response<-EngineResponse{
			Err:   nil,
			Value: nil,
		}
	}
	

	Engine() //repeat
}