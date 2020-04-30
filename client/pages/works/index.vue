<template>
  <div class="container" v-if="reload">
    <Loading v-if="isLoading" />
    <div class="container__side">
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
      APIURL: ""
    };
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
.container {
  display: flex;
  width: 100%;
  min-height: 81vh;
}
.container__side {
  padding-top: 70px;
  min-height: 100%;
  width: 400px;
  background-color: #fdeaa0;
}
.container__main {
  padding-top: 100px;
  min-height: 100%;
  width: 100%;
  background-color: lighten(rgb(221, 209, 209), 5%);
}
</style>
