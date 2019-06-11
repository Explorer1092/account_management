<template>
  <div class="note" :style ="note">
    <el-form :model="ruleForm2" :rules="rules2" ref="ruleForm2" label-position="left" label-width="0px" class="demo-ruleForm login-container">
      <h3 class="title">系统登录</h3>
      <el-form-item prop="account">
        <el-input type="text" v-model="ruleForm2.phone_num"  auto-complete="off" placeholder="账号"></el-input>
      </el-form-item>
      <el-form-item prop="checkPass">
        <el-input type="password" v-model="ruleForm2.login_passwd" auto-complete="off" placeholder="密码"></el-input>
      </el-form-item>
      <el-checkbox v-model="checked"  class="remember">记住密码</el-checkbox>
      <el-form-item style="width:100%;">

        <el-button type="primary" style="width:100%; color:yellow; bgcolor:black" @click.native.prevent="register" :loading="registered">注册</el-button>
        <br>

        <el-button type="success" style="width:100%;" @click.native.prevent="handleSubmit2" :loading="logining">登录</el-button>
        <br>
        <el-button @click.native.prevent="handleReset2">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>


<script>
  //import NProgress from 'nprogress'
  import {registerUser, requestLogin} from "../api/api";
  import * as $route from "express/lib/router/layer";

  export default {
    data() {
      return {
        logining: false,
        registered:false,
        ruleForm2: {
          phone_num:'18710681922',
          login_passwd:''
        },
        rules2: {
          phone_num: [
            { required: true, message: '请输入账号', trigger: 'blur' },
            //{ validator: validaePass }
          ],
          login_passwd: [
            { required: true, message: '请输入密码', trigger: 'blur' },
            //{ validator: validaePass2 }
          ]
        },
        checked: false,

        note:{
          backgroundImage:"url(" + require("../assets/background.jpg") + ")",
          backgroundRepeat: "no-repeat",
          backgroundSize: "2000px auto",
          height:"645px",
          marginTop: "0px",

        }
      };
    },


    methods: {
      register(){
        var _this = this;
        this.$refs.ruleForm2.validate((valid) => {
          if (valid) {
            //_this.$router.replace('/table');
            //NProgress.start();
            var loginParams = {
              username: this.ruleForm2.phone_num,
              password: this.ruleForm2.login_passwd,
            };


            registerUser(loginParams).then(data => {
              //NProgress.done();
              let { code, msg, user } = data;
              if (code !== 200) {
                this.$message({
                  message: msg,
                  type: 'error'
                });
              }else{
                this.$message({
                  message:msg,
                  type:'success'
                })
              }
            })
          }
          else{
            console.log('error submit!!');
            return false;
          }
        })
      },


      handleReset2() {
        this.ruleForm2.phone_num=''
        this.ruleForm2.login_passwd=''
      },


      handleSubmit2(ev) {
        var _this = this;
        this.$refs.ruleForm2.validate((valid) => {
          if (valid) {
            //_this.$router.replace('/table');
            this.logining = true;
            //NProgress.start();
            var loginParams = {
              phone_num: this.ruleForm2.phone_num,
              login_passwd: this.ruleForm2.login_passwd
            };

            requestLogin(loginParams).then(data => {

              this.logining = false;
              //NProgress.done();
              let { code, msg, user_token } = data;
              var abc;
              if (code !== 200) {
                this.$message({
                  message: msg,
                  type: 'error'
                });
              } else {
                this.$message({
                  message: msg,
                  type: 'success'
                })
                localStorage.setItem('phone_num', this.ruleForm2.phone_num)

                this.$router.push({path: '/dist/table'})
              }
            }),function () {
              console.log("send failed!")
            };
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      }
    }
  }

</script>

<style lang="scss" scoped>
  .login-container {
    /*box-shadow: 0 0px 8px 0 rgba(0, 0, 0, 0.06), 0 1px 0px 0 rgba(0, 0, 0, 0.02);*/
    -webkit-border-radius: 5px;
    border-radius: 5px;
    -moz-border-radius: 5px;
    background-clip: padding-box;
    position:absolute;
    margin:100px 430px;
    width: 350px;
    padding: 35px 35px 15px 35px;
    background: #fff;
    border: 1px solid #eaeaea;
    box-shadow: 0 0 25px #cac6c6;
    .title {
      margin: 0px auto 40px auto;
      text-align: center;
      color: #505458;
    }
    .remember {
      margin: 0px 0px 35px 0px;
    }
  }

</style>
