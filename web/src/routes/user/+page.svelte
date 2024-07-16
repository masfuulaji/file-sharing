<script>
    import axios from "axios";

    import Button from "../../components/items/button.svelte";

    const fetchUser = (async () => {
        const response = await axios.get("http://localhost:3000/user");
        return response.data;
    })();
</script>

{#await fetchUser}
    <p>...waiting</p>
{:then data}

<Button url="/user/form" text="Create User" />
<div class="relative overflow-x-auto mt-4">
    <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
                <th scope="col" class="px-6 py-3">
                    Username
                </th>
                <th scope="col" class="px-6 py-3">
                    Email
                </th>
            </tr>
        </thead>
        <tbody>
            {#each data as user}
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                    <td class="px-6 py-4">
                        {user.username}
                    </td>
                    <td class="px-6 py-4">
                        {user.email}
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>

{:catch error}
    <p>An error occurred!</p>
    <p>{error.message}</p>
{/await}
