package bewilderinggrace

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func loadJSON(filename string, dest interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, dest)
	if err != nil {
		return err
	}

	return nil
}

// Randomize ...
func Randomize(saveSlot int, seed int64) error {
	var (
		characters []Character
		jobs       []Job
		missions   []Mission
		offsets    Offsets
	)

	if err := loadJSON("assets/characters.json", &characters); err != nil {
		return err
	}
	if err := loadJSON("assets/jobs.json", &jobs); err != nil {
		return err
	}
	if err := loadJSON("assets/missions.json", &missions); err != nil {
		return err
	}
	if err := loadJSON("assets/offsets.json", &offsets); err != nil {
		return err
	}

	r := NewRandomizer(seed, characters, jobs, missions, offsets)
	if err := r.Randomize(saveSlot); err != nil {
		return err
	}

	return nil
}
