<template>
  <div class="nav-header">
      <el-menu
    :default-active="activeIndex"
    mode="horizontal"
    :ellipsis="false"
    @select="handleSelect"
    background-color="#fd780f"
    text-color="#fff"
    active-text-color="#ffd04b"
  >
    <el-menu-item index="0">
      <div class="logo-box">
        <img class="logo" src="../assets/img/logo-no-background.png" alt="">
      </div>
    </el-menu-item>
    <div class="flex-grow" />
    <el-menu-item index="1" @click="$router.replace('/diary')">啤酒日记</el-menu-item>
    <el-menu-item index="2" @click="$router.replace('/find')">喝一杯</el-menu-item>
    <el-menu-item index="3" @click="$router.replace('/my')">喝过的</el-menu-item>
    <el-menu-item index="4" @click="$router.replace('/wiki')">能喝啥</el-menu-item>
    <el-sub-menu index="5">
      <template #title>登录/注册</template>
      <el-menu-item index="5-1" @click="$router.replace('/login')">登录</el-menu-item>
      <el-menu-item index="5-2" @click="$router.replace('/login')">注册</el-menu-item>
    </el-sub-menu>
  </el-menu>
      <!-- <el-menu
      :default-active="activeIndex2"
      class="el-menu"
      mode="horizontal"
      background-color="#fd780f"
      text-color="#fff"
      active-text-color="#ffd04b"
      @select="handleSelect"
    >
    <el-menu-item>
      <div class="logo-box">
        <img class="logo" src="../assets/img/logo-no-background.png" alt="">
      </div>
    </el-menu-item>
      <router-link to="diary">
        <el-menu-item index="1">
            啤酒日记
        </el-menu-item>
      </router-link>
  
      <router-link to="find">
        <el-menu-item index="2">
          喝一杯
        </el-menu-item>
      </router-link>
  
      <router-link to="my">
        <el-menu-item index="3">喝过的</el-menu-item>
      </router-link>
  
      <router-link to="wiki">
        <el-menu-item index="4">有什么</el-menu-item>
      </router-link>
      </el-menu>
      <div class="user-info"> -->
        <!-- <el-dropdown>
         <span class="el-dropdown-link">
      登录 / 注册
      <el-icon class="el-icon--right">
        <arrow-down />
      </el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item>Action 1</el-dropdown-item>
        <el-dropdown-item>Action 2</el-dropdown-item>
        <el-dropdown-item>Action 3</el-dropdown-item>
        <el-dropdown-item disabled>Action 4</el-dropdown-item>
        <el-dropdown-item divided>Action 5</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
      </div> -->
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
.nav-header {
    width: 100%;
    z-index: 100;
}

.content {
  background-color: #fd780f;

  display: flex;
  justify-content: space-between;
  align-items: center;
  flex: 1;
}
.logo {
  width: 100px;
}
.el-menu >>> .el-menu-item {

}

.flex-grow {
  flex-grow: 1;
}

.el-dropdown-link {
  display: flex;
  align-items: center;
  font-weight: 700;
  padding: 20px;
  color: white;
  cursor: pointer;
}

</style>