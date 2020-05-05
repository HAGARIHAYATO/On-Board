<template>
  <div class="info__container">
    <div class="info__layout">
      <div class="info__search">
        <input type="text" v-model="search" @change="searchMail" placeholder="Please Keyword & Enter">
      </div>
      <p class="info__text">{{lists.length}}件表示</p>
      <div class="info__rows">
        <div class="info__bar">
          <div class="info" v-for="(mail, index) in lists" :key="index" @click="getMail(mail)">
            <p class="title"><span v-if="mail.UserID > 0">●</span>{{ mail.Title }}</p>
            <p class="date">{{ mail.CreatedAt }}</p>
            <p class="message">{{ cutSTR(mail.Message) }}</p>
          </div>
        </div>
      </div>
    </div>
    <div class="layer" v-if="appearMailBox" @click="resetMail">
      <div class="mail__box">
        <div class="mail__bar">
          <p class="title"><span v-if="selectedMail.UserID > 0">●</span>{{ selectedMail.Title }}</p>
          <p class="date">{{ selectedMail.CreatedAt }}</p>
          <p class="message">{{ selectedMail.Message }}</p>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      count: 3,
      search: "",
      lists: [],
      APIURL: "",
      mails: [],
      selectedMail: {},
      appearMailBox: false
    }
  },
  head () {
    return {
      title: "OnBoard / おしらせ",
      meta: [
        { hid: "information", name: "おしらせ　", content: "日々のおしらせ。" }
      ]
    }
  },
  mounted() {
    this.APIURL = this.GetURL();
    const uid = this.isLoginUser()
    this.$axios
      .get(this.APIURL + "/users/" + uid + "/information")
      .then(response => {
        this.mails = response.data.Info;
        this.lists = this.mails
      })
      .catch(response => console.error(response));
  },
  methods: {
    resetMail: function() {
      this.selectedMail = {}
      this.appearMailBox = false
    },
    getMail: function(mail) {
      this.selectedMail = mail
      this.appearMailBox = true
    },
    isLoginUser: function() {
      if (this.$auth.user || this.$auth.user.loggedIn) {
        return this.$auth.user.ID
      }
      return 0
    },
    searchMail: function(e) {
      this.lists = this.mails.filter(
        mail =>
          mail.Title.indexOf(e.target.value) != -1
      );
    },
    cutSTR: function(str) {
      if (str && str.length > 15) {
        return str.slice( 0, 14 ) + "..."
      }
      return str
    }
  },
}
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
* {
  box-sizing: border-box;
}
.layer{
  position: fixed;
  z-index: 3;
  top: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0.9);
}
.mail__bar{
  width: 460px;
  margin: 20px 30px 20px 12px;
  word-break: break-all;
  & p {
    margin-bottom: 20px;
  }
}
.mail__box {
  overflow-y: scroll;
  z-index: 4;
  background-color: white;
  border: solid 3px grey;
  border-radius: 25px;
  height: 400px;
  width: 500px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.info__container{
  padding: 110px 0 40px 0;
}
.info__layout{
  margin: 0 auto;
  width: 500px;
  height: 100%;
  border: solid 3px $bg-main;
}
.info__search{
  height: 43px;
  background-color: $bg-main;
  text-align: right;
  & input {
    height: 40px;
    width: 160px;
    border-radius: 20px;
    border: solid 3px grey;
    outline: none;
    padding: 0 15px;
  }
}
.info__text{
  height: 30px;
  line-height: 30px;
  text-align: center;
  font-size: 10px;
  color: grey;
}
.info__rows{
  height: 550px;
  overflow-y: scroll;
}
.info__bar{
  color: darken(grey, 10%);
  & .info {
    border-top: solid 1px lightgrey;
    padding: 10px;
  }
  &:last-child {
    border-bottom: solid 1px lightgrey;
  }
}
.title{
  & span {
    color: red;
  }
  padding: 3px 0 0 10px;
  font-weight: bold;
}
.date{
  font-size: 8px !important;
  text-align: right !important;
}
.message{
  padding-left: 10px;
}
</style>