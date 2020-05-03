<template>
  <div class="pagenation">
    <p class="pagenation__main__number">{{page}}</p>
    <div class="pagenate__main">
      <p class="pagenation__item" @click="goPrev()">prev</p>
      <p class="pagenation__item" @click="goNext()">next</p>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {};
  },
  methods: {
    goPrev: function() {
      this.$emit("goPrev");
    },
    goNext: function() {
      this.$emit("goNext");
    }
  },
  props: {
    page: Number
  }
};
</script>
<style lang="scss" scopoed>
@import "assets/scss/app";
.pagenation{
}
.pagenate__main {
  display: flex;
  justify-content: center;
}
.pagenation__item {
  margin: 0 20px !important;
  text-align: center !important;
  height: 30px;
  line-height: 30px;
  font-size: 22px;
  font-weight: bold;
  color: $bg-main;
  cursor: pointer;
  transition: all 0.3s;
  border-radius: 20px;
  padding: 4px 2%;
  &:hover {
    color: silver;
  }
  &:active {
    transition: all 0.2s;
    color: silver;
  }
}
.pagenation__main__number {
  margin: 0 auto;
  text-align: center;
  width: 50px;
  height: 50px;
  line-height: 50px;
  font-size: 30px;
  font-weight: bold;
  border-radius: 50%;
  color: $bg-main;
  background-color: white;
  box-shadow: 0 0 2px black;
}
</style>