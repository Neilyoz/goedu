<template>
  <div class="container">
    <div>
      <form class="register-form" v-on:submit.prevent="submit()" autocomplete="off">
        <div class="form-items">
          <label for="email">邮箱：</label>
          <input type="text" id="email" v-model="postData.email" />
        </div>

        <div class="form-items">
          <label for="password">密码：</label>
          <input type="password" id="password" v-model="postData.password" />
        </div>

        <div class="form-items">
          <label></label>
          <button type="submit">登录</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      postData: {
        email: "",
        password: ""
      }
    };
  },

  methods: {
    async submit() {
      let response = await this.$http.post("/login", { ...this.postData });
      if (response.data.code === 200) {
        localStorage.setItem("token", response.data.data.token);
        this.$store.commit("setToken", response.data.data.token);

        localStorage.setItem("isLogin", true);
        this.$store.commit("setLoginStatus", true);

        localStorage.setItem("userId", response.data.data.user_id)
        this.$store.commit("setUserId", response.data.data.user_id);

        this.$router.push("/");
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.register-form {
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 15px;

  .form-items:not(:last-child) {
    margin-bottom: 10px;
  }

  .form-items {
    label {
      width: 80px;
      display: inline-block;
      text-align: right;
    }

    input {
      height: 25px;
      border-radius: 5px;
      border: 1px solid #ccc;
      padding: 0 5px;
      outline: none;
    }

    button {
      border: 1px solid #ccc;
      padding: 8px 10px;
      border-radius: 5px;
      cursor: pointer;
    }
  }
}
</style>
