<template>
  <div class="validation__container" v-if="isExistMessage()">
    <p class="validation__count">
      {{messages.length}}件のエラーがあります。
    </p>
    <p class="validation__message" v-for="(m,i) in messages" :key="i">
      ・{{m}}
    </p>
  </div>
</template>
<script>
export default {
  methods: {
    isExistMessage: function() {
      if (this.messages.length === 0) {
        return false
      }
      return true
    }
  },
  props: {
    messages: {
      type: Array
    }
  }
}
</script>
<style lang="scss" scoped>
.validation__container{
  color: red;
  width: 600px;
  margin: 0 auto;
  font-size: 12px;
}
.validation__message{
  margin: 0 0 0 10px;
}
</style>