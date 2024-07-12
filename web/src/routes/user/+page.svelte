<script>
    import axios from "axios";

    const fetchUser = (async () => {
        const response = await axios.get("http://localhost:3000/user");
        return response.data;
    })();
</script>

{#await fetchUser}
    <p>...waiting</p>
{:then data}
    {#each data as user}
        <p>{user.username}</p>
    {/each}
{:catch error}
    <p>An error occurred!</p>
    <p>{error.message}</p>
{/await}
