<template>
  <div class="container" v-if="reload">
    <Loading v-if="isLoading" />
    <div
      class="container__side"
      :class="slider()"
    >
      <p
        class="drawBtn"
        @click="drawSlider()"
      >＜</p>
      <search-form
        @search="search($event)"
        @cancel="cancel()"
        :isSearch="isSearch"
        searchType="users"
      />
      <pagenate :page="page" @goPrev="goPrev()" @goNext="goNext()" v-if="!isSearch" />
    </div>
    <div class="container__main">
      <h1 class="searchAlert" v-if="isSearch && returnUsers.length == 0">検索結果はありません</h1>
      <div class="user__container">
        <div v-for="(user, index) in returnUsers" :key="index" class="user__bar__wrapper">
          <nuxt-link :to="'/users/' + user.ID">
            <div class="user__bar">
              <img :src="returnURL(user.ImageURL)" :alt="user.Name" />
              <p>{{ user.Name }}</p>
              <p class="counts">{{ user.WorksCount }}件投稿</p>
            </div>
          </nuxt-link>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import SearchForm from "~/components/SearchForm.vue";
import Pagenate from "~/components/Pagenate.vue";
import Loading from "~/components/Loading.vue";
export default {
  components: {
    SearchForm,
    Pagenate,
    Loading
  },
  data() {
    return {
      page: 1,
      usersList: [],
      reload: true,
      isSearch: false,
      users: [],
      maxDataCount: 6,
      isLoading: false,
      APIURL: "",
      slide: false
    };
  },
  head () {
    return {
      title: "OnBoard / 投稿者一覧",
      meta: [
        { hid: "users", name: "投稿者一覧", content: "投稿者の一覧画面です。 好きな投稿者をフォローしましょう。" }
      ]
    }
  },
  async mounted() {
    this.APIURL = this.GetURL();
    this.$nextTick(async () => {
      await this.showBubble();
      await this.$axios
        .get(this.APIURL + "/users")
        .then(response => {
          this.users = response.data.Users;
        })
        .catch(response => console.error(response));
      await this.init();
      await setTimeout(() => this.showBubble(), 500);
    });
  },
  computed: {
    returnUsers: function() {
      return this.usersList;
    }
  },
  methods: {
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
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    returnURL: function(url) {
      return url ? url : "/NO_IMAGE.jpeg";
    },
    init: function() {
      this.usersList = [];
      if (!this.users) return
      if (this.users.length > 0) {
        for (let i = 0; i < this.maxDataCount; i++) {
          if (this.users[i]) {
            this.usersList.push(this.users[i]);
          } else {
            continue;
          }
        }
      }
      this.page = 1;
    },
    search: function(e) {
      if (this.users[0]) {
        this.usersList = this.users.filter(
          user => user.Name.indexOf(e.input) != -1
        );
      }
      this.isSearch = true;
    },
    cancel: function() {
      this.isSearch = false;
      this.reload = false;
      this.init();
      this.reload = true;
    },
    goPrev: function() {
      if (this.page > 1) {
        this.usersList = [];
        for (
          let i = (this.page - 1) * this.maxDataCount;
          i > (this.page - 2) * this.maxDataCount;
          i--
        ) {
          if (this.users[(this.page - 1) * this.maxDataCount - i]) {
            this.usersList.push(
              this.users[(this.page - 1) * this.maxDataCount - i]
            );
          }
        }
        this.page--;
      }
    },
    goNext: function() {
      if (this.users.length - this.page * this.maxDataCount > 0) {
        this.usersList = [];
        for (let i = 0; i < this.maxDataCount; i++) {
          if (this.users[this.page * this.maxDataCount + i]) {
            this.usersList.push(this.users[this.page * this.maxDataCount + i]);
          }
        }
        this.page++;
      }
    }
  }
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
* {
  text-decoration: none;
}
.container {
  display: flex;
  width: 100%;
  min-height: 81vh;
}
.counts{
  font-size: 10px;
  color: #777777;
}
.container__side {
  min-height: 100%;
  padding-top: 70px;
  width: 400px;
  background-color: $bg-yellow;
  box-shadow: 0 0 5px $bg-main;
}
.container__main {
  overflow: hidden;
  padding: 100px 0 70px 0;
  min-height: 100%;
  width: 100%;
  background-color: $bg-color;
}
.user__bar__wrapper {
  padding: 10px 5px;
  background-color: white;
  margin: 20px;
  box-shadow: 0 2px 3px $bg-main;
  width: inherit;
  border-radius: 40px;
  &:hover {
    box-shadow: 0 0 1px $bg-main;
    & div {
      transition: all 0.3s;
    }
  }
}
.user__container{
  margin: 0 auto;
  width: 500px;
  min-height: 580px;
}
.user__bar {
  height: 60px;
  width: 90%;
  font-weight: bold;
  color: $bg-main !important;
  display: flex;
  justify-content: space-between;
  & p {
    line-height: 60px;
    margin-left: 20px;
  }
  & img {
    display: inline-block;
    width: 60px;
    height: 60px;
    border-radius: 50%;
    margin: 0 10px 0 10px;
  }
  &:hover {
    & p {
      color: grey;
    }
  }
}
.searchAlert {
  width: 100%;
  color: white;
  text-align: center;
  line-height: 51vh;
  font-size: 40px;
}
.drawBtn{
  cursor: pointer;
  position: absolute;
  top: 100px;
  left: -25px;
  border-radius: 50%;
  background-color: $bg-yellow;
  color: white;
  height: 50px;
  width: 50px;
  line-height: 50px;
  text-align: center;
  font-size: 20px;
  font-weight: bold;
  border-top: solid 2px $bg-main;
  border-bottom: solid 2px white;
  border-left: solid 2px white;
  border-right: solid 2px $bg-main;
  box-shadow: 0 0 5px $bg-main;
}
.slider__off{
  & p {
    transition: all .5s;
  }
  border-left: solid 2px white;
  transition: all .5s;
  z-index: 1;
  position: fixed;
  top: 0;
  right: -350px;
}
.slider__on{
  & p {
    transition: all .5s;
    transform: rotate(-540deg);
    color: $bg-main;
    // border: solid 2px $bg-main;
  }
  border-left: solid 2px $bg-main;
  transition: all .5s;
  z-index: 1;
  position: fixed;
  top: 0;
  right: 0;
}
@media screen and (max-width: $PhoneSize) {
  .user__container {
    width: 98%;
  }
  .user__bar__wrapper {
    margin: 20px auto;
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
    border-left: solid 2px $bg-main;
    transition: all .5s;
    z-index: 1;
    position: fixed;
    top: 0;
    right: 0;
  }
  .container__side {
    width: 90%;
  }
}
</style>
