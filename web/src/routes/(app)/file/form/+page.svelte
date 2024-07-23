<script>
  import { page } from "$app/stores";
  import axios from "axios";
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  /**
   * @type {string | null}
   */
  let id;

  $: {
    const url = new URL($page.url);
    id = url.searchParams.get("id");
  }

  onMount(async () => {
    if (id) {
      const response = await axios.get(`http://localhost:3000/file/${id}`);
    }
  });

  async function saveUser() {
    const fileInput = document.getElementById("file");
    // @ts-ignore
    const file = fileInput.files[0];

    const formData = new FormData();
    formData.append("file", file);
    const user = formData;
    if (id) {
      const response = await axios.put(
        `http://localhost:3000/user/${id}`,
        user,
      );
    } else {
      const response = await axios.post(
        "http://localhost:3000/upload",
        formData,
        {
          withCredentials: true,
          headers: {
            "Content-Type": "multipart/form-data",
          },
        },
      );
    }
    alert("Saved");
    goto("/file");
  }
</script>

<form class=" p-4 mx-auto">
  <div class="mb-5">
    <label for="email" class="block mb-2 text-sm font-medium text-gray-900"
      >File Upload</label
    >
    <input
      type="file"
      id="file"
      class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
      required
    />
  </div>
  <button
    type="submit"
    on:click={() => saveUser()}
    class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
    >Save</button
  >
  <button
    type="button"
    on:click={() => goto("/user")}
    class="focus:outline-none text-white bg-purple-700 hover:bg-purple-800 focus:ring-4 focus:ring-purple-300 font-medium rounded-lg text-sm px-5 py-2.5 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
    >Back</button
  >
</form>
