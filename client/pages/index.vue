<template>
  <div>
    <div class="container">
      <div>
        <h1 class="container__title">未知の作品と出会おう</h1>
        <button class="container__button">
          <nuxt-link to="/works">作品を探す</nuxt-link>
        </button>
      </div>
    </div>
    <div class="container">
      <div class="block__content shadow__none">
        <Doughnut-Chart :dataset="dataset" text="" :height="calcSize()" :width="calcSize()" />
      </div>
      <div class="block__content">
        <h2 class="content__title">作品使用スキル内訳</h2>
        <p class="content__paragraph">現在投稿されているすべての投稿から使用されているスキルセットを抽出しています</p>
      </div>
    </div>
    <div class="container contain__end">
      <div class="block__content">
        <h2 class="content__title">日別作品投稿推移</h2>
        <p class="content__paragraph">過去5日間の投稿数を抽出しています</p>
      </div>
      <div class="block__content shadow__none">
        <Line-Chart :dataset="dateset" text="" :height="calcSize()" :width="calcSize()"/>
      </div>
    </div>
  </div>
</template>

<script>
import DoughnutChart from "~/components/DoughnutChart.vue"
import LineChart from "~/components/LineChart.vue"
export default {
  components: {
    DoughnutChart,
    LineChart
  },
  head () {
    return {
      title: "OnBoard",
      meta: [
        { hid: "top", name: "トップ", content: "日々新しい作品が作られていきます。" }
      ]
    }
  },
  data() {
    return {
      APIURL: "",
      dataset: {},
      dateset: {
        "4日前": 0,
        "3日前": 0,
        "2日前": 0,
        "昨日": 0,
        "今日": 0,
      }
    }
  },
  methods: {
    getSkills: async function() {
      this.APIURL = this.GetURL();
      await this.$axios
        .get(this.APIURL + "/skills")
        .then(response => {
          this.dataset = response.data.Skills
        })
        .catch(response => {
        });
    },
    getPostByDay: async function() {
      this.APIURL = this.GetURL();
      await this.$axios
        .get(this.APIURL + "/works_per_day")
        .then(response => {
          this.dateset["今日"] = response.data.Today
          this.dateset["昨日"] = response.data.Yesterday
          this.dateset["2日前"] = response.data.DBY
          this.dateset["3日前"] = response.data.TDA
          this.dateset["4日前"] = response.data.FDA
        })
        .catch(response => {
        });
    },
    calcSize: function() {
      // if (window.innerWidth > 400) {
      //   return 400
      // } else {
      //   return window.innerWidth * 0.9
      // }
      return 400
    }
  },
  async mounted() {
    await this.getSkills()
    await this.getPostByDay()
  },
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
.container {
  padding-top: 70px;
  margin: 0 auto;
  min-height: 88vh;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  align-items: center;
  text-align: center;
  width: 100%;
}
.container__title {
  color: lightgrey;
}
.container__button {
  margin-top: 20px;
  width: 200px;
  height: 50px;
  border-radius: 10px;
  border: solid 3px $bg-main;
  outline: 0;
  & a {
    font-size: 24px;
    font-weight: bold;
    text-decoration: none;
    color: $bg-main;
  }
  &:hover {
    & a {
      color: $bg-yellow;
    }
    animation-name: bg-change;
    animation-fill-mode: forwards;
    animation-duration: 1s;
    animation-timing-function: linear;
    animation-delay: 0s;
    animation-direction: normal;
  }
}
.block__content{
  border-radius: 5px;
  box-shadow: 0 1px 2px black;
  width: 40%;
  min-height: 100px;
  margin: 20px;
}
.shadow__none{
  box-shadow: 0 0 0 black;
}
.content__title{
  color: grey;
  margin-top: 20px;
}
.content__paragraph{
  color: $bg-main;
  font-size: 12px;
}
.contain__end{
  margin-bottom: 100px; 
}
@keyframes bg-change {
  0% {
    background-color: lightgrey;
  }
  100% {
    background-color: $bg-main;
  }
}
@media screen and (max-width: $PhoneSize) {
  .container__title {
    font-size: 20px;
  }
  .block__content {
    width: 90%;
  }
}
</style>
