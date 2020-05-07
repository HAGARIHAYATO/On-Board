<template>
  <div class="show__wrapper">
    <Loading v-if="isLoading" />
    <div class="show__container">
      <div class="show__main">
        <a :href="work.URL" class="show__container__image">
          <img :src="returnURL(work.ImageURL)" :alt="work.Name" class="show__image" />
        </a>
      </div>
      <h2>Information - 作品情報</h2>
      <div class="show__bar">
        <div class="show__info">
          <h3 class="info__title"><span class="grey_span">TITLE : </span>{{ work.Name }}</h3>
          <p class="info__title__sub">{{ work.URL }}</p>
          <div class="show__user">
            <nuxt-link :to="'/users/' + work.UserID" class="info__name">
              <img :src="returnURL(work.UserImageURL)" class="user__icon">
              <p>{{ work.UserName }}</p>
            </nuxt-link>
          </div>
          <div class="show__description">{{ isNone(work.Description) }}</div>
        </div>
        <div class="show__subInfo">
          <p
            class="skill__tag"
            v-for="(skill, index) in Skills"
            :key="index"
          >
            {{skill}}
          </p>
        </div>
      </div>
      <h2>Item - 作品詳細</h2>
      <div class="select__item" v-if="isArrayExist(work.WorkItems)">
        <div class="select__item__image" v-if="selectItem">
          <img :src="returnURL(selectItem.ImageURL)" alt="作品画像" />
        </div>
        <div class="select__item__content">
          <p>{{ isNone(selectItem.Body) }}</p>
        </div>
      </div>
      <div class="show__container__main">
        <div
          v-for="(item, index) in work.WorkItems"
          :key="index"
          class="item__wrapper"
          @click="showWorksItem(item)"
          :class="isSelected(item.ID)"
        >
          <p v-if="isSelected(item.ID)">▼</p>
          <p v-else>-</p>
          <img :src="returnURL(item.ImageURL)" alt="作品画像" />
        </div>
      </div>
      <h2>Composition - 構図</h2>
      <div class="">
        <img
          v-if="isExist(work.CacooURL)"
          class="cacoo__frame"
          :src="isExist(work.CacooURL)"
        />
        <div
         v-else
         class="cacoo__frame"
        >
         <p class="nothing__alert">構図はありません。</p>
        </div>
      </div>
      <p class="operation" v-if="isMine">
        <nuxt-link :to="'/works/' + work.ID + '/edit'">作品編集</nuxt-link>|
        <nuxt-link :to="'/works/' + work.ID + '/edit_item'">アイテム編集</nuxt-link>
      </p>
      <button class="deleteBtn" v-if="isMine" @click="openDeleteModal">削除する</button>
      <Delete-Modal
        v-if="isOpenDeleteModal"
        @delete="deleteWork"
        @back="closeDeleteModal"
        confirmStr="削除してしまうと復元することはできません。よろしいですか?"
        btnStr="同意して削除する"
      />
    </div>
  </div>
</template>
<script>
import Loading from "~/components/Loading.vue";
import DeleteModal from "~/components/DeleteModal.vue";
export default {
  components: {
    Loading,
    DeleteModal
  },
  data() {
    return {
      selectId: "",
      selectItem: {},
      work: {},
      isLoading: false,
      APIURL: "",
      isOpenDeleteModal: false,
      /////////////////////////
      Skills: [
        "Docker",
        "AWS",
        "Kubernetes",
        "Azure",
        "Next.js",
        "Rust"
      ],
    };
  },
  head () {
    return {
      title: "OnBoard / " + this.work.Name,
      meta: [
        { hid: this.work.ID, name: this.work.Name, content: this.work.Description }
      ]
    }
  },
  computed: {
    isMine: function() {
      if (!this.$auth.user || !this.work.UserID) return;
      if (this.work.UserID === this.$auth.user.ID) {
        return true;
      }
      return false;
    }
  },
  methods: {
    isExist: function(str) {
      return str ? str : false
    },
    isNone: function(str) {
      return str == "" ? "本文はありません。" : str
    },
    closeDeleteModal: function() {
      this.isOpenDeleteModal = false;
    },
    openDeleteModal: function() {
      this.isOpenDeleteModal = true;
    },
    deleteWork: function() {
      this.$nextTick(async () => {
        this.APIURL = this.GetURL();
        await this.showBubble();
        await this.$axios
          .delete(this.APIURL + this.$route.path)
          .then(response => {
            this.$router.push("/works");
          })
          .catch(response => console.error(response));
        await setTimeout(() => this.showBubble(), 1000);
      });
      this.isOpenDeleteModal = false;
    },
    isArrayExist: function(array) {
      if (array) {
        return array.length > 0 ? true : false;
      }
    },
    isSelected: function(id) {
      if (this.selectItem.ID === id) {
        return "selected-item";
      }
    },
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    returnURL: function(url) {
      return url ? url : "/NO_IMAGE.jpeg";
    },
    showWorksItem: function(item) {
      if (typeof item != "object") return;
      this.selectId = item.ID;
      this.selectItem = item;
    }
  },
  async mounted() {
    this.$nextTick(async () => {
      this.APIURL = this.GetURL();
      await this.showBubble();
      await this.$axios
        .get(this.APIURL + this.$route.path)
        .then(response => {
          this.work = response.data;
          this.selectItem = response.data.WorkItems[0]
        })
        .catch(response => console.error(response));
      await setTimeout(() => this.showBubble(), 1000);
    });
  }
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
*,
body {
  box-sizing: border-box;
}
h2 {
  margin: 50px;
  color: $bg-main;
  font-size: 18px;
  border-bottom: solid 2px $bg-main;
}
.show__wrapper {
  padding: 100px 0 50px 0;
  width: 100%;
  min-height: 81vh;
  background-color: $bg-color;
}
.show__container{
  margin: 30px;
  box-shadow: 0 2px 5px grey;
  background-color: white;
  border-radius: 10px;
  padding: 4% 2%;
}
.show__bar {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  width: 70%;
  margin: 30px auto 50px auto;
}
.show__description {
  font-size: 10px;
  text-align: left !important;
  padding: 20px;
  min-height: 20vh;
  min-width: 300px;
  color: grey;
  word-break: break-all;
  background-color: white;
  box-shadow: 0 0 2px grey;
}
.show__container__image {
  min-width: 200px;
  height: 300px;
}
.show__image {
  display: block;
  margin: 0 auto;
  max-width: 600px;
  height: 300px;
  box-shadow: 0 0 5px grey;
  &:hover{
    box-shadow: 0 0 0 grey;
  }
}
.show__user {
  text-align: left;
  margin: 0 0 34px 0;
  max-width: 400px;
}
.info__name {
  display: inline-flex;
  min-width: 40%;
  background-color: white;
  border-radius: 25px !important;
  box-shadow: 0 0 1px grey;
  padding: 0 4px;
  margin: 0 auto;
  height: 40px;
  word-break: break-all !important;
  text-decoration: none !important;
  color: $bg-main;
  & p {
    margin: 0 10px;
    line-height: 40px;
  }
  &:hover {
    box-shadow: 0 0 5px grey;
  }
}
.user__icon{
  margin: 5px 2px;
  width: 30px;
  height: 30px;
  border-radius: 50%;
  border: solid 1px grey;
}
.show__container__main {
  height: 130px;
  display: flex;
  justify-content: center;
  border-radius: 5px;
  & img {
    border-radius: 50%;
    width: 70px;
    height: 70px;
    margin: 0 20px;
  }
}
.item__wrapper {
  height: 130px;
  & p {
    transition: all .3s;
    color: white;
    text-align: center;
  }
  & img {
    cursor: pointer;
    box-shadow: 0 0 4px grey;
  }
  &:hover {
    & p {
      color: black;
    }
    & img {
      box-shadow: 0 0 7px grey;
    }
  }
  &:active {
    & p {
      color: black;
    }
    & img {
      box-shadow: 0 0 1px grey;
      transform: scale(0.97);
    }
  }
}
.select__item {
  margin: 30px auto 0 auto;
  width: 70%;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  border-radius: 5px;
  transition: all 0.3s;
}
.select__item__image {
  display: inline-block;
  height: 350px;
  & img {
    height: 350px;
  }
}
.select__item__content {
  background-color: white;
  border-radius: 3px;
  width: 70%;
  min-height: 50px;
  padding: 10px 2%;
  & p {
    word-break: break-all;
    font-size: 10px;
    color: grey;
    padding: 2%;
    margin: 0 auto;
    min-width: 99.5%;
    min-height: 99%;
  }
}
.selected-item {
  & img {
    box-shadow: 0 0 2px grey;
    transform: scale(0.99);
    border: solid 2px black;
  }
  &:hover {
    & img {
      box-shadow: 0 0 2px grey;
      transform: scale(0.97);
    }
  }
}
.close-modal {
  transform: scale(0);
}
.operation {
  text-align: center;
  font-weight: bold;
}
.deleteBtn {
  margin: 10px auto;
  height: 40px;
  width: 180px;
  display: block;
  outline: none;
  background-color: red;
  color: white;
  border-radius: 5px;
  font-weight: bold;
  &:hover {
    box-shadow: 0 0 5px grey;
  }
}
.info__title {
  text-align: left;
  font-size: 24px;
  margin: 0 0 0 0;
}
.info__title__sub {
  text-align: left;
  margin: 0 0 10px 0;
  font-size: 8px;
  color: grey;
}
.show__main, .show__info{
  margin: 0 25px;
}
.grey_span{
  color: grey;
  font-size: 16px;
}
.show__subInfo{
  border-radius: 5px;
  width: 300px;
  margin: 0 20px 0 20px;
  padding: 10px;
}
.skill__tag{
  display: inline-block;
  background-color: $tag-color;
  border-radius: 20px;
  color: white;
  padding: 4px 15px;
  font-weight: bold;
  box-shadow: 0 0 5px grey;
  margin: 4px;
}
.cacoo__frame{
  display: block;
  width: 70%;
  margin: 50px auto;
}
.nothing__alert{
  text-align: center;
  line-height: 528px;
  color: grey;
  font-weight: bold;
  font-size: 30px;
}
</style>
