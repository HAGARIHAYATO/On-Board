<template>
  <a :href="url">
    <div class="item__container">
      <div class="pedestal">
        <img :src="isHubOrLab()">
      </div>
      <p class="item__title">{{title}}</p>
    </div>
  </a>
</template>
<script>
export default {
  props: {
    platform: {
      type: Number
    },
    title: {
      type: String
    },
    url: {
      type: String
    }
  },
  data() {
    return {
    }
  },
  methods: {
    isHubOrLab: function(){
      let image = ""
      switch (this.platform) {
        case 0:
        image = "/github.png";
        break;
        case 1:
        image = "/gitlab.png";
        break;
      }
      return image;
    } 
  },
}
</script>
<style lang="scss" scoped>
* {
  box-sizing: border-box;
  text-decoration: none !important;
}
.item__container{
  border-radius: 25px;
  height: 40px;
  min-width: 150px !important;
  background-color: lighten(lightgrey, 10%);
  box-shadow: 0 0 3px black;
  line-height: 40px;
  display: inline-flex;
  &:hover{
    box-shadow: 0 0 5px black;
  }
  &:active{
    border: solid 2px lightgreen;
    box-shadow: 0 0 0 black;
    & div {
      margin: 1px;
    }
    & p {
      margin: -2px 20px 0 10px;
    }
  }
}
.pedestal{
  border-radius: 50%;
  width: 34px;
  height: 34px;
  background-color: white;
  padding: 2px;
  box-shadow: 0 0 2px black;
  margin: 3px;
  & img {
    height: 30px;
    width: 30px;
  }
}
.item__title{
  margin: 0 20px 0 10px;
  height: 40px;
  line-height: 40px;
  color: black;
}
</style>