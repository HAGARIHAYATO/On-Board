<template>
  <div class="info__container">
    <div class="info__layout">
      <div class="info__search">
        <input type="text" v-model="search" @change="searchMail" placeholder="Please Keyword & Enter">
      </div>
      <p class="info__text">{{lists.length}}件表示</p>
      <div class="info__rows">
        <div class="info__bar">
          <div class="info" v-for="(mail, index) in lists" :key="index">
            <p class="title">{{ mail.Title }}</p>
            <p class="date">{{ mail.CreatedAt }}</p>
            <p class="message">{{ cutSTR(mail.Message) }}</p>
          </div>
        </div>
      </div>
    </div>
    <div class="layer"></div>
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
      mails: []
    }
  },
  mounted() {
    this.APIURL = this.GetURL();
    this.$axios
      .get(this.APIURL + "/users/" + 0 + "/information")
      .then(response => {
        this.mails = response.data.Info;
        this.lists = this.mails
      })
      .catch(response => console.error(response));
  },
  methods: {
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
  height: 40px;
  background-color: $bg-main;
  text-align: right;
  & input {
    height: 30px;
    width: 120px;
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
    height: 67px;
  }
  &:last-child {
    border-bottom: solid 1px lightgrey;
  }
}
.title{
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