<template>
  <div class="pagenation">
    <div class="pagenate__main">
      <p class="pagenation__item left" @click="goPrev()">&#8810;</p>
      <p class="pagenation__main__number">{{page}}</p>
      <p class="pagenation__item right" @click="goNext()">&#8811;</p>
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
.pagenate__main {
  display: flex;
  justify-content: center;
  margin: 50px;
}
.pagenation__item {
  text-align: center !important;
  font-size: 22px;
  font-weight: bold;
  color: $bg-main;
  width: 40px;
  cursor: pointer;
  transition: all 0.3s;
  border: solid 1px grey;
  &:hover {
    color: silver;
  }
  &:active {
    transition: all 0.2s;
    color: silver;
  }
}
.pagenation__main__number {
  text-align: center;
  font-size: 20px;
  width: 30px;
  font-weight: bold;
  color: $bg-main;
  background-color: white;
  border: solid 1px grey;
}
.left {
  border-radius: 10px 0 0 10px;
}
.right {
  border-radius: 0 10px 10px 0;
}
@media screen and (max-width: $PhoneSize) {
  .pagenation__item {
    width: 25px;
    font-size: 12px;
  }
  .pagenation__main__number {
    width: 15px;
    font-size: 12px;
  }
  .left {
  border-radius: 5px 0 0 5px;
  }
  .right {
    border-radius: 0 5px 5px 0;
  }
}
</style>