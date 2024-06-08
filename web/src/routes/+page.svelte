<script>
  const fetchImage = (async () => {
    const response = await fetch("http://localhost:3000/json");
    return await response.json();
  })();
</script>

{#await fetchImage}
  <p>...waiting</p>
{:then data}
  <p>{data.message}</p>
  <p>{data.status}</p>
{:catch error}
  <p>An error occurred!</p>
{/await}
