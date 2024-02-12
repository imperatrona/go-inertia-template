import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  publicDir: "./frontend/public",
  resolve: {
    alias: {
      $components: path.resolve(__dirname, "./frontend/components"),
      $layouts: path.resolve(__dirname, "./frontend/layouts/"),
    },
  },
});
