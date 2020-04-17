<template>
  <div>
    <Loading v-if="isLoading" />
    <div class="confirmModal" v-if="confirmModal">
      <div class="confirm">
        <button @click="doLogout()">ログアウトする</button>
        <button @click="cancelLogout()">ログアウトせずに戻る</button>
      </div>
    </div>
    <div class="header">
      <div class="infoModal">
        <div class="infoModal__item" v-if="openModal">
          <template v-if="!isLogin">
            <nuxt-link to="/">トップ</nuxt-link>
            <nuxt-link to="/signup">無料登録</nuxt-link>
            <nuxt-link to="/login">ログイン</nuxt-link>
          </template>
          <template v-else>
            <nuxt-link :to="'/users/3'">マイページ</nuxt-link>
            <nuxt-link to="/works/new">新規投稿</nuxt-link>
            <nuxt-link to="/config">設定</nuxt-link>
            <a @click="showConfirm()">ログアウト</a>
          </template>
        </div>
      </div>
      <div class="header__wrapper">
        <div class="header__logo" @click="drawModal"></div>
        <div class="header__navigater">
          <div class="header__navigater__tab">
            <nuxt-link to="/works">作品</nuxt-link>
          </div>
          <div class="header__navigater__tab">
            <nuxt-link to="/users">ユーザー</nuxt-link>
          </div>
        </div>
      </div>
    </div>
    <nuxt />
    <footer>
      <div class="footer__logo"></div>
    </footer>
  </div>
</template>
<script>
import Loading from "~/components/Loading.vue";
export default {
  components: {
    Loading
  },
  data() {
    return {
      openModal: false,
      confirmModal: false,
      isLoading: false
    };
  },
  // TODO
  computed: {
    isLogin: function() {
      return this.$store.$auth.loggedIn ? true : false;
    }
  },
  methods: {
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    drawModal: function() {
      this.openModal = !this.openModal;
    },
    showConfirm: function() {
      this.confirmModal = true;
    },
    cancelLogout: function() {
      this.confirmModal = false;
    },
    doLogout: function() {
      this.$nextTick(async () => {
        await this.showBubble();
        // TODO ログアウト
        this.$auth.logout();
        this.confirmModal = false;
        this.$router.push("/");
        await setTimeout(() => this.showBubble(), 1000);
      });
    }
  }
};
</script>
<style lang="scss" scoped>
html {
  margin: 0 !important;
}
*,
*:before,
*:after {
  z-index: 0;
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  text-decoration: none;
}
.header__logo {
  background-image: url("/logo.jpeg");
  background-position: center;
  background-size: cover;
  margin-left: 60px;
  width: 100px;
  height: 70px;
  box-shadow: 0 32px 30px -30px black;
  border-radius: 0 0 6px 6px;
  z-index: 1;
  &:hover {
    cursor: pointer;
    transition-delay: 0.5s;
  }
}
.header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 1;
}
.header__wrapper {
  display: flex;
  z-index: 1;
}
.infoModal {
  width: 100%;
  background-color: #192b3d;
  display: flex;
  justify-content: flex-end;
  z-index: 1;
  box-shadow: 0 32px 30px -30px black;
  z-index: 1;
}
.infoModal__item {
  height: 100px;
  padding: 0 10%;
  z-index: 1;
  & a,
  span {
    color: lighten($color: lightgrey, $amount: 10%);
    font-weight: bold;
    transition: all 0.3s;
    line-height: 100px;
    margin-left: 20px;
    &:hover {
      color: #fdeaa0;
      transition: all 0.4s;
    }
  }
  animation-name: easeAppear;
  animation-fill-mode: forwards;
  animation-duration: 0.2s;
  animation-timing-function: linear;
  animation-delay: 0s;
  animation-direction: normal;
}
.header__navigater {
  display: flex;
  & a {
    display: inline-block;
    color: #192b3d;
    font-weight: bold;
    transition: all 0.4s;
    line-height: 70px;
    z-index: 1;
  }
}
.header__navigater__tab {
  text-align: center;
  min-width: 100px;
  height: 70px;
  border: solid 3px #192b3d;
  border-top: none;
  box-shadow: 0 32px 30px -30px black;
  border-radius: 0 0 6px 6px;
  background-color: white;
  z-index: 1;
  margin-left: 5px;
  &:hover {
    & a {
      color: #fdeaa0;
    }
    animation-name: colorChange;
    animation-fill-mode: forwards;
    animation-duration: 0.3s;
    animation-timing-function: linear;
    animation-delay: 0s;
    animation-direction: normal;
  }
}
.confirmModal {
  z-index: 3;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0.9);
}
.confirm {
  padding: 10px;
  height: 130px;
  width: 100%;
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  & button {
    margin: 0 auto 10px auto;
    display: block;
    outline: 0;
    height: 50px;
    width: 250px;
    font-size: 16px;
    background: #192b3d;
    font-weight: bold;
    color: #fff;
    border: none;
    border-radius: 10px;
    transition: all 0.2s;
    &:hover {
      box-shadow: 0 0 5px black;
      color: silver;
    }
    &:active {
      box-shadow: 0 0 0 black;
      color: black;
    }
  }
}
footer {
  background-color: #192b3d;
  min-height: 150px;
}
.footer__logo {
  background-image: url("/logo.jpeg");
  background-position: center;
  background-size: cover;
  margin-left: 60px;
  width: 200px;
  height: 150px;
}
@keyframes colorChange {
  0% {
    background-color: white;
  }
  100% {
    background-color: #192b3d;
  }
}
@keyframes easeAppear {
  0% {
    height: 0;
  }
  100% {
    height: 100px;
    & a {
      display: inline;
    }
  }
}
</style>
