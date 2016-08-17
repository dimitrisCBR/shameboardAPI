package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrAlreadyExists = errors.New("album already exists")
)

// The DB interface defines methods to manipulate the albums.
type DB interface {
	Get(id int) *Shame
	GetAll() []*Shame
	Find(user, description string, year int) []*Shame
	Add(a *Shame) (int, error)
	Update(a *Shame) error
	Delete(id int)
}

// Thread-safe in-memory map of albums.
type shamesDB struct {
	sync.RWMutex
	m   map[int]*Shame
	seq int
}

// The one and only database instance.
var db DB

func init() {
	db = &shamesDB{
		m: make(map[int]*Shame),
	}
	// Fill the database
	db.Add(&Shame{Id: 1, User: "Juan", Description: "Sucks", Year: 1986})
	db.Add(&Shame{Id: 2, User: "William", Description: "Sucks 2", Year: 1990})
	db.Add(&Shame{Id: 3, User: "Stefan", Description: "Sucks 3", Year: 1975})
}

// GetAll returns all albums from the database.
func (db *shamesDB) GetAll() []*Shame {
	db.RLock()
	defer db.RUnlock()
	if len(db.m) == 0 {
		return nil
	}
	ar := make([]*Shame, len(db.m))
	i := 0
	for _, v := range db.m {
		ar[i] = v
		i++
	}
	return ar
}

// Find returns albums that match the search criteria.
func (db *shamesDB) Find(user, description string, year int) []*Shame {
	db.RLock()
	defer db.RUnlock()
	var res []*Shame
	for _, v := range db.m {
		if v.User == user || user == "" {
			if v.Description == description || description == "" {
				if v.Year == year || year == 0 {
					res = append(res, v)
				}
			}
		}
	}
	return res
}

// Get returns the album identified by the id, or nil.
func (db *shamesDB) Get(id int) *Shame {
	db.RLock()
	defer db.RUnlock()
	return db.m[id]
}

// Add creates a new album and returns its id, or an error.
func (db *shamesDB) Add(a *Shame) (int, error) {
	db.Lock()
	defer db.Unlock()
	// Return an error if band-title already exists
	if !db.isUnique(a) {
		return 0, ErrAlreadyExists
	}
	// Get the unique ID
	db.seq++
	a.Id = db.seq
	// Store
	db.m[a.Id] = a
	return a.Id, nil
}

// Update changes the album identified by the id. It returns an error if the
// updated album is a duplicate.
func (db *shamesDB) Update(a *Shame) error {
	db.Lock()
	defer db.Unlock()
	if !db.isUnique(a) {
		return ErrAlreadyExists
	}
	db.m[a.Id] = a
	return nil
}

// Delete removes the album identified by the id from the database. It is a no-op
// if the id does not exist.
func (db *shamesDB) Delete(id int) {
	db.Lock()
	defer db.Unlock()
	delete(db.m, id)
}

// Checks if the album already exists in the database, based on the Band and Title
// fields.
func (db *shamesDB) isUnique(a *Shame) bool {
	for _, v := range db.m {
		if v.User == a.User && v.Description == a.Description && v.Id != a.Id {
			return false
		}
	}
	return true
}

// The Album data structure, serializable in JSON, XML and text using the Stringer interface.
type Shame struct {
	XMLName     xml.Name `json:"-" xml:"shame"`
	Id          int      `json:"id" xml:"id,attr"`
	User        string   `json:"user" xml:"user"`
	Description string   `json:"description" xml:"description"`
	Year        int      `json:"year" xml:"year"`
}

func (a *Shame) String() string {
	return fmt.Sprintf("%s - %s (%d)", a.User, a.Description, a.Year)
}
