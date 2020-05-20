<template>
  <div class="loading-wrapper">
    <div class="loading-container">
      <div class="loading"></div>
      <div class="loading__bar">
        <div class="bar pattern1"></div>
        <div class="bar pattern2"></div>
        <div class="bar pattern3"></div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  methods: {}
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
* {
  box-sizing: border-box;
}
.loading-wrapper {
  overflow: hidden;
  z-index: 4;
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: $bg-main;
}
.loading-container {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.loading {
  background-image: url("/logo.jpeg");
  background-position: center;
  background-size: contain;
  width: 200px;
  height: 200px;
}
.loading__bar {
  width: 100px;
  height: 60px;
  margin: 0 auto;
  display: flex;
  justify-content: space-around;
}
.bar {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  margin: auto 0;
  background-color: $bg-yellow;
  animation-name: loading;
  animation-duration: 1s;
  animation-iteration-count: infinite;
  animation-timing-function: linear;
  animation-direction: alternate;
}
.pattern1 {
  animation-delay: 0s;
}
.pattern2 {
  animation-delay: 0.33s;
}
.pattern3 {
  animation-delay: 0.67s;
}
@keyframes loading {
  0% {
    transform: scale(0);
  }
  25% {
    transform: scale(0.4);
  }
  50% {
    transform: scale(1);
  }
  75% {
    transform: scale(0.4);
  }
  100% {
    transform: scale(0);
  }
}
</style>
