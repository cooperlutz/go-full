import type {} from 'vitest/config';
import { fileURLToPath, URL } from 'node:url';

import tailwindcss from '@tailwindcss/vite';
import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import autoprefixer from 'autoprefixer';
import { defineConfig } from 'vite';
import { configDefaults } from 'vitest/config'

export default defineConfig({
  test: {
    globals: true,
    environment: 'happy-dom',
    coverage: {
      reporter: ['text', 'html'],
      include: [
        'src/**/*.{js,ts,vue}' // Include all Vue, JS, and TS files in src
      ],
      exclude: [
        ...configDefaults.coverage.exclude ?? [],
        '**/api/**',
      ],
      thresholds: {
        statements: 80,
        branches: 80,
        functions: 80,
        lines: 80,
      },
    },
    reporters: 'dot',
    exclude: [
      ...(configDefaults.exclude), 
      '**/api/**',
    ]
  },
  base: '/',
  appType: 'spa',
  plugins: [tailwindcss(), vue(), vueJsx()],
  css: {
    postcss: {
      plugins: [autoprefixer()],
    },
  },
  resolve: {
    alias: {
      '~': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
},);
