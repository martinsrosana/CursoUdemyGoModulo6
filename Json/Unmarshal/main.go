package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome": "Rex", "raca": "poodle", "idade": 3}`

	var c cachorro
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorroDoisEmJSON := `{"nome": "Toto", "raca": "vira-lata"}`
	
	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorroDoisEmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}	
	fmt.Println(c2)
	
}
