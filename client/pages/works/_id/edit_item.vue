<template>
  <div class="item__wrapper">
    <h1 class="work__name">{{ work.Name }}</h1>
    <form class="form" @submit.prevent="submit">
      <div class="item__container">
        <div class="item">
          <label for="image1">
            <input id="image1" type="file" @change="setImage($event, data1)" />
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
            <input id="image2" type="file" @change="setImage($event, data2)" />
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
            <input id="image3" type="file" @change="setImage($event, data3)" />
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
            <input id="image4" type="file" @change="setImage($event, data4)" />
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
            <input id="image5" type="file" @change="setImage($event, data5)" />
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
            <input id="image6" type="file" @change="setImage($event, data6)" />
            <p>クリックor画像をドロップ</p>
          </label>
          <div class="set__file" v-if="data6.image">
            <p @click="resetImage($event, data6)">×</p>
            <img :src="data6.image" />
          </div>
          <textarea v-model="data6.body"></textarea>
        </div>
      </div>
      <input type="submit" @submit="submit" value="保存" />
    </form>
    <p class="operation">
      <nuxt-link :to="'/works/' + work.ID">戻る</nuxt-link> |
      <nuxt-link :to="'/works/' + work.ID + '/edit'">作品編集</nuxt-link>
    </p>
  </div>
</template>
<script>
export default {
  // middleware: ["auth"],
  data() {
    return {
      work: {},
      data1: {
        image: "",
        body: ""
      },
      data2: {
        image: "",
        body: ""
      },
      data3: {
        image: "",
        body: ""
      },
      data4: {
        image: "",
        body: ""
      },
      data5: {
        image: "",
        body: ""
      },
      data6: {
        image: "",
        body: ""
      },
      APIURL: "http://localhost:8080/api/v1"
    };
  },
  async mounted() {
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
  },
  methods: {
    setImage: function(e, data) {
      if (data) {
        const fileImg = e.target.files[0];
        if (fileImg.type.startsWith("image/")) {
          data.image = window.URL.createObjectURL(fileImg);
        }
      }
    },
    resetImage(e, data) {
      if (data) {
        e.path[2].childNodes[0].children[0].value = "";
        data.image = "";
      }
    },
    async submit() {
      try {
        const data = new FormData();
        data.append("file1", this.data1.image);
        data.append("body1", this.data1.body);
        data.append("file2", this.data2.image);
        data.append("body2", this.data2.body);
        data.append("file3", this.data3.image);
        data.append("body3", this.data3.body);
        data.append("file4", this.data4.image);
        data.append("body4", this.data4.body);
        data.append("file5", this.data5.image);
        data.append("body5", this.data5.body);
        data.append("file6", this.data6.image);
        data.append("body6", this.data6.body);
        const headers = { "content-type": "multipart/form-data" };
        await this.$axios
          .post(this.APIURL + "/works/" + this.work.ID, data, {
            headers
          })
          .then(res => console.log(res));
      } catch (error) {
        // handling
      }
    },
    insertData: function(data, i) {
      const item = this.work.WorkItems;
      if (item[i]) {
        data.image = item[i].ImageURL;
        data.body = item[i].Body;
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
