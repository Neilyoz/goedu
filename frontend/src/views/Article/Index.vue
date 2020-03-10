<template>
  <div class="container">
    <div>
      <div v-for="article in articles" :key="article.id">
        <div>标题：{{article.title}}</div>
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
