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

// Home Page Handler
func HomePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Get the query from the url parameters
		var query string = strings.ToLower(r.URL.Query().Get("q"))

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
		var marshalData, _ = json.Marshal(synonyms)

		// If the synonyms result is equal to seven
		if len(synonyms) > 0 {

			// Set the query and the marshal data in the cache
			Cache.Set(query, marshalData)
		}

		// Write the synonyms to the response
		w.Write(marshalData)
	}
}
