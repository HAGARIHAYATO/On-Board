<template>
  <form @submit.prevent="submit" class="edit__wrapper">
    <Loading v-if="isLoading" />
    <div class="edit__user__bar">
      <div class="edit__user__bar__top">
        <div class="edit__url">
          <div class="edit__container__image">
            <label for="edit__image" class="edit__image">
              クリックor画像をドラッグ
              <span>+</span>
              <input
                name="work_image"
                id="edit__image"
                type="file"
                ref="file"
                accept="image/png, image/jpg, image/jpeg"
                @change="setImage"
              />
            </label>
            <div class="edit__file" v-if="data.image">
              <p @click="resetImage">×</p>
              <img :src="data.image" />
            </div>
          </div>
          <div class="edit__url__info">
            <p class="info__name__sub">作品名</p>
            <input name="work_name" type="text" class="info__name" v-model="workName" />
            <p class="info__name__sub">作品URL</p>
            <input name="work_url" type="url" class="info__name" v-model="workURL" />
          </div>
        </div>
      </div>
      <textarea
        class="edit__description"
        name="work_description"
        cols="30"
        rows="10"
        v-model="workDesc"
      ></textarea>
    </div>
    <div class="edit__btn">
      <input type="submit" value="保存" />
    </div>
    <p class="operation">
      <nuxt-link :to="'/works/' + work.ID">戻る</nuxt-link>|
      <nuxt-link :to="'/works/' + work.ID + '/edit_item'">アイテム編集</nuxt-link>
    </p>
  </form>
</template>
<script>
import Loading from "~/components/Loading.vue";
export default {
  components: {
    Loading
  },
  middleware: ["auth"],
  data() {
    return {
      data: {
        image: "",
        name: ""
      },
      workName: "",
      workURL: "",
      workDesc: "",
      selectWindow: false,
      selectId: "",
      selectItem: {},
      work: {},
      APIURL: "",
      isLoading: false
    };
  },
  async mounted() {
    this.$nextTick(async () => {
      await this.showBubble();
      this.APIURL = this.GetURL();
      const url = this.$route.path.slice(0, -5);
      await this.$axios
        .get(this.APIURL + url)
        .then(response => {
          this.work = response.data;
          this.workName = response.data.Name;
          this.workURL = response.data.URL;
          this.workDesc = response.data.Description;
          if (response.data.ImageURL) {
            this.data.image = response.data.ImageURL;
            this.data.name = response.data.Name;
          }
        })
        .catch(response => console.error(response));
      await setTimeout(() => this.showBubble(), 1000);
    });
  },
  methods: {
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    submit: async function() {
      this.$nextTick(async () => {
        await this.showBubble();
        try {
          const data = new FormData();
          data.append("name", this.workName);
          data.append("description", this.workDesc);
          data.append("url", this.workURL);
          data.append("file", this.data.name);
          const headers = { "content-type": "multipart/form-data" };
          await this.$axios
            .put(this.APIURL + "/works/" + this.work.ID, data, {
              headers
            })
            .then(res => {
              if (res.data) {
                this.$router.push("/works/" + res.data.ID);
              }
            });
        } catch (error) {
          // handling
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
*,
body {
  box-sizing: border-box;
}
.edit__wrapper {
  overflow-x: hidden;
  padding: 100px 0 50px 0;
  width: 100%;
  min-height: 81vh;
}
.edit__user__bar {
  width: 600px;
  text-align: center;
  margin: 0 auto;
}
.edit__description {
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
.edit__container__image {
  margin: 0 auto;
  min-width: 200px;
  height: 200px;
}
.edit__url {
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
      & .edit__image {
        box-shadow: 0 0 5px grey;
      }
    }
    &:active {
      & .edit__image {
        box-shadow: 0 0 1px grey;
      }
    }
  }
  text-align: center;
}
.edit__url__info {
  text-align: left;
  width: 400px;
  padding: 2% 0 2% 2%;
}
.info__name__sub {
  margin-left: 10px;
  font-size: 8px;
  font-weight: bold;
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
    color: #192b3d;
    &:hover {
      font-weight: bold;
    }
  }
}
.edit__image {
  display: block;
  margin: 0 auto;
  height: 200px;
  outline: 0;
  border-radius: 5px;
  padding: 40% 0;
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
.edit__file {
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
      color: #fdeaa0;
    }
    &:active {
      color: silver;
    }
  }
}
.edit__btn {
  margin: 0 50%;
  transform: translateX(-50%);
  display: inline-block;
  & input {
    width: 100px;
    font-weight: bold;
    font-size: 18px;
    height: 50px;
    border-radius: 3px;
    background-color: #192b3d;
    color: white;
  }
}
.operation {
  text-align: center;
  font-weight: bold;
}
</style>
