package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/realTristan/QuickWrite/server/cache"
	"github.com/realTristan/QuickWrite/server/words"
)

// Declare global cache
var Cache *cache.Cache = cache.Init()

// The isAllLetters() function is used to determine
// whether all the characters in a string are either
// upper case letters or lower case letters.
func isAllLetters(s string) bool {
	for i := 0; i < len(s); i++ {
		// Use bytes to determine whether the character is a letter
		if !(s[i] >= 97 && s[i] <= 122) || !(s[i] >= 65 && s[i] <= 90) {
			return false
		}
	}
	return true
}

// The IsQueryErrors() function is used to return any errors
// in the provided url query.
func IsQueryErrors(w http.ResponseWriter, query string) bool {
	// Query is invalid
	if len(query) < 1 {
		w.Write([]byte(`{"error": "No query provided"}`))
		return true
	}

	// Query is too short
	if len(query) < 4 {
		w.Write([]byte(`{"error": "Query too short"}`))
		return true
	}

	// Query is too large
	if len(query) > 15 {
		w.Write([]byte(`{"error": "Query too large"}`))
		return true
	}

	// Query contains any non-alpha characters
	if !isAllLetters(query) {
		w.Write([]byte(`{"error": "Query has non-alpha characters"}`))
		return true
	}
	// There are no errors
	return false
}

// Synonym Page Handler
func SynonymsPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Get the query from the url parameters
		var query string = strings.ToLower(r.URL.Query().Get("q"))
		if IsQueryErrors(w, query) {
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
			synonymRequestBody string = words.SynonymRequest(query)

			// Get the array of scraped synonyms
			synonyms []string = words.GetSynonyms(synonymRequestBody)
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
