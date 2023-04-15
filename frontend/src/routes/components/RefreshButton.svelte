<script lang="ts">
    // Variable Imports
    export let divTextEditor: any;
    export let querySynonyms: any;

    // Error message variable
    let errorMessage: string = "";
</script>

<main>
	<div class="container">
        <div class="btn">
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <div class="get_name" 
				on:click={() => {
                    // Make sure the content is under or equal to 750 characters
                    if (divTextEditor.innerText.length <= 750) {
                        // Reset the error message
                        errorMessage = "";

                        // Store current innerText
					    let c = divTextEditor.innerText;

                        // Reset the innerText
                        divTextEditor.innerHTML = "";

                        // Query for synonyms using previous innerText
                        querySynonyms(c.split(" "));
                    } 

                    // Set the error message
                    else {
                        errorMessage = `Too many characters! ${divTextEditor.innerText.length}/750`
                    }
				}}
			> Refresh</div>
            <div class="snowflake-grid to-left">
                <div class="snowflake-item border-bottom border-right">
                    <div class="sub-items border-right border-bottom pull-down">
                    <div class="m-w-15 m-h-15 border-right border-bottom m-3"></div>
                    </div>
                </div>
                <div class="snowflake-item border-bottom border-left">
                    <div class="sub-items border-right border-bottom r-90 pull-down-right">
                    <div class="m-w-15 m-h-15 border-right border-bottom m-3"></div>
                    </div>
                </div>
                <div class="snowflake-item border-top border-right">
                    <div class="sub-items border-right border-bottom r-270 pull-right">
                    <div class="m-w-15 m-h-15 border-right border-bottom m-3"></div>
                    </div>
                </div>
                <div class="snowflake-item border-top border-left">
                    <div class="sub-items border-right border-bottom r-180 pull-left">
                    <div class="m-w-15 m-h-15 border-right border-bottom m-3"></div>
                    </div>
                </div>
            </div>
            <div class="snowflake-grid to-right">
                <div class="snowflake-item border-bottom border-right">
                    <div class="sub-items border-right border-bottom pull-down">
                    <div class="m-w-15 m-h-15 border-right border-bottom m-3"></div>
                    </div>
                </div>
                <div class="snowflake-item border-bottom border-left">
                    <div class="sub-items border-right border-bottom r-90 pull-down-right">
                    <div class="m-w-15 m-h-15 border-right border-bottom m-3"></div>
                    </div>
                </div>
                <div class="snowflake-item border-top border-right">
                    <div class="sub-items border-right border-bottom r-270 pull-right">
                    <div class="m-w-15 m-h-15 border-right border-bottom m-3"></div>
                    </div>
                </div>
                <div class="snowflake-item border-top border-left">
                    <div class="sub-items border-right border-bottom r-180 pull-left">
                    <div class="m-w-15 m-h-15 border-right border-bottom m-3"></div>
                    </div>
                </div>
            </div>
        </div>
        </div>
        <div class="container">
            <p style="color: #ff5353">{errorMessage}</p>
        </div>
</main>

<style>
    .container {
        display: flex;
        justify-content: center;
        margin: 0 auto;
    }
    .btn {
        margin-top: 3%;
        display: flex;
        justify-content: center;
        align-items: center;
        width: 220px;
        height: 60px;
        background-color: var(--light-theme);
        cursor: pointer;
        font-size: 24px;
        color: var(--color);
        transition: all 0.3s;
        position: absolute;
        text-align: center;
        overflow: hidden;
        border-radius: 5px;
        box-shadow: 0 6px 30px -10px rgba(204, 204, 204, 1);
    }
    .btn:hover {
        transform: translateX(5px) translateY(-7px);
    }
    .get_name {
        position: relative;
    }
    .snowflake-grid {
        display: inline-grid;
        grid-template-columns: 1fr 1fr;
        grid-auto-rows: 50px;
        gap: 5px;
        width: 100px;
        transform: scale(0.4);
    }
    .snowflake-grid div {
        border-radius: 5px;
    }
    .snowflake-grid .snowflake-item {
        position: relative;
    }
    .to-left {
        position: absolute;
        top: -4px;
        left: -35px;
        animation: swing 3s infinite;
    }
    .to-left div {
        animation: flash 3s infinite;
    }
    .to-right {
        position: absolute;
        top: -25px;
        right: -35px;
        animation: swing 2.5s infinite;
    }
    .to-right div {
        animation: flash 2.5s infinite;
    }
    .border-left {
        border-left: 4px solid #5858ff;
    }
    .border-right {
        border-right: 4px solid #5858ff;
    }
    .border-bottom {
        border-bottom: 4px solid #5858ff;
    }
    .border-top {
        border-top: 4px solid #5858ff;
    }
    .sub-items {
        height: 28px;
        width: 28px;
    }
    .m-w-15 {
        width: 15px;
    }
    .m-h-15 {
        height: 15px;
    }
    .r-270 {
        transform: rotate(270deg);
    }
    .r-180 {
        transform: rotate(180deg);
    }
    .r-90 {
        transform: rotate(90deg);
    }
    .pull-down {
        position: absolute;
        bottom: 5px;
        right: 5px;
    }
    .pull-down-right {
        position: absolute;
        bottom: 5px;
        left: 5px;
    }
    .pull-right {
        position: absolute;
        right: 5px;
        top: 5px;
    }
    .pull-left {
        position: absolute;
        left: 5px;
        top: 5px;
    }
    .m-3 {
        margin: 3px;
    }
    @keyframes swing {
        50% {
            transform: rotateZ(10deg) scale(0.4);
    }
    }
    @keyframes flash {
        50% {
            border-color: #485563;
    }
    }
</style>
