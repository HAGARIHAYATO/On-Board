<template>
  <div class="container" v-if="reload">
    <Loading v-if="isLoading" />
    <div class="container__main">
      <div class="user__works">
        <h2>User</h2>
        <div class="user">
          <div class="user__side">
            <div class="image-wrapper">
              <img :src="returnURL(user.ImageURL)" :alt="user.Name" />
            </div>
          </div>
          <div class="user__main">
            <p class="user__name">{{ user.Name }}</p>
            <a class="user__url" :href="user.URL">{{ user.URL }}</a>
          </div>
        </div>
        <div class="user__intro">{{ user.Introduction }}</div>
        <h2>Works</h2>
        <works :works="returnWorks" :isSearch="false" :isUserShow="true"/>
        <pagenate :page="page" @goPrev="goPrev()" @goNext="goNext()" v-if="this.works && this.works.length > 8" />
      </div>
      <Git-Hub-Field :hubs="hubs" :ghUser="ghUser" v-if="hubs.length > 0"/>
    </div>
  </div>
</template>
<script>
import Works from "~/components/Works.vue";
import SearchForm from "~/components/SearchForm.vue";
import Pagenate from "~/components/Pagenate.vue";
import Loading from "~/components/Loading.vue";
import GitHubField from "~/components/GitHubField.vue";
export default {
  components: {
    Works,
    SearchForm,
    Pagenate,
    Loading,
    GitHubField
  },
  head () {
    return {
      title: "OnBoard / " + this.user.Name,
      meta: [
        { hid: this.user.ID, name: this.user.Name, content: this.user.Introduction }
      ]
    }
  },
  async mounted() {
    this.APIURL = this.GetURL();
    this.$nextTick(async () => {
      await this.showBubble();
      await this.$axios
        .get(this.APIURL + this.$route.path)
        .then(response => {
          this.user = response.data;
          this.works = response.data.Works;
          this.ghToken = response.data.GitHubToken
        }).then(() => {
          this.gitHubInit()
        })
        .catch(response => console.error(response));
      await this.init();
      await setTimeout(() => this.showBubble(), 1000);
    });
  },
  computed: {
    returnWorks: function() {
      return this.worksList;
    }
  },
  methods: {
    gitHubInit: async function() {
      if (this.ghToken) {
        const headers = { 
          "content-type": "application/json",
          "Authorization": "",
        };
        const github = "https://api.github.com/users/" + this.ghToken
        await this.$axios.get(github, {
          headers
        }).then(res => {
          this.ghUser = res.data
          if (res.data.type !== "User") {
            this.author = 1
          }
        })
        const endpoint = await this.FetchGitInfo(this.author, this.ghToken)
        await this.$axios.get(endpoint, {
          headers
        }).then(res => {
          this.hubs = res.data
        })
      }
    },
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    returnURL: function(url) {
      return url ? url : "/NO_IMAGE.jpeg";
    },
    init: function() {
      this.worksList = [];
      if (this.works) {
        for (let i = 0; i < this.maxCardCount; i++) {
          if (this.works[i]) {
            this.worksList.push(this.works[i]);
          } else {
            continue;
          }
        }
      }
      this.page = 1;
    },
    goPrev: function() {
      if (this.page > 1) {
        this.worksList = [];
        for (
          let i = (this.page - 1) * this.maxCardCount;
          i > (this.page - 2) * this.maxCardCount;
          i--
        ) {
          if (this.works[(this.page - 1) * this.maxCardCount - i]) {
            this.worksList.push(
              this.works[(this.page - 1) * this.maxCardCount - i]
            );
          }
        }
        this.page--;
      }
    },
    goNext: function() {
      if (this.works) {
        if (this.works.length - this.page * this.maxCardCount > 0) {
          this.worksList = [];
          for (let i = 0; i < this.maxCardCount; i++) {
            if (this.works[this.page * this.maxCardCount + i]) {
              this.worksList.push(
                this.works[this.page * this.maxCardCount + i]
              );
            }
          }
          this.page++;
        }
      }
    }
  },
  data() {
    return {
      page: 1,
      worksList: [],
      reload: true,
      maxCardCount: 8,
      user: {},
      works: [],
      isLoading: false,
      APIURL: "",
      ghToken: "",
      author: 0,
      ghUser: {},
      hubs: []
    };
  }
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
.container {
  width: 100%;
  min-height: 81vh;
}
.container__main {
  padding: 90px 0;
  min-height: 100%;
  width: 100%;
  background-color: $bg-color;
}
.user{
  display: flex;
}
.user__works{
  position: relative;
  margin: 30px;
  background-color: white;
  border-radius: 10px;
  border: solid .5px lightgrey;
  padding: 2% 2% 3% 2%;
  & h2 {
    margin: 0 10px 10px 10px;
    color: $bg-main;
    font-size: 18px;
  }
}
.user__name {
  margin: 10px 0 0 0;
  text-align: left;
  font-weight: bold;
  color: $bg-main;
  font-size: 24px;
  line-height: 24px;
}
.user__url {
  color:grey !important;
  text-decoration: none !important;
  font-size: 10px;
}
.user__intro {
  margin: 0 20px;
  font-size: 10px;
  padding: 10px;
  color:grey;
  width: 100%;
  max-height: 360px;
}
.image-wrapper {
  margin: 0 10px 0 30px;
  width: 60px;
  height: 60px;
  & img {
    width: 60px;
    height: 60px;
  }
}
</style>
