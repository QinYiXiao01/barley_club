<template>
  <div class="login-page">
      <el-row>
          <el-col :span="12">
            <el-card class="box-card">
            <template #header>
              <div class="card-header">
                <h1>🎇欢迎来到Barley Club🎇</h1>
              </div>
            </template>
            <div class="text item">
              在这里，你能够：
            </div>
            <div class="text item">
              <div class="title">
                <el-icon size="50px"><Notebook /></el-icon>
                <strong style="font-size: larger;">记录喝过的美酒佳酿</strong><br>
                <div>Take down all the good brews</div>
              </div>
            </div>
            <div class="text item">
              <div class="title">
                <el-icon size="50px"><Mug /></el-icon>
                <strong style="font-size: larger;">探索无穷的啤酒世界</strong><br>
                <div>Explore the infinite beer world</div>
              </div>
            </div>
            <div class="text item">
              <div class="title">
                <el-icon size="50px"><MapLocation /></el-icon>
                <strong style="font-size: larger;">随时随地找到目的地</strong><br>
                <div>Find your destination right away</div>
              </div>
            </div>
            <div class="text item">
              <div class="title">
                <el-icon size="50px"><Medal /></el-icon>
                <strong style="font-size: larger;">收集你的专属徽章</strong><br>
                <div>Collect your achievement medals</div>
              </div>
            </div>
            
          </el-card>
          </el-col>
          <el-col :span="12">
            <register-panel></register-panel>
            <login-panel></login-panel>
          </el-col>
      </el-row>
  </div>
</template>

<script>
import { defineComponent } from "vue"
import RegisterPanel from "../components/RegisterPanel.vue";
import LoginPanel from "../components/LoginPanel.vue";

export default defineComponent({
  components: {
      RegisterPanel,
      LoginPanel
  },

  setup() {
        return {
          RegisterPanel,
          LoginPanel
        }
  }
})
</script>

<style scoped lang="scss">
.login-page {
  display: flex;
}
.ep-row {
  flex:1;
  flex-wrap: wrap;
}

.ep-col {
  flex:auto;
}
.box-card {
  margin: 40px;
  margin-left: 100px;
  background-color: #f5f5f5;
}
.card-header {
display: flex;
justify-content: space-between;
align-items: center;
}

.text {
  font-size: 20px;
}

.item {
  margin-bottom: 30px;
}

.box-card {
  // width: 480px;
}

.title {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
}

.ep-icon {
  transition: all 0.2s;
  margin: 0 10px;
  &:hover {
    color: #fd780f;
  }
}


</style>