<template>
  <div class="field__container">
    <h2 class="field__title" v-if="ghUser.type === 'User'">GitHub Account</h2>
    <h2 class="field__title" v-else>GitHub Team</h2>
    <p class="gh__user"><img :src="ghUser.avatar_url"><a :href="ghUser.html_url">{{ghUser.login}}</a></p>
    <h2 class="field__title">Github Analysis</h2>
    <div class="chart__box">
      <Chart :dataset="dataset" text="言語別使用率" :height="calcSize()" :width="calcSize()" />
    </div>
    <h2 class="field__title">GitHub Projects <span class="hubs__count">{{ghUser.public_repos}}件</span></h2>
    <div class="field">
      <div v-for="(item, index) in hubs" :key="index" class="field__item">
        <GitHubItem :platform="0" :title="item.name" :url="item.html_url" />
      </div>
    </div>
  </div>
</template>
<script>
import GitHubItem from "~/components/GitHubItem.vue";
import Chart from "~/components/Chart.vue";
export default {
  components: {
    GitHubItem,
    Chart
  },
  props: {
    hubs: {
      type: Array
    },
    ghUser: {
      type: Object
    }
  },
  data() {
    return {
      dataset: {}
    }
  },
  mounted() {
    this.chartInit()
  },
  methods: {
    chartInit: function() {
      let obj = {}
      if (this.hubs) {
        for (const hub of this.hubs) {
          if (obj[hub.language]) {
            obj[hub.language] ++
          } else {
            obj[hub.language] = 1
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
  box-shadow: 0 2px 5px grey;
}
.field{
  display: flex;
  justify-content: flex-start;
  flex-wrap: wrap;
}
.field__item {
  margin :6px;
}
.field__title{
  margin: 0 10px 10px 10px;
  color: $bg-main;
  font-size: 18px;
}
.gh__user{
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
.hubs__count{
  color: #777;
  font-size: 12px;
  margin-left: 10px;
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
  .gh__user{
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
}
</style>