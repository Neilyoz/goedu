<template>
  <div id="app">
    <div class="nav" id="nav">
      <ul>
        <li>
          <router-link to="/">首页</router-link>
        </li>
        <li>|</li>
        <!-- <li>
          <router-link to="/course">课程</router-link>
        </li>
        <li>|</li> -->
        <!-- <li>
          <router-link to="/article">文章</router-link>
        </li>
        <li>|</li> -->
        <li>
          <router-link to="/about">关于</router-link>
        </li>
      </ul>

      <div class="top_right">
        <template v-if="!this.isLogin">
          <router-link to="/register">注册</router-link>
          <router-link to="/login">登录</router-link>
        </template>
        <template v-else>
          <router-link to="/article/create">写文章</router-link>
          <router-link v-if="$store.getters.userId == 1" to="/page">设置关于页</router-link>
          <a href="javascript:void(0);" @click="logout">退出</a>
        </template>
      </div>
    </div>
    <router-view />
  </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  computed: {
    ...mapGetters(["isLogin"])
  },
  methods: {
    logout() {
      localStorage.removeItem("isLogin");
      localStorage.removeItem("token");
      localStorage.removeItem("userId");
      this.$store.commit("setLoginStatus", false);
      this.$router.push({ name: "Login" });
    }
  }
};
</script>
