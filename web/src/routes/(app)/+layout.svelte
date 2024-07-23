<script>
  import Sidebar from "../../components/partials/sidebar.svelte";
  import Header from "../../components/partials/header.svelte";
  import { onMount } from "svelte";
  import axios from "axios";
  import { goto } from "$app/navigation";

  onMount(async () => {
    await axios
      .get("http://localhost:3000/isLogin", { withCredentials: true })
      .catch((error) => {
        goto("/login");
      });
  });
</script>

<div>
  <div class="flex h-screen bg-gray-200">
    <div
      class="fixed inset-0 z-20 transition-opacity bg-black opacity-50 lg:hidden"
    ></div>
    <Sidebar />

    <div class="flex flex-col flex-1 overflow-hidden">
      <Header />
      <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-200">
        <div class="container py-6 px-6">
          <slot />
        </div>
      </main>
    </div>
  </div>
</div>

