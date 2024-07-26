<script>
  import axios from "axios";
  import { page } from "$app/stores";

  import ButtonPrimary from "../../../../components/items/buttonPrimary.svelte";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  /**
   * @type {any[]}
   */
  let permissions = [];
  let file;
  // @ts-ignore
  let is_public;

  // @ts-ignore
  let id;

  $: {
    const url = new URL($page.url);
    id = url.searchParams.get("id");
  }

  onMount(async () => {
    permissions = await fetchPermission();
    file = await getFile();
    if (file.is_public) {
      is_public = "Public";
    } else {
      is_public = "Private";
    }
  });

  const getFile = async () => {
    const response = await axios.get(`http://localhost:3000/file/${id}`);
    return response.data;
  };
  const fetchPermission = async () => {
    const response = await axios.get(
      `http://localhost:3000/permission/user/${id}`,
    );
    return response.data;
  };

  const switchPublic = async () => {
    const response = await axios.put(
      `http://localhost:3000/file/switch_public/${id}`,
    );
    goto("/file");
  };

  const copyLink = () => {
    navigator.clipboard
      .writeText(`http://localhost:3000/download/${file.id}`)
      .then(() => {
        alert("Link copied");
      });
  };

  // @ts-ignore
  const editUser = (id) => {
    goto(`/user/form?id=${id}`);
  };

  // @ts-ignore
  const deleteUser = async (id) => {
    await axios.delete(`http://localhost:3000/user/${id}`);
    permissions = await fetchPermission();
  };
</script>

<div>
  <ButtonPrimary url="/file/permission/form" text="Create Permission" />
  {#if is_public == "Public"}
    <button
      type="button"
      on:click={() => switchPublic()}
      class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
    >
      Public
    </button>
    <button
      type="button"
      on:click={() => copyLink()}
      class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
    >
      Copy Link
    </button>
  {:else}
    <button
      type="button"
      on:click={() => switchPublic()}
      class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900"
      >Private</button
    >
  {/if}
  <div class="relative overflow-x-auto mt-4">
    <table
      class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400"
    >
      <thead
        class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
      >
        <tr>
          <th scope="col" class="px-6 py-3"> Name </th>
          <th scope="col" class="px-6 py-3"> Permission </th>
          <th scope="col" class="px-6 py-3"> Action </th>
        </tr>
      </thead>
      <tbody>
        {#if permissions != null && permissions.length > 0}
          {#each permissions as permission}
            <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
              <td class="px-6 py-4">
                {permission.username}
              </td>
              <td class="px-6 py-4">
                <button
                  type="button"
                  on:click={() => editUser(permission.id)}
                  class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                  >Edit</button
                >
                <button
                  type="button"
                  on:click={() => deleteUser(permission.id)}
                  class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900"
                  >Delete</button
                >
              </td>
            </tr>
          {/each}
        {:else}
          <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
            <td colspan="3" class="px-6 py-4 text-center"> No data </td>
          </tr>
        {/if}
      </tbody>
    </table>
  </div>
</div>
