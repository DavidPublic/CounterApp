import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { UserConfig } from 'vite';

export default defineConfig({
  plugins: [react()],
  test: {
    globals: true,
    environment: 'jsdom',
  },
} as UserConfig);
