import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';
import { TanStackRouterVite } from '@tanstack/router-plugin/vite';
import tsconfigPaths from 'vite-tsconfig-paths';

// https://vitejs.dev/config/
export default ({ _mode }) => {
  const devServerPort = process.env.PROXY_PORT;
  const devServerHost = process.env.PROXY_HOST;
  const devViteServerPort = Number(process.env.VITE_PORT);
  return defineConfig({
    plugins: [react(), tsconfigPaths(), TanStackRouterVite()],
    server: {
      port: devViteServerPort,
      // host: "127.0.0.1",
      strictPort: true,
      proxy: {
        // 外部サーバー
        '/api1': {
          target: 'https://yesno.wtf/api',
          changeOrigin: true,
          secure: false,
          rewrite: (path) => path.replace(/^\/api1/, ''),
        },
        // ローカルの開発環境APIサーバー
        '/api': {
          target: `http://${devServerHost}:${devServerPort}`,
          changeOrigin: true,
          secure: false,
          // rewrite: (path) => path.replace(/^\/api/, ''),
        },
      },
    },
    build: {
      minify: true,
      outDir: 'dist',
    },
  });
};
