<template>
    <div class="register-panel">
      <el-form
      ref="ruleFormRef" 
      status-icon
      label-width="120px"
      class="register"
      >
      <el-form-item>
      <div class="title-1" style="font-size: small; margin-right: 20px; color:#4b4b4b;">还没有账号❓</div>
      <div class="title-2" style="font-size: medium; font-weight: 700; color:#4b4b4b">现在加入barley-club🍺</div>
      </el-form-item>
      <el-form-item label="Phone" prop="checkPass">
      <el-input
      />
      </el-form-item>
      <el-form-item label="Password" prop="pass">
      <el-input type="password" autocomplete="off" />
      </el-form-item>
      
      <el-form-item label="Username" prop="age">
      <el-input/>
      </el-form-item>

      <el-form-item>
      <el-button type="primary" style="border-radius: 100px;"
        >Submit</el-button
      >
      <el-button style="border-radius: 100px;">Reset</el-button>
      </el-form-item>
      </el-form>
    </div>
  
</template>

<script>
export default {
  setup() {
    
  }

}
</script>

<style scoped lang="scss">
.register-panel {
  transition: all 0.2s;
  width: 400px;
  background-color: #f5f5f5;
  padding: 20px 60px 20px 2px;
  margin: 40px;
  border-radius: 10px;
}

.register-panel:focus-within{
  box-shadow: 0 2px 4px 3px #fee5cf;
}


</style>