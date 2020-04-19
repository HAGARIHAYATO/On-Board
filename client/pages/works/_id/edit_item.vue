<template>
  <div class="item__wrapper">
    <Loading v-if="isLoading" />
    <h1 class="work__name">{{ work.Name }}</h1>
    <form class="form" @submit.prevent="submit">
      <div class="item__container">
        <div class="item">
          <label for="image1">
            <input
              id="image1"
              accept="image/png, image/jpg, image/jpeg"
              type="file"
              @change="setImage($event, data1)"
            />
            <p>クリックor画像をドロップ</p>
          </label>
          <div class="set__file" v-if="data1.image">
            <p @click="resetImage($event, data1)">×</p>
            <img :src="data1.image" />
          </div>
          <textarea v-model="data1.body"></textarea>
        </div>
        <div class="item">
          <label for="image2">
            <input
              id="image2"
              type="file"
              accept="image/png, image/jpg, image/jpeg"
              @change="setImage($event, data2)"
            />
            <p>クリックor画像をドロップ</p>
          </label>
          <div class="set__file" v-if="data2.image">
            <p @click="resetImage($event, data2)">×</p>
            <img :src="data2.image" />
          </div>
          <textarea v-model="data2.body"></textarea>
        </div>
        <div class="item">
          <label for="image3">
            <input
              id="image3"
              type="file"
              accept="image/png, image/jpg, image/jpeg"
              @change="setImage($event, data3)"
            />
            <p>クリックor画像をドロップ</p>
          </label>
          <div class="set__file" v-if="data3.image">
            <p @click="resetImage($event, data3)">×</p>
            <img :src="data3.image" />
          </div>
          <textarea v-model="data3.body"></textarea>
        </div>
        <div class="item">
          <label for="image4">
            <input
              id="image4"
              type="file"
              accept="image/png, image/jpg, image/jpeg"
              @change="setImage($event, data4)"
            />
            <p>クリックor画像をドロップ</p>
          </label>
          <div class="set__file" v-if="data4.image">
            <p @click="resetImage($event, data4)">×</p>
            <img :src="data4.image" />
          </div>
          <textarea v-model="data4.body"></textarea>
        </div>
        <div class="item">
          <label for="image5">
            <input
              id="image5"
              type="file"
              accept="image/png, image/jpg, image/jpeg"
              @change="setImage($event, data5)"
            />
            <p>クリックor画像をドロップ</p>
          </label>
          <div class="set__file" v-if="data5.image">
            <p @click="resetImage($event, data5)">×</p>
            <img :src="data5.image" />
          </div>
          <textarea v-model="data5.body"></textarea>
        </div>
        <div class="item">
          <label for="image6">
            <input
              id="image6"
              type="file"
              accept="image/png, image/jpg, image/jpeg"
              @change="setImage($event, data6)"
            />
            <p>クリックor画像をドロップ</p>
          </label>
          <div class="set__file" v-if="data6.image">
            <p @click="resetImage($event, data6)">×</p>
            <img :src="data6.image" />
          </div>
          <textarea v-model="data6.body"></textarea>
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
      data1: {
        id: null,
        image: "",
        name: "",
        body: ""
      },
      data2: {
        id: null,
        image: "",
        name: "",
        body: ""
      },
      data3: {
        id: null,
        image: "",
        name: "",
        body: ""
      },
      data4: {
        id: null,
        image: "",
        name: "",
        body: ""
      },
      data5: {
        id: null,
        image: "",
        name: "",
        body: ""
      },
      data6: {
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
          this.insertData(this.data1, 0);
          this.insertData(this.data2, 1);
          this.insertData(this.data3, 2);
          this.insertData(this.data4, 3);
          this.insertData(this.data5, 4);
          this.insertData(this.data6, 5);
        })
        .catch(response => console.error(response));
      await setTimeout(() => this.showBubble(), 1000);
    });
  },
  methods: {
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    setImage: function(e, data) {
      if (data) {
        const fileImg = e.target.files[0];
        if (fileImg.size > 3000000) {
          alert(this.AlertMessage());
          return;
        }
        if (fileImg.type.startsWith("image/")) {
          data.image = window.URL.createObjectURL(fileImg);
          data.name = e.target.files[0];
        }
      }
    },
    resetImage(e, data) {
      if (data) {
        e.path[2].childNodes[0].children[0].value = "";
        data.image = "";
        data.name = "";
      }
    },
    async pushData(data) {
      const json = JSON.stringify(
        {
          id: data.id,
          image: data.name,
          body: data.body
        },
        undefined,
        1
      );
      return json;
    },
    async submit() {
      this.$nextTick(async () => {
        await this.showBubble();
        try {
          const data = new FormData();
          data.append("data1", this.pushData(this.data1));
          data.append("data2", this.pushData(this.data2));
          data.append("data3", this.pushData(this.data3));
          data.append("data4", this.pushData(this.data4));
          data.append("data5", this.pushData(this.data5));
          data.append("data6", this.pushData(this.data6));
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
    },
    insertData: function(data, i) {
      const item = this.work.WorkItems;
      if (item[i]) {
        data.image = item[i].ImageURL;
        data.body = item[i].Body;
        data.id = item[i].ID;
      }
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
  min-height: 65vh;
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
</style>
