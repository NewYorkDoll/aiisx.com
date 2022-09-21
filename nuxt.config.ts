import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { NaiveUiResolver } from "unplugin-vue-components/resolvers";
import IconsResolver from "unplugin-icons/resolver";
import Icons from "unplugin-icons/vite";

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  typescript: {
    strict: true,
  },
  buildModules: ["@nuxtjs/tailwindcss"],
  modules: [
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
        resolvers: [
          IconsResolver({
            componentPrefix: "icon",
            enabledCollections: ["mdi", "logos"],
          }),
        ],
        eslintrc: {
          enabled: true,
        },
      }),
      Icons({
        autoInstall: true,
        defaultClass: "icon",
      }),
      Components({
        dts: true,
        directives: true,
        directoryAsNamespace: false,
        resolvers: [
          NaiveUiResolver(),
          IconsResolver({
            componentPrefix: "i",
            enabledCollections: ["mdi", "logos"],
          }),
        ],
      }),
    ],
  },
});
