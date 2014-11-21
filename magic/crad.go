package crad

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)

type Crad struct {
	Name       string   `json:"name"`
	ManaCost   string   `json:"manaCost"`
	Colors     []string `json:"colors"`
	Type       string   `json:"type"`
	Rarity     string   `json:"rarity"`
	Text       string   `json:"text"`
	Supertypes []string `json:"supertypes"`
	Types      []string `json:"types"`
	Subtypes   []string `json:"subtypes"`
	Cmc        float64  `json:"cmc"`
}

func GetCrads() (map[string]Crad, map[float64][]Crad) {
	cradList, err := ioutil.ReadFile("./AllCards-x.json")
	if err != nil {
		log.Fatal("opening config file", err.Error())
	}

	// jsonParser := json.NewDecoder(cradList)
	crads := make(map[string]Crad)
	err = json.Unmarshal(cradList, &crads)
	if err != nil {
		log.Fatal("parsing config file", err.Error())
	}

	// now we need to parse each field!
	cmcs := indexCmc(crads)

	return crads, cmcs
}

func indexCmc(crads map[string]Crad) map[float64][]Crad {
	cmcs := make(map[float64][]Crad)
	for _, crad := range crads {
		cmcs[crad.Cmc] = append(cmcs[crad.Cmc], crad)
		// fmt.Printf("CRAD: %#v", crad.Cmc)
	}
	for _, cmcNew := range cmcs {
		fmt.Printf("CRADS: %#v", cmcNew)
	}

	return cmcs
}
