import Vue from "vue";

Vue.mixin({
  methods: {
    GetURL: function() {
      const local = "/api/v1";
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
    }
  }
});
