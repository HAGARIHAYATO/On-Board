<template>
  <div class="container" v-if="reload">
    <Loading v-if="isLoading" />
    <div class="container__side">
      <div class="image-wrapper">
        <img :src="returnURL(user.ImageURL)" :alt="user.Name" />
      </div>
      <p class="user__name">{{ user.Name }}</p>
      <p class="user__url">
        <a :href="user.URL">{{ user.URL }}</a>
      </p>
      <pagenate :page="page" @goPrev="goPrev()" @goNext="goNext()" />
      <div class="user__intro">{{ user.Introduction }}</div>
    </div>
    <div class="container__main">
      <div class="user__works">
        <works :works="returnWorks" :isSearch="false" />
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
.container {
  display: flex;
  width: 100%;
  min-height: 81vh;
}
.container__side {
  padding-top: 140px;
  min-height: 100%;
  width: 400px;
  background-color: #fdeaa0;
}
.container__main {
  padding-top: 70px;
  min-height: 100%;
  width: 100%;
  background-color: lighten(rgb(221, 209, 209), 5%);
}
.image-wrapper {
  margin: 0 auto;
  width: 200px;
  height: 200px;
  & img {
    box-shadow: 0 0 5px grey;
    width: 200px;
    height: 200px;
  }
}
.user__name {
  margin-top: 10px;
  text-align: center;
  font-weight: bold;
  color: #192b3d;
  font-size: 24px;
  line-height: 24px;
}
.user__url {
  text-align: center;
  margin-bottom: 50px;
}
.user__intro {
  box-shadow: 0 0 5px grey;
  font-size: 10px;
  margin: 40px auto 60px auto;
  padding: 3%;
  border: solid 3px #192b3d;
  border-radius: 5px;
  background-color: white;
  width: 80%;
  min-height: 120px;
  max-height: 360px;
}
</style>
