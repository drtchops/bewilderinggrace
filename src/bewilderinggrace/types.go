package bewilderinggrace

type Character struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Money int    `json:"money"`
}

type Job struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	HP        int            `json:"hp"`
	MP        int            `json:"mp"`
	Equipment map[string]int `json:"equipment"`
	Items     []struct {
		ItemID int `json:"item_id"`
		Count  int `json:"count"`
	} `json:"items"`
}

type Mission struct {
	CharacterID int    `json:"character_id"`
	Progress    int    `json:"progress"`
	BGM         int    `json:"bgm"`
	MapID       int    `json:"map_id"`
	MapIndex    int    `json:"map_index"`
	Region      int    `json:"region"`
	LevelName   string `json:"level_name"`
	LevelID     int    `json:"level_id"`
	Position    struct {
		X float32 `json:"x"`
		Y float32 `json:"y"`
		Z float32 `json:"z"`
	} `json:"position"`
}
