<template>
  <div class="main-content">
    <el-row>
      <el-input class="search-bar" v-model="input4" placeholder="搜索你想知道的...">
          <template #prefix>
            <el-icon class="el-input__icon"><search /></el-icon>
          </template>
      </el-input>
    </el-row>
    <el-row>
      <el-col
        :span="8"
        :offset="index > 0 ? 2 : 0"
      >
        <el-card :body-style="{ padding: '0px' }">
          <img
            src="../assets/img/beer-type.jpg"
            class="image"
          />
          <div style="padding: 14px">
            <span>啤酒种类大不同🔍</span>
            <div class="bottom">
              <time class="time">{{ currentDate }}</time>
              <el-button text class="button">阅读详情</el-button>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col
        :span="8"
        :offset="index > 0 ? 2 : 0"
      >
      <el-card :body-style="{ padding: '0px' }">
        <img
          src="../assets/img/logo-color.png"
          class="image"
          style="width:100%"
        />
        <div style="padding: 14px">
          <span>来到Barley Club的第一天🚩</span>
          <div class="bottom">
            <time class="time">{{ currentDate }}</time>
            <el-button text class="button">Operating</el-button>
          </div>
        </div>
      </el-card>
      </el-col>
    </el-row>
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


<style scoped lang="scss">


.main-content{
  // display: flex;
  background: url('../assets/img/main-beer.jpg');
  background-size: cover;
  height: 100vh;
  position: relative;

  &::after {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: black;
    opacity: 0.3;
  }
}

.ep-row {
  flex: 1;
  align-items: center;
  justify-content: center;
  margin-bottom: 30px;
  z-index:100;
}

.ep-card {
  height: 400px;
  margin: 20px 20px;
  border-radius: 20px;
}
.eq-input {
  width: 500px;
}
.ep-input {
    width: 400px;
    height: 80px;
    margin-top: 40px;
    margin-bottom: 30px;
    border-radius: 20px;
    overflow: hidden;

    z-index:10;
    font-size: 20px;

    :deep(&__wrapper) {
      border-radius: 20px;
    }
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.button {
  padding: 0;
  min-height: auto;
}

.image {
  width: 100%;
  height: 300px;
  display: block;
}



</style>


