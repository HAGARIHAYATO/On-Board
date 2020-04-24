<template>
  <div class="field__container">
    <div v-for="(item, index) in hubs" :key="index" class="field__item">
      <GitHubItem :platform="0" :title="item.name" :url="item.html_url"/>
    </div>
  </div>
</template>
<script>
import GitHubItem from "~/components/GitHubItem.vue";
export default {
  components: {
    GitHubItem
  },
  props: {
    hubs: {
      type: Array
    }
  },
  data() {
    return {
    }
  }
}
</script>
<style lang="scss" scoped>
.field__container{
  padding: 2%;
  display: flex;
  justify-content: flex-start;
  flex-wrap: wrap;
  background-color: white;
  border-radius: 10px;
  border: solid .5px lightgrey;
  margin: 30px;
}
.field__item {
  margin :5px;
}
</style>