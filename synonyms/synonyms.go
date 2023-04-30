package synonyms

// Import packages
import (
	"strings"

	"github.com/realTristan/quickwrite.xyz/requests"
	"github.com/valyala/fasthttp"
)

// Request Client used for sending http requests
var RequestClient *fasthttp.Client = &fasthttp.Client{}

// The CleanSpecialCharacters() function replaces
// all percentage characters (eg: %20) with their
// designated byte character (eg: %20 -> " ")
func CleanSpecialCharacters(s string) string {
	s = strings.ReplaceAll(s, "%20", " ")
	s = strings.ReplaceAll(s, "%2F", "/")
	return s
}

// The GetFromThesaurus() function is used to send
// an http request to the thesaurus website. Once the
// request has been sent, it returns the response body
func GetFromThesaurus(word string) string {
	// Create the http request object
	var req = &requests.HttpRequest{
		Client: RequestClient,
		Url:    "https://www.thesaurus.com/browse/" + word,
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

// The Parse function is used to parse the scraped synonyms
// from the thesaurus.com website into a string array.
//
// Using the provided body, the function splits it by
// the "data-linkid". After said split, it limits it's result
// to only seven rows. For each of the seven rows,
// split the row by the href tag to get the synonym name
//
// Once the synonym name has been acquired, append the synonym
// to the result slice
func Parse(body string) []string {
	var result []string = make([]string, 7)

	// Get the synonym rows
	var synonyms []string = strings.Split(body, `" data-linkid="nn1ov4"`)
	if len(synonyms) > 7 {
		synonyms = synonyms[:7]
	}

	// Iterate over the scraped synonym rows
	for i := 0; i < len(synonyms); i++ {
		var split []string = strings.Split(
			synonyms[i], `<a font-weight="inherit" href="/browse/`)

		// Check split length
		if len(split) <= 1 {
			continue
		}

		// Verify that the synonym is between 4 and 20 characters
		if len(split[1]) >= 4 && len(split[1]) <= 20 {
			// Append the synonym to the result slice
			var clean string = CleanSpecialCharacters(split[1])
			result = append(result, clean)
		}
	}
	// Return the result
	return result[7:]
}
