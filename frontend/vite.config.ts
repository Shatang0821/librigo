import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    watch: {
      usePolling: true, // ファイル変更を定期的にチェックする設定
    },
    host: true, // Docker コンテナ外からのアクセスを許可
    strictPort: true,
    port: 5173,
  },
})
