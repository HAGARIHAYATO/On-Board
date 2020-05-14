<template>
  <div class="signup-container">
    <Loading v-if="isLoading" />
    <Validation :messages="errors"/>
    <div class="signup-form-wrapper">
      <form @submit.prevent="submit">
        <div class="signup__form">
          <p>
            <label for="name">名前</label>
            <input type="text" name="name" v-model="name" autocomplete="on" />
          </p>
          <p>
            <label for="email">メール</label>
            <input v-model="email" name="email" type="email" autocomplete="on" />
          </p>
          <p>
            <label for="password">パスワード</label>
            <input
              name="password"
              type="password"
              autocomplete="on"
              v-model="password"
            />
          </p>
          <p>
            <label for="password">パスワード確認</label>
            <input
              name="password"
              type="password"
              autocomplete="on"
              v-model="passwordConfirm"
            />
          </p>
        </div>
        <p>
          <input class="signup__form__btn" type="submit" value="無料登録" />
        </p>
      </form>
    </div>
  </div>
</template>
<script>
import Loading from "~/components/Loading.vue";
import Validation from "~/components/Validation.vue";
export default {
  components: {
    Loading,
    Validation
  },
  data() {
    return {
      errors: [],
      isLoading: false,
      name: "",
      email: "",
      password: "",
      passwordConfirm: "",
      APIURL: ""
    };
  },
  mounted() {
    this.APIURL = this.GetURL();
  },
  methods: {
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    valCheck: function() {
      this.errors = []
      if (this.name === "") {
        const empty = "作品名は必須です。"
        this.errors.push(empty)
      }
      if (this.email === "") {
        const empty = "メールアドレスは必須です。"
        this.errors.push(empty)
      }
      if (this.password === "" || this.password !== this.passwordConfirm) {
        const check = "パスワードを確認してください。"
        this.errors.push(check)
      }
      if (this.password.length < 5 || this.password.length > 20) {
        const check = "パスワードは6文字以上20字以下です。"
        this.errors.push(check)
      }
    },
    async submit() {
      this.valCheck()
      if (this.errors.length !== 0) {
        return
      }
      try {
        const data = new FormData();
        data.append("name", this.name);
        data.append("email", this.email);
        data.append("password", this.password);
        const headers = { "content-type": "application/x-www-form-urlencoded" };
        this.$nextTick(async () => {
          await this.showBubble();
          await this.$axios
            .post(this.APIURL + "/signup", data, {
              headers
            })
            .then(res => {
              if (res.data.ID) {
                if (res.data.Token) {
                  this.$auth.setUserToken(res.data.Token)
                }
                this.$router.push("/users/" + res.data.ID);
              }
            });
          await setTimeout(() => this.showBubble(), 1000);
        });
      } catch (error) {
        // handling show message
      }
    }
  }
};
</script>
<style lang="scss" scoped>
@import "assets/scss/app";
.signup-form-wrapper {
  background-color: white;
  border: solid 3px $bg-main;
  border-radius: 20px;
  height: 400px;
  width: 610px;
  margin: 10px auto;
  position: relative;
}
.signup-container {
  background-color: $bg-color;
  padding-top: 120px;
  margin: 0 auto;
  min-height: 81vh;
  width: 100%;
}
.signup__form__btn {
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
.signup__form {
  padding: 20px 40px 40px 40px;
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
      border: solid 3px $bg-main;
      border-radius: 20px;
      padding: 0 20px;
      font-size: 18px;
      height: 30px;
      width: 100%;
      font-weight: bold;
    }
  }
}
@media screen and (max-width: $PhoneSize) {
  input[type='text'],
  input[type='email'],
  input[type='password'] {
    width: 100% !important;
  }
  .signup-form-wrapper {
    width: 98%;
    border: none;
  }
  .signup__form {
    padding: 40px;
    & p {
      width: 100%;
    }
  }
  .signup-container {
    overflow: hidden;
  }
}
</style>
