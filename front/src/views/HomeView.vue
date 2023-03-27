<template>
    <div class="common-layout">
      <el-container class="header">
        <el-header style="padding: 0">
            <nav-header></nav-header>
        </el-header>
      </el-container>

      <el-container class="content">
        <el-main style="padding: 0">
            <router-view></router-view>
        </el-main>
      </el-container>
    </div>
</template>

<script>
import { defineComponent } from "vue"
// import MainContent from "../components/MainContent.vue"
import NavHeader from "../components/NavHeader.vue"

export default defineComponent({
    components: {
        NavHeader
    },

    setup() {
        return{
            NavHeader
        }
    }
})
</script>

<style scoped>
.common-layout {
    top: 0;
    left: 0;
    position: fixed;
    width: 100%;
    height: 100%;
}

</style>