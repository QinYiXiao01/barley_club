<template>
  <div class="nav-header">
    <div class="content">
      <div class="logo-box">
        <img class="logo" src="../assets/img/logo-no-background.png" alt="">
      </div>
      <el-menu
      :default-active="activeIndex2"
      class="el-menu"
      mode="horizontal"
      background-color="#fd780f"
      text-color="#fff"
      active-text-color="#ffd04b"
      @select="handleSelect"
    >
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
      <div class="user-info">
        <el-dropdown>
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
      </div>
    </div>
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
  width: 80px;
  margin: 20px 20px;
}

.el-menu {
  flex: 1;
}

.el-menu >>> .el-menu-item {
  font-size: large;
  font-weight: 700;
  padding: 40px 20px;
  margin: 0 10px;
}

a {
  text-decoration: none;
}

.router-link-active {
  text-decoration: none;
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