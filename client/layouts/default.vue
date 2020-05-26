<template>
  <div class="layout">
    <Loading v-if="isLoading" />
    <div class="confirmModal" v-if="confirmModal">
      <div class="confirm">
        <button @click="doLogout()">ログアウトする</button>
        <button @click="cancelLogout()">ログアウトせずに戻る</button>
      </div>
    </div>
    <div class="header">
      <div class="infoModal">
        <div class="infoModal__item" v-if="true">
          <div class="infoModal__nav static-bar">
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
          <div class="infoModal__nav" :class="drawed">
            <template v-if="!isLogin">
              <nuxt-link to="/">トップ</nuxt-link>
              <nuxt-link to="/information">お知らせ</nuxt-link>
              <nuxt-link to="/signup">無料登録</nuxt-link>
              <nuxt-link to="/login">ログイン</nuxt-link>
            </template>
            <template v-else>
              <nuxt-link :to="'/users/'+$auth.user.ID">マイページ</nuxt-link>
              <nuxt-link to="/information">お知らせ</nuxt-link>
              <nuxt-link to="/works/new">新規投稿</nuxt-link>
              <nuxt-link to="/config">設定</nuxt-link>
              <nuxt-link to="/admins/management" v-if="checkAdmin">管理画面</nuxt-link>
              <a @click="showConfirm()">ログアウト</a>
            </template>
          </div>
        </div>
      </div>
    </div>
    <nuxt class="cover__main" />
    <footer>
      <div class="footer__logo"></div>
      <p class="footer__content">
        <nuxt-link to="/contact">お問い合わせ</nuxt-link>
      </p>
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
    },
    checkAdmin: function() {
      return this.$store.$auth.user.IsAdmin ? true : false;
    },
    drawed: function() {
      return this.openModal ? "drawed" : "betodraw"
    }
  },
  methods: {
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    drawModal: function() {
      if (window.parent.screen.width > 767) {
        this.$router.push("/")
      } else {
        this.openModal = !this.openModal;
      }
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
@import "assets/scss/app";
html {
  margin: 0 !important;
  overflow: hidden !important;
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
  margin: 0 10px;
  width: 100px;
  height: 50px;
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
  z-index: 2;
}
.header__wrapper {
  display: flex;
  z-index: 1;
}
.infoModal {
  width: 100%;
  background-color: $bg-main;
  z-index: 1;
}
.infoModal__item {
  justify-content: space-between;
  display: flex;
  height: 50px;
  padding: 0 5px;
  z-index: 1;
  & a,
  span {
    color: #FFFFFF;
    font-weight: bold;
    font-size: 14px;
    transition: all 0.3s;
    line-height: 50px;
    margin: 0 10px;
    &:hover {
      color: $bg-yellow;
      transition: all 0.4s;
    }
  }
  & a {
    display: inline-block;
  }
}
.infoModal__nav{
  display: flex;
}
.header__navigater {
  display: flex;
}
.confirmModal {
  z-index: 3;
  position: fixed;
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
    background: $bg-main;
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
  background-color: $bg-main;
  min-height: 150px;
  width: 100%;
  margin-bottom: 0;
  display: flex;
  justify-content: space-between;
}
.footer__logo {
  background-image: url("/logo.jpeg");
  background-position: center;
  background-size: cover;
  margin-left: 60px;
  width: 200px;
  height: 150px;
}
.cover__main {
  min-height: 82vh !important;
}
.footer__content {
  margin: 0 50px;
  & a {
    font-weight: bold;
    font-size: 18px;
    line-height: 150px;
    text-decoration: none;
    color: white;
    &:hover {
      color: $bg-yellow;
    }
  }
}
@keyframes colorChange {
  0% {
    background-color: white;
  }
  100% {
    background-color: $bg-main;
  }
}
@keyframes easeAppear {
  0% {
    height: 0;
  }
  100% {
    height: 100px;
  }
}
@keyframes ankerAppear {
  0% {
    transform: scaleY(0);
  }
  100% {
    transform: scaleY(1);
  }
}
@media screen and (max-width: $PhoneSize) {
  .infoModal__item {
    width: 100%;
    padding: 0 0;
    z-index: 1;
    display: block;
  }
  .header__logo {
    margin-left: 10px;
  }
  .header__navigater {
    & a {
      font-size: 18px;
    }
  }
  .infoModal__nav {
    display: block;
    width: 100%;
    background-color: $bg-main;
    & a {
      display: block;
    }
    transition: all .3s;
  }
  .static-bar {
    display: flex;
  }
  .drawed {
    animation-name: draw;
    animation-duration: .1s;
  }
  .betodraw {
    height: 0;
    & a {
      display: none;
      animation-name: draw;
      animation-duration: .1s;
    }
  }
  @keyframes draw {
    0% {
      height: 0;
    }
    100% {
      height: auto;
    }
  }
  @keyframes close {
    0% {
      height: auto;
    }
    100% {
      height: 0;
    }
  }
}
</style>
