<template>
  <div class="search-container">
    <div class="search-form-wrapper">
      <input v-model="inputValue" class="search-form__box" type="search" placeholder="キーワード検索" />
      <input
        type="submit"
        @click="search"
        class="search-form__button"
        value="検索"
        minlength="1"
        :disabled="inputValue==''"
      />
    </div>
    <div class="search-form-wrapper" v-if="searchType=='works'">
      <p>
        <input class="checkbtn" type="checkbox" checked v-model="isChecked" />
        <span>ユーザー名も含めて検索</span>
      </p>
    </div>
    <div class="search-form-wrapper" v-if="isSearch">
      <button @click="cancel()" class="search-form__button__reset">
        <p>検索状況を解除</p>
      </button>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    isSearch: {
      type: Boolean
    },
    searchType: {
      type: String
    }
  },
  computed: {},
  methods: {
    search: function() {
      this.$emit("search", { input: this.inputValue, check: this.isChecked });
    },
    cancel: function() {
      this.inputValue = "";
      this.$emit("cancel");
    }
  },
  data() {
    return {
      inputValue: "",
      isChecked: true
    };
  }
};
</script>
<style lang="scss" scoped>
input[type="checkbox"] {
  transform: scale(1.5);
}
.search-container {
  margin: 100px auto;
}
.search-form-wrapper {
  z-index: 0;
  position: relative;
  width: 270px;
  height: 50px;
  margin: 10px auto;
  & span,
  .checkbtn {
    text-align: center;
    margin-left: 15px;
    font-weight: bold;
    color: #192b3d;
    line-height: 50px;
  }
}
.search-form__box {
  outline: 0;
  height: 50px;
  padding: 0 10px;
  position: absolute;
  width: 200px;
  left: 0;
  top: 0;
  background: #eee;
  font-weight: bold;
}
.search-form__button {
  outline: 0;
  height: 50px;
  width: 70px;
  position: absolute;
  left: 200px;
  top: 0;
  background: #192b3d;
  color: #fff;
  border: none;
  font-weight: bold;
  transition: all 0.2s;
  &:hover {
    color: #fdeaa0;
  }
  &:active {
    transition: all 0.2s;
    color: #192b3d;
  }
}
.search-form__button__reset {
  outline: 0;
  display: block;
  width: 150px;
  height: 50px;
  margin: 100px auto 0 auto !important;
  background-color: red;
  color: white;
  font-weight: bold;
  border-radius: 25px;
}
.checkbtn {
  background-color: #192b3d;
  color: white;
}
</style>
