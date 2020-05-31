<template>
  <div class="container" v-if="reload">
    <Loading v-if="isLoading" />
    <div
      class="container__side"
      :class="slider()"
    >
      <search-form
        @search="search($event)"
        @cancel="cancel()"
        :isSearch="isSearch"
        searchType="works"
      />
    </div>
    <div class="container__main">
      <works :works="returnWorks" :isSearch="isSearch" />
      <pagenate :page="page" @goPrev="goPrev()" @goNext="goNext()" v-if="!isSearch" />
    </div>
    <p class="searchBtn" @click="drawSlider()"><img src="/search.svg" alt="search"></p>
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
        await setTimeout(() => this.showBubble(), 500);
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
            work.Name.toLowerCase().indexOf(e.input.toLowerCase()) != -1 ||
            work.UserName.toLowerCase().indexOf(e.input.toLowerCase()) != -1
        );
      } else if (this.works[0] && !e.check) {
        this.worksList = this.works.filter(
          work => work.Name.toLowerCase().indexOf(e.input.toLowerCase()) != -1
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
  margin-top: 70px;
  padding: 10px 20px;
  background-color: lightgrey;
  border-radius: 5px;
  box-shadow: 0 0 1px black;
}
.slider__off{
  & p {
    transition: all .5s;
  }
  transition: all .5s;
  z-index: 1;
  position: fixed;
  top: 0;
  right: -400px;
}
.slider__on{
  & p {
    transition: all .5s;
    transform: rotate(-540deg);
    color: $bg-main;
  }
  transition: all .5s;
  z-index: 1;
  position: fixed;
  top: 0;
  right: 20px;
}
.container__main {
  padding-top: 100px;
  min-height: 100%;
  width: 100%;
  background-color: $bg-color;
}
.searchBtn {
  cursor: pointer;
  width: 60px;
  height: 60px;
  border-radius: 50%;
  position: fixed;
  bottom: 200px;
  right: 100px;
  background-color: lightgray;
  text-align: center;
  line-height: 60px;
  font-size: 22px;
  font-weight: bold;
  box-shadow: 0 1px 3px black;
  z-index: 3;
  & img {
    height: 16px;
    margin: 22px 0;
  }
}
@media screen and (max-width: $PhoneSize) {
  .slider__off{
    & p {
      transition: all .5s;
      left: -50px;
    }
    transition: all .5s;
    z-index: 1;
    position: fixed;
    top: 0;
    right: -90%;
  }
  .slider__on{
    & p {
      transition: all .5s;
      transform: rotate(-540deg);
      color: $bg-main;
    }
    transition: all .5s;
    z-index: 1;
    position: fixed;
    top: 0;
    right: 0;
  }
  .container__side {
    width: 90%;
  }
  .searchBtn {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    position: fixed;
    bottom: 20px;
    right: 80px;
    background-color: lightgray;
    text-align: center;
    line-height: 40px;
    font-size: 12px;
    font-weight: bold;
    box-shadow: 0 1px 3px black;
    z-index: 3;
    & img {
      height: 10px;
      margin: 15px 0;
    }
  }
}
</style>
