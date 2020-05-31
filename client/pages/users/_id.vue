<template>
  <div class="container" v-if="reload">
    <Loading v-if="isLoading" />
    <div
      class="container__side"
      :class="slider()"
    >
      <search-form
        @search="search($event)"
        @cancel="cancel()"
        :isSearch="isSearch"
        searchType="works_but_nouser"
      />
    </div>
    <div class="container__main">
      <div class="user__works">
        <h2>User - ユーザー情報</h2>
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
        <div class="user">
          <p class="is__published" v-if="hubs.length > 0">Github連携済み</p>
          <p class="is__published" v-else>Github連携なし</p>
          <p class="is__published" v-if="articles.length > 0">Qiita連携済み</p>
          <p class="is__published" v-else>Qiita連携なし</p>
        </div>
        <div class="csv">
          <p class="user__csv" @click="downloadCSV">スキルCSV取得 <img src="/download.svg"></p>
          <p class="user__csv" @click="downloadCSVWorks">作品CSV取得 <img src="/download.svg"></p>
        </div>
        <div class="user__intro">{{ user.Introduction }}</div>
        <h2>Works - 作品一覧</h2>
        <works :works="returnWorks" :isSearch="false" :isUserShow="true"/>
        <pagenate :page="page" @goPrev="goPrev()" @goNext="goNext()" v-if="!isSearch" />
      </div>
      <Git-Hub-Field :hubs="hubs" :ghUser="ghUser" v-if="hubs.length > 0"/>
      <Qiita-Field :articles="articles" :qiitaUser="qiitaUser" v-if="articles.length > 0"/>
    </div>
    <p class="searchBtn" @click="drawSlider()"><img src="/search.svg" alt="search"></p>
  </div>
</template>
<script>
import axios from "axios"
import Works from "~/components/Works.vue";
import SearchForm from "~/components/SearchForm.vue";
import Pagenate from "~/components/Pagenate.vue";
import Loading from "~/components/Loading.vue";
import GitHubField from "~/components/GitHubField.vue";
import QiitaField from "~/components/QiitaField.vue";
export default {
  components: {
    Works,
    SearchForm,
    Pagenate,
    Loading,
    GitHubField,
    QiitaField
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
      await this.$axios
        .get(this.APIURL + this.$route.path)
        .then(response => {
          this.user = response.data;
          this.works = response.data.Works;
          this.ghToken = response.data.GitHubToken
          this.qiita = response.data.QiitaName
          this.skillData = this.getSkills(response.data.Works)
        }).then(() => {
          this.gitHubInit()
          this.qiitaInit()
        })
        .catch(response => console.error(response));
      await this.init();
      await setTimeout(() => this.showBubble(), 500);
    });
  },
  computed: {
    returnWorks: function() {
      return this.worksList;
    }
  },
  methods: {
    getSkills: function(array) {
      let skillList = []
      for(const el of array) {
        if (el["Skills"]) {
          for (const skill of el["Skills"]) {
            const lowerSkill = skill.Name.toLowerCase()
            skillList.push(lowerSkill)
          }
        }
      }
      let counts = {};
      for(let i=0;i< skillList.length;i++) {
        let key = skillList[i];
        counts[key] = (counts[key])? counts[key] + 1 : 1 ;
      }
      let skills = Object.keys(counts)
      skillList = []
      for (const index in skills) {
        let obj = {}
        obj["name"] = skills[index]
        obj["count"] = counts[skills[index]]
        skillList.push(obj)
      }
      return skillList
    },
    downloadCSV () {  
      var csv = '\ufeff' + 'スキル名,スキル使用数\n'
      this.skillData.forEach((el, index) => {
        var line = el['name'] + ',' + el['count'] + '\n'
        csv += line
      })
      let blob = new Blob([csv], { type: 'text/csv' })
      let link = document.createElement('a')
      link.href = window.URL.createObjectURL(blob)
      link.download = `${this.user.Name}_skills.csv`
      link.click()
    },
    downloadCSVWorks () {
      var csv = '\ufeff' + 'ID,作品名,画像,作者名,作者画像\n'
      this.works.forEach(el => {
        var line = el['ID'] + ',' + el['Name'] + ',' + el['ImageURL'] + ',' + el['UserName'] + ',' + el['UserImageURL'] + '\n'
        csv += line
      })
      let blob = new Blob([csv], { type: 'text/csv' })
      let link = document.createElement('a')
      link.href = window.URL.createObjectURL(blob)
      link.download = `${this.user.Name}_works.csv`
      link.click()
    },
    slider: function() {
      const off = "slider__off"
      const on = "slider__on"
      if (this.slide) {
        return on
      } else {
        return off
      }
    },
    drawSlider: function() {
      this.slide = !this.slide
    },
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
    qiitaInit: async function() {
      // authorizationを外すために通常のnuxtにbindされていないaxiosを使用
      if (this.qiita) {
        const headers = { 
          "content-type": "application/json",
        };
        const url = "https://qiita.com/api/v2/users/" + this.qiita + "/items" + "?page=1&per_page=20"
        await axios.get(url, {
          headers
        }).then(res => {
          if(res.data && res.data.length > 0) {
            this.articles = res.data
            this.qiitaUser = res.data[0].user
          }
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
    cancel: function() {
      this.isSearch = false;
      this.reload = false;
      this.init();
      this.reload = true;
    },
    search: function(e) {
      if (this.works[0] && e.check) {
        this.worksList = this.works.filter(
          work =>
            work.Name.toLowerCase().indexOf(e.input.toLowerCase()) != -1
        );
      } else if (this.works[0] && !e.check) {
        this.worksList = this.works.filter(
          work => work.Name.toLowerCase().indexOf(e.input.toLowerCase()) != -1
        );
      }
      this.isSearch = true;
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
      isLoading: true,
      APIURL: "",
      ghToken: "",
      author: 0,
      ghUser: {},
      hubs: [],
      isSearch: false,
      slide: false,
      qiita: "",
      articles: [],
      qiitaUser: {},
      skillData: []
    };
  }
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
.container {
  width: 100%;
  min-height: 81vh;
  overflow: hidden;
}
.container__side {
  margin-top: 70px;
  padding: 10px 20px;
  background-color: lightgrey;
  border-radius: 5px;
  box-shadow: 0 0 1px black;
}
.slider__off{
  & p {
    transition: all .5s;
  }
  transition: all .5s;
  z-index: 1;
  position: fixed;
  top: 0;
  right: -400px;
}
.slider__on{
  & p {
    transition: all .5s;
    transform: rotate(-540deg);
    color: $bg-main;
    left: -25px;
  }
  transition: all .5s;
  z-index: 1;
  position: fixed;
  top: 0;
  right: 20px;
}
.searchBtn {
  cursor: pointer;
  width: 60px;
  height: 60px;
  border-radius: 50%;
  position: fixed;
  bottom: 200px;
  right: 100px;
  background-color: lightgray;
  text-align: center;
  line-height: 60px;
  font-size: 22px;
  font-weight: bold;
  box-shadow: 0 1px 3px black;
  z-index: 3;
  & img {
    height: 16px;
    margin: 22px 0;
  }
}
.container__main {
  padding: 90px 0;
  min-height: 100%;
  width: 100%;
  background-color: $bg-color;
}
.user__works{
  position: relative;
  margin: 30px;
  background-color: white;
  border-radius: 10px;
  border: solid .5px lightgrey;
  padding: 2% 2% 3% 2%;
  // box-shadow: 0 1px 2px grey;
  & h2 {
    margin: 50px;
    color: $bg-main;
    font-size: 18px;
    border-bottom: solid 2px $bg-main;
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
  word-break: break-all;
  margin: 0 0 0 100px;
  font-size: 10px;
  color:grey;
  width: 100%;
  max-height: 360px;
}
.user{
  padding: 0 100px;
  display: flex;
}
.csv{
  display: flex;
  padding: 0 90px;
}
.image-wrapper {
  margin: 0 10px 0 0;
  width: 60px;
  height: 60px;
  & img {
    width: 60px;
    height: 60px;
  }
}
.is__published{
  display: inline-block;
  padding:  3px 15px;
  background-color: $bg-main;
  color: $bg-yellow;
  border-radius: 20px;
  margin: 20px 5px;
}
.user__csv{
  font-weight: bold;
  border-radius: 5px;
  box-shadow: 0 1px 2px grey;
  margin: 0 10px 18px 10px;
  padding: 0 20px;
  cursor: pointer;
  color: grey;
  height: 30px;
  line-height: 30px;
  & img {
    display: inline-block;
    line-height: 30px;
    height: 16px;
    transform: translate(15%, 15%);
  }
  &:active{
    box-shadow: 0 0 1px black;
  }
  &:hover {
    box-shadow: 0 0 1px grey;
  }
}
@media screen and (max-width: $PhoneSize) {
  h2 {
    margin: 40px 5px !important;
  }
  .user__works {
    width: 98%;
    margin: 30px auto;
  }
  .user__intro {
    margin: 0;
  }
  .user__url{
    font-size: 6px;
  }
  .user__name {
    font-size: 14px;
  }
  .user {
    padding: 0;
  }
  .slider__off{
    & p {
      transition: all .5s;
      left: -50px;
    }
    transition: all .5s;
    z-index: 1;
    position: fixed;
    top: 0;
    right: -90%;
  }
  .slider__on{
    & p {
      transition: all .5s;
      transform: rotate(-540deg);
      color: $bg-main;
    }
    transition: all .5s;
    z-index: 1;
    position: fixed;
    top: 0;
    right: 0;
  }
  .container__side {
    width: 90%;
  }
  .user__csv{
    display: none;
  }
  .searchBtn {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    position: fixed;
    bottom: 20px;
    right: 80px;
    background-color: lightgray;
    text-align: center;
    line-height: 40px;
    font-size: 12px;
    font-weight: bold;
    box-shadow: 0 1px 3px black;
    z-index: 3;
    & img {
      height: 10px;
      margin: 15px 0;
    }
  }
}
</style>
