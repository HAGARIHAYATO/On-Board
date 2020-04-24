<template>
  <div class="works__wrapper">
    <h1 class="searchAlert" v-if="isSearch && works.length == 0">
      検索結果はありません
    </h1>
    <h1 class="searchAlert" v-if="!isSearch && works.length == 0">
      投稿はありません
    </h1>
    <div v-for="(work, index) in works" :key="index" class="works__card">
      <nuxt-link :to="'/works/' + work.ID">
        <img :src="returnURL(work.ImageURL)" :alt="work.Name" />
      </nuxt-link>
      <nuxt-link :to="'/users/' + work.UserID">
        <div class="works__card__user">
          <img :src="returnURL(work.UserImageURL)" :alt="work.UserName" />
          <p>{{ work.UserName }}</p>
        </div>
      </nuxt-link>
      <nuxt-link :to="'/works/' + work.ID">
        <p class="work__card__title">{{ work.Name }}</p>
      </nuxt-link>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    works: {
      type: Array
    },
    isSearch: {
      type: Boolean
    }
  },
  methods: {
    returnURL: function(url) {
      return url ? url : "/NO_IMAGE.jpeg";
    }
  }
};
</script>
<style lang="scss" scoped>
* {
  text-decoration: none;
  box-sizing: border-box;
}
.works__wrapper {
  margin: 0 auto 20px auto;
  max-width: 960px;
  min-height: 640px;
  display: flex;
  flex-wrap: wrap;
}
.works__card {
  position: relative;
  & img {
    width: 200px;
    height: 200px;
    border-radius: 7px 7px 0 0;
  }
  margin: 1% auto;
  width: 200px;
  height: 300px;
  border-radius: 7px;
  background-color: white;
  box-shadow: 0 0 5px grey;
  transition: all 0.5s;
  &:hover {
    transition: all 0.5s;
    box-shadow: 0 0 12px darken(grey, 20%);
  }
}
.works__card__user {
  position: absolute;
  top: 205px;
  height: 40px;
  font-size: 10px;
  font-weight: bold;
  color: #192b3d !important;
  display: flex;
  & p {
    line-height: 40px;
    &:hover {
      text-shadow: 0 0 2px grey;
    }
  }
  & img {
    display: inline-block;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    margin: 0 10px 0 10px;
  }
  &:hover {
    & p {
      color: grey;
    }
  }
}
.work__card__title {
  position: absolute;
  top: 260px;
  left: 50%;
  transform: translateX(-50%);
  font-weight: bold;
  word-break: break-all;
  color: #192b3d !important;
  font-size: 12px;
  text-align: center;
  &:hover {
    text-shadow: 0 0 2px grey;
  }
}
.searchAlert {
  width: 100%;
  color: white;
  text-align: center;
  line-height: 80vh;
  font-size: 40px;
}
</style>
