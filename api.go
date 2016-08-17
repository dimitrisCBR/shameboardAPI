package main


import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

// GetShames returns the list of shames (possibly filtered).
func GetShames(r *http.Request, enc Encoder, db DB) string {
	// Get the query string arguments, if any
	qs := r.URL.Query()
	user, description, yrs := qs.Get("user"), qs.Get("description"), qs.Get("year")
	yri, err := strconv.Atoi(yrs)
	if err != nil {
		// If year is not a valid integer, ignore it
		yri = 0
	}
	if user != "" || description != "" || yri != 0 {
		// At least one filter, use Find()
		return Must(enc.Encode(toIface(db.Find(user, description, yri))...))
	}
	// Otherwise, return all albums
	return Must(enc.Encode(toIface(db.GetAll())...))
}

// GetShame returns the requested shame.
func GetShame(enc Encoder, db DB, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	al := db.Get(id)
	if err != nil || al == nil {
		// Invalid id, or does not exist
		return http.StatusNotFound, Must(enc.Encode(
			NewError(ErrCodeNotExist, fmt.Sprintf("the shame with id %s does not exist", parms["id"]))))
	}
	return http.StatusOK, Must(enc.Encode(al))
}

// AddAlbum creates the posted album.
func AddShame(w http.ResponseWriter, r *http.Request, enc Encoder, db DB) (int, string) {
	al := getPostShame(r)
	id, err := db.Add(al)
	switch err {
	case ErrAlreadyExists:
		// Duplicate
		return http.StatusConflict, Must(enc.Encode(
			NewError(ErrCodeAlreadyExists, fmt.Sprintf("the shame '%s' from '%s' already exists", al.Description, al.User))))
	case nil:
		// TODO : Location is expected to be an absolute URI, as per the RFC2616
		w.Header().Set("Location", fmt.Sprintf("/albums/%d", id))
		return http.StatusCreated, Must(enc.Encode(al))
	default:
		panic(err)
	}
}

// UpdateShame changes the specified shame.
func UpdateShame(r *http.Request, enc Encoder, db DB, parms martini.Params) (int, string) {
	al, err := getPutShame(r, parms)
	if err != nil {
		// Invalid id, 404
		return http.StatusNotFound, Must(enc.Encode(
			NewError(ErrCodeNotExist, fmt.Sprintf("the shame with id %s does not exist", parms["id"]))))
	}
	err = db.Update(al)
	switch err {
	case ErrAlreadyExists:
		return http.StatusConflict, Must(enc.Encode(
			NewError(ErrCodeAlreadyExists, fmt.Sprintf("the shame '%s' from '%s' already exists", al.Description, al.User))))
	case nil:
		return http.StatusOK, Must(enc.Encode(al))
	default:
		panic(err)
	}
}

// Parse the request body, load into an Shame structure.
func getPostShame(r *http.Request) *Shame {
	user, description, yrs := r.FormValue("user"), r.FormValue("description"), r.FormValue("year")
	yri, err := strconv.Atoi(yrs)
	if err != nil {
		yri = 0 // Year is optional, set to 0 if invalid/unspecified
	}
	return &Shame{
		User:        user,
		Description: description,
		Year:        yri,
	}
}

// Like getPostShame, but additionnally, parse and store the `id` query string.
func getPutShame(r *http.Request, parms martini.Params) (*Shame, error) {
	al := getPostShame(r)
	id, err := strconv.Atoi(parms["id"])
	if err != nil {
		return nil, err
	}
	al.Id = id
	return al, nil
}

// Martini requires that 2 parameters are returned to treat the first one as the
// status code. Delete is an idempotent action, but this does not mean it should
// always return 204 - No content, idempotence relates to the state of the server
// after the request, not the returned status code. So I return a 404 - Not found
// if the id does not exist.
func DeleteShame(enc Encoder, db DB, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	al := db.Get(id)
	if err != nil || al == nil {
		return http.StatusNotFound, Must(enc.Encode(
			NewError(ErrCodeNotExist, fmt.Sprintf("the shame with id %s does not exist", parms["id"]))))
	}
	db.Delete(id)
	return http.StatusNoContent, ""
}

func toIface(v []*Shame) []interface{} {
	if len(v) == 0 {
		return nil
	}
	ifs := make([]interface{}, len(v))
	for i, v := range v {
		ifs[i] = v
	}
	return ifs
}
