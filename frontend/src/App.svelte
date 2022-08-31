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
		fetch("http://127.0.0.1:8000/synonyms?q=" + query)
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
				    <span class="hover-underline-animation">{word} </span>
				    <div class="dropdown-content">

						<!-- For every synonym in the queries WordMap array -->
					    {#each WordMap[word] as synonym, n}
						    <div style="margin-bottom: 5px;">
                                <!-- svelte-ignore a11y-invalid-attribute -->
                                <a href="#" style="cursor: pointer; color: #7c3aed;"
									on:click={() => {
										SplitDivTextInput[i] = WordMap[word][n]
										GetSynonyms(WordMap[word][n])
                                	}}>
									{synonym}
								</a>
                            </div>
					    {/each}
				    </div>
			    </div>

			<!-- Use a plain word with no special hover effects-->
			{:else}
				<span>{word} </span>
			{/if}
		{/each}
	</div>

	<!-- Refresh Button -->
	<RefreshButton 
		DivTextEditor={DivTextEditor} 
		QuerySynonyms={QuerySynonyms} 
	/>
</main>

<style>
	/* Main */
	main {
		padding: 1em;
		margin: 0 auto;
	}

	/* The Text Input Div */
    .main_text_div {
		text-align: left;
		box-shadow: 0.5px 0.5px 20px 0.5px rgba(0, 0, 0, 0.1);
        white-space: pre-wrap;
		margin-top: 3%;
        margin-left: 20%; 
        margin-right: 20%; 
        color: black; 
        font-weight: 400; 
        font-size: 13px;
        padding: 5%;
        border-radius: 15px;
        outline: 0px solid transparent;
        background-color: white;
    }

	/* Text Hover Underline Slide */
	.hover-underline-animation {
		display: inline-block;
		position: relative;
		color: #7c3aed;
	}
	.hover-underline-animation:after {
		content: '';
		position: absolute;
		width: 100%;
		transform: scaleX(0);
		height: 1.5px;
		bottom: 0;
		left: 0;
		background-color: #7c3aed;
		transform-origin: bottom right;
		transition: transform 0.25s ease-out;
	}
	.hover-underline-animation:hover:after {
		transform: scaleX(1);
		transform-origin: bottom left;
	}

    /* The Synonym List Dropdown */
	.dropdown {
		margin-bottom: 3%;
		display: inline-block;
	}
	.dropdown-content {
		display: none;
		position: absolute;
		background-color: white;
		box-shadow: 0.5px 0.5px 20px 0.5px rgba(0, 0, 0, 0.1);
		padding: 12px 16px;
		z-index: 1;
	}
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
