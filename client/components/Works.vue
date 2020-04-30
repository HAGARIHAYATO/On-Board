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
      <img class="user__image" v-if="!isUserShow" :src="returnURL(work.UserImageURL)" :alt="work.UserName" />
      <p class="speech__bubble">
        <nuxt-link :to="'/works/' + work.ID">
          {{ hideLongTitle(work.Name) }}
        </nuxt-link>
      </p>
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
    },
    isUserShow: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    returnURL: function(url) {
      return url ? url : "/NO_IMAGE.jpeg";
    },
    hideLongTitle: function(title) {
      if (title.length > 4) {
        return title.slice( 0, 3 ) + "..."
      }
      return title
    }
  }
};
</script>
<style lang="scss" scoped>
* {
  text-decoration: none;
  box-sizing: border-box;
}
.speech__bubble{
  position: absolute;
  min-width: 80px;
  text-align: center;
  left: 55%;
  bottom: -5px;
  display: none;
  border: solid 3px #192b3d;
  border-radius: 20px;
  padding: 0 3%;
  background-color: lightgreen;
  & a {
    color: white !important;
  }
}
.user__image{
  position: absolute;
  bottom: -5px;
  left: 0;
  display: block;
  width: 50px !important;
  height: 50px !important;
  border: solid 3px #192b3d;
  border-radius: 20px;
}
.works__wrapper {
  margin: 0 auto 20px auto;
  max-width: 870px;
  min-height: 435px;
  display: flex;
  flex-wrap: wrap;
}

.works__card {
  & img {
    width: 100%;
    height: 100%;
    border-radius: 50%;
  }
  position: relative;
  margin: 1%;
  width: 200px;
  height: 200px;
  border-radius: 50%;
  border: solid 3px #192b3d;
  background-color: white;
  box-shadow: 0 0 5px #192b3d;
  transition: all 0.1s;
  &:hover {
    transition: all 0.5s;
    box-shadow: 0 0 0 black;
    & .speech__bubble {
      display: inline-block !important;
    }
  }
}
.searchAlert {
  width: 100%;
  color: lightgrey;
  text-align: center;
  line-height: 40vh;
  font-size: 40px;
}
</style>
