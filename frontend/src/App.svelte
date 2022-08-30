<script>
	// Track input last key up
	var queryTimer;
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
		console.log(1)
		// Clear previous timer
		clearTimeout(queryTimer);

		queryTimer = setTimeout(() => {
			// Store the previous split length
			let splitPreviousLength = splitInput.length

			// Split the input
			splitInput = s.split(" ")

			// Slice the split input to avoid duplicates
			let _split = splitInput.slice(splitPreviousLength)

			// Iterate over the sliced split
			for (let i = 0; i < _split.length; i++) {

				// If the word is all lowercase and it's length is within 4-15
				if (isAllLower(_split[i]) && _split[i].length >= 4 && _split[i].length <= 15) {

					// And the word doesn't already exist in the wordMap
					if (wordMap[_split[i]] === undefined) {

						// Get the synonyms for said word (_split[i])
						GetSynonyms(_split[i]);
					}
				}
			}
		}, 320); // 320 milliseconds
	}

</script>

<main>
	<div style="width:50%; margin:0 auto;">
		<!-- svelte-ignore a11y-autofocus -->
		<input
			type="text"
			class="text_input"
			placeholder="Text"
			on:keyup={({ target: { value } }) => Query(value)} 
		/>
	</div>

	<div oninput="Query()" style="margin-left: 20%; margin-right: 20%" contenteditable>
		{#each splitInput as word, i}
			<div class="dropdown">
				{#if wordMap[word] !== undefined && wordMap[word].length > 0}
					<span style="color: blue;">{word}&nbsp;</span>
					<div class="dropdown-content">
						{#each wordMap[word] as synonym, n}
							<div style="margin-bottom: 5px">
								<!-- svelte-ignore a11y-invalid-attribute -->
								<a href="#" on:click={() => {
									splitInput[i] = wordMap[word][n]
								}}>{synonym}</a>
							</div>
						{/each}
					</div>
				{:else}
					<span>{word}&nbsp;</span>
				{/if}
			</div>
		{/each}
	</div>

</main>

<style>
	.text_input {
		width: 80%;
		height: 50%;
		padding-bottom: 30%;
	}
	.dropdown {
		position: relative;
		display: inline-block;
	}

	.dropdown-content {
		display: none;
		position: absolute;
		background-color: #f9f9f9;
		box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
		padding: 12px 16px;
		z-index: 1;
	}

	.dropdown:hover .dropdown-content {
		display: block;
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
