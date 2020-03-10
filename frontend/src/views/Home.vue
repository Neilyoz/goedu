<template>
  <div class="container">
    <div>
      <div class="article_wrap" v-for="article in articles" :key="article.id">
        <div class="title">
          标题：{{ article.title }}
          <span class="sub" v-if="$store.getters.isLogin && $store.getters.userId == article.user_id">
            [<router-link
            :to="{ name: 'Article/Edit', params: { id: article.id } }"
            >编辑</router-link
          >]
            </span>
        </div>
        <hr />
        <div v-html="article.html_content"></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      articles: []
    };
  },

  created() {
    this.getArticles();
  },

  methods: {
    async getArticles() {
      const response = await this.$http.get("/article");
      if (response.data.code === 200) {
        this.articles = response.data.data;
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.article_wrap {
  border: 1px solid #ccc;
  margin-bottom: 10px;
  padding: 15px;

  .title {
    font-size: 18px;
    margin-bottom: 10px;

    .sub {
      font-size: 12px;
    }
  }
}
</style>
