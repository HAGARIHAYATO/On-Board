export default {
  mode: "universal",
  /*
   ** Headers of the page
   */
  head: {
    title: process.env.npm_package_name || "",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "description",
        content: process.env.npm_package_description || ""
      }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },
  /*
   ** Customize the progress-bar color
   */
  loading: {
    color: "yellowgreen",
    height: "6px"
  },
  /*
   ** Global CSS
   */
  css: [{ src: "~/assets/scss/app.scss", lang: "scss" }],
  /*
   ** Plugins to load before mounting the App
   */
  // nuxtAuth
  axios: {
    baseURL: process.env.API_URL,
    browserBaseURL: process.env.API_URL_BROWSER
  },
  auth: {
    redirect: {
      login: "/",
      logout: "/",
      callback: false,
      home: "/"
    },
    strategies: {
      local: {
        endpoints: {
          login: {
            url: "http://localhost:8080/api/v1/login",
            method: "post",
            propertyName: "Token"
          },
          user: {
            url: "http://localhost:8080/api/v1/credential",
            method: "get",
            propertyName: false
          },
          logout: false
        }
      }
    }
  },
  proxy: {
    "/api": {
      target: "http://localhost:8080",
      pathRewrite: {
        "^/api": "/api"
      }
    }
  },
  // nuxtAuth
  plugins: ["@/plugins/common"],
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [],
  /*
   ** Nuxt.js modules
   */
  modules: ["@nuxtjs/axios", "@nuxtjs/auth", "@nuxtjs/proxy"],
  env: {
    API_URL: "http://api:8080/",
    API_URL_BROWSER: "http://localhost:8080"
  },
  /*
   ** Build configuration
   */
  build: {
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {}
  }
};
