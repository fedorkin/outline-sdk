import { defineConfig } from "vite";

export default defineConfig({
  build: {
    outDir: './output/frontend',
  },
  server: {
    host: true
  }
})
