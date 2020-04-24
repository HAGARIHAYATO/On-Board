<template>
  <div class="item__wrapper">
    <Loading v-if="isLoading" />
    <h1 class="work__name">{{ work.Name }}</h1>
    <form class="form" @submit.prevent="submit">
      <div class="image__container">
        <div class="image__item" :class="isSelectedItem(index)" @click="selectWorkItem(index)" v-for="(item,index) in items" :key="index">
          <img :src="returnURL(item.ImageURL)">
        </div>
      </div>
      <div class="item__container">
        <div class="item">
          <label for="image">
            <input
              id="image"
              accept="image/png, image/jpg, image/jpeg"
              type="file"
              ref="file"
              @change="setImage"
            />
            <p>クリックor画像をドロップ</p>
          </label>
          <div class="set__file" v-if="data.image">
            <p @click="resetImage">×</p>
            <img :src="data.image" />
          </div>
          <textarea v-model="data.body"></textarea>
        </div>
      </div>
      <input type="submit" value="保存" />
    </form>
    <p class="operation">
      <nuxt-link :to="'/works/' + work.ID">戻る</nuxt-link>|
      <nuxt-link :to="'/works/' + work.ID + '/edit'">作品編集</nuxt-link>
    </p>
  </div>
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
      work: {},
      items: [],
      selectedIndex: 0,
      data: {
        id: null,
        image: "",
        name: "",
        body: ""
      },
      APIURL: "",
      isLoading: false
    };
  },
  async mounted() {
    this.$nextTick(async () => {
      await this.showBubble();
      this.APIURL = this.GetURL();
      const url = this.$route.path.slice(0, -10);
      await this.$axios
        .get(this.APIURL + url)
        .then(response => {
          this.work = response.data;
          this.items = response.data.WorkItems
          if (this.data) {
            this.data.id = this.items[0].ID
            this.data.body = this.items[0].Body
            this.data.image = this.items[0].ImageURL
          }
        })
        .catch(response => console.error(response));
      await setTimeout(() => this.showBubble(), 1000);
    });
  },
  methods: {
    isSelectedItem: function(i) {
      return this.selectedIndex === i ? "item__selected" : ""
    },
    returnURL: function(url) {
      return url ? url : "/NO_IMAGE.jpeg";
    },
    selectWorkItem: function(i) {
      this.selectedIndex = i
      this.data.id = this.items[i].ID
      this.data.image = this.items[i].ImageURL
      this.data.body = this.items[i].Body
      this.data.name = ""
    },
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    setImage: function(e) {
      if (this.data) {
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
      }
    },
    resetImage(e) {
        this.$refs.file.value = "";
        this.data.name = "";
        this.data.image = "";
    },
    async submit() {
      this.$nextTick(async () => {
        await this.showBubble();
        try {
          const data = new FormData();
          data.append("id", this.data.id);
          data.append("body", this.data.body);
          data.append("uid", this.work.UserID);
          data.append("file", this.data.name);
          const headers = { "content-type": "multipart/form-data" };
          await this.$axios
            .put(this.APIURL + "/works/" + this.work.ID + "/edit_item", data, {
              headers
            })
            .then(res => {
              this.$router.push("/works/" + res.data.ID);
            })
            .catch(res => console.error(res));
        } catch (error) {
          // handling
        }
        await setTimeout(() => this.showBubble(), 1000);
      });
    }
  }
};
</script>
<style lang="scss" scoped>
*,
body {
  box-sizing: border-box;
  text-decoration: none;
}
.item__wrapper {
  margin: 90px 0 90px 0;
}
.work__name {
  text-align: center;
}
.item__container {
  max-width: 1000px;
  margin: 0 auto;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}
.form {
  margin: 20px auto 0 auto;
}
.operation {
  text-align: center;
  font-weight: bold;
}
.item {
  position: relative;
  width: 300px;
  border-radius: 20px;
  box-shadow: 0 0 5px black;
  margin: 20px auto;
  padding: 2%;
  & label {
    text-align: center;
    line-height: 60px;
    height: 60px;
    width: 100%;
    border: dotted 1px black;
    display: block;
    & input {
      display: none;
    }
  }
  & textarea {
    margin: 20px auto 0 auto;
    width: 100%;
    min-height: 120px;
    border-radius: 3px;
    padding: 1%;
  }
}
input[type="submit"] {
  display: block;
  margin: 20px auto 0 auto;
  width: 100px;
  font-weight: bold;
  font-size: 18px;
  height: 50px;
  border-radius: 3px;
  background-color: #192b3d;
  color: white;
}
.set__file {
  z-index: 2;
  top: 20px;
  position: absolute;
  margin: 0 !important;
  & img {
    z-index: 2;
    min-width: 60px;
    height: 60px;
    border-radius: 5px;
  }
  & p {
    margin: 0 !important;
    position: absolute;
    top: -5px;
    left: -5px;
    width: 20px;
    height: 20px;
    text-align: center;
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
.image__container{
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  width: 200px;
  margin: 40px auto;
}
.image__item{
  cursor: pointer;
  &:hover{
    & img{
      box-shadow: 0 0 5px grey;
    }
  }
  & img{
    height: 60px;
  }
}
.item__selected{
  & img {
    border: solid 3px black;
    transform: scale(0.95);
    transition: all .2s;
  }
  &:hover{
    & img {
      box-shadow: 0 0 0 black !important;
    }
  }
}
</style>
