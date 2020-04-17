<template>
  <div class="config-container">
    <div class="form-wrapper">
      <form :action="returnEndPoint" method="post" enctype="multipart/form-data">
        <div v-if="generalWindow" class="general__form">
          <p>
            <label for="name">名前</label>
            <input
              type="text"
              name="name"
              autocomplete="on"
              placeholder="名前"
              :value="user.name"
              required
            />
          </p>
          <p>
            <label for="url">サイトURL</label>
            <input
              type="text"
              name="url"
              autocomplete="on"
              :value="user.url"
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
            <textarea name="introduction" type="text" rows="4" />
          </p>
        </div>
        <div v-else class="secret__form">
          <p>
            <label for="email">メール</label>
            <input name="email" type="email" autocomplete="on" pattern="^[0-9A-Za-z.@/:-_]+$" />
          </p>
          <p>
            <label for="password">現在のパスワード</label>
            <input name="password" type="password" autocomplete="on" pattern="^[0-9A-Za-z.@/:-_]+$" />
          </p>
          <p>
            <label for="new-password">新しいパスワード</label>
            <input name="new-password" type="password" pattern="^[0-9A-Za-z.@/:-_]+$" />
          </p>
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
  </div>
</template>
<script>
export default {
  // middleware: ["auth"],
  data() {
    return {
      data: {
        image: "",
        name: ""
      },
      returnFile: null,
      generalWindow: false,
      user: {
        name: "tanaka",
        url: "hsssss.com",
        ImageURL: "https://images.app.goo.gl/Td6Ab2G5EFEJLKeh9"
      }
    };
  },
  mounted() {
    if (this.user.ImageURL) {
      this.data.image = this.user.ImageURL;
    }
  },
  computed: {
    returnEndPoint: function() {
      return this.generalWindow ? "/users/update" : "/users/password/refresh";
    }
  },
  methods: {
    selectGeneral: function() {
      this.generalWindow = true;
    },
    selectSecret: function() {
      this.generalWindow = false;
    },
    setImage(e) {
      const files = this.$refs.file;
      const fileImg = files.files[0];
      if (fileImg.type.startsWith("image/")) {
        this.data.image = window.URL.createObjectURL(fileImg);
        this.data.name = fileImg.name;
        this.data.type = fileImg.type;
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
</style>
