package yugioh_master_duel_card_art

import (
	_ "embed"
	"encoding/json"
	"log"
	"strings"
)

var AllowChars = make(map[rune]bool)

func init() {
	// diff range string vs range []rune: https://stackoverflow.com/a/49062341/4097963,
	// TLDR: if don't care about rune index, same results
	for _, char := range "abcdefghijklmnopqrstuvwxyz_0123456789" {
		AllowChars[char] = true
	}
}

// NormalizeName keeps alphanumeric chars, replace others with underscore _
func NormalizeName(s string) string {
	var ret []rune
	for _, char := range strings.ToLower(s) {
		if !AllowChars[char] {
			ret = append(ret, '_')
		} else {
			ret = append(ret, char)
		}
	}
	return string(ret)
}

//go:embed konami_db.json
var allCardDataKonami []byte

/*
CardKonami is data crawled from Konami "db.yugioh-card.com", e.g.:

	{
		"CardName": "Blue-Eyes White Dragon",
		"CardType": "Monster",
		"CardSubtype": "MonsterNormal",
		"CardEffect": "This legendary dragon is a powerful engine of destruction. Virtually invincible, very few have faced this awesome creature and lived to tell the tale.",
		"CardArt": "",
		"MonsterAttribute": "LIGHT",
		"MonsterType": "Dragon",
		"MonsterLevelRankLink": 8,
		"MonsterATK": 3000,
		"MonsterATKStr": "3000",
		"MonsterDEF": 2500,
		"MonsterDEFStr": "2500",
		"MonsterAbilities": null,
		"MonsterLinkArrows": null,
		"IsNonEffectMonster": true,
		"IsPendulum": false,
		"PendulumScale": 0,
		"PendulumEffect": "",
		"MiscKonamiSet": "LOB-001",
		"MiscKonamiCardID": "4007",
		"MiscYear": "2002",
		"MiscCreator": ""
	}
*/
type CardKonami struct {
	// Konami cardID, same as file name in Master Duel data, equal to old Card.Cid
	CardID     string `json:"MiscKonamiCardID"`
	CardName   string
	MonsterATK float64
	AltArtID   string
}

func ReadAllCardDataKonami() map[string]CardKonami {
	var cardList []CardKonami
	err := json.Unmarshal(allCardDataKonami, &cardList)
	if err != nil {
		log.Fatalf("error json.Unmarshal: %v", err)
	}
	cards := make(map[string]CardKonami)
	for _, v := range cardList {
		cards[v.CardID] = v
	}

	// map alt art ID to original art ID
	altArts := map[string]string{
		"3801":  "4007",  // Blue-Eyes White Dragon
		"3863":  "4041",  // Dark Magician
		"3868":  "4998",  // Obelisk the Tormentor
		"3869":  "4999",  // Slifer the Sky Dragon
		"3401":  "5000",  // The Winged Dragon of Ra
		"20040": "5328",  // Reinforcement of the Army
		"3881":  "6653",  // Elemental HERO Neos
		"19077": "7696",  // Junk Warrior
		"3882":  "7734",  // Stardust Dragon
		"19943": "9122",  // Tuning
		"3895":  "11257", // El Shaddoll Winda
		"3894":  "11258", // El Shaddoll Construct
		"3891":  "12950", // Ash Blossom & Joyous Spring
		"3892":  "13587", // Ghost Belle & Haunted Mansion
		"3421":  "13601", // Knightmare Unicorn
		"3899":  "13668", // Sky Striker Ace - Kagari
		"3433":  "13669", // Sky Striker Ace - Shizuku
		"21227": "13670", // Sky Striker Ace - Raye
		"3434":  "13671", // Sky Striker Mobilize - Engage!
		"21228": "13763", // Sky Striker Ace - Hayate
		"3411":  "14496", // Apollousa, Bow of the Goddess
		"3415":  "14676", // I:P Masquerena
		"3423":  "15123", // Eldlich the Golden Lord
		"3437":  "15626", // Evil★Twin Ki-sikil
		"3438":  "15627", // Evil★Twin Lil-la
	}
	for alt, origin := range altArts {
		v := cards[origin]
		v.AltArtID = alt
		cards[alt] = v
	}

	log.Printf("ok ReadAllCardDataKonami, len(cards): %v", len(cards))
	return cards
}
