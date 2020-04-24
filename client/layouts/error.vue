<template>
  <div class="background">
    <div class="container">
      <h1>{{ error.statusCode }}</h1>
      <p v-if="error.statusCode === 404">
      ページが見つかりません
      </p>
      <p v-else>
      エラーが発生しました
      </p>
      <p>{{time}}秒後にリダイレクトされます。</p>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      time: 5 
    }
  },
  props: {
    error: {
      type: Object,
      default: null
    }
  },
  mounted() {
    setInterval(() => { this.time-- }, 1000)
    setTimeout(() => this.$router.push("/"), 5000);
  }
}
</script>
<style lang="scss" scoped>
.background{
  z-index: 5;
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  background-color: white;
  width: 100%;
}
.container{
  width: 500px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  text-align: center;
}
</style>