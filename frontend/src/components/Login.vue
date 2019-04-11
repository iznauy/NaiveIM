<template>
  <div class="panel">
    <h2>登录</h2>
    <el-form class="form" label-width="60px">
      <el-form-item label="用户名">
        <el-input v-model="name" size="small"></el-input>
      </el-form-item>
      <el-form-item label-width="0">
        <el-button type="success" @click="login" size="small">登录</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>

  import {loginAPI} from "@/api/login";

  export default {
    name: "Login",
    data() {
      return {
        name: ""
      }
    },
    methods: {
      login() {
        loginAPI(this.name, res => {
          this.$store.dispatch("updateName", this.name);
          this.$router.push('/room')
        }, err => {
          console.log("err");
          this.$alert("登录出现异常", "", {
            confirmButtonText: '确定'})
        })
      }
    }
  }
</script>

<style scoped>
  .panel {
    width: 300px;
    text-align: center;
    margin: 100px auto;
  }
  .form {
    width: 200px;
    margin: 50px auto 0 auto;
  }
</style>
