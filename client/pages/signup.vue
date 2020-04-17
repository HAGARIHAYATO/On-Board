<template>
  <div class="signup-container">
    <Loading v-if="isLoading" />
    <div class="signup-form-wrapper">
      <form @submit.prevent="submit">
        <div class="signup__form">
          <p>
            <label for="name">名前</label>
            <input type="text" name="name" v-model="name" autocomplete="on" required />
          </p>
          <p>
            <label for="email">メール</label>
            <input v-model="email" name="email" type="email" autocomplete="on" required />
          </p>
          <p>
            <label for="password">パスワード</label>
            <input
              name="password"
              type="password"
              autocomplete="on"
              minlength="6"
              required
              v-model="password"
            />
          </p>
          <p>
            <label for="password">パスワード確認</label>
            <input
              name="password"
              type="password"
              autocomplete="on"
              minlength="6"
              v-model="passwordConfirm"
              required
            />
          </p>
        </div>
        <p>
          <input class="signup__form__btn" type="submit" value="無料登録" :disabled="returnCheck" />
        </p>
      </form>
    </div>
  </div>
</template>
<script>
import Loading from "~/components/Loading.vue";
export default {
  components: {
    Loading
  },
  data() {
    return {
      isLoading: false,
      name: "",
      email: "",
      password: "",
      passwordConfirm: "",
      APIURL: "http://localhost:8080/api/v1"
    };
  },
  computed: {
    returnCheck: function() {
      if (this.password != this.passwordConfirm) return true;
      return false;
    }
  },
  methods: {
    showBubble: function() {
      this.isLoading = !this.isLoading;
    },
    async submit() {
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
.signup-form-wrapper {
  background-color: white;
  border: solid 3px #192b3d;
  border-radius: 20px;
  height: 400px;
  width: 50%;
  position: relative;
}
.signup-container {
  background-color: lighten(rgb(221, 209, 209), 5%);
  padding-top: 40px;
  margin: 0 auto;
  min-height: 81vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
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
  background-color: #192b3d;
  &:hover {
    color: #fdeaa0;
  }
  &:active {
    color: silver;
  }
}
.signup__form {
  padding: 15px;
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
