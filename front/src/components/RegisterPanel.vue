<template>
    <div class="register-panel">
      <el-form
      ref="ruleFormRef" 
      status-icon
      label-width="120px"
      class="register"
      :rules="registerFormRules"
      >
      <el-form-item>
      <div class="title-1" style="font-size: small; margin-right: 20px; color:#4b4b4b;">还没有账号❓</div>
      <div class="title-2" style="font-size: medium; font-weight: 700; color:#4b4b4b">现在加入barley-club🍺</div>
      </el-form-item>
      <el-form-item label="Phone" prop="telephone">
      <el-input v-model="user.telephone"
      />
      </el-form-item>
      <el-form-item label="Password" prop="password">
      <el-input type="password" v-model="user.password" autocomplete="off" />
      </el-form-item>
      
      <el-form-item label="Username">
      <el-input  v-model="user.name"/>
      </el-form-item>

      <el-form-item>
      <el-button type="primary" style="border-radius: 100px;"
        >Register</el-button
      >
      <el-button style="border-radius: 100px;">Reset</el-button>
      </el-form-item>
      </el-form>
    </div>
  
</template>

<script setup>
import { reactive } from 'vue'

const user = reactive({
  telephone: '',
  password: '',
  name: ''
})

const registerFormRules =  reactive({
        // 验证用户名是否合法
        telephone: [
          { required: true, message: '请输入手机号', trigger: 'blur' },
          { min: 11, max: 11, message: '手机号应为11位', trigger: 'blur' }
        ],
        // 验证密码是否合法
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 15, message: '密码应为 6 到 15 个字符', trigger: 'blur' }
        ]
      })
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