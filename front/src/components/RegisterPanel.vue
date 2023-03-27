<template>
    <div class="register-panel">
      <el-form
      ref="ruleFormRef" 
      status-icon
      label-width="120px"
      class="register"
      >
      <el-form-item label="Password" prop="pass">
      <el-input class="input" type="password" autocomplete="off" />
      </el-form-item>
      <el-form-item label="Confirm" prop="checkPass">
      <el-input
        type="password"
        autocomplete="off"
      />
      </el-form-item>
      <el-form-item label="Age" prop="age">
      <el-input/>
      </el-form-item>

      <el-form-item>
      <el-button type="primary"
        >Submit</el-button
      >
      <el-button >Reset</el-button>
      </el-form-item>
      </el-form>
    </div>
  
</template>

<script>
export default {

}
</script>

<style scoped>
.register-panel {
  width: 500px;
  background-color: #f5f5f5;
  padding: 20px;
}


</style>