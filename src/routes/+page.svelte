<script>
	import axios from "axios";
    import { Button, Checkbox, Dropdown, DropdownItem, Range, Label } from 'flowbite-svelte';

    const conferences = [
        {name: "ACC", code: "ACC", type: "P4", checked: false},
        {name: "Big 12", code: "B12", type: "P4", checked: false},
        {name: "Big Ten", code: "B1G", type: "P4", checked: false},
        {name: "SEC", code: "SEC", type: "P4", checked: false},
        {name: "PAC-12", code: "PAC", type: "FBS", checked: false},
        {name: "American Athletic", code: "AAC", type: "FBS", checked: false},
        {name: "Conference USA", code: "CUSA", type: "FBS", checked: false},
        {name: "Mid-American", code: "MAC", type: "FBS", checked: false},
        {name: "Mountain West", code: "MWC", type: "FBS", checked: false},
        {name: "Sun Belt", code: "SBC", type: "FBS", checked: false},
        {name: "FBS Independents", code: "Ind", type: "FBS", checked: false}
    ];
    let conference = "";
    let status = {
        text: "",
        positive: false
    };
    const teamStats = {
        kickReturns: 0,
        games: 0,
        passingTDs: 0,
        rushingTDs: 0,
        interceptionYards: 0,
        tacklesForLoss: 0,
        kickReturnTDs: 0,
        rushingYards: 0,
        fourthDownConversions: 0,
        possessionTime: 0,
        penalties: 0,
        puntReturnYards: 0,
        totalYards: 0,
        interceptionTDs: 0,
        puntReturnTDs: 0,
        kickReturnYards: 0,
        firstDowns: 0,
        sacks: 0,
        passesIntercepted: 0,
        puntReturns: 0,
        fumblesLost: 0,
        passCompletions: 0,
        netPassingYards: 0,
        fourthDowns: 0,
        turnovers: 0,
        passAttempts: 0,
        interceptions: 0,
        fumblesRecovered: 0,
        thirdDowns: 0,
        penaltyYards: 0,
        thirdDownConversions: 0,
        rushingAttempts: 0
    };
    let response = [];

    function callBackend() {
        console.log(conferences);
        if (!conference) {
            status = {text: "Please enter a code.", positive: false};
            return;
        }
        axios.post(`http://localhost:8080/data/${conference}`, newForm(), {headers: {'Content-Type': 'multipart/form-data'}})
        .then(res => {
            response = res.data;
            console.log(response)
            status = {text: "Received response.", positive: true};
        }).catch(err => {
            console.log(err);
            status = {text: "Error: " + err.message, positive: false};
        }).finally(() => {
            console.log("Received response.")
        });
    }

    function newForm() {
        let formData = new FormData();
        formData.append("weights", JSON.stringify(teamStats));
        return formData;
    }
</script>

<h1>CFB Auto Power Rankings</h1>
<button class="bg-gray-600">Dropdown checkbox</button>
<Dropdown class="w-44 p-3 space-y-3 text-sm bg-gray-600">
    {#each conferences as conference, i}
        <li id={conference.code} class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
            <Checkbox on:change={() => conferences[i].checked = !conferences[i].checked}>{conference.name}</Checkbox>
        </li>
    {/each}
</Dropdown>

<p style="color: {status.positive ? 'green' : 'red'}">{status.text}</p>
<button on:click={callBackend}>Click to call</button>

{#if response}
    <ul class="max-w-md space-y-1 list-decimal list-inside">
    {#each response as team}
        <li>{team.Team} (Score: {team.Score})</li>
    {/each}
    </ul>
{:else}
    <p>Waiting for response.</p>
{/if}

<ul>
    {#each new Map(Object.entries(teamStats)) as stat}
    <Label>{stat[0]}</Label>
    <Range id={stat[0]} min="0" max="100" bind:value={teamStats[stat[0]]}/>    
    <li>{stat[1]}</li>
    {/each}
</ul>
