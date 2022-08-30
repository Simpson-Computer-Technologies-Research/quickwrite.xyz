package words

// Import packages
import (
	"strings"

	"github.com/realTristan/QuickWrite/server/requests"
	"github.com/valyala/fasthttp"
)

// The CleanSpecialCharacters() function replaces
// all percentage characters (eg: %20) with their
// designated byte character (eg: %20 -> " ")
func CleanSpecialCharacters(s string) string {
	s = strings.ReplaceAll(s, "%20", " ")
	s = strings.ReplaceAll(s, "%2F", "/")
	return s
}

// The SynonymRequest() function is used to send
// an http request to the thesaurus website. Once the
// request has been sent, it returns the response body
func SynonymRequest(synonym string) string {
	// Create the http request object
	var req = &requests.HttpRequest{
		Client: &fasthttp.Client{},
		Url:    "https://www.thesaurus.com/browse/" + synonym,
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36",
		},
	}

	// Send the http request
	var resp, _ = req.Send()

	// Return the response body
	return string(resp.Body())
}

// The GetSynonyms function is used to webscrape synonyms
// from the thesaurus.com website.
//
// Using the provided respBody, the function splits it by
// the "data-linkid". After said split, it limits it's result
// to only seven rows. For each of the seven rows,
// split the row by the href tag to get the synonym name
//
// Once the synonym name has been acquired, append the synonym
// to the result slice

// The GetSynonyms() function is used to get synonyms
// using the thesaurus response body
func GetSynonyms(respBody string) []string {
	var result []string = make([]string, 7)

	// Get the synonym rows
	var synonyms []string = strings.Split(respBody, `" data-linkid="nn1ov4"`)
	if len(synonyms) > 7 {
		synonyms = synonyms[:7]
	}

	// Iterate over the scraped synonym rows
	for i := 0; i < len(synonyms); i++ {
		var synonymSplit []string = strings.Split(
			synonyms[i], `<a font-weight="inherit" href="/browse/`)

		// Check split length
		if len(synonymSplit) > 1 {
			// Make sure the synonym is valid
			if len(synonymSplit[1]) >= 4 && len(synonymSplit[1]) <= 20 {

				// Append the synonym to the result slice
				result = append(result, CleanSpecialCharacters(synonymSplit[1]))
			}
		}
	}
	// Return the result
	return result[7:]
}
