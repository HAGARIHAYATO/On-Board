import Vue from "vue";

Vue.mixin({
  methods: {
    GetURL: function() {
      const local = "http://localhost:8080/api/v1";
      return local;
    },
    AlertMessage: function() {
      const message = "画像サイズが大きすぎます。(最大3MB)";
      return message;
    },
    FetchGitInfo: function(platform, token) {
      if (token === "") return ""
      const github = "https://api.github.com/users/"
      let endpoint = ""
      switch (platform) {
        case 0:
        endpoint = github + token + "/repos"
        break;
      }
      return endpoint
    }
  }
});
