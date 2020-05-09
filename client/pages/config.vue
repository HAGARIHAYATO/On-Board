<template>
  <div class="config-container">
    <Loading v-if="isLoading" />
    <Validation :messages="errors"/>
    <div class="form-wrapper">
      <form @submit.prevent="submit">
        <div v-if="generalWindow" class="general__form">
          <p>
            <label for="name">名前</label>
            <input type="text" name="name" autocomplete="on" v-model="user.name"/>
          </p>
          <p>
            <label for="email">メール</label>
            <input name="email" type="email" autocomplete="on" v-model="user.email"/>
          </p>
          <p>
            <label for="url">サイトURL</label>
            <input
              type="text"
              name="url"
              autocomplete="on"
              v-model="user.url"
              min="1"
              pattern="^[0-9A-Za-z.@/:-]+$"
            />
          </p>
          <div class="file-wrapper">
            <label for="file" class="file">
              クリックor画像をドラッグ
              <span>+</span>
              <input
                name="file"
                id="file"
                type="file"
                ref="file"
                accept="image/png, image/jpeg"
                @change="setImage"
              />
            </label>
            <div class="file__image-wrapper" v-if="data.image">
              <p @click="resetImage">×</p>
              <img :src="data.image" />
            </div>
          </div>
          <p>
            <label for="text">紹介文</label>
            <textarea name="introduction" type="text" rows="4" v-model="user.introduction" />
          </p>
          <p>
            <label for="github">GitHubユーザー名</label>
            <input name="github" type="text" autocomplete="on" v-model="user.github" />
          </p>
          <p>
            <label for="github">Qiitaユーザー名</label>
            <input name="github" type="text" autocomplete="on" v-model="user.qiita" />
          </p>
        </div>
        <div v-else class="secret__form">
          <p>
            <label for="password">現在のパスワード</label>
            <input name="password" type="password" autocomplete="on" v-model="user.password" />
          </p>
          <p>
            <label for="new-password">新しいパスワード</label>
            <input name="new-password" type="password" v-model="user.newPass" />
          </p>
          <p>
            <label for="new-password">新しいパスワード確認</label>
            <input name="new-password" type="password" v-model="user.newPassConfirm" />
          </p>
          <p class="deleteModalBtn" @click="openDeleteModal">アカウントを削除</p>
        </div>
        <div class="form__transform">
          <p class="form__transform__btn1" @click="selectGeneral()">▲</p>
          <p class="form__transform__btn2" @click="selectSecret()">▼</p>
        </div>
        <p>
          <input class="form__btn" type="submit" value="設定を更新" />
        </p>
      </form>
    </div>
    <Delete-Modal
      v-if="isOpenDeleteModal"
      @delete="deleteWork"
      @back="closeDeleteModal"
      confirmStr="削除してしまうと復元することはできません。よろしいですか?"
      btnStr="同意してアカウントを削除"
    />
  </div>
</template>
<script>
import Loading from "~/components/Loading.vue";
import DeleteModal from "~/components/DeleteModal.vue";
import Validation from "~/components/Validation.vue";
export default {
  components: {
    DeleteModal,
    Loading,
    Validation
  },
  middleware: ["auth"],
  data() {
    return {
      errors: [],
      data: {
        image: "",
        name: ""
      },
      returnFile: null,
      generalWindow: false,
      user: {
        id: "",
        name: "",
        url: "",
        email: "",
        password: "",
        newPass: "",
        newPassConfirm: "",
        github: "",
        qiita: ""
      },
      APIURL: "",
      isOpenDeleteModal: false,
      isLoading: false
    };
  },
  async mounted() {
    this.$nextTick(async () => {
      await this.showBubble();
      try {
        this.APIURL = this.GetURL();
        if (this.$auth.user) {
          await this.$axios
            .get(this.APIURL + "/users/" + this.$auth.user.ID)
            .then(r => {
              this.data.image = r.data.ImageURL;
              this.user.id = r.data.ID;
              this.user.url = r.data.URL;
              this.user.name = r.data.Name;
              this.user.introduction = r.data.Introduction;
              this.user.email = r.data.Email;
              this.user.github = r.data.GitHubToken;
              this.user.qiita = r.data.QiitaName;
            })
            .catch(response => console.error(response));
            }
      } catch (error) {
        // handling
      }
      await setTimeout(() => this.showBubble(), 1000);
    });
  },
  methods: {
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    closeDeleteModal: function() {
      this.isOpenDeleteModal = false;
    },
    openDeleteModal: function() {
      this.isOpenDeleteModal = true;
    },
    deleteWork: function() {
      this.$nextTick(async () => {
        await this.showBubble();
        try {
          await this.$axios
            .delete(this.APIURL + "/delete_account")
            .then(res => {
              this.isOpenDeleteModal = false;
              this.$auth.logout();
              this.$router.push("/");
            });
        } catch (error) {}
        await setTimeout(() => this.showBubble(), 1000);
      });
    },
    valCheck: function(bool) {
      this.errors = []
      if (bool) {
        if (this.user.name === "") {
          const empty = "ユーザー名は必須です。"
          this.errors.push(empty)
        }
        if (this.user.email === "") {
          const empty = "メールアドレスは必須です。"
          this.errors.push(empty)
        }
        if (this.user.name.length > 300) {
          const over = "文字数は最大20字です。"
          this.errors.push(over)
        }
      } else {
        if (this.user.password === "" || this.user.newPass === "") {
          const empty = "パスワードは必須です。"
          this.errors.push(empty)
        }
        if (this.user.newPass !== this.user.newPassConfirm) {
          const check = "新しいパスワードを確認してください。"
          this.errors.push(check)
        }
      }
    },
    async submit() {
      this.$nextTick(async () => {
        this.valCheck(this.generalWindow)
        if (this.errors.length !== 0) {
          return
        }
        await this.showBubble();
        try {
          const data = new FormData();
          if (this.generalWindow) {
            data.append("user_id", this.user.id);
            data.append("name", this.user.name);
            data.append("email", this.user.email);
            data.append("github", this.user.github);
            data.append("qiita", this.user.qiita);
            data.append("url", this.user.url);
            data.append("introduction", this.user.introduction);
            data.append("file", this.data.name);
            data.append("windowOpt", true);
          } else {
            data.append("user_id", this.user.id);
            data.append("password", this.user.password);
            data.append("new_password", this.user.newPass);
            data.append("windowOpt", false);
          }
          const headers = { "content-type": "multipart/form-data" };
          await this.$axios
            .put(this.APIURL + "/users/" + this.user.id, data, {
              headers
            })
            .then(res => {
              if (res.data.ID) {
                this.$router.push("/users/" + res.data.ID);
              }
            });
        } catch (e) {}
        await setTimeout(() => this.showBubble(), 1000);
      });
    },
    selectGeneral: function() {
      this.generalWindow = true;
    },
    selectSecret: function() {
      this.generalWindow = false;
    },
    setImage(e) {
      const files = this.$refs.file;
      const fileImg = files.files[0];
      if (fileImg.size > 3000000) {
        alert(this.AlertMessage());
        return;
      }
      if (fileImg.type.startsWith("image/")) {
        this.data.image = window.URL.createObjectURL(fileImg);
        this.data.name = e.target.files[0];
      }
    },
    resetImage(e) {
      this.$refs.file.value = "";
      this.data.name = "";
      this.data.image = "";
    }
  }
};
</script>
<style lang="scss" scopped>
@import "assets/scss/app";
* {
  box-sizing: border-box;
  margin: 0;
}
input,
textarea,
.file {
  &:hover {
    box-shadow: 0 0 5px black;
  }
}
.config-container {
  padding-top: 90px;
  margin: 0 auto;
  min-height: 81vh;
  width: 100%;
}
.form-wrapper {
  box-shadow: 0 0 5px grey;
  background-color: white;
  border: solid 3px $bg-main;
  border-radius: 5px;
  min-height: 600px;
  padding-bottom: 50px; 
  margin: 10px auto 50px auto;
  width: 610px;
  position: relative;
  text-align: center;
}
.form__btn {
  outline: 0;
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  height: 50px;
  width: 100px;
  color: white;
  font-weight: bold;
  font-size: 14px;
  border-radius: 15px 15px 0 0;
  background-color: $bg-main;
  &:hover {
    color: $bg-yellow;
  }
  &:active {
    color: silver;
  }
}
.form__transform {
  display: inline-block;
  width: 60px;
  height: 100px;
  position: absolute;
  right: 0;
  top: 50%;
  transform: translateY(-50%);
  & p {
    line-height: 50px;
    outline: 0;
    width: 50px;
    height: 50px;
    color: white;
    font-size: 20px;
  }
  & .form__transform__btn1 {
    background-color: #f8d54a;
    border-radius: 20px 20px 0 0;
    &:hover {
      box-shadow: 0 0 5px black;
      color: silver;
    }
    &:active {
      color: grey;
    }
  }
  & .form__transform__btn2 {
    background-color: $bg-main;
    border-radius: 0 0 20px 20px;
    &:hover {
      box-shadow: 0 0 5px black;
      color: silver;
    }
    &:active {
      color: grey;
    }
  }
}
.general__form {
  padding-top: 30px;
  & p,
  .file-wrapper {
    font-weight: bold;
    width: 300px;
    margin: 10px auto;
    & label {
      display: block;
      text-align: left;
    }
    & input[type="text"],
    input[type="email"] {
      outline: 0;
      border: solid 3px $bg-main;
      border-radius: 20px;
      padding: 0 20px;
      font-size: 18px;
      height: 30px;
      width: 300px;
      font-weight: bold;
    }
    & .image__icon {
      margin: 10px 0 0 10;
      height: 70px;
    }
    & .file {
      border: solid 3px $bg-main;
      outline: 0;
      border-radius: 5px;
      padding: 10px;
      height: 80px;
      width: 150px;
      background-color: $bg-main;
      color: white;
      font-size: 8px;
      & span {
        margin: 2px auto;
        display: block;
        width: 20px;
        height: 20px;
        font-size: 14px;
        font-weight: bold;
        line-height: 18px;
        text-align: center;
        border-radius: 50%;
        transition: all 0.3s;
        background-color: rgb(150, 248, 4);
        box-shadow: 0 0 3px black;
        &:hover {
          color: rgb(192, 192, 192);
          background-color: darken(rgb(150, 248, 4), 10%);
        }
        &:active {
          box-shadow: 0 0 0 black;
        }
      }
      & input {
        display: none;
      }
    }
    & textarea {
      outline: 0;
      width: 300px;
      min-height: 100px;
      border: solid 3px $bg-main;
      border-radius: 5px;
      padding: 20px;
      font-weight: bold;
      font-size: 14px;
    }
  }
}
.file-wrapper {
  position: relative;
  display: flex;
  height: 80px;
  & img {
    height: 80px;
    border-radius: 10px;
  }
}
.file__image-wrapper {
  position: relative;
  margin: 0 0 0 10px !important;
  & p {
    margin: 0 !important;
    position: absolute;
    top: -5px;
    right: -5px;
    width: 20px;
    height: 20px;
    border-radius: 50%;
    background-color: grey;
    color: white;
    font-weight: bold;
    line-height: 18px;
    cursor: pointer;
    &:hover {
      color: $bg-yellow;
    }
    &:active {
      color: silver;
    }
  }
}
.secret__form {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-weight: bold;
  & p {
    width: 340px;
    margin: 20px auto;
    & label {
      display: block;
      text-align: left;
    }
    & input {
      margin: 0;
      outline: 0;
      border: solid 3px $bg-main;
      border-radius: 20px;
      padding: 0 20px;
      font-size: 18px;
      height: 30px;
      width: 300px;
      font-weight: bold;
    }
  }
}
.deleteModalBtn {
  margin: 10px 0 0 0;
  outline: 0;
  background-color: red;
  color: white;
  border-radius: 20px;
  padding: 0 20px;
  font-size: 12px;
  height: 30px;
  width: 300px !important;
  line-height: 30px;
  font-weight: bold;
}
</style>
