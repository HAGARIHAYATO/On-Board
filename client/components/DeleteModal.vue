<template>
  <div class="overlay">
    <div class="modal__container">
      <p class="strMessage">{{confirmStr}}</p>
      <button class="btnMessage" @click="submit">{{btnStr}}</button>
      <button class="backMessage" @click="back">戻る</button>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    confirmStr: {
      type: String
    },
    btnStr: {
      type: String
    }
  },
  methods: {
    async back() {
      this.$emit("back");
    },
    async submit() {
      this.$emit("delete");
    }
  }
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
.overlay {
  position: fixed;
  z-index: 3;
  top: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0.9);
}
.modal__container {
  height: 200px;
  width: 500px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.btnMessage {
  margin: 20px auto 2px auto;
  height: 40px;
  width: 180px;
  display: block;
  outline: none;
  background-color: red;
  color: white;
  border-radius: 5px;
  font-weight: bold;
  &:hover {
    box-shadow: 0 0 5px grey;
  }
}
.backMessage {
  margin: 0 auto 20px auto;
  height: 40px;
  width: 180px;
  display: block;
  outline: none;
  background-color: $bg-main;
  color: white;
  border-radius: 5px;
  font-weight: bold;
  &:hover {
    box-shadow: 0 0 5px grey;
  }
}
.strMessage {
  margin: 10px auto 10px auto;
  text-align: center;
  color: $bg-main;
  word-wrap: break;
}
</style>