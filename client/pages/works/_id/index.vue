<template>
  <div class="show__wrapper">
    <Loading v-if="isLoading" />
    <Assess-Modal v-if="isOpenAssessModal" @close="viewAssess(true)" @loading="postLoading()" />
    <div class="toPostModal" v-if="isClose">
      <p class="closeBtn" @click="migrateModal()">×</p>
      <button class="modalBtn" v-if="$store.$auth.loggedIn" @click="viewAssess(true)">評価する</button>
      <button class="modalBtn" v-else @click="viewAssess(false)">ログインして評価する</button>
    </div>
    <p class="plusBtn" v-else @click="migrateModal()">＋</p>
    <div class="show__container">
      <div class="show__main">
        <a :href="work.URL" class="show__container__image">
          <img :src="returnURL(work.ImageURL)" :alt="work.Name" class="show__image" />
        </a>
      </div>
      <h2>Information - 作品情報</h2>
      <div class="show__bar">
        <div class="show__info">
          <h3 class="info__title">{{ work.Name }}</h3>
          <p class="info__title__sub">{{ work.URL }}</p>
          <div>
            <p v-if="work.IsPublished" class="is__published">公開済み</p>
            <p v-else class="is__published">作成中</p>
            <p v-if="work.GHR" class="is__published">Github連携済み</p>
            <p v-else class="is__published">Github連携なし</p>
          </div>
          <div class="show__user">
            <nuxt-link :to="'/users/' + work.UserID" class="info__name">
              <img :src="returnURL(work.UserImageURL)" class="user__icon">
              <p>{{ work.UserName }}</p>
            </nuxt-link>
          </div>
          <div class="show__description">{{ isNone(work.Description) }}</div>
        </div>
        <div class="show__subInfo">
          <p
            class="skill__tag"
            v-for="(skill, index) in work.Skills"
            :key="index"
          >
            {{skill.Name}}
          </p>
        </div>
      </div>
      <h2>Item - 作品詳細</h2>
      <div class="select__item" v-if="isArrayExist(work.WorkItems)">
        <div class="select__item__image" v-if="selectItem">
          <img :src="returnURL(selectItem.ImageURL)" alt="作品画像" />
        </div>
        <div class="select__item__content">
          <p>{{ isNone(selectItem.Body) }}</p>
        </div>
      </div>
      <div class="show__container__main">
        <div
          v-for="(item, index) in work.WorkItems"
          :key="index"
          class="item__wrapper"
          @click="showWorksItem(item)"
          :class="isSelected(item.ID)"
        >
          <template v-if="item.ImageURL != '' || item.Body != ''">
            <p v-if="isSelected(item.ID)">▼</p>
            <p v-else>-</p>
            <img :src="returnURL(item.ImageURL)" alt="作品画像" />
          </template>
        </div>
      </div>
      <h2>Assessment - 作品評価</h2>
      <div v-if="isEvaluate">
        <Radar-Chart :dataset="radarData" :height="170" :width="calcSize()"/>
        <h3>総合スコア   {{calcDataPoint}}点 / 50点</h3>
        <Users-Ball v-if="work.AssessmentUsers && work.AssessmentUsers.length > 0" :users="work.AssessmentUsers" />
      </div>
      <div v-else>
        <p class="nothing__alert">評価はありません。</p>
      </div>
      <h2>Composition - 構図</h2>
      <div class="">
        <img
          v-if="isExist(work.CacooURL)"
          class="cacoo__frame"
          :src="isExist(work.CacooURL)"
        />
        <div
         v-else
         class="cacoo__frame"
        >
         <p class="nothing__alert">構図はありません。</p>
        </div>
      </div>
      <h2>Github - Github連携情報</h2>
      <div v-if="work.GHR" class="gh-box">
        <div class="gh-info">
          <h3><a :href="ghRepository.html_url">{{work.GHR}}</a></h3>
          <p>（※Githubプロジェクト名）</p>
          <p>
            By
            <img
              :src="ghRepository.owner.avatar_url"
              :alt="ghRepository.owner.login"
            />
            <a :href="ghRepository.owner.html_url">{{ghRepository.owner.login}}</a>
          </p>
          <h3>{{RewriteTime(ghRepository.created_at)}} ~ {{RewriteTime(ghRepository.pushed_at)}}</h3>
          <p>（※リポジトリ作成 ~ 最終push）</p>
          <div>
            <p class="gh-title__sub" @click="changeGraph(language, true)">＞ 言語使用状況 <img class="checkbox" src="/check.svg" v-if="isPie"></p>
            <p
              v-for="(d, n) in Object.keys(language)"
              :key="n"
              class="gh-check"
            >
              {{d}}
            </p>
          </div>
          <div>
            <p class="gh-title__sub" @click="changeGraph(score, false)">＞ スコア <img class="checkbox" src="/check.svg" v-if="!isPie"></p>
            <p
              v-for="(s, nu) in Object.keys(score)"
              :key="nu"
              class="gh-check"
            >
              {{s}}
            </p>
          </div>
        </div>
        <div class="gh-graph">
          <Chart :dataset="dataset" text="プロジェクト内言語別使用率" v-if="isPie" :height="calcSize()" :width="calcSize()" />
          <Bar-Chart :dataset="dataset" v-else :height="calcSize()" :width="calcSize()" />
        </div>
      </div>
      <p v-else class="nothing__alert">情報はありません。</p>
      <p class="pdfOpen" @click="OpenPDF">PDFを取得 <img src="/download.svg"></p>
      <p class="operation" v-if="isMine">
        <nuxt-link :to="'/works/' + work.ID + '/edit'">作品編集</nuxt-link>|
        <nuxt-link :to="'/works/' + work.ID + '/edit_item'">アイテム編集</nuxt-link>
      </p>
      <button class="deleteBtn" v-if="isMine" @click="openDeleteModal">削除する</button>
      <Delete-Modal
        v-if="isOpenDeleteModal"
        @delete="deleteWork"
        @back="closeDeleteModal"
        confirmStr="削除してしまうと復元することはできません。よろしいですか?"
        btnStr="同意して削除する"
      />
    </div>
  </div>
</template>
<script>
import Loading from "~/components/Loading.vue";
import DeleteModal from "~/components/DeleteModal.vue";
import Chart from "~/components/Chart.vue";
import BarChart from "~/components/BarChart.vue";
import RadarChart from "~/components/RadarChart.vue";
import AssessModal from "~/components/AssessModal.vue";
import UsersBall from "~/components/UsersBall.vue";
export default {
  components: {
    Loading,
    DeleteModal,
    Chart,
    BarChart,
    RadarChart,
    AssessModal,
    UsersBall
  },
  data() {
    return {
      selectId: "",
      selectItem: {},
      work: {},
      isLoading: true,
      APIURL: "",
      isOpenDeleteModal: false,
      isPie: true,
      ghRepository: {
        owner: {
          login: "",
          avatar_url: "",
          html_url: ""
        },
        html_url: ""
      },
      language: {},
      score:{
        stars: 0,
        subscribers: 0,
        watchers: 0,
        forks: 0
      },
      dataset: {},
      radarData: {
        "機能の充実度": 0,
        "UI/UX": 0,
        "不都合な動作の少なさ": 0,
        "内容の斬新さ": 0,
        "言語やFWのモダンさ": 0
      },
      isEvaluate: 0,
      isOpenAssessModal: false,
      isClose: true
    };
  },
  head () {
    return {
      title: "OnBoard / " + this.work.Name,
      meta: [
        { hid: this.work.ID, name: this.work.Name, content: this.work.Description }
      ]
    }
  },
  computed: {
    calcDataPoint: function() {
      const points = Object.values(this.radarData)
      const sum = points.reduce(function(a, x){return a + ((x || 0) - 0);}, 0);
      return sum
    },
    isMine: function() {
      if (!this.$auth.user || !this.work.UserID) return;
      if (this.work.UserID === this.$auth.user.ID) {
        return true;
      }
      return false;
    }
  },
  methods: {
    migrateModal: function() {
      this.isClose = !this.isClose
    },
    viewAssess: function(bool) {
      if (bool) {
        this.isOpenAssessModal = !this.isOpenAssessModal
      } else {
        this.$router.push("/login");
      }
    },
    OpenPDF: function() {
      this.$router.push("/works/" + this.work.ID + "/download")
    },
    changeGraph: function(obj, bool) {
      if (obj) {
        this.dataset = obj;
        if (bool) {
          this.isPie = true
        } else {
          this.isPie = false
        }
      }
    },
    isExist: function(str) {
      return str ? str : false
    },
    isNone: function(str) {
      return str == "" ? "本文はありません。" : str
    },
    closeDeleteModal: function() {
      this.isOpenDeleteModal = false;
    },
    openDeleteModal: function() {
      this.isOpenDeleteModal = true;
    },
    deleteWork: function() {
      this.$nextTick(async () => {
        this.APIURL = this.GetURL();
        await this.showBubble();
        await this.$axios
          .delete(this.APIURL + this.$route.path)
          .then(response => {
            this.$router.push("/works");
          })
          .catch(response => console.error(response));
        await setTimeout(() => this.showBubble(), 10);
      });
      this.isOpenDeleteModal = false;
    },
    isArrayExist: function(array) {
      if (array) {
        return array.length > 0 ? true : false;
      }
    },
    isSelected: function(id) {
      if (this.selectItem.ID === id) {
        return "selected-item";
      }
    },
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    returnURL: function(url) {
      return url ? url : "/NO_IMAGE.jpeg";
    },
    showWorksItem: function(item) {
      if (typeof item != "object") return;
      this.selectId = item.ID;
      this.selectItem = item;
    },
    FetchGitRepository: async function(userName, repoName) {
      if (!userName || !repoName) return
      let obj = {}
      const headers = { 
        "content-type": "application/json",
        "Authorization": "",
      };
      const url = "https://api.github.com/repos/"
      await this.$axios
        .get(url + userName + "/" + repoName, {
          headers
        })
        .then(res => {
          this.ghRepository = res.data
          this.score.stars = this.ghRepository.stargazers_count
          this.score.watchers = this.ghRepository.watchers_count
          this.score.subscribers = this.ghRepository.subscribers_count
          this.score.forks = this.ghRepository.forks_count
        })
        .catch(response => console.error(response));
      await this.$axios
        .get(url + userName + "/" + repoName + "/languages", {
          headers
        })
        .then(res => {
          this.language = res.data
          this.dataset = this.language
        })
        .catch(response => console.error(response));
    },
    assessmentInit: function(assess) {
      this.radarData["機能の充実度"] = assess.Function
      this.radarData["不都合な動作の少なさ"] = assess.BugSafe
      this.radarData["UI/UX"] = assess.UIX
      this.radarData["内容の斬新さ"] = assess.Content
      this.radarData["言語やFWのモダンさ"] = assess.MDN
    },
    postLoading: function() {
      this.isOpenAssessModal = !this.isOpenAssessModal
      this.showBubble();
      this.initDisplay()
    },
    initDisplay: function() {
      this.$nextTick(async () => {
        this.APIURL = this.GetURL();
        // await this.showBubble();
        await this.$axios
          .get(this.APIURL + this.$route.path)
          .then(response => {
            this.work = response.data;
            this.selectItem = response.data.WorkItems[0]
            this.assessmentInit(response.data.Assessment)
            this.isEvaluate = response.data.AssessmentUserCount
          })
          .catch(response => {
            if (response.response.status === 404) {
              this.$router.push("/error.html")
            }
          });
        await this.FetchGitRepository(this.work.UserGithubToken, this.work.GHR)
        await setTimeout(() => this.showBubble(), 1000);
      });
    }
  },
  async mounted() {
    this.initDisplay()
  }
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
*,
body {
  box-sizing: border-box;
}
h2 {
  margin: 50px;
  color: $bg-main;
  font-size: 18px;
  border-bottom: solid 2px $bg-main;
}
h3 {
  margin: 35px;
  color: $bg-main;
  font-size: 16px;
  text-align: center;
}
.show__wrapper {
  overflow: hidden;
  padding: 100px 0 50px 0;
  width: 100%;
  min-height: 81vh;
  background-color: $bg-color;
}
.show__container{
  margin: 30px;
  // box-shadow: 0 2px 5px grey;
  background-color: white;
  border-radius: 10px;
  padding: 4% 2%;
}
.show__bar {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  width: 70%;
  margin: 30px auto 50px auto;
}
.show__description {
  font-size: 10px;
  text-align: left !important;
  padding: 20px;
  min-height: 20vh;
  width: 300px;
  color: grey;
  word-break: break-all;
  background-color: white;
  box-shadow: 0 0 1px grey;
}
.show__container__image {
  min-width: 200px;
  height: 300px;
}
.show__image {
  display: block;
  margin: 0 auto;
  max-width: 600px;
  height: 300px;
  box-shadow: 0 1px 2px grey;
  &:hover{
    box-shadow: 0 0 1px grey;
  }
}
.show__user {
  text-align: left;
  margin: 0 0 34px 0;
  max-width: 400px;
}
.info__name {
  display: inline-flex;
  min-width: 40%;
  background-color: white;
  border-radius: 25px !important;
  box-shadow: 0 1px 2px grey;
  padding: 0 4px;
  margin: 0 auto;
  height: 40px;
  word-break: break-all !important;
  text-decoration: none !important;
  color: $bg-main;
  & p {
    margin: 0 10px;
    line-height: 40px;
  }
  &:hover {
    box-shadow: 0 0 1px grey;
  }
}
.user__icon{
  margin: 5px 2px;
  width: 30px;
  height: 30px;
  border-radius: 50%;
  border: solid 1px grey;
}
.show__container__main {
  height: 130px;
  display: flex;
  justify-content: center;
  border-radius: 5px;
  & img {
    border-radius: 50%;
    width: 70px;
    height: 70px;
    margin: 0 20px;
  }
}
.item__wrapper {
  height: 130px;
  & p {
    transition: all .3s;
    color: white;
    text-align: center;
  }
  & img {
    cursor: pointer;
    box-shadow: 0 1px 2px grey;
  }
  &:hover {
    & p {
      color: black;
    }
    & img {
      box-shadow: 0 0 1px grey;
    }
  }
  &:active {
    & p {
      color: black;
    }
    & img {
      box-shadow: 0 0 1px grey;
      transform: scale(0.97);
    }
  }
}
.select__item {
  margin: 30px auto 0 auto;
  width: 70%;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  border-radius: 5px;
  transition: all 0.3s;
}
.select__item__image {
  display: inline-block;
  height: 350px;
  & img {
    height: 350px;
  }
}
.select__item__content {
  background-color: white;
  border-radius: 3px;
  width: 70%;
  min-height: 50px;
  padding: 10px 2%;
  & p {
    word-break: break-all;
    font-size: 10px;
    color: grey;
    padding: 2%;
    margin: 0 auto;
    min-width: 99.5%;
    min-height: 99%;
  }
}
.selected-item {
  & img {
    transform: scale(0.99);
    border: solid 2px black;
  }
  &:hover {
    & img {
      transform: scale(0.97);
    }
  }
}
.close-modal {
  transform: scale(0);
}
.operation {
  text-align: center;
  font-weight: bold;
}
.deleteBtn {
  margin: 10px auto;
  height: 40px;
  width: 180px;
  display: block;
  outline: none;
  background-color: red;
  color: white;
  border-radius: 5px;
  font-weight: bold;
  box-shadow: 0 1px 2px grey;
  &:hover {
    box-shadow: 0 0 1px grey;
  }
}
.info__title {
  text-align: left;
  font-size: 24px;
  margin: 0 0 0 0;
}
.info__title__sub {
  text-align: left;
  margin: 0 0 10px 0;
  font-size: 8px;
  color: grey;
}
.show__main, .show__info{
  margin: 0 25px;
}
.grey_span{
  color: grey;
  font-size: 16px;
}
.show__subInfo{
  border-radius: 5px;
  width: 300px;
  margin: 0 20px 0 20px;
  padding: 10px;
}
.skill__tag{
  display: inline-block;
  background-color: $tag-color;
  border-radius: 20px;
  color: white;
  padding: 4px 15px;
  font-weight: bold;
  box-shadow: 0 1px 2px grey;
  margin: 4px;
}
.cacoo__frame{
  display: block;
  width: 70%;
  margin: 50px auto;
}
.nothing__alert{
  text-align: center;
  line-height: 528px;
  color: grey;
  font-weight: bold;
  font-size: 30px;
}
.is__published{
  display: inline-block;
  padding:  3px 15px;
  background-color: $bg-main;
  color: $bg-yellow;
  border-radius: 20px;
  margin: 10px 0 20px 0;
}
.gh-box{
  display: flex;
  justify-content: center;
  margin: 0 0 50px 0;
  flex-wrap: wrap;
}
.gh-info{
  width: 400px;
  border-radius: 10px;
  box-shadow: 0 1px 2px grey;
  margin: 20px;
  transition: all .2s;
  padding: 20px;
  font-size: 12px;
  color: #777777;
  & h3 {
    margin: 10px 0 0 0;
  }
  & a {
    text-decoration: none !important;
    color: $tag-color !important;
  }
  & img {
    width: 15px;
    height: 15px;
  }
  &:hover{
    box-shadow: 0 0 1px grey;
  }
}
.gh-graph{
  padding: 0 0 10px 0;
  width: 400px;
  border-radius: 10px;
  box-shadow: 0 1px 2px grey;
  margin: 20px;
  transition: all .2s;
  &:hover{
    box-shadow: 0 0 1px grey;
  }
}
.gh-title__sub{
  color: $tag-color;
  font-size: 16px;
  font-weight: bold;
  margin: 20px 0 5px 0;
  transition: all .2s;
  cursor: pointer;
  &:hover{
    margin: 20px 0 5px 5px;
  }
}
.gh-check{
  display: inline-block;
  margin: 5px;
}
.checkbox{
  width: 16px;
}
.pdfOpen{
  text-align: center;
  cursor: pointer;
  margin: 10px auto;
  height: 40px;
  line-height: 40px;
  width: 180px;
  display: block;
  background-color: white;
  color: grey;
  border-radius: 5px;
  font-weight: bold;
  box-shadow: 0 1px 2px grey;
  & img {
    display: inline-block;
    line-height: 40px;
    height: 16px;
    transform: translate(15%, 15%);
  }
  &:hover {
    box-shadow: 0 0 1px grey;
  }
}
.toPostModal {
  position: fixed;
  top: 100px;
  right: 50px;
  width: 200px;
  border-radius: 5px;
  box-shadow: 0 2px 3px grey;
  background-color: white;
}
.closeBtn {
  position: absolute;
  top: -10px;
  right: -10px;
  border-radius: 50px;
  color: white;
  background-color: grey;
  height: 28px;
  width: 28px;
  cursor: pointer;
  font-weight: bold;
  font-size: 20px;
  line-height: 28px;
  text-align: center;
}
.modalBtn {
  background-color: green;
  color: white;
  width: 90%;
  margin: 20px auto;
  display: block;
  height: 40px;
  border-radius: 5px;
  font-weight: bold;
  font-size: 14px;
  outline: none;
  cursor: pointer;
}
.plusBtn {
  position: fixed;
  top: 100px;
  right: 50px;
  box-shadow: 0 1px 2px grey;
  background-color: green;
  border-radius: 50%;
  width: 30px;
  height: 30px;
  margin: 5px;
  cursor: pointer;
  font-weight: bold;
  font-size: 20px;
  color: white;
  outline: none;
  text-align: center;
}
@media screen and (max-width: $PhoneSize) {
  h2 {
    margin: 40px 5px !important;
  }
  .show__image {
    max-width: 100%;
    height: auto;
    margin-top: 30px; 
  }
  .select__item__image {
    height: auto;
    & img {
      max-width: 100%;
      height: auto;
      margin-top: 30px; 
    }
  }
  .show__container{
    width: 98%;
    margin: 30px auto;
  }
  .show__container__main {
    overflow-x: scroll;
  }
  .item__wrapper {
    &:nth-child(1) {
      margin-left: 300px;
    }
  }
  .gh-graph{
    width: 100%;
    margin: 20px 0;
  }
  .gh-info{
    width: 100%;
    margin: 20px 0;
  }
  .nothing__alert{
    font-size: 16px;
  }
  .pdfOpen{
    display: none;
  }
  .toPostModal,
  .plusBtn {
    display: none;
  }
}
</style>
