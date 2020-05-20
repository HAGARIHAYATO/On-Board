const domain = "https://on-board-project.com/"
const targetID = "UA-165370681-1"
import axios from 'axios';
export default {
  mode: "universal",
  /*
   ** Headers of the page
   */
  generate: {
    fallback: true,
    // async routes () {
    //   const generates = []
    //   await axios.get("https://api.on-board-project.com/api/v1/works_ids")
    //     .then((res) => {
    //       res.data.IDs.map((work) => {
    //         generates.push('works/'+work)
    //       })
    //     })
    //   await axios.get("https://api.on-board-project.com/api/v1/users_ids")
    //     .then((res) => {
    //       res.data.IDs.map((user) => {
    //         generates.push('users/'+user)
    //       })
    //     })
    //   return generates
    // }
  },
  head: {
    title: "OnBoard | みんなの作品・ポートフォリオ置場",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "OnBoard | みんなの作品・ポートフォリオ置場",
        content: "ポートフォリオを投稿して評価を受けたりスキルや作品のCSV資料を取得したりしよう。"
      }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/logo.jpeg" }]
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
    API_URL: "https://api.on-board-project.com",
    API_URL_BROWSER: "https://api.on-board-project.com"
  },
  axios: {
    baseURL: process.env.API_URL | "https://api.on-board-project.com",
    browserBaseURL: process.env.API_URL_BROWSER | "https://api.on-board-project.com"
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
            url: "https://api.on-board-project.com/api/v1/login",
            method: "post",
            propertyName: "Token"
          },
          user: {
            url: "https://api.on-board-project.com/api/v1/credential",
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
      target: "https://api.on-board-project.com",
      pathRewrite: {
        "^/api": "/"
      }
    }
    : {
      target: "https://api.on-board-project.com",
      pathRewrite: {
        "^/api": "/"
      }
    }
  },
  // nuxtAuth
  plugins: [{ src: "~/plugins/common", ssr: false }],
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    ['@nuxtjs/google-adsense', { 
      id: 'ca-pub-4712100556499049',
      pageLevelAds: true,
      analyticsUacct: targetID,
      analyticsDomainName: domain 
    }],
    ['@nuxtjs/google-analytics', {
      id: targetID
    }]
  ],
  /*
   ** Nuxt.js modules
   */
  modules: ["@nuxtjs/axios", "@nuxtjs/auth", "@nuxtjs/proxy", 
    ['@nuxtjs/google-adsense', { 
      id: 'ca-pub-4712100556499049',
      pageLevelAds: true,
      analyticsUacct: targetID,
      analyticsDomainName: domain
    }]
  ],
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
