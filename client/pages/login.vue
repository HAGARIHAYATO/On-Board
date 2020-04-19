<template>
  <div class="login-container">
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
  </div>
</template>
<script>
export default {
  data() {
    return {
      form: {
        email: "",
        password: ""
      }
    };
  },
  methods: {
    async submit() {
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
.login-form-wrapper {
  background-color: white;
  border: solid 3px #192b3d;
  border-radius: 20px;
  height: 300px;
  width: 50%;
  position: relative;
}
.login-container {
  background-color: lighten(rgb(221, 209, 209), 5%);
  padding-top: 70px;
  margin: 0 auto;
  min-height: 81vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
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
  background-color: #192b3d;
  &:hover {
    color: #fdeaa0;
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
      border: solid 3px #192b3d;
      border-radius: 20px;
      padding: 0 20px;
      font-size: 18px;
      height: 30px;
      width: 300px;
      font-weight: bold;
    }
  }
}
</style>
