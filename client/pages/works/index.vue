<template>
  <div class="container" v-if="reload">
    <Loading v-if="isLoading" />
    <div
      class="container__side"
      :class="slider()"
    >
      <p
        class="drawBtn"
        @click="drawSlider()"
      >＜</p>
      <search-form
        @search="search($event)"
        @cancel="cancel()"
        :isSearch="isSearch"
        searchType="works"
      />
      <pagenate :page="page" @goPrev="goPrev()" @goNext="goNext()" v-if="!isSearch" />
    </div>
    <div class="container__main">
      <works :works="returnWorks" :isSearch="isSearch" />
    </div>
  </div>
</template>
<script>
import Works from "~/components/Works.vue";
import SearchForm from "~/components/SearchForm.vue";
import Pagenate from "~/components/Pagenate.vue";
import Loading from "~/components/Loading.vue";
export default {
  components: {
    Works,
    SearchForm,
    Pagenate,
    Loading
  },
  data() {
    return {
      page: 1,
      works: [],
      worksList: [],
      reload: true,
      maxCardCount: 8,
      isSearch: false,
      isLoading: false,
      APIURL: "",
      slide: false
    };
  },
  head () {
    return {
      title: "OnBoard / 作品一覧",
      meta: [
        { hid: "works", name: "作品一覧", content: "作品の一覧画面です。 面白い作品をチェックしましょう。" }
      ]
    }
  },
  mounted() {
    this.APIURL = this.GetURL();
    try {
      this.$nextTick(async () => {
        await this.showBubble();
        await this.$axios
          .get(this.APIURL + "/works")
          .then(response => {
            this.works = response.data.Works;
          })
          .catch(response => console.error(response)),
          await this.init();
        await setTimeout(() => this.showBubble(), 1000);
      });
    } catch (error) {}
  },
  computed: {
    returnWorks: function() {
      return this.worksList;
    }
  },
  methods: {
    slider: function() {
      const off = "slider__off"
      const on = "slider__on"
      if (this.slide) {
        return on
      } else {
        return off
      }
    },
    drawSlider: function() {
      this.slide = !this.slide
      console.log(this.slide)
    },
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    init: function() {
      this.worksList = [];
      if (!this.works) return;
      if (this.works.length > 0) {
        for (let i = 0; i < this.maxCardCount; i++) {
          if (this.works[i]) {
            this.worksList.push(this.works[i]);
          } else {
            continue;
          }
        }
      }
      this.page = 1;
    },
    search: function(e) {
      if (this.works[0] && e.check) {
        this.worksList = this.works.filter(
          work =>
            work.Name.indexOf(e.input) != -1 ||
            work.UserName.indexOf(e.input) != -1
        );
      } else if (this.works[0] && !e.check) {
        this.worksList = this.works.filter(
          work => work.Name.indexOf(e.input) != -1
        );
      }
      this.isSearch = true;
    },
    cancel: function() {
      this.isSearch = false;
      this.reload = false;
      this.init();
      this.reload = true;
    },
    goPrev: function() {
      if (this.page > 1) {
        this.worksList = [];
        for (
          let i = (this.page - 1) * this.maxCardCount;
          i > (this.page - 2) * this.maxCardCount;
          i--
        ) {
          if (this.works[(this.page - 1) * this.maxCardCount - i]) {
            this.worksList.push(
              this.works[(this.page - 1) * this.maxCardCount - i]
            );
          }
        }
        this.page--;
      }
    },
    goNext: function() {
      if (this.works.length - this.page * this.maxCardCount > 0) {
        this.worksList = [];
        for (let i = 0; i < this.maxCardCount; i++) {
          if (this.works[this.page * this.maxCardCount + i]) {
            this.worksList.push(this.works[this.page * this.maxCardCount + i]);
          }
        }
        this.page++;
      }
    }
  }
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
.container {
  display: flex;
  width: 100%;
  min-height: 81vh;
}
.container__side {
  padding-top: 70px;
  min-height: 100%;
  width: 400px;
  background-color: $bg-yellow;
  border-left: solid 2px white;
  box-shadow: 0 0 5px $bg-main;
}
.slider__off{
  & p {
    transition: all .5s;
  }
  transition: all .5s;
  z-index: 1;
  position: fixed;
  top: 0;
  right: -350px;
}
.slider__on{
  & p {
    transition: all .5s;
    transform: rotate(-540deg);
    color: $bg-main;
  }
  border-left: solid 2px $bg-main;
  transition: all .5s;
  z-index: 1;
  position: fixed;
  top: 0;
  right: 0;
}
.container__main {
  padding-top: 100px;
  min-height: 100%;
  width: 100%;
  background-color: $bg-color;
}
.drawBtn{
  position: absolute;
  top: 100px;
  left: -25px;
  border-radius: 50%;
  background-color: $bg-yellow;
  color: white;
  height: 50px;
  width: 50px;
  line-height: 46px;
  text-align: center;
  font-size: 20px;
  font-weight: bold;
  border-top: solid 2px $bg-main;
  border-bottom: solid 2px white;
  border-left: solid 2px white;
  border-right: solid 2px $bg-main;
  box-shadow: 0 0 5px $bg-main;
}
</style>
