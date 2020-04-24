<template>
  <div class="field__container">
    <h2 class="field__title">GitHub Account</h2>
    <p class="gh__user"><img :src="ghUser.avatar_url"><a :href="ghUser.html_url">{{ghUser.login}}</a></p>
    <h2 class="field__title">GitHub Projects</h2>
    <div class="field">
      <div v-for="(item, index) in hubs" :key="index" class="field__item">
        <GitHubItem :platform="0" :title="item.name" :url="item.html_url"/>
      </div>
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
    },
    ghUser: {
      type: Object
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
  background-color: white;
  border-radius: 10px;
  border: solid .5px lightgrey;
  margin: 30px;
}
.field{
  display: flex;
  justify-content: flex-start;
  flex-wrap: wrap;
}
.field__item {
  margin :5px;
}
.field__title{
  margin: 0 10px 10px 10px;
  color: #192b3d;
  font-size: 18px;
}
.gh__user{
  margin: 0 0 0 50px;
  height: 50px;
  & img {
    height: 30px;
    width: 30px;
  }
  & a {
    margin: 0 10px;
    font-size: 30px;
    font-weight: bold;
    text-decoration: none !important;
  }
}
</style>