package bewilderinggrace

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
)

const maxItems = 25

type Randomizer struct {
	save       []byte
	random     *rand.Rand
	characters []Character
	jobs       []Job
	missions   []Mission
	offsets    Offsets
}

func NewRandomizer(seed int64, characters []Character, jobs []Job, missions []Mission, offsets Offsets) *Randomizer {
	return &Randomizer{
		random:     rand.New(rand.NewSource(seed)),
		characters: characters,
		jobs:       jobs,
		missions:   missions,
		offsets:    offsets,
	}
}

func (r *Randomizer) Randomize(saveSlot int) error {
	if err := r.loadFile(); err != nil {
		return fmt.Errorf("cannot load default save: %s", err)
	}

	r.random.Shuffle(len(r.characters), func(i, j int) { r.characters[i], r.characters[j] = r.characters[j], r.characters[i] })
	r.random.Shuffle(len(r.jobs), func(i, j int) { r.jobs[i], r.jobs[j] = r.jobs[j], r.jobs[i] })
	r.random.Shuffle(len(r.missions), func(i, j int) { r.missions[i], r.missions[j] = r.missions[j], r.missions[i] })

	fmt.Printf("Your four travelers are %s, %s, %s, and %s!\n", r.characters[0].Name, r.characters[1].Name, r.characters[2].Name, r.characters[3].Name)

	if err := r.randomizeSave(); err != nil {
		return fmt.Errorf("cannot randomize save: %s", err)
	}

	if err := r.saveFile(saveSlot); err != nil {
		return fmt.Errorf("cannot create new save: %s", err)
	}

	return nil
}

func (r *Randomizer) randomizeSave() error {
	mainCharacter := r.characters[0]
	mainJob := r.jobs[0]

	var mainMission Mission
	// for _, m := range r.missions {
	// 	if m.CharacterID == mainCharacter.ID {
	// 		mainMission = m
	// 		break
	// 	}
	// }
	mainMission = r.missions[0]

	if err := r.writeValue(r.offsets.Money, mainCharacter.Money); err != nil {
		return err
	}

	for i := 0; i < maxItems; i++ {
		itemOffset := r.offsets.Items.Base + (r.offsets.Items.Size * uint64(i))

		if i < len(mainJob.Items) {
			item := mainJob.Items[i]
			r.writeValue(itemOffset+r.offsets.Items.ItemID, item.ItemID)
			r.writeValue(itemOffset+r.offsets.Items.Count, item.Count)
		} else {
			r.writeValue(itemOffset+r.offsets.Items.ItemID, 0)
			r.writeValue(itemOffset+r.offsets.Items.Count, 0)
		}
	}

	for i, character := range r.characters {
		job := r.jobs[i]
		offsets := r.offsets.Characters[i]

		r.writeValue(offsets.ID, character.ID)
		r.writeValue(offsets.RawHP, job.HP)
		r.writeValue(offsets.RawMP, job.MP)
		r.writeValue(offsets.FirstJobID, job.ID)
		r.writeValue(offsets.SecondJobID, -1)

		if e, ok := job.Equipment["sword"]; ok {
			r.writeValue(offsets.SlotSword, e)
		} else {
			r.writeValue(offsets.SlotSword, -1)
		}
		if e, ok := job.Equipment["lance"]; ok {
			r.writeValue(offsets.SlotLance, e)
		} else {
			r.writeValue(offsets.SlotLance, -1)
		}
		if e, ok := job.Equipment["dagger"]; ok {
			r.writeValue(offsets.SlotDagger, e)
		} else {
			r.writeValue(offsets.SlotDagger, -1)
		}
		if e, ok := job.Equipment["axe"]; ok {
			r.writeValue(offsets.SlotAxe, e)
		} else {
			r.writeValue(offsets.SlotAxe, -1)
		}
		if e, ok := job.Equipment["bow"]; ok {
			r.writeValue(offsets.SlotBow, e)
		} else {
			r.writeValue(offsets.SlotBow, -1)
		}
		if e, ok := job.Equipment["rod"]; ok {
			r.writeValue(offsets.SlotRod, e)
		} else {
			r.writeValue(offsets.SlotRod, -1)
		}
		if e, ok := job.Equipment["shield"]; ok {
			r.writeValue(offsets.SlotShield, e)
		} else {
			r.writeValue(offsets.SlotShield, -1)
		}
		if e, ok := job.Equipment["head"]; ok {
			r.writeValue(offsets.SlotHead, e)
		} else {
			r.writeValue(offsets.SlotHead, -1)
		}
		if e, ok := job.Equipment["body"]; ok {
			r.writeValue(offsets.SlotBody, e)
		} else {
			r.writeValue(offsets.SlotBody, -1)
		}
		if e, ok := job.Equipment["acc1"]; ok {
			r.writeValue(offsets.SlotAcc1, e)
		} else {
			r.writeValue(offsets.SlotAcc1, -1)
		}
		if e, ok := job.Equipment["acc2"]; ok {
			r.writeValue(offsets.SlotAcc2, e)
		} else {
			r.writeValue(offsets.SlotAcc2, -1)
		}

		if mainMission.CharacterID == character.ID {
			r.writeValue(offsets.MissionProgress, 0)
		} else {
			r.writeValue(offsets.MissionProgress, 0)
		}
	}

	r.writeValue(r.offsets.MainMemberID, 0)
	r.writeValue(r.offsets.MainMemberCharacterID, mainCharacter.ID)
	r.writeValue(r.offsets.MainMemberJobID, mainCharacter.ID)
	r.writeValue(r.offsets.MainMissionStart, mainMission.CharacterID)
	r.writeValue(r.offsets.FirstCharacterID, mainCharacter.ID)

	for i, o := range r.offsets.VisitedMap {
		if i == mainMission.MapIndex {
			r.writeValue(o, mainMission.MapID)
		} else {
			r.writeValue(o, 0)
		}
	}

	r.writeValue(r.offsets.VisitedRegion, mainMission.Region)
	// r.writeValue(r.offsets.LevelName, mainMission.LevelName)
	r.writeValue(r.offsets.LevelX, mainMission.Position.X)
	r.writeValue(r.offsets.LevelY, mainMission.Position.Y)
	r.writeValue(r.offsets.LevelZ, mainMission.Position.Z)
	r.writeValue(r.offsets.LevelID1, mainMission.LevelID)
	r.writeValue(r.offsets.LevelID2, mainMission.LevelID)
	r.writeValue(r.offsets.LastBGMID1, mainMission.BGM)
	r.writeValue(r.offsets.LastBGMID2, mainMission.BGM)

	return nil
}

func (r *Randomizer) writeValue(offset uint64, data interface{}) error {
	if offset == 0 {
		return nil
	}

	var bs []byte

	if i, ok := data.(int); ok {
		bs = make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, uint32(i))
	}
	if s, ok := data.(string); ok {
		bs = []byte(s)
	}
	if f, ok := data.(float32); ok {
		bs = make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, math.Float32bits(f))
	}
	if len(bs) == 0 {
		return fmt.Errorf("unknown data: %v", data)
	}

	for i, b := range bs {
		r.save[offset+uint64(i)] = b
	}

	return nil
}

func (r *Randomizer) loadFile() error {
	f, err := os.Open("assets/start.sav")
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	r.save = bytes
	return nil
}

func (r *Randomizer) saveFile(slot int) error {
	octoDir := os.ExpandEnv("$USERPROFILE\\Documents\\my games\\Octopath_Traveler")

	var steamDir string
	files, err := ioutil.ReadDir(octoDir)
	if err != nil {
		return err
	}

	for _, fi := range files {
		if _, err := strconv.Atoi(fi.Name()); fi.IsDir() && err == nil {
			steamDir = fi.Name()
			break
		}
	}

	filename := fmt.Sprintf("%s\\%s\\SaveGames\\SaveData%d.sav", octoDir, steamDir, slot)
	fmt.Printf("Saving to %s\n", filename)
	return ioutil.WriteFile(filename, r.save, 0644)
}
