import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    proxy: {
      '/devstoreaccount1': {
        target: 'http://127.0.0.1:10001',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/devstoreaccount1/, ''),
        configure: (proxy, options) => {
          proxy.on('proxyReq', (proxyReq, req, res) => {
            proxyReq.setHeader('origin', 'http://localhost:5173');
          });
        },
      },
    },
  },
});
