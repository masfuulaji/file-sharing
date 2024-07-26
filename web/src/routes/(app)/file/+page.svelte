<script>
  import axios from "axios";

  import ButtonPrimary from "../../../components/items/buttonPrimary.svelte";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  /**
   * @type {any[]}
   */
  let files = [];

  onMount(async () => {
    files = await fetchFile();
  });

  const fetchFile = async () => {
    const response = await axios.get("http://localhost:3000/file");
    return response.data;
  };

  // @ts-ignore
  const editFile = (id) => {
    goto(`/file/form?id=${id}`);
  };

  // @ts-ignore
  const permissionFile = (id) => {
    goto(`/file/permission?id=${id}`);
  };

  // @ts-ignore
  const deleteFile = async (id) => {
    await axios.delete(`http://localhost:3000/file/${id}`);
    files = await fetchFile();
  };

  // @ts-ignore
  const downloadFile = async (filePath) => {
    await axios
      .get(`http://localhost:3000/download/${filePath}`, {
        withCredentials: true,
        responseType: "blob",
      })
      .then((response) => {
        // create file link in browser's memory
        const href = URL.createObjectURL(response.data);

        // create "a" HTML element with href to file & click
        const link = document.createElement("a");
        link.href = href;
        link.setAttribute("download", "file.pdf"); //or any other extension
        document.body.appendChild(link);
        link.click();

        // clean up "a" element & remove ObjectURL
        document.body.removeChild(link);
        URL.revokeObjectURL(href);
      })
      .catch((err) => {
        alert("Unauthorized");
      });
  };
</script>

<ButtonPrimary url="/file/form" text="Create File" />
<div class="relative overflow-x-auto mt-4">
  <table
    class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400"
  >
    <thead
      class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
    >
      <tr>
        <th scope="col" class="px-6 py-3"> Name </th>
        <th scope="col" class="px-6 py-3"> Action </th>
      </tr>
    </thead>
    <tbody>
      {#if files != null && files.length > 0}
        {#each files as file}
          <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
            <td class="px-6 py-4" width="50%">
              {file.file_name}
            </td>
            <td class="px-6 py-4">
              <button
                type="button"
                on:click={() => downloadFile(file.id)}
                class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >Download</button
              >
              <button
                type="button"
                on:click={() => permissionFile(file.id)}
                class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >Permission</button
              >
              <button
                type="button"
                on:click={() => editFile(file.id)}
                class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >Edit</button
              >
              <button
                type="button"
                on:click={() => deleteFile(file.id)}
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
