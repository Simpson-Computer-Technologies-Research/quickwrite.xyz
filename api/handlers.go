package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/realTristan/quickwrite.xyz/cache"
	"github.com/realTristan/quickwrite.xyz/synonyms"
)

// Declare global cache
var Cache *cache.Cache = cache.Init()

// The isAllAlpha() function is used to determine
// whether all the characters in a string are either
// upper case letters or lower case letters.
func isAllAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		// Use bytes to determine whether the character is a letter
		// uppercase: 65 - 90
		// lowercase: 97 - 122
		if !(s[i] >= 97 && s[i] <= 122) && !(s[i] >= 65 && s[i] <= 90) {
			return false
		}
	}
	return true
}

// The IsQueryErrors() function is used to return any errors
// in the provided url query.
func existsQueryErrors(w http.ResponseWriter, query string) error {
	// Query is invalid
	if len(query) < 1 {
		return errors.New(`{"error": "No query provided"}`)
	}

	// Query is too short
	if len(query) < 4 {
		return errors.New(`{"error": "Query too short"}`)
	}

	// Query is too large
	if len(query) > 15 {
		return errors.New(`{"error": "Query too large"}`)
	}

	// Query contains any non-alpha characters
	if !isAllAlpha(query) {
		return errors.New(`{"error": "Query has non-alpha characters"}`)
	}

	// There are no errors
	return nil
}

// Synonym Page Handler
func SynonymsPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Get the query from the url parameters
		var query string = strings.ToLower(r.URL.Query().Get("q"))

		// If there's errors in the query
		if err := existsQueryErrors(w, query); err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// If the query exists in the cache
		if Cache.ExistsInData(query) {
			// Return the synonyms from the cache
			w.Write(Cache.Get(query))
			return
		}

		// Define Synonym Variables
		var (
			// Get the synonyms request body from thesaurus.com
			body string = synonyms.GetFromThesaurus(query)

			// Get the array of scraped synonyms
			synonyms []string = synonyms.Parse(body)
		)

		// Marshal the data and return it
		var marshalData, err = json.Marshal(synonyms)

		// If the synonyms result is equal to seven
		if len(synonyms) > 0 && err == nil {

			// Set the query and the marshal data in the cache
			Cache.Set(query, marshalData)
		}

		// Write the synonyms to the response
		w.Write(marshalData)
	}
}
