<script setup>
import HomeView from './views/HomeView.vue';
</script>

<template>
  <el-config-provider namespace="ep">
    <HomeView></HomeView>
  </el-config-provider>
</template>

<style scoped>
#app {
  /* color: var(--ep-text-color-primary); */
}


</style>
