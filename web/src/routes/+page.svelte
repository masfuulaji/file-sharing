<script>
  import axios from "axios";

  const fetchImage = (async () => {
    const response = await axios.get("http://localhost:3000/");
    return response.data;
  })();
</script>

{#await fetchImage}
  <p>...waiting</p>
{:then data}
  <p class="underline">{data.message}</p>
{:catch error}
  <p>An error occurred!</p>
  <p>{error.message}</p>
{/await}