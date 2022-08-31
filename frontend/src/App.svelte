<script>
	import RefreshButton from "./components/RefreshButton.svelte";

	var wordMap = {}
	let splitInput = [];

	function isAllLower(s) {
		for (let i = 0; i < s.length; i++) {
			if (s[i] != s[i].toLowerCase()) {
				return false
			}
		}
		return true
	}

	function GetSynonyms(query) {
		// Send the http request to the golang api
		fetch("http://127.0.0.1:8000?q=" + query)
			.then((response) => response.json())
			.then((data) => {
				wordMap[query] = data;
			})
	}

	function Query(s) {
		// Iterate over the sliced split
		for (let i = 0; i < s.length; i++) {
			// If the word is all lowercase and it's length is within 4-15
			if (isAllLower(s[i]) && s[i].length >= 4 && s[i].length <= 15) {
				// And the word doesn't already exist in the wordMap
				if (wordMap[s[i]] === undefined) {
					// Get the synonyms for said word (s[i])
					GetSynonyms(s[i]);
				}
			}
			if (s[i].length > 0) {
				splitInput = [...splitInput, s[i]]
			}
		}
	}
	let editor;
</script>

<main>
	<!-- svelte-ignore a11y-autofocus -->
	<div
		autofocus="true"
		contenteditable="true"
		bind:this = {editor}
		style="
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
		"
	>
		{#each splitInput as word, i}
			{#if word.length > 0 && wordMap[word] !== undefined && wordMap[word].length > 0}
				<div class="dropdown">
				<span style="color: #6366f1; font-weight: 800; font-size: 20px;">{word} </span>
				<div class="dropdown-content">
					{#each wordMap[word] as synonym, n}
						<div style="margin-bottom: 5px">
							<!-- svelte-ignore a11y-invalid-attribute -->
							<a style="color: #6366f1; font-weight: 500; font-size: 16px; cursor: pointer;" href="#" on:click={() => {
								splitInput[i] = wordMap[word][n]
								GetSynonyms(wordMap[word][n])
							}}>{synonym}</a>
						</div>
					{/each}
				</div>
			</div>
			{:else}
				<span style="color: white; font-weight: 700; font-size: 20px;">{word} </span>
			{/if}
		{/each}
	</div>
	<RefreshButton Query={Query} editor={editor}/>
</main>

<style>
	.dropdown {
		display: inline-block;
	}

	.dropdown-content {
		display: none;
		position: absolute;
		background-color: #FDFDFD;
		padding: 12px 16px;
		z-index: 1;
	}

	.dropdown:hover .dropdown-content {
		display: block;
		border-radius: 15px;
	}

	:global(body) {
		background-color: #2B2B2B;
	}

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
</style>
