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
  env: {
    API_URL: "http://backend:8080",
    API_URL_BROWSER: "http://localhost:8080"
  },
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
            url: "/api/v1/login",
            method: "post",
            propertyName: "Token"
          },
          user: {
            url: "/api/v1/credential",
            method: "get",
            propertyName: false
          },
          logout: false
        }
      }
    }
  },
  proxy: {
    "/api":
    process.env.NODE_ENV === 'development'
    ? {
      target: "http://backend:8080",
      pathRewrite: {
        "^/api": "/api"
      }
    }
    : {
      "*": "http://[::1]:8080"
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
