<template>
  <div class="search-container">
    <p class="validation">{{error}}</p>
    <div class="search-form-wrapper">
      <input v-model="inputValue" class="search-form__box" type="search" placeholder="キーワード検索" />
      <input
        type="submit"
        @click="search"
        class="search-form__button"
        value="検索"
        minlength="1"
      />
    </div>
    <p v-if="searchType=='works'">
      <input class="checkbtn" type="checkbox" checked v-model="isChecked" />
      <span>ユーザー名も含めて検索</span>
    </p>
    <button v-if="isSearch" @click="cancel()" class="search-form__button__reset">
      <p>検索を解除</p>
    </button>
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
      this.error = ""
      if (this.inputValue === "") {
        this.error = "キーワードを入力してください"
        return
      }
      if (this.inputValue.length > 20) {
        this.error = "入力は最大20字です。"
        return
      }
      this.$emit("search", { input: this.inputValue, check: this.isChecked });
    },
    cancel: function() {
      this.error = ""
      this.inputValue = "";
      this.$emit("cancel");
    }
  },
  data() {
    return {
      inputValue: "",
      isChecked: true,
      error: ""
    };
  }
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
input, button {
  &:hover {
    box-shadow: none; 
  }
}
input[type="checkbox"] {
  transform: scale(1.5);
  box-shadow: none;
}
.search-container {
  // margin: 100px auto;
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
    color: $bg-main;
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
  border-radius: 10px 0 0 10px;
  border: none;
  background: #eee;
  font-weight: bold;
}
.search-form__button {
  border-radius: 0 10px 10px 0;
  outline: 0;
  height: 50px;
  width: 70px;
  position: absolute;
  left: 200px;
  top: 0;
  background: $bg-main;
  color: #fff;
  border: none;
  font-weight: bold;
  transition: all 0.2s;
  &:hover {
    color: $bg-yellow;
  }
  &:active {
    transition: all 0.2s;
    color: $bg-main;
  }
}
.search-form__button__reset {
  outline: 0;
  display: block;
  padding: 0 20px;
  height: 30px;
  border: none;
  margin: 0 auto !important;
  background-color: red;
  color: white;
  font-weight: bold;
  border-radius: 25px;
}
.checkbtn {
  background-color: $bg-main;
  color: white;
  margin-left: 8px;
}
.validation{
  color: red;
  margin: 0 0 0 22px;
  font-size: 10px;
}
</style>
