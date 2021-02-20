<template>
  <md-content>
    <h1 class="md-display-3">Что-то пошло не так</h1>
    <h3 class="md-display-2">{{ error }}</h3>
    <p class="md-display-1" v-if="message">{{ message }}</p>
  </md-content>
</template>

<script>
export default {
  name: "ErrorPage",
  props: {
    error: {
      type: String,
      required: true,
    },
    message: {
      type: String,
      default: '',
    },
  },
  methods: {
    goBack() {
      this.$router.back();
    },
  }
}
</script>