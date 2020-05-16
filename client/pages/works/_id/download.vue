<template>
  <div class="pdfModal">
    <Loading v-if="isLoading" />
    <p class="arrowBack" @click="OpenPDF">←</p>
    <p class="pdfDownload" @click="download"><img src="/download.svg"></p>
    <div id="print" class="A4">
      <div class="pageout"><PrintElement ref="print" :work="work" :items="items" /></div>
      <div class="preview"><PrintElement :work="work" :items="items" /></div>
    </div>
  </div>
</template>
<script>
import PrintElement from "~/components/Print.vue";
import Loading from "~/components/Loading.vue";
export default {
  components: {
    PrintElement,
    Loading
  },
  data() {
    return {
      items: [],
      work: {},
      isLoading: false,
      APIURL: ""
    }
  },
  methods: {
    download() {
      this.createPdfFromHtml(this.$refs.print.$el);
    },
    OpenPDF: function() {
      const route = this.$route.path.split("/")
      this.$router.push({ path: "/" + route[1] + "/" + route[2] })
    },
    showBubble: function() {
      this.isLoading = !this.isLoading;
    }
  },
  mounted() {
    const route = this.$route.path.split("/")
    this.$nextTick(async () => {
      this.APIURL = this.GetURL();
      await this.$axios
        .get(this.APIURL + "/works/" + route[2])
        .then(response => {
          this.work = response.data;
          this.items = response.data.WorkItems;
        })
        .catch(response => {
        });
    });
  }
}
</script>
<style lang="scss" scoped>
.pdfModal{
  background-color: #000000;
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 2;
  padding: 1%;
}
.arrowBack{
  box-shadow: 0 0 2px grey;
  border-radius: 5px;
  line-height: 45px;
  height: 45px;
  width: 45px;
  text-align: center;
  color: grey;
  font-weight: bold;
  font-size: 30px;
  position: fixed;
  background-color: white;
  top: 10px;
  left: 10px;
  cursor: pointer;
  &:hover{
    text-shadow: 0 0 3px grey;
  }
}
.pdfDownload{
  box-shadow: 0 0 2px grey;
  border-radius: 5px;
  line-height: 45px;
  height: 45px;
  width: 45px;
  text-align: center;
  color: white;
  text-shadow: 0 0 3px grey;
  font-weight: bold;
  font-size: 30px;
  position: fixed;
  background-color: white;
  top: 70px;
  left: 10px;
  cursor: pointer;
  & img {
    height: 30px;
    margin: 7.5px 0;
  }
  &:active{
    background-color: black;
  }
}
.pageout {
  position: fixed;
  top: 200vh;
}
.preview{
  height: 98%;
}
</style>