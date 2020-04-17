<template>
  <div>
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
.pagenate__main {
  display: flex;
  justify-content: center;
}
.pagenation__item {
  margin: 0 20px;
  text-align: center;
  height: 30px;
  line-height: 30px;
  font-size: 22px;
  font-weight: bold;
  color: white;
  cursor: pointer;
  transition: all 0.3s;
  &:hover {
    color: #192b3d;
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
  color: #192b3d;
  background-color: white;
  box-shadow: 0 0 5px black;
}
</style>