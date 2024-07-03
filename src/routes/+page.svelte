<script>
	import axios from "axios";
    import { Button, Range, Label } from 'flowbite-svelte';

    let conference = "";
    let status = {
        text: "",
        positive: false
    };
    let values = {
        ppa: 50,
    }

    function callBackend() {
        if (!conference) {
            status = {text: "Please enter a code.", positive: false};
            return;
        }
        axios.post(`http://localhost:8080/data/${conference}`, newForm(), {headers: {'Content-Type': 'multipart/form-data'}})
        .then(res => {
            console.log(res);
            status = {text: "Received response.", positive: true};
        }).catch(err => {
            console.log(err);
            status = {text: "Error: " + err.message, positive: false};
        }).finally(() => {
            console.log("Received response.")
            console.log(values)
        });
    }

    function newForm() {
        let formData = new FormData();
        formData.append("weights", JSON.stringify(values));
        return formData;
    }
</script>

<h1>CFB Auto Power Rankings</h1>
<input type="text" bind:value={conference}/>
<p>{conference}</p>
<p style="color: {status.positive ? 'green' : 'red'}">{status.text}</p>
<button on:click={callBackend}>Click to call</button>
<h1 class="text-3xl font-bold overline">
    Hello world!
</h1>
<Button>
    asdasd
</Button>

<Label>Points Per Attempt</Label>
<Range id="ppa" min="0" max="100" bind:value={values.ppa} />
<p>Points per Attempt: {values.ppa}</p>