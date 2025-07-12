<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { register, login } from '@/api/account'
const accountInput = ref('postman')
const passwordInput = ref('postman')

const loginFail = (msg)=>{
    ElMessage({
        showClose: true,
        message: msg,
        type: 'warning',
    })
}

async function loginClick(){
    await login(accountInput.value, passwordInput.value).then((resp)=>{
        var userId = resp.data.user_id
        if(userId == undefined || userId == null || userId == ""){
            loginFail(resp.data.message)
        }
    })
}

async function registerClick(){
    await register(accountInput.value, passwordInput.value).then((resp)=>{
        ElMessage({
            showClose: true,
            message: resp.data.message,
            type: "success"
        })
    })
}

</script>

<template>
    <span>account:</span>
    <el-input v-model="accountInput" style="width: 240px" placeholder="Please input" />
    <span>password:</span>
    <el-input v-model="passwordInput" type="password" style="width: 240px" placeholder="Please input" show-password/>
    <span></span>
    <el-button @click="loginClick">Login</el-button>
    <el-button @click="registerClick">Register</el-button>
</template>

<style scoped>
</style>
