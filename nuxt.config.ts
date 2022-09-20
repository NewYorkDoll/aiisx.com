import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    buildModules: ['@nuxtjs/tailwindcss'],
    css:[
      "@/assets/css/main.css"
    ],
    build: {
        transpile:
          process.env.NODE_ENV === 'production'
            ? [
                'naive-ui',
                'vueuc',
                '@css-render/vue3-ssr',
                '@juggle/resize-observer'
              ]
            : ['@juggle/resize-observer']
      },
      vite: {
        optimizeDeps: {
          include:
            process.env.NODE_ENV === 'development'
              ? ['naive-ui', 'vueuc', 'date-fns-tz/esm/formatInTimeZone']
              : []
        },
        plugins:[
            AutoImport({
              imports: [
                'vue',
                {
                  'naive-ui': [
                    'useDialog',
                    'useMessage',
                    'useNotification',
                    'useLoadingBar'
                  ]
                }
              ]
            }),
            Components({
              dts: true,
              directives: true,
              directoryAsNamespace: false,
              resolvers: [NaiveUiResolver()]
            })
        ]
      }
})
