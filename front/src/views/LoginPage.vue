<template>
    <div class="login-page">
        <div>来壶高的</div>
        <register-panel></register-panel>
        <!-- <el-row>
            <el-col :span="12"><div class="grid-content ep-bg-purple" /></el-col>
            <el-col :span="12"><div class="grid-content ep-bg-purple-light" /></el-col>
        </el-row> -->
    </div>
  </template>
  
  <script>
  import { defineComponent } from "vue"
  import RegisterPanel from "../components/RegisterPanel.vue";
  
  export default defineComponent({
    components: {
        RegisterPanel
    },

    setup() {
          return {
            RegisterPanel
          }
    }
  })
  </script>
  
  <style scoped>
  
  </style>