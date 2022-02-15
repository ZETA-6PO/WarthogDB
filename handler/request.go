package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"warthogdb/engine"
)

// Routines qui gère la connection
// en cas d'erreur la connection est rompue.
// Une requête authorisé par connection.
func HandleRequest(conn net.Conn) {
	log.Println("[Handler] Handling request from ", conn.RemoteAddr(), ".")
	buf := make([]byte, 1024) //max size
	l, err := conn.Read(buf)
	if err != nil {
		conn.Write([]byte("Error reading byte."))
		conn.Close()
		return
	}
	//redimensionnnement pour éviter `invalid character '\x00' after top-level value`
	buf = buf[:l]
	// prossesing et gestion des erreurs
	err, request := engine.ProcessRequest(buf)
	if (err!=nil) {
		conn.Write([]byte(fmt.Sprint("Error parsing JSON. ", err)))
		conn.Close()
		return
	}
	//création du channel de réponse
	c_response := make(chan engine.EngineResponse)
	request.Response = c_response
	//ajout à la queu
	engine.AddToQueu(request)
	//attente d'une réponse
	response := <-c_response
	//gestion des erreurs et mashalisation en JSON
	if response.Err != nil {
		r, _ := json.Marshal(map[string]interface{}{"error" : response.Err.Error()})
		conn.Write(r)
		conn.Close()
		return
	}
	//succes et Marshalisation
	r, _ := json.Marshal(map[string]interface{}{"value" : response.Value})
	conn.Write(r)
	conn.Close()
  }