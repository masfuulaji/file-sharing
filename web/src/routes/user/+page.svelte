<script>
    import axios from "axios";

    import ButtonPrimary from "../../components/items/buttonPrimary.svelte";
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";

    /**
     * @type {any[]}
     */
    let users = [];

    onMount(async () => {
        users = await fetchUser();
    });

    const fetchUser = async () => {
        const response = await axios.get("http://localhost:3000/user");
        return response.data;
    };

    // @ts-ignore
    const editUser = (id) => {
        goto(`/user/form?id=${id}`);
    };

    // @ts-ignore
    const deleteUser = async (id) => {
        await axios.delete(`http://localhost:3000/user/${id}`);
        users = await fetchUser();
    };
</script>

<ButtonPrimary url="/user/form" text="Create User" />
<div class="relative overflow-x-auto mt-4">
    <table
        class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400"
    >
        <thead
            class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
        >
            <tr>
                <th scope="col" class="px-6 py-3"> Username </th>
                <th scope="col" class="px-6 py-3"> Email </th>
                <th scope="col" class="px-6 py-3"> Action </th>
            </tr>
        </thead>
        <tbody>
            {#if users != null && users.length > 0}
                {#each users as user}
                    <tr
                        class="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
                    >
                        <td class="px-6 py-4">
                            {user.username}
                        </td>
                        <td class="px-6 py-4">
                            {user.email}
                        </td>
                        <td class="px-6 py-4">
                            <button
                                type="button"
                                on:click={() => editUser(user.id)}
                                class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                                >Edit</button
                            >
                            <button
                                type="button"
                                on:click={() => deleteUser(user.id)}
                                class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900"
                                >Delete</button
                            >
                        </td>
                    </tr>
                {/each}
            {:else}
                <tr
                    class="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
                >
                    <td colspan="3" class="px-6 py-4 text-center"> No data </td>
                </tr>
            {/if}
        </tbody>
    </table>
</div>
