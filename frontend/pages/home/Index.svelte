<script lang="ts">
  import Base from "$layouts/Base.svelte";
  import { router } from "@westacks/inertia-svelte";

  export let random: number;
</script>

<Base>
  <h1>Hi!</h1>
  <p>Random number: {random}</p>

  <button on:click={() => router.reload({ only: ["random"] })} type="button">
    New random
  </button>

  <p>
    * this button uses <a
      href="https://inertiajs.com/partial-reloads"
      target="_blank"
    >
      partial reload
    </a>
    to update only the random number
  </p>
</Base>
