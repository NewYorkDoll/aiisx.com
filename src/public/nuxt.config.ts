import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { NaiveUiResolver } from "unplugin-vue-components/resolvers";
import presetIcons from "@unocss/preset-icons";

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  typescript: {
    strict: true,
  },
  app: {
    head: {
      title: "Website of 绎紫洛英",
      link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.png" }],
      script: [
        {
          src: "https://umami.aiisx.com/umami.js",
          async: true,
          defer: true,
          "data-website-id": "f9b403eb-6d5f-4712-9ef3-50f73ac015c3",
        },
      ],
    },
  },
  plugins: [{ src: "@/plugins/baidu", mode: "client" }],
  modules: [
    "@nuxtjs/tailwindcss",
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
      GQL_HOST: "https://aiisx.com/query", // overwritten by process.env.GQL_HOST
      gql: {
        watch: true,
      },
    },
  },
  unocss: {
    // presets
    uno: false, // enabled "@unocss/preset-uno"
    icons: true, // enabled "@unocss/preset-icons"
    attributify: false, // enabled "@unocss/preset-attributify"
    // core options
    shortcuts: [],
    safelist: [
      "i-mdi-home",
      "i-mdi-book-open-page-variant-outline",
      "i-mdi-image-edit-outline",
      "i-mdi-link",
      "i-mdi-logout",
    ],
    rules: [],
    presets: [
      presetIcons({
        extraProperties: {
          display: "inline-block",
          "vertical-align": "middle",
          // ...
        },
        collections: {
          mdi: () =>
            import("@iconify-json/mdi/icons.json").then((i) => i.default),
        },
      }),
    ],
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
  },
  webpack: {
    postcss: {
      postcssOptions: require("./postcss.config.js"),
    },
  },
  vite: {
    ssr: {
      noExternal: ["codemirror"],
    },
    optimizeDeps: {
      exclude: ["fsevents"],
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
