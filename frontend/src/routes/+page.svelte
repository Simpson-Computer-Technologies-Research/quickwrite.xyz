<script lang="ts">
  // Component Imports
  import RefreshButton from "./components/RefreshButton.svelte";
  import SiteHeader from "./components/SiteHeader.svelte";

  // Main Text Div text editor
  let divTextEditor: any;

  // Word Map holding synonyms
  let wordCache: any = {};

  // splitDivTextInput array holding the div text words
  let splitDivTextInput: string[] = [];

  // The isAllLower() function is used to check
  // whether or not all the characters in the
  // provided string are lowercase.
  //
  // This function is used to prevent passing
  // invalid words to the golang api
  function isAllLower(s: string): boolean {
    // Iterate over the string
    for (let i = 0; i < s.length; i++) {
      // If the string index doesn't equal it's lowercase self
      if (s[i] != s[i].toLowerCase()) {
        return false;
      }
    }
    return true;
  }

  // The GetSynonyms() function is used to send an
  // http request to the golang api. This request
  // returns a response containing all the synonyms
  // for the provided word.
  function GetSynonyms(query: string): void {
    fetch("http://127.0.0.1:8000/synonyms?q=" + query)
      .then((response) => response.json())
      .then((data) => {
        // If the wordCache is too large, delete the first key
        if (wordCache.length > 1000) {
          delete wordCache[Object.keys(wordCache)[0]];
        }

        // If the data is not empty
        if (data.length > 0) {
          wordCache[query] = data
        }
      });
  }

  // The querySynonyms() function is used to get all
  // the words that will be used for getting synonyms.
  function querySynonyms(s: string): void {
    // Iterate over the sliced split
    for (let i = 0; i < s.length; i++) {
      // If the word is all lowercase and it's length is within 4-15
      if (isAllLower(s[i]) && s[i].length >= 4 && s[i].length <= 15) {
        // And the word doesn't already exist in the wordCache
        if (wordCache[s[i]] === undefined) {
          // Get the synonyms for said word (s[i])
          GetSynonyms(s[i]);
        }
      }

      // Set the splitDivTextInput values
      if (s[i].length > 0) {
        splitDivTextInput = [...splitDivTextInput, s[i]];
      }
    }
  }
</script>

<!-- Site Header -->
<SiteHeader />

<!-- Main Text Div -->
<!-- svelte-ignore a11y-autofocus -->
<div contenteditable="true" class="main_text_div" bind:this={divTextEditor}>
  <!-- For each word in the provided text -->
  {#each splitDivTextInput as word, i}
    <!-- If there are no synonyms available -->
    {#if word.length <= 0 || wordCache[word] === undefined}
      {word}
    
    <!-- Otherwise, there are synonyms available -->
    {:else}
      <!-- Create the synonym dropdown div -->
      <div class="dropdown">
        <span class="hover-underline-animation">{word} </span>
        <div class="dropdown-content">
          <!-- For every synonym in the queries wordCache array -->
          {#each wordCache[word] as synonym, n}
            <div style="margin-bottom: 5px;">
              <a 
                href="#{word}"
                style="cursor: pointer; color: #7c3aed; margin-bottom: 5px;"
                on:click={() => {
                  splitDivTextInput[i] = wordCache[word][n];
                  GetSynonyms(wordCache[word][n]);
                }}>{synonym}</a>
            </div>
          {/each}
        </div>
      </div>
    {/if}
  {/each}
</div>

<!-- Refresh Button -->
<RefreshButton {divTextEditor} {querySynonyms} />

<style>
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
    content: "";
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
