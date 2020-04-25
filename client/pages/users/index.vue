<template>
  <div class="container" v-if="reload">
    <Loading v-if="isLoading" />
    <div class="container__side">
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
      <div v-for="(user, index) in returnUsers" :key="index" class="user__bar__wrapper">
        <nuxt-link :to="'/users/' + user.ID">
          <div class="user__bar">
            <img :src="returnURL(user.ImageURL)" :alt="user.Name" />
            <p>{{ user.Name }}</p>
            <p>{{ user.WorksCount }}件投稿</p>
          </div>
        </nuxt-link>
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
      APIURL: ""
    };
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
      await setTimeout(() => this.showBubble(), 1000);
    });
  },
  computed: {
    returnUsers: function() {
      return this.usersList;
    }
  },
  methods: {
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
* {
  text-decoration: none;
}
.container {
  display: flex;
  width: 100%;
  min-height: 81vh;
}
.container__side {
  min-height: 100%;
  padding-top: 70px;
  width: 400px;
  background-color: #fdeaa0;
}
.container__main {
  padding: 140px 0 70px 0;
  min-height: 100%;
  width: 100%;
  background-color: lighten(rgb(221, 209, 209), 5%);
}
.user__bar__wrapper {
  &:first-child {
    border-top: solid 3px #192b3d;
  }
  padding-top: 10px;
  background-color: white;
  margin-left: 20px;
  border-bottom: solid 3px #192b3d;
  margin: auto 0;
  &:hover {
    & div {
      border-left: solid 5px rgb(9, 185, 9);
      transition: all 0.3s;
    }
  }
}
.user__bar {
  margin-bottom: 10px;
  height: 60px;
  width: 70%;
  font-weight: bold;
  color: #192b3d !important;
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
</style>
