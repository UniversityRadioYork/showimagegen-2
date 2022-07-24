/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package persistence

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

const bufferSize int = 5

type persistenceEntry struct {
	showID        int
	Title         string
	Filepath      string
	UploadedPhoto string
	Datetime      time.Time
}

// Engine contains the internal representation of our persistence, along with
// methods to read and write this from a file
type Engine struct {
	filepath     string
	lockFilepath string
	entries      chan persistenceEntry
	state        map[int][]persistenceEntry
}

// CreatePersistenceEngine will return a default Engine, along with an error
// if it encountered one reading the state from the file
func CreatePersistenceEngine() (*Engine, error) {
	engine := Engine{
		filepath:     "showimagegenstate",
		lockFilepath: ".showimagegenstate.lock",
		entries:      make(chan persistenceEntry, bufferSize),
	}

	if err := engine.read(); err != nil {
		return nil, err
	}

	return &engine, nil
}

func (e *Engine) waitForUnlock() error {
	var err error
	for {
		// TODO - this should timeout
		_, err = os.Stat(e.lockFilepath)
		if err != nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}

func (e *Engine) read() error {
	if err := e.waitForUnlock(); err != nil {
		return err
	}

	if _, err := os.Create(e.lockFilepath); err != nil {
		return err
	}
	defer os.Remove(e.lockFilepath)

	data, err := os.ReadFile(e.filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			e.state = make(map[int][]persistenceEntry)
			return nil
		}

		return err
	}

	if err := json.Unmarshal(data, &e.state); err != nil {
		return err
	}

	return nil
}

func (e *Engine) write() error {
	if err := e.waitForUnlock(); err != nil {
		return err
	}

	if _, err := os.Create(e.lockFilepath); err != nil {
		return err
	}
	defer os.Remove(e.lockFilepath)

	data, err := json.Marshal(e.state)
	if err != nil {
		return err
	}

	if err = os.WriteFile(e.filepath, data, 0644); err != nil {
		return err
	}

	return nil
}

// Daemon allows the persistence engine to run, waiting for entries to write
// and storing them
func (e *Engine) Daemon() {
	for {
		entry, open := <-e.entries

		if !open {
			break
		}

		e.state[entry.showID] = append(e.state[entry.showID], entry)

		if err := e.write(); err != nil {
			panic(err) // TODO
		}
	}
}

// AddEntry will add a title/filepath pair to the show ID in the persistence
func (e *Engine) AddEntry(showID int, title, filepath, uploadedPhoto string) {
	e.entries <- persistenceEntry{
		showID:        showID,
		Title:         title,
		Filepath:      filepath,
		UploadedPhoto: uploadedPhoto,
		Datetime:      time.Now(),
	}
}

// Close should be deferred and will tidy up the persistence engine
func (e *Engine) Close() {
	close(e.entries)
}
