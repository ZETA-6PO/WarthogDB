package engine

import (
	"errors"
	"log"
	"reflect"
)

var table map[string]interface{}

// initialisation et création de la map table
func init() {
	log.Println("[Engine] Loading table to memory.")
	table = make(map[string]interface{})
}

func getValue(key string) (error, interface{}) {
	e := false //variable d'existance
	for k := range table {
		if (k==key) {
			e = true;
		}
	}
	if !e {
		return errors.New("cannot find value at this key"), nil
	}
	return nil, table[key]
}

func setValue(key string, value interface{}) error {
	e := false //variable d'existance
	var t reflect.Value //variable type de la valeur
	for k, v := range table {
		if (k==key) {
			e = true;
			t = reflect.ValueOf(v)
		}
	}
	if !e { //si n'existe pas
		return errors.New("cannot find value at this key") //renvoie erreur
	} 
	
	if reflect.ValueOf(value) != t { //si pas le même type, renvoie erreur
		return errors.New("wrong type")
	}

	table[key] = value //set value


	return nil 
}

func createValue(key string, value interface{}) error{
	e := false //variable d'existancer
	for k := range table {
		if (k==key) {
			e = true;
		}
	}
	if e {
		return errors.New("existing value at this key, cannot override it")
	}
	table[key] = value
	return nil
}

func deleteValue(key string) error{
	e := false //variable d'existancer
	for k := range table {
		if (k==key) {
			e = true;
		}
	}
	if !e {
		return errors.New("cannot find value at this key")
	}
	delete(table, key)
	return nil
}

// fonction en prévision
func doesExist(key string) bool {
	e := false //variable d'existancer
	for k := range table {
		if (k==key) {
			e = true;
		}
	}
	return e
}