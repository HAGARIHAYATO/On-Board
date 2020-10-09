<template>
<div class="contact__form__wrapper">
  <h2 class="page-title">お問い合わせ</h2>
  <Validation :messages="errors" />
  <p :class="success ? 'success-card' : 'invisible'">送信完了<span @click="() => this.success = false">×</span></p>
  <form @submit.prevent="submit" class="contact__form">
    <p>
      <label for="name">名前</label>
      <input name="name" v-model="form.name" type="text" autocomplete="name" />
    </p>
    <p>
      <label for="email">ご連絡用メールアドレス※返信をご希望の方</label>
      <input name="email" v-model="form.email" type="email" autocomplete="email" />
    </p>
    <p>
      <label for="text">お問い合わせ内容</label>
      <textarea name="contact" cols="40" rows="20" v-model="form.content"></textarea>
    </p>
    <p>
      <input class="form__btn" type="submit" value="送信" />
    </p>
  </form>
</div>
</template>

<script>
import axios from "axios"
import Validation from "~/components/Validation.vue";
import * as env from "~/env.json"
export default {
  components: {
    Validation
  },
  middleware: ["auth"],
  data() {
    return {
      form: {
        name: "",
        content: "",
        email: ""
      },
      errors: [],
      success: false
    }
  },
  methods: {
    valCheck: function () {
      this.errors = []
      if (this.form.name === "") {
        const empty = "名前は必須です。"
        this.errors.push(empty)
      }
      if (this.form.content === "") {
        const empty = "お問い合わせ内容は必須です。"
        this.errors.push(empty)
      }
    },
    async slack(payload) {
      const webhookUrl = env.SLACK_WEB_HOOK ? env.SLACK_WEB_HOOK : ""
      const res = await axios.post(webhookUrl, JSON.stringify(payload))
      return res.data
    },
    async submit() {
      this.valCheck()
      if (this.errors.length !== 0) {
        return
      }
      try {
        this.slack({
          text: "お問い合わせメッセージを受け付けました。",
          attachments: [{
            color: "#2eb886",
            author_name: this.$store.$auth.loggedIn ? this.$auth.user.Name : this.form.name,
            author_link: this.$store.$auth.loggedIn ? `https://on-board-project.com/users/${this.$auth.user.ID}` : "",
            author_icon: this.$store.$auth.loggedIn ? this.$auth.user.ImageURL : "",
            fields: [{
                title: this.$store.$auth.loggedIn ? `UserID : ${this.$auth.user.ID}` : "未ログインユーザー",
                value: this.form.content,
                short: false
              },
              {
                title: this.form.email ? `Email : ${this.form.email}` : "メールアドレスは指定されていません。"
              }
            ]
          }]
        }).then(() => {
          this.success = true
          this.form = {}
        })
      } catch (e) {
        this.errors = []
        const invalid = "送信に失敗しました。"
        this.errors.push(invalid)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import "assets/scss/app";

.contact__form__wrapper {
  width: 600px;
  padding: 120px 0 50px 0;
  margin: 0 auto;
}

.contact__form {
  position: relative;
  border-radius: 10px;
  border: solid 1px black;
  padding: 20px 20px 40px 20px;
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
      border-radius: 10px;
      padding: 0 10px;
      font-size: 18px;
      height: 30px;
      width: 100%;
      font-weight: bold;
    }

    & textarea {
      margin: 0;
      outline: 0;
      border: solid 1px $bg-main;
      border-radius: 10px;
      padding: 5px 10px;
      font-size: 18px;
      height: 200px;
      width: 100%;
      font-weight: bold;
    }
  }
}

.form__btn {
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

.page-title {
  text-align: center;
  font-size: 18px;
  color: grey;
}

.success-card {
  z-index: 2;
  background-color: lightgray;
  color: green;
  text-align: center;
  border: solid 1px black;
  line-height: 100px;
  width: 300px;
  height: 100px;
  position: fixed;
  bottom: 200px;
  left: 50px;
  font-size: 20px;
  font-weight: bold;

  & span {
    cursor: pointer;
    position: absolute;
    color: white;
    background-color: black;
    border-radius: 50%;
    height: 20px;
    width: 20px;
    line-height: 17px;
    font-weight: bold;
    font-size: 18px;
    top: 5px;
    right: 5px;
  }
}

.invisible {
  display: none;
}

@media screen and (max-width: $PhoneSize) {
  .contact__form__wrapper {
    padding: 100px 0;
    width: 90%;
  }

  .contact__form {
    & p {
      width: 95%;
    }
  }

  .success-card {
    width: 80%;
  }
}
</style>
