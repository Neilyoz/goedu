<template>
  <div class="container">
    <div>
      <div v-html="postData.html_content"></div>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return {
      postData: {}
    }
  },
  async created() {
    let response = await this.$http.get('about/about')
    this.postData = response.data.data
  }
}
</script>
