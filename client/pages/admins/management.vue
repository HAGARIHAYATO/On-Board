<template>
  <div class="admin__container">
    <Loading v-if="isLoading" />
    <p>管理者画面</p>
    <div class="admin__info__container">
      <p>お知らせ</p>
      <p class="info-message">{{infoStatus}}</p>
      <p><input type="text" v-model="informTitle"></p>
      <p>
        <textarea v-model="inform" cols="100" rows="10"></textarea>
      </p>
      <p>
      　<button @click="sendInfo">お知らせ送信</button>
      </p>
    </div>
    <div class="manage__account">
      <div class="admin__layout">
        <div class="admin__search">
          <input type="text" v-model="search" @change="searchUser" placeholder="Please Keyword & Enter">
        </div>
        <p class="admin__text">{{lists.length}}件表示</p>
        <div class="admin__rows">
          <div class="admin__bar">
            <div class="admin" v-for="(user, index) in lists" :key="index" @click="userSelect($event, user)">
              <p class="title">{{ user.Name }}</p>
              <p class="date">{{ user.ID }}</p>
            </div>
          </div>
        </div>
      </div>
      <div class="select__account" v-if="selectedUser.ID">
        <p>ID: {{selectedUser.ID}}</p>
        <p> NAME: {{selectedUser.Name}}</p>
        <p class="red-message">{{status}}</p>
        <p>
          <button @click="openDeleteModal">ユーザーを削除</button>
          <button @click="sendAlert">警告メッセージを送信</button>
        </p>
        <p><input type="text" v-model="alertTitle"></p>
        <p>
          <textarea v-model="alert" cols="90" rows="10">
          </textarea>
        </p>
      </div>
    </div>
    <Delete-Modal
      v-if="isOpenDeleteModal"
      @delete="deleteUser"
      @back="closeDeleteModal"
      confirmStr="削除してしまうと復元することはできません。よろしいですか?"
      btnStr="同意してアカウントを削除"
    />
  </div>
</template>
<script>
import Loading from "~/components/Loading.vue";
import DeleteModal from "~/components/DeleteModal.vue";
export default {
  components: {
    DeleteModal,
    Loading
  },
  middleware: ["auth"],
  async mounted() {
    if (!this.$auth.user.IsAdmin) {
      await this.$router.push("/")
    }
    this.APIURL = this.GetURL();
    this.$nextTick(async () => {
      await this.showBubble();
      await this.$axios
        .get(this.APIURL + "/users")
        .then(response => {
          this.users = response.data.Users;
          this.lists = this.users
        })
        .catch(response => console.error(response));
      await setTimeout(() => this.showBubble(), 1000);
    });
  },
  data() {
    return {
      inform: "",
      informTitle: "",
      users: [],
      lists: [],
      search: "",
      isLoading: false,
      selectedUser: {},
      status: "",
      infoStatus: "",
      alert: "",
      alertTitle: "",
      isOpenDeleteModal: false,
      APIURL: ""
    }
  },
  head () {
    return {
      title: "OnBoard / 管理者",
      meta: [
        { hid: "admin", name: "管理者", content: "管理者専用ページです。" }
      ]
    }
  },
  methods: {
    userSelect: function(e, user) {
      this.selectedUser = user
    },
    closeDeleteModal: function() {
      this.isOpenDeleteModal = false;
    },
    openDeleteModal: function() {
      this.isOpenDeleteModal = true;
      if (this.$auth.user.ID === this.selectedUser.ID) {
        this.isOpenDeleteModal = false; 
        this.status = "このユーザーは削除できません。" 
      }
    },
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    searchUser: function(e) {
      this.lists = this.users.filter(
        user =>
          user.Name.indexOf(e.target.value) != -1
      );
    },
    cutSTR: function(str) {
      if (str && str.length > 15) {
        return str.slice( 0, 14 ) + "..."
      }
      return ""
    },
    sendAlert: function(e) {
      if (this.alertTitle === "" || this.alert === "") {
        this.status = "入力を確認してください。"
        return
      }
      const data = {
        uid: this.selectedUser.ID,
        message: this.alert,
        title: this.alertTitle
      }
      const headers = {"content-type": "application/json"}
      this.$axios
        .post(this.APIURL + "/information", data, {
          headers
        })
        .then(res => {
          this.status = "送信OK"
          this.alert = ""
        })
        .catch(response => {
          console.error(response)
          this.status = "送信できませんでした。"
        });
    },
    sendInfo: function(e) {
      if (this.inform === "" || this.informTitle === "") {
        this.infoStatus = "入力を確認してください。"
        return
      }
      const data = {
        uid: 0,
        message: this.inform,
        title: this.informTitle
      }
      const headers = {"content-type": "application/json"}
      this.$axios
        .post(this.APIURL + "/information", data, {
          headers
        })
        .then(res => { 
          this.infoStatus = "送信OK"
          this.inform = ""
        })
        .catch(response => {
          console.error(response)
          this.infoStatus = "送信できませんでした。"
        });
    },
    deleteUser: function() {
      this.$nextTick(async () => {
        const data = {
          uid: this.selectedUser.ID
        }
        const headers = {"content-type": "application/json"}
        await this.showBubble();
        try {
          await this.$axios
            .post(this.APIURL + "/execute_account", data, {
              headers
            })
            .then(res => {
              this.status = res.data.status
            });
        } catch (error) {}
        await setTimeout(() => this.showBubble(), 1000);
      });
    },
  },
}
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
*{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
button{
  padding: 1%;
  border-radius: 5px;
  border: solid 2px red;
  cursor: pointer;
  outline: none;
  &:active{
    border: solid 2px blue;
  }
}
textarea{
  border: solid 2px grey;
  border-radius: 5px;
  padding: 1%;
  outline: none;
}
input[type="text"] {
  border: solid 2px grey;
  border-radius: 5px;
  padding: 1%;
  margin: 5px 0;
  outline: none;
  width: 200px;
}
.manage__account{
  margin: 40px 0;
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-start;
}
.select__account{
  margin: 1%;
  width: 48%;
  min-height: 350px;
  border: solid 3px $bg-main;
  padding: 20px;
  background-color: white;
  & p {
    margin: 5px 0;
    width: 100%;
    & input {
      width: 100%;
    }
  }
}
.admin__container{
  min-height: 100vh;
  padding: 100px;
  background-color: lightgray;
}
.admin__layout{
  margin: 1%;
  width: 48%;
  height: 100%;
  border: solid 3px $bg-main;
  background-color: white;
}
.admin__search{
  height: 43px;
  background-color: $bg-main;
  text-align: right;
  & input {
    margin: 0;
    height: 40px;
    width: 160px;
    border-radius: 20px;
    border: solid 3px grey;
    outline: none;
    padding: 0 15px;
    &:hover {
      box-shadow: none;
    }
  }
}
.admin__text{
  height: 30px;
  line-height: 30px;
  text-align: center;
  font-size: 10px;
  color: grey;
}
.admin__rows{
  height: 550px;
  overflow-y: scroll;
}
.admin__bar{
  color: darken(grey, 10%);
  & .admin {
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
.red-message{
  color: blue !important;
  font-size: 10px;
}
.info-message{
  color: green !important;
  font-size: 14px;
}
textarea {
    width: 100%;
  }
@media screen and (max-width: $PhoneSize) {
  .admin__container {
    padding: 150px 10px;
  }
  textarea {
    width: 100%;
  }
  .admin__layout{
    width: 100%;
  }
  .select__account{
    width: 100%;
  }
  .admin__rows {
    height: 400px;
  }
}
</style>
