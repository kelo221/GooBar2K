import { defineConfig } from 'vite'
import TailwindCss from '@tailwindcss/vite'
import Vue from '@vitejs/plugin-vue'
import * as CompilerSFC from '@vue/compiler-sfc'
import * as path from "node:path";
export default defineConfig({
  plugins: [
    Vue({
      compiler: CompilerSFC,
      include: [/\.vue$/, /\.md$/],
    }),
    TailwindCss(),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
    },
  },
})
