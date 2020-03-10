<template>
  <div class="container">
    <div>
      <form class="register-form" v-on:submit.prevent="submit()" autocomplete="off">
        <div class="form-items">
          <label for="username">用户名：</label>
          <input type="text" id="username" v-model="postData.username" />
        </div>

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
          <button type="submit">注册</button>
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
        username: "",
        email: "",
        password: ""
      }
    };
  },

  methods: {
    async submit() {
      const response = await this.$http.post("/register", { ...this.postData });

      if (response.data.code === 200) {
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
