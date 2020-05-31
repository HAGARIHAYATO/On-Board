<template>
  <div>
    <div class="userball__box">
      <p v-for="(user, i) in users" :key="i" @click="moveToUser(user.UserID)" class="userball">
        <img class="userball__image" :src="returnURL(user.ImageURL)" :alt="user.Name">
        <span class="userball__username">{{user.Name}}<span class="speech__buble"></span></span>
      </p>
    </div>
    <h2>{{users.length}}名が評価しています。</h2>
  </div>
</template>
<script>
export default {
  props: {
    users: {
      type: Array,
      require: true
    } 
  },
  methods: {
    returnURL: function(url) {
      return url ? url : "/NO_IMAGE.jpeg";
    },
    moveToUser: function(id) {
      this.$router.push("/users/" + id)
    }
  },
}
</script>
<style lang="scss" scoped>
h2 {
  text-align: center;
  font-size: 12px;
  color: grey;
}
.userball__box {
  width: 600px;
  margin: 20px auto;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
}
.userball__image {
  width: 50px;
  height: 50px;
  margin: 3px;
  border-radius: 50%;
  transition: all .3s;
  cursor: pointer;
  &:hover {
    border: solid 1px black;
  }
}
.userball {
  position: relative;
  &:hover {
    & .userball__username {
      transform: scale(1);
    }
  } 
}
.userball__username {
  transition: all .3s;
  transform: scale(0);
  position: absolute;
  top: -32px;
  left: 0;
  background-color: lightgray;
  border-radius: 20px;
  border: solid 1px grey;
  padding: 1px 10px;
  min-width: 100px;
  text-align: center;
  & .speech__buble {
    position: absolute;
    width: 10px;
    height: 10px;
    bottom: -6px;
    left: 22px;
    background-color: lightgray;
    border-left: solid 1px grey;
    border-bottom: solid 1px grey;
    transform: rotate(-45deg);
  }
}
</style>