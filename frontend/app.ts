import "./app.css";
import { createInertiaApp } from "@westacks/inertia-svelte";

createInertiaApp({
  // @ts-ignore
  resolve: (name) => {
    const pages = import.meta.glob("./pages/**/*.svelte", { eager: true });
    return pages[`./pages/${name}.svelte`];
  },
  setup({ el, App, props }) {
    new App({ target: el, props });
  },
});
