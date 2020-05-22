<template>
  <div class="login-container">
  <h2 class="page-title">ログイン画面</h2>
  <Validation :messages="errors"/>
    <div class="login-form-wrapper">
      <form @submit.prevent="submit">
        <div class="login__form">
          <p>
            <label for="email">メール</label>
            <input name="email" v-model="form.email" type="email" autocomplete="on" />
          </p>
          <p>
            <label for="password">パスワード</label>
            <input name="password" type="password" v-model="form.password" autocomplete="on" />
          </p>
        </div>
        <p>
          <input class="login__form__btn" type="submit" value="ログイン" />
        </p>
      </form>
    </div>
    <p class="navigate__signup"><nuxt-link to="/signup">ご登録はこちらから</nuxt-link></p>
  </div>
</template>
<script>
import Validation from "~/components/Validation.vue";
export default {
  components: {
    Validation
  },
  data() {
    return {
      errors: [],
      form: {
        email: "",
        password: ""
      }
    };
  },
  methods: {
    valCheck: function() {
      this.errors = []
      if (this.form.email === "") {
        const empty = "メールアドレスは必須です。"
        this.errors.push(empty)
      }
      if (this.form.password === "") {
        const empty = "パスワードは必須です。"
        this.errors.push(empty)
      }
    },
    async submit() {
      this.valCheck()
      if (this.errors.length !== 0) {
        return
      }
      try {
        await this.$auth.loginWith("local", {
          data: this.form
        });
      } catch (e) {
        //
      }
    }
  }
};
</script>
<style lang="scss" scopped>
@import "assets/scss/app";
.login-form-wrapper {
  background-color: white;
  border: solid 1px $bg-main;
  border-radius: 20px;
  height: 300px;
  width: 610px;
  margin: 10px auto;
  position: relative;
}
.login-container {
  background-color: $bg-color;
  padding-top: 150px;
  margin: 0 auto;
  min-height: 81vh;
  width: 100%;
}
.login__form__btn {
  outline: 0;
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  height: 50px;
  width: 100px;
  color: white;
  font-weight: bold;
  font-size: 14px;
  border-radius: 15px 15px 0 0;
  background-color: $bg-main;
  &:hover {
    color: $bg-yellow;
  }
  &:active {
    color: silver;
  }
}
.login__form {
  padding: 40px;
  font-weight: bold;
  & p {
    width: 340px;
    margin: 20px auto;
    & label {
      display: block;
      text-align: left;
    }
    & input {
      margin: 0;
      outline: 0;
      border: solid 1px $bg-main;
      border-radius: 20px;
      padding: 0 20px;
      font-size: 18px;
      height: 30px;
      width: 100%;
      font-weight: bold;
    }
  }
}
.navigate__signup {
  text-align: center;
  & a {
    font-size: 12px;
    text-decoration: none;
    color: blue;
  }
}
.page-title {
  text-align: center;
  font-size: 18px;
  color: grey;
}
@media screen and (max-width: $PhoneSize) {
  .login-form-wrapper {
    width: 98%;
    border: none;
  }
  .login__form {
    & p {
      width: 100%;
    }
  }
  .login-container {
    overflow: hidden;
  }
}
</style>
