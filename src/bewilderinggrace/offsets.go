package bewilderinggrace

import (
	"encoding/json"
	"strconv"
)

type Offsets struct {
	Money                 uint64
	Items                 itemsOffset
	Characters            []characterOffset
	MainMemberID          uint64
	MainMemberCharacterID uint64
	MainMemberJobID       uint64
	MainMissionStart      uint64
	FirstCharacterID      uint64
	VisitedMap            []uint64
	VisitedRegion         uint64
	LevelName             uint64
	LevelX                uint64
	LevelY                uint64
	LevelZ                uint64
	LevelID1              uint64
	LevelID2              uint64
	LastBGMID1            uint64
	LastBGMID2            uint64
}

type itemsOffset struct {
	Base   uint64
	Size   uint64
	ItemID uint64
	Count  uint64
}

type characterOffset struct {
	ID              uint64
	RawHP           uint64
	RawMP           uint64
	FirstJobID      uint64
	SecondJobID     uint64
	SlotSword       uint64
	SlotLance       uint64
	SlotDagger      uint64
	SlotAxe         uint64
	SlotBow         uint64
	SlotRod         uint64
	SlotShield      uint64
	SlotHead        uint64
	SlotBody        uint64
	SlotAcc1        uint64
	SlotAcc2        uint64
	MissionProgress uint64
}

type offsetJSONAlias struct {
	Money                 string                     `json:"money"`
	Items                 itemsOffsetJSONAlias       `json:"items"`
	Characters            []characterOffsetJSONAlias `json:"characters"`
	MainMemberID          string                     `json:"main_member_id"`
	MainMemberCharacterID string                     `json:"main_member_character_id"`
	MainMemberJobID       string                     `json:"main_member_job_id"`
	MainMissionStart      string                     `json:"main_mission_start"`
	FirstCharacterID      string                     `json:"first_character_id"`
	VisitedMap            []string                   `json:"visited_map"`
	VisitedRegion         string                     `json:"visited_region"`
	LevelName             string                     `json:"level_name"`
	LevelX                string                     `json:"level_x"`
	LevelY                string                     `json:"level_y"`
	LevelZ                string                     `json:"level_z"`
	LevelID1              string                     `json:"level_id_1"`
	LevelID2              string                     `json:"level_id_2"`
	LastBGMID1            string                     `json:"last_bgm_1"`
	LastBGMID2            string                     `json:"last_bgm_2"`
}

type itemsOffsetJSONAlias struct {
	Base   string `json:"base"`
	Size   string `json:"size"`
	ItemID string `json:"item_id"`
	Count  string `json:"count"`
}

type characterOffsetJSONAlias struct {
	ID              string `json:"id"`
	RawHP           string `json:"raw_hp"`
	RawMP           string `json:"raw_mp"`
	FirstJobID      string `json:"first_job_id"`
	SecondJobID     string `json:"second_job_id"`
	SlotSword       string `json:"slot_sword"`
	SlotLance       string `json:"slot_lance"`
	SlotDagger      string `json:"slot_dagger"`
	SlotAxe         string `json:"slot_axe"`
	SlotBow         string `json:"slot_bow"`
	SlotRod         string `json:"slot_rod"`
	SlotShield      string `json:"slot_shield"`
	SlotHead        string `json:"slot_head"`
	SlotBody        string `json:"slot_body"`
	SlotAcc1        string `json:"slot_acc1"`
	SlotAcc2        string `json:"slot_acc2"`
	MissionProgress string `json:"mission_progress"`
}

func (o *Offsets) MarshalJSON() ([]byte, error) {
	a := &offsetJSONAlias{
		Money: strconv.FormatUint(o.Money, 16),
		Items: itemsOffsetJSONAlias{
			Base:   strconv.FormatUint(o.Items.Base, 16),
			Size:   strconv.FormatUint(o.Items.Size, 16),
			ItemID: strconv.FormatUint(o.Items.ItemID, 16),
			Count:  strconv.FormatUint(o.Items.Count, 16),
		},
		MainMemberID:          strconv.FormatUint(o.MainMemberID, 16),
		MainMemberCharacterID: strconv.FormatUint(o.MainMemberCharacterID, 16),
		MainMemberJobID:       strconv.FormatUint(o.MainMemberJobID, 16),
		MainMissionStart:      strconv.FormatUint(o.MainMissionStart, 16),
		FirstCharacterID:      strconv.FormatUint(o.FirstCharacterID, 16),
		VisitedRegion:         strconv.FormatUint(o.VisitedRegion, 16),
		LevelName:             strconv.FormatUint(o.LevelName, 16),
		LevelX:                strconv.FormatUint(o.LevelX, 16),
		LevelY:                strconv.FormatUint(o.LevelY, 16),
		LevelZ:                strconv.FormatUint(o.LevelZ, 16),
		LevelID1:              strconv.FormatUint(o.LevelID1, 16),
		LevelID2:              strconv.FormatUint(o.LevelID2, 16),
		LastBGMID1:            strconv.FormatUint(o.LastBGMID1, 16),
		LastBGMID2:            strconv.FormatUint(o.LastBGMID2, 16),
	}

	a.Characters = nil
	for _, c := range o.Characters {
		a.Characters = append(a.Characters, characterOffsetJSONAlias{
			ID:              strconv.FormatUint(c.ID, 16),
			RawHP:           strconv.FormatUint(c.RawHP, 16),
			RawMP:           strconv.FormatUint(c.RawMP, 16),
			FirstJobID:      strconv.FormatUint(c.FirstJobID, 16),
			SecondJobID:     strconv.FormatUint(c.SecondJobID, 16),
			SlotSword:       strconv.FormatUint(c.SlotSword, 16),
			SlotLance:       strconv.FormatUint(c.SlotLance, 16),
			SlotDagger:      strconv.FormatUint(c.SlotDagger, 16),
			SlotAxe:         strconv.FormatUint(c.SlotAxe, 16),
			SlotBow:         strconv.FormatUint(c.SlotBow, 16),
			SlotRod:         strconv.FormatUint(c.SlotRod, 16),
			SlotShield:      strconv.FormatUint(c.SlotShield, 16),
			SlotHead:        strconv.FormatUint(c.SlotHead, 16),
			SlotBody:        strconv.FormatUint(c.SlotBody, 16),
			SlotAcc1:        strconv.FormatUint(c.SlotAcc1, 16),
			SlotAcc2:        strconv.FormatUint(c.SlotAcc2, 16),
			MissionProgress: strconv.FormatUint(c.MissionProgress, 16),
		})
	}

	a.VisitedMap = nil
	for _, m := range o.VisitedMap {
		a.VisitedMap = append(a.VisitedMap, strconv.FormatUint(m, 16))
	}

	return json.Marshal(a)
}

func (o *Offsets) UnmarshalJSON(data []byte) error {
	var a offsetJSONAlias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	o.Money, _ = strconv.ParseUint(a.Money, 16, 64)
	o.Items.Base, _ = strconv.ParseUint(a.Items.Base, 16, 64)
	o.Items.Size, _ = strconv.ParseUint(a.Items.Size, 16, 64)
	o.Items.ItemID, _ = strconv.ParseUint(a.Items.ItemID, 16, 64)
	o.Items.Count, _ = strconv.ParseUint(a.Items.Count, 16, 64)
	o.MainMemberID, _ = strconv.ParseUint(a.MainMemberID, 16, 64)
	o.MainMemberCharacterID, _ = strconv.ParseUint(a.MainMemberCharacterID, 16, 64)
	o.MainMemberJobID, _ = strconv.ParseUint(a.MainMemberJobID, 16, 64)
	o.MainMissionStart, _ = strconv.ParseUint(a.MainMissionStart, 16, 64)
	o.FirstCharacterID, _ = strconv.ParseUint(a.FirstCharacterID, 16, 64)
	o.VisitedRegion, _ = strconv.ParseUint(a.VisitedRegion, 16, 64)
	o.LevelName, _ = strconv.ParseUint(a.LevelName, 16, 64)
	o.LevelX, _ = strconv.ParseUint(a.LevelX, 16, 64)
	o.LevelY, _ = strconv.ParseUint(a.LevelY, 16, 64)
	o.LevelZ, _ = strconv.ParseUint(a.LevelZ, 16, 64)
	o.LevelID1, _ = strconv.ParseUint(a.LevelID1, 16, 64)
	o.LevelID2, _ = strconv.ParseUint(a.LevelID2, 16, 64)
	o.LastBGMID1, _ = strconv.ParseUint(a.LastBGMID1, 16, 64)
	o.LastBGMID2, _ = strconv.ParseUint(a.LastBGMID2, 16, 64)

	o.Characters = nil
	for _, ac := range a.Characters {
		c := characterOffset{}
		c.ID, _ = strconv.ParseUint(ac.ID, 16, 64)
		c.RawHP, _ = strconv.ParseUint(ac.RawHP, 16, 64)
		c.RawMP, _ = strconv.ParseUint(ac.RawMP, 16, 64)
		c.FirstJobID, _ = strconv.ParseUint(ac.FirstJobID, 16, 64)
		c.SecondJobID, _ = strconv.ParseUint(ac.SecondJobID, 16, 64)
		c.SlotSword, _ = strconv.ParseUint(ac.SlotSword, 16, 64)
		c.SlotLance, _ = strconv.ParseUint(ac.SlotLance, 16, 64)
		c.SlotDagger, _ = strconv.ParseUint(ac.SlotDagger, 16, 64)
		c.SlotAxe, _ = strconv.ParseUint(ac.SlotAxe, 16, 64)
		c.SlotBow, _ = strconv.ParseUint(ac.SlotBow, 16, 64)
		c.SlotRod, _ = strconv.ParseUint(ac.SlotRod, 16, 64)
		c.SlotShield, _ = strconv.ParseUint(ac.SlotShield, 16, 64)
		c.SlotHead, _ = strconv.ParseUint(ac.SlotHead, 16, 64)
		c.SlotBody, _ = strconv.ParseUint(ac.SlotBody, 16, 64)
		c.SlotAcc1, _ = strconv.ParseUint(ac.SlotAcc1, 16, 64)
		c.SlotAcc2, _ = strconv.ParseUint(ac.SlotAcc2, 16, 64)
		c.MissionProgress, _ = strconv.ParseUint(ac.MissionProgress, 16, 64)
		o.Characters = append(o.Characters, c)
	}

	o.VisitedMap = nil
	for _, am := range a.VisitedMap {
		m, _ := strconv.ParseUint(am, 16, 64)
		o.VisitedMap = append(o.VisitedMap, m)
	}

	return nil
}
