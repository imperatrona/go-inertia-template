<script lang="ts">
  import { Link } from "@westacks/inertia-svelte";
</script>

<nav class="nav">
  <Link href="/">home</Link>
  <Link href="/about">about</Link>
</nav>

<style>
  .nav {
    display: flex;
    gap: 8px;
  }
</style>
