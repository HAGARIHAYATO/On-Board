<template>
  <div class="works__wrapper">
    <h1 class="searchAlert" v-if="isSearch && works.length == 0">
      検索結果はありません
    </h1>
    <h1 class="searchAlert" v-if="!isSearch && works.length == 0">
      投稿はありません
    </h1>
    <nuxt-link
      :to="'/works/' + work.ID"
      v-for="(work, index) in works"
      :key="index"
      class="works__card"
    >
      <div>
        <p class="work__image"><img :src="returnURL(work.ImageURL)" :alt="work.Name" /></p>
        <div class="work__top">
          <p class="work__title">
            {{work.Name}}
          </p>
          <div class="user__box" v-if="!isUserShow">
            <img class="user__image" :src="returnURL(work.UserImageURL)" :alt="work.UserName" />
            <p class="user__name">{{work.UserName}}</p>
          </div>
        </div>
        <div class="work__bottom">
          <p
            v-for="(s, i) in work.Skills"
            :key="i"
            class="skill__tag"
          >{{s.Name}}</p>
          <div class="work__bottom__layer"></div>
        </div>
      </div>
    </nuxt-link>
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
@import "assets/scss/app";
* {
  text-decoration: none;
  box-sizing: border-box;
}
.works__wrapper {
  margin: 0 auto 20px auto;
  max-width: 870px;
  min-height: 435px;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
.works__card {
  & .work__image{
    & img {
      height: 100px;
      max-width: 200px;
      border-radius: 5px;
      display: block;
      margin: 0 0 10px 0;
    }
  }
  border: solid 3px white;
  overflow: hidden;
  padding: 1%;
  position: relative;
  margin: 2%;
  width: 400px;
  height: 250px;
  border-radius: 5px;
  background-color: white;
  box-shadow: 0 2px 3px grey;
  transition: all 0.3s;
  &:hover {
    transition: all 0.3s;
    box-shadow: 0 0 1px grey;
  }
}
.searchAlert {
  width: 100%;
  color: lightgrey;
  text-align: center;
  line-height: 435px;
  font-size: 40px;
}
.work__title{
  font-weight: bold;
  color: #777777 !important;
  font-size: 20px;
  height: 40px;
  line-height: 40px;
  margin-right: 10px;
}
.user__image{
  display: block;
  width: 30px !important;
  height: 30px !important;
  border-radius: 50% !important;
}
.user__name{
  color: #777777 !important;
  font-size: 20px;
  height: 30px;
  line-height: 30px;
  margin: 0 10px;
}
.user__box{
  display: inline-flex;
  border-radius: 25px;
  padding: 5px;
  box-shadow: 0 0 1px grey;
}
.work__top{
  display: flex;
  flex-wrap: wrap;
  margin: 0 0 0 10px;
}
.work__bottom{
  margin: 10px;
  position: relative;
  height: 90px;
}
.work__bottom__layer{
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background: linear-gradient(to top, rgba(255, 255, 255, 1), rgba(255, 255, 255, 0));
}
.skill__tag{
  display: inline-block;
  background-color: $tag-color;
  border-radius: 20px;
  color: white;
  padding: 4px 20px;
  font-weight: bold;
  box-shadow: 0 1px 2px grey;
  margin: 5px;
}
</style>
