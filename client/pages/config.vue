<template>
  <div class="config-container">
    <div class="form-wrapper">
      <form @submit.prevent="submit">
        <div v-if="generalWindow" class="general__form">
          <p>
            <label for="name">名前</label>
            <input type="text" name="name" autocomplete="on" v-model="user.name" required />
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
        </div>
        <div v-else class="secret__form">
          <p>
            <label for="email">メール</label>
            <input name="email" type="email" autocomplete="on" v-model="user.email" />
          </p>
          <p>
            <label for="password">現在のパスワード</label>
            <input name="password" type="password" autocomplete="on" v-model="user.password" />
          </p>
          <p>
            <label for="new-password">新しいパスワード</label>
            <input name="new-password" type="password" v-model="user.newPass" />
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
import DeleteModal from "~/components/DeleteModal.vue";
export default {
  components: {
    DeleteModal
  },
  middleware: ["auth"],
  data() {
    return {
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
        newPass: ""
      },
      APIURL: "",
      isOpenDeleteModal: false
    };
  },
  mounted() {
    this.APIURL = this.GetURL();
    if (this.$auth.user) {
      this.data.image = this.$auth.user.ImageURL;
      this.user.id = this.$auth.user.ID;
      this.user.url = this.$auth.user.URL;
      this.user.name = this.$auth.user.Name;
      this.user.introduction = this.$auth.user.Introduction;
      this.user.email = this.$auth.user.Email;
    }
  },
  methods: {
    closeDeleteModal: function() {
      this.isOpenDeleteModal = false;
    },
    openDeleteModal: function() {
      this.isOpenDeleteModal = true;
    },
    deleteWork: function() {
      ///user削除処理
      this.isOpenDeleteModal = false;
      this.$auth.logout();
      this.$router.push("/");
    },
    async submit() {
      try {
        const data = new FormData();
        if (this.generalWindow) {
          data.append("user_id", this.user.id);
          data.append("name", this.user.name);
          data.append("url", this.user.url);
          data.append("introduction", this.user.introduction);
          data.append("file", this.data.name);
          data.append("windowOpt", true);
        } else {
          data.append("user_id", this.user.id);
          data.append("email", this.user.email);
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
  padding-top: 70px;
  margin: 0 auto;
  min-height: 81vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  width: 100%;
}
.form-wrapper {
  box-shadow: 0 0 5px grey;
  background-color: white;
  border: solid 3px #192b3d;
  border-radius: 5px;
  height: 500px;
  width: 60%;
  position: relative;
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
  background-color: #192b3d;
  &:hover {
    color: #fdeaa0;
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
    background-color: #192b3d;
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
    & input[type="text"] {
      outline: 0;
      border: solid 3px #192b3d;
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
      border: solid 3px #192b3d;
      outline: 0;
      border-radius: 5px;
      padding: 10px;
      height: 80px;
      width: 150px;
      background-color: #192b3d;
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
      border: solid 3px #192b3d;
      border-radius: 5px;
      padding: 20px;
      font-weight: bold;
      font-size: 14px;
    }
  }
}
.file-wrapper {
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
      color: #fdeaa0;
    }
    &:active {
      color: silver;
    }
  }
}
.secret__form {
  font-weight: bold;
  padding-top: 100px;
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
      border: solid 3px #192b3d;
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
