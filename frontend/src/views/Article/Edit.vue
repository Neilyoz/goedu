<template>
  <div class="container">
    <form v-on:submit.prevent="submit()">
      <div>
        <label for>标题</label>
        <input type="text" v-model="postData.title" />
      </div>
      <div>
        <label for>文章</label>
        <textarea cols="30" rows="10" v-model="postData.plain_content"></textarea>
      </div>

      <div>
        <label for></label>
        <button type="submit">添加</button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      postData: {
        title: "",
        plain_content: ""
      }
    };
  },
  methods: {
    async submit() {
      const response = await this.$http.post("/article/create", {
        ...this.postData
      });

      if (response.data.code === 200) {
        this.$router.push({ name: "Article" });
      }
    }
  }
};
</script>
