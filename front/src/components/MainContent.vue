<template>
  <div class="main-content">
    <el-input class="search-bar w-50 m-2" v-model="input4" placeholder="Type something">
        <template #prefix>
          <el-icon class="el-input__icon"><search /></el-icon>
        </template>
    </el-input>
  </div>
</template>

<script>
import { defineComponent } from "vue"

export default defineComponent({
    setup() {
        return {}
    }
})
</script>


<style scoped>
.search-bar >>> .el-input__wrapper {
    border-radius: 20px;
}

.main-content{
    background: url('../assets/img/main-beer.jpg');
    background-size: cover;
    height: 100vh;
}

.search-bar {
    position: absolute;
    width: 400px;
    height: 80px;

    border-radius: 20px;
    overflow: hidden;
    z-index:10;
    top: 20%;
    left: 50%;
    transform: translate(-50%, -50%);
    font-size: 20px;
}



</style>


