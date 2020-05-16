<template>
  <div class="sheet padding-10mm">
    <div>
      <p class="title">{{work.Name}}</p>
      <p><a :href="work.URL">{{ work.URL }}</a></p>
    </div>
    <div>
      <p class="title">画像</p>
      <p class="url"><a :href="work.ImageURL">{{work.ImageURL}}</a></p>
      <p class="url" v-for="(item,i) in items" :key="i"><a :href="item.ImageURL">{{ item.ImageURL }}</a></p>
    </div>
    <p class="title">スキル</p>
    <div class="row__box">
      <p class="column" v-for="(skill,index) in work.Skills" :key="index">{{skill.Name}}</p>
    </div>
    <div>
      <p class="title">by {{work.UserName}}</p>
      <p class="url"><a :href="work.UserImageURL">{{ work.UserImageURL }}</a></p>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    work: {
      type: Object,
      require: true
    },
    items: {
      type: Array,
      require: true
    }
  },
  data() {
    return {
    }
  },
  mounted() {
    //画像入れたい
  },
  methods: {
  },
}
</script>
<style lang="scss" scoped>
.sheet {
  margin: 0 auto;
  height: 90vh;
  background-color: white;
  overflow: hidden;
  position: relative;
  box-sizing: border-box;
  page-break-after: always;
  width: 142mm;
  height: 200mm;
}
.title{
  font-weight: bold;
}
.url{
  font-size: 6px !important;
  width: 300px;
  word-break: break-all;
}
.row__box{
  display: flex;
  flex-wrap: wrap;
}
.column{
  display: inline-block;
  font-size: 6px !important;
  word-break: break-all;
  margin: 5px;
}
</style>