import Vue from "vue";

Vue.mixin({
  methods: {
    calcSize: function() {
      if (window.innerWidth > 400) {
        return "400px"
      } else {
        return window.innerWidth * 0.9 + "px"
      }
    },
    GetURL: function() {
      // const local = "/api/v1";
      const local = "https://api.on-board-project.com/api/v1";
      return local;
    },
    AlertMessage: function() {
      const message = "画像サイズが大きすぎます。(最大3MB)";
      return message;
    },
    FetchGitInfo: function(author, token) {
      if (token === "") return ""
      const github = "https://api.github.com"
      const user = "/users/"
      const orgs = "/orgs/"
      let endpoint = ""
      switch (author) {
        case 0:
        endpoint = github + user + token + "/repos"
        break;
        case 1:
        endpoint = github + orgs + token + "/repos"
        break;
      }
      return endpoint
    },
    RewriteTime: function(time) {
      if (time) {
        let fixedTime = ""
        fixedTime = time.split("T")
        fixedTime = fixedTime[0].split("-")
        fixedTime = fixedTime[0] + "年" + fixedTime[1] + "月" + fixedTime[2] + "日"
        return fixedTime
      }
    }
  }
});
