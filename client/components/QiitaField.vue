<template>
  <div class="field__container">
    <h2 class="field__title">Qiita Account</h2>
    <p class="qiita__user"><img :src="qiitaUser.profile_image_url"><a :href="'https://qiita.com/' + qiitaUser.id">{{qiitaUser.name}}</a></p>
    <h2 class="field__title">Qiita Analysis</h2>
    <div class="chart__box">
      <Chart :dataset="dataset" text="タグ使用率"  :height="calcSize()" :width="calcSize()"/>
    </div>
    <h2 class="field__title">Qiita Articles <span class="articles__count">{{qiitaUser.items_count}}件</span></h2>
    <div class="field">
      <div v-for="(article, index) in articles" :key="index" class="field__item">
        <p class="field__article"><a :href="article.url">＞ {{article.title}}</a></p>
        <p
          v-for="(tag, i) in article.tags"
          :key="i"
          class="field__tags"
        >{{tag.name}}</p>
        <p class="articles__count"><img src="/star.svg" alt="star" />{{article.likes_count}}</p>
      </div>
    </div>
  </div>
</template>
<script>
import Chart from "~/components/Chart.vue";
export default {
  components: {
    Chart
  },
  props: {
    articles: {
      type: Array
    },
    qiitaUser: {
      type: Object
    }
  },
  data() {
    return {
      dataset: {}
    }
  },
  async mounted() {
    this.chartInit()
  },
  methods: {
    chartInit: function() {
      let obj = {}
      if (this.articles) {
        for (const article of this.articles) {
          for (const tag of article.tags) {
            if (obj[tag.name]) {
              obj[tag.name] ++
            } else {
              obj[tag.name] = 1
            }
          }
        }
        this.dataset = obj
      }
    }
  },
}
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
.field__container{
  padding: 2%;
  background-color: white;
  border-radius: 10px;
  border: solid .5px lightgrey;
  margin: 30px;
  // box-shadow: 0 2px 5px grey;
}
.field__item {
  margin : 10px 30px;
  display: flex;
  & a {
    color: #777777 !important;
    text-decoration: none;
  }
}
.field__tags{
  border-radius: 20px;
  color: $bg-yellow;
  background-color: $bg-main;
  padding: 0 10px;
  margin: 0 5px;
}
.field__article{
  font-weight: bold;
  transition: all .3s;
  &:hover{
    margin: 0 0 0 10px;
  }
}
.field__title{
  margin: 0 10px 10px 10px;
  color: $bg-main;
  font-size: 18px;
}
.qiita__user{
  margin: 0 0 0 50px;
  height: 50px;
  & img {
    height: 30px;
    width: 30px;
  }
  & a {
    color: #777777 !important;
    margin: 0 10px;
    font-size: 30px;
    font-weight: bold;
    text-decoration: none !important;
  }
}
.articles__count{
  color: #777;
  font-size: 12px;
  margin-left: 10px;
  border-radius: 20px;
  font-weight: bold;
  & img {
    width: 22px;
    height: 22px;
  }
}
.chart__box{
  display: flex;
  flex-wrap: wrap;
  padding: 0 20px 20px 20px;
  & canvas {
    width: 100%;
    height: 100%;
  }
}
@media screen and (max-width: $PhoneSize) {
  .field__container{
    width: 98%;
    margin: 30px auto;
  }
  .qiita__user{
    margin:  5px 10px;
    & img {
      height: 20px;
      width: 20px;
    }
    & a {
      font-size: 14px;
      word-break: break-all;
    }
  }
  .chart__box{
    width: 100%;
    padding: 0;
    & canvas {
      width: 98% !important;
    }
  }
  .field__tags {
    display: none;
  }
  .articles__count{
    color: #777;
    font-size: 10px;
    & img {
      width: 10px;
      height: 10px;
    }
  }
}
</style>