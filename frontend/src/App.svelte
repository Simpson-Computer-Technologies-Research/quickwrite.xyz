<script>
	// Component Imports
	import RefreshButton from "./components/RefreshButton.svelte";
	import SiteHeader from "./components/SiteHeader.svelte";

	// Main Text Div text editor
	let DivTextEditor;

	// Word Map holding synonyms
	let WordMap = {}

	// SplitDivTextInput array holding the div text words
	let SplitDivTextInput = [];

	// The isAllLower() function is used to check
	// whether or not all the characters in the
	// provided string are lowercase.
	//
	// This function is used to prevent passing
	// invalid words to the golang api
	function isAllLower(s) {
		// Iterate over the string
		for (let i = 0; i < s.length; i++) {
			// If the string index doesn't equal it's
			// lowercase self
			if (s[i] != s[i].toLowerCase()) {
				return false
			}
		}
		return true
	}

	// The GetSynonyms() function is used to send an
	// http request to the golang api. This request
	// returns a response containing all the synonyms
	// for the provided word.
	function GetSynonyms(query) {
		// Send the http request to the golang api
		fetch("http://127.0.0.1:8000?q=" + query)
			.then((response) => response.json())
			.then((data) => {
				// Add the data to the wordmap
				WordMap[query] = data;
			})
	}

	// The QuerySynonyms() function is used to get all
	// the words that will be used for getting synonyms.
	function QuerySynonyms(s) {
		// Iterate over the sliced split
		for (let i = 0; i < s.length; i++) {
			// If the word is all lowercase and it's length is within 4-15
			if (isAllLower(s[i]) && s[i].length >= 4 && s[i].length <= 15) {
				// And the word doesn't already exist in the wordMap
				if (WordMap[s[i]] === undefined) {
					// Get the synonyms for said word (s[i])
					GetSynonyms(s[i]);
				}
			}

			// Set the SplitDivTextInput values
			if (s[i].length > 0) {
				SplitDivTextInput = [...SplitDivTextInput, s[i]]
			}
		}
	}
</script>

<main>
	<!-- Site Header -->
	<SiteHeader/>

	<!-- Main Text Div -->
	<!-- svelte-ignore a11y-autofocus -->
	<div
		autofocus="true"
		contenteditable="true"
		class="main_text_div"
		bind:this = {DivTextEditor}
	>
		<!-- For each word in the provided text -->
		{#each SplitDivTextInput as word, i}

            <!-- Check to make sure the word is valid -->
			{#if word.length > 0 && WordMap[word] !== undefined && WordMap[word].length > 0}

                <!-- Create the synonym dropdown div -->
				<div class="dropdown">
				    <span style="color: #6366f1; font-weight: 800; font-size: 20px;">{word} </span>
				    <div class="dropdown-content">

						<!-- For every synonym in the queries WordMap array -->
					    {#each WordMap[word] as synonym, n}
						    <div style="margin-bottom: 5px">
                                <!-- svelte-ignore a11y-invalid-attribute -->
                                <a href="#"
									style="color: #6366f1; font-weight: 500; font-size: 16px; cursor: pointer;" 
									on:click={() => {
										SplitDivTextInput[i] = WordMap[word][n]
										GetSynonyms(WordMap[word][n])
                                	}}>
								{synonym}</a>
                            </div>
					    {/each}
				    </div>
			    </div>
			{:else}
				<span style="color: white; font-weight: 700; font-size: 20px;">{word} </span>
			{/if}
		{/each}
	</div>

	<!-- Refresh Button -->
	<RefreshButton DivTextEditor={DivTextEditor} QuerySynonyms={QuerySynonyms}/>
</main>

<style>
	/* Main */
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}
	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}

	/* The Text Input Div */
    .main_text_div {
        white-space: pre-wrap;
        margin-left: 20%; 
        margin-right: 20%; 
        color: white; 
        font-weight: 700; 
        font-size: 20px;
        height: 100%;
        padding: 5%;
        border: 5px solid black;
        border-color: #6366f1;
        border-radius: 15px;
        outline: 0px solid transparent;
        background-color: #363636;
    }

    /* The Synonym Dropdown */
	.dropdown {
		display: inline-block;
	}
    /* The Synonym Dropdown Content */
	.dropdown-content {
		display: none;
		position: absolute;
		background-color: #FDFDFD;
		padding: 12px 16px;
		z-index: 1;
	}
    /* The Synonym Dropdown On Hover */
	.dropdown:hover .dropdown-content {
		display: block;
		border-radius: 15px;
	}

	/* Scroll Bar */
	:root::-webkit-scrollbar {
		width: 20px;
	}
	:root::-webkit-scrollbar-track {
		background: #f1f1f1; 
	}
	:root::-webkit-scrollbar-thumb {
		background: #6366f1;
	}
	:root::-webkit-scrollbar-thumb:hover {
		background: #474af2;
	}
</style>
