package types

type Stats struct {
	Intelligence  uint32 `json:"intelligence,string,omitempty"`
	Strength      uint32 `json:"strength,string,omitempty"`
	Speed         uint32 `json:"speed,string,omitempty"`
	Durability    uint32 `json:"durability,string,omitempty"`
	Power         uint32 `json:"power,string,omitempty"`
	Combat        uint32 `json:"combat,string,omitempty"`
	ActualStamina uint8  `json:"actual_stamina,string,omitempty"`
}

type Hero struct {
	Id         uint32 `json:"id,string,omitempty"`
	Name       string `json:"name,omitempty"`
	PowerStats Stats  `json:"powerstats,omitempty"`
	Biography  Bio    `json:biography,omitempty`
}

type Bio struct {
	Alignment string `json:"alignment,omitempty"`
}
