<template>
  <div class="container">
    <div>
      <h1 class="container__title">未知の作品と出会おう</h1>
      <button class="container__button">
        <nuxt-link to="/works">作品を探す</nuxt-link>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  components: {}
};
</script>
<style lang="scss" scoped>
.container {
  padding-top: 70px;
  margin: 0 auto;
  min-height: 81vh;
  display: flex;
  justify-content: center;
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
  border: solid 3px #192b3d;
  outline: 0;
  & a {
    font-size: 24px;
    font-weight: bold;
    text-decoration: none;
    color: #192b3d;
  }
  &:hover {
    & a {
      color: #fdeaa0;
    }
    animation-name: bg-change;
    animation-fill-mode: forwards;
    animation-duration: 1s;
    animation-timing-function: linear;
    animation-delay: 0s;
    animation-direction: normal;
  }
}
@keyframes bg-change {
  0% {
    background-color: lightgrey;
  }
  100% {
    background-color: #192b3d;
  }
}
</style>
