<template>
  <form @submit.prevent="submit" class="new__wrapper">
    <Loading v-if="isLoading" />
    <Validation :messages="errors"/>
    <div class="new__user__bar">
      <div class="new__user__bar__top">
        <div class="new__url">
          <div class="new__container__image">
            <label for="new__image" class="new__image">
              クリックor画像をドラッグ
              <span>+</span>
              <input
                name="work_image"
                id="new__image"
                type="file"
                ref="file"
                accept="image/png, image/jpeg"
                @change="setImage"
              />
            </label>
            <div class="new__file" v-if="data.image">
              <p @click="resetImage">×</p>
              <img :src="data.image" />
            </div>
          </div>
          <div class="new__url__info">
            <p class="info__name__sub">作品名</p>
            <input name="work_name" type="text" class="info__name" v-model="workName" />
            <p class="info__name__sub">作品URL</p>
            <input name="work_url" type="url" class="info__name" v-model="workURL" />
          </div>
        </div>
      </div>
      <textarea
        class="new__description"
        name="work_description"
        cols="30"
        rows="10"
        v-model="workDesc"
      ></textarea>
      <p class="info__name__sub">Cacoo画像URL</p>
      <img class="info__image" src="/cacoo__teach.png" alt="cacoo使い方">
      <p class="info__teach">
        <a href="https://cacoo.com">クリックしてCacooサイトへ遷移</a>した後、シートを作成して上画像の通りURLを取得
      </p>
      <input type="text" class="info__name__cacoo" v-model="cacooURL" />
      <select v-model="is_published">
        <option :value="true">公開済み</option>
        <option :value="false">作成中</option>
      </select>
    </div>
    <div class="new__btn">
      <input type="submit" value="保存" />
    </div>
  </form>
</template>
<script>
import Loading from "~/components/Loading.vue";
import Validation from "~/components/Validation.vue";
export default {
  components: {
    Loading,
    Validation
  },
  middleware: ["auth"],
  data() {
    return {
      data: {
        image: "",
        name: ""
      },
      errors: [],
      workName: "",
      workURL: "",
      workDesc: "",
      APIURL: "",
      cacooURL: "",
      is_published: true,
      isLoading: false
    };
  },
  mounted() {
    this.APIURL = this.GetURL();
  },
  methods: {
    valCheck: function() {
      this.errors = []
      if (this.workName === "") {
        const empty = "作品名は必須です。"
        this.errors.push(empty)
      }
      if (this.workDesc.length > 300) {
        const over = "文字数は最大300字です。"
        this.errors.push(over)
      }
    },
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    async submit() {
      this.$nextTick(async () => {
        this.valCheck()
        if (this.errors.length !== 0) {
          return
        }
        await this.showBubble();
        try {
          const data = new FormData();
          data.append("user_id", this.$auth.user.ID);
          data.append("name", this.workName);
          data.append("description", this.workDesc);
          data.append("url", this.workURL);
          data.append("file", this.data.name);
          data.append("cacoo_url", this.cacooURL);
          data.append("is_published", this.is_published);
          const headers = { "content-type": "multipart/form-data" };
          await this.$axios
            .post(this.APIURL + "/works", data, {
              headers
            })
            .then(res => {
              if (res.data.ID) {
                this.$router.push("/works/" + res.data.ID);
              }
            });
        } catch (error) {
          // handling show message
        }
        await setTimeout(() => this.showBubble(), 1000);
      });
    },
    returnURL: function(url) {
      return url ? url : "/NO_IMAGE.jpeg";
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
<style lang="scss" scoped>
@import "assets/scss/app";
*,
body {
  box-sizing: border-box;
}
.new__wrapper {
  overflow-x: hidden;
  padding: 100px 0 50px 0;
  width: 100%;
  min-height: 81vh;
  background-color: $bg-color;
}
.new__user__bar {
  width: 600px;
  text-align: left;
  margin: 0 auto;
}
.new__description {
  font-size: 10px;
  text-align: left !important;
  padding: 20px;
  min-height: 20vh;
  width: 600px;
  margin: 20px auto;
  word-break: break-all;
  border-radius: 3px;
  background-color: white;
  box-shadow: 0 0 2px grey;
}
.new__container__image {
  margin: 0 auto;
  min-width: 200px;
  height: 200px;
}
.new__url {
  display: flex;
  max-width: 600px;
  & a {
    display: inline-block;
    background-color: white;
    text-decoration: none;
    transition: all 0.2s;
    box-shadow: 0 0 2px grey;
    border-radius: 2px;
    &:hover {
      & .new__image {
        box-shadow: 0 0 5px grey;
      }
    }
    &:active {
      & .new__image {
        box-shadow: 0 0 1px grey;
      }
    }
  }
  text-align: center;
}
.new__url__info {
  text-align: left;
  width: 400px;
  padding: 2% 0 2% 2%;
}
.info__name__sub {
  margin-left: 10px;
  font-size: 8px;
  font-weight: bold;
  text-align: left !important; 
}
.info__name {
  width: 95%;
  background-color: white;
  border-radius: 2px;
  box-shadow: 0 0 5px grey;
  padding: 0 20px;
  margin: 10px 0 10px 10px;
  height: 60px;
  line-height: 60px;
  word-break: break-all !important;
  & a {
    box-shadow: none;
    color: $bg-main;
    &:hover {
      font-weight: bold;
    }
  }
}
.info__name__cacoo{
  width: 100%;
  background-color: white;
  border-radius: 2px;
  box-shadow: 0 0 5px grey;
  padding: 0 20px;
  margin: 10px 0 20px 0;
  height: 60px;
  line-height: 60px;
  word-break: break-all !important;
}
.info__image{
  width: 60%;
}
.info__teach{
  font-size: 12px;
}
.new__image {
  display: block;
  margin: 0 auto;
  height: 200px;
  outline: 0;
  border-radius: 5px;
  padding: 40% 0;
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
.new__file {
  top: -200px;
  position: relative;
  margin: 0 !important;
  & img {
    max-width: 200px;
    height: 200px;
    border-radius: 5px;
  }
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
.new__btn {
  margin: 0 50%;
  transform: translateX(-50%);
  display: inline-block;
  & input {
    cursor: pointer;
    line-height: 50px;
    text-align: center;
    width: 100px;
    font-weight: bold;
    font-size: 18px;
    height: 50px;
    border-radius: 3px;
    background-color: $bg-main;
    color: white;
    &:hover {
      box-shadow: 0 0 5px grey;
    }
    &:active {
      box-shadow: 0 0 0 grey;
    }
  }
}
</style>
