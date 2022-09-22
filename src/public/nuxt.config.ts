import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { NaiveUiResolver } from "unplugin-vue-components/resolvers";
import IconsResolver from "unplugin-icons/resolver";
import Icons from "unplugin-icons/vite";
import UnoCSS from "unocss";
import UnocssIcons from "@unocss/preset-icons";
// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  typescript: {
    strict: true,
  },
  buildModules: ["@nuxtjs/tailwindcss"],
  modules: [
    "@unocss/nuxt",
    "nuxt-graphql-client",
    [
      "@pinia/nuxt",
      {
        autoImports: [
          // automatically imports `defineStore`
          "defineStore", // import { defineStore } from 'pinia'
          // automatically imports `defineStore` as `definePiniaStore`
          ["defineStore", "definePiniaStore"], // import { defineStore as definePiniaStore } from 'pinia'
        ],
      },
    ],
  ],
  runtimeConfig: {
    public: {
      GQL_HOST: "http://localhost:8080/query", // overwritten by process.env.GQL_HOST
    },
  },
  unocss: {
    // presets
    uno: false, // enabled "@unocss/preset-uno"
    icons: true, // enabled "@unocss/preset-icons"
    attributify: false, // enabled "@unocss/preset-attributify"
    // core options
    shortcuts: [],
    rules: [],
  },
  css: ["@/assets/css/main.css"],
  build: {
    transpile:
      process.env.NODE_ENV === "production"
        ? [
            "naive-ui",
            "vueuc",
            "@css-render/vue3-ssr",
            "@juggle/resize-observer",
          ]
        : ["@juggle/resize-observer"],
    postcss: {
      postcssOptions: require("./postcss.config.js"),
    },
  },
  vite: {
    optimizeDeps: {
      include:
        process.env.NODE_ENV === "development"
          ? ["naive-ui", "vueuc", "date-fns-tz/esm/formatInTimeZone"]
          : [],
    },
    plugins: [
      AutoImport({
        imports: ["vue", "@vueuse/core"],
        eslintrc: {
          enabled: true,
        },
      }),
      Components({
        dts: true,
        directives: true,
        directoryAsNamespace: false,
        resolvers: [NaiveUiResolver()],
      }),
    ],
  },
});
