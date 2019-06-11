<template>
	<div class="all-container">
		<div class="all-container-padding bg" >
			<el-tabs v-model="activeName" @tab-click="handleClick">

				<el-tab-pane label="修改昵称" name="first">
					<el-form :model="userlist" :rules="rules" ref="EditorUserForms">


						<el-form-item label="上传头像" prop="avatar_url" :label-width="formLabelWidth" >
                        <el-upload action="http://39.108.156.137:8000/api/InputImg"  :data="uploadParm" accept="image/jpeg,image/png,image/jpg"  :limit="1" :on-exceed="handleExceed"
									   :on-error="handleError" :on-success="handleSuccess" :on-change="onUploadChange">
									<input type="file" name ="file">
							</el-upload>
						</el-form-item>


						<el-form-item label="用户昵称" prop="full_name" :label-width="formLabelWidth">
							<el-col :span="5">
								<el-input v-model="userlist.full_name"  placeholder="请输入修改的昵称" ></el-input>
							</el-col>
						</el-form-item>
					</el-form>


					<div class="grid-content bg-purple">
						<el-button type="primary" @click="EditorUserClick('userlist')" >保存修改</el-button>
					</div>


				</el-tab-pane>
				<el-tab-pane label="修改密码" name="second">
					<el-form :model="ruleForm" :rules="rules" ref="ruleForm">
						<el-form-item label="原密码" prop="pass" :label-width="formLabelWidth">
							<el-col :span="8">  <el-input v-model="ruleForm.pass" placeholder="请输入原密码" type="password"></el-input></el-col>
						</el-form-item>
						<el-form-item label="新密码" prop="newpass" :label-width="formLabelWidth">
							<el-col :span="8"><el-input v-model="ruleForm.newpass" placeholder="请输入新密码" id="newkey" type="password"></el-input></el-col>
						</el-form-item>
						<el-form-item label="重复新密码" prop="checknewpass" :label-width="formLabelWidth">
							<el-col :span="8">  <el-input v-model="ruleForm.checknewpass" placeholder="请再次输入新密码" id='newkey1' type="password"></el-input></el-col>
						</el-form-item>
					</el-form>
					<div class="grid-content bg-purple">
						<el-button type="primary" @click="submitForm('ruleForm')">保存</el-button>
					</div>



				</el-tab-pane>

				<el-tab-pane label="导入账单" name="third">
					<span>上传时间会根据账单大小而定，请耐心等待...</span>
					<br><br>
					<el-upload
							 class="upload-demo"
							 ref="upload"
                             action="http://39.108.156.137:8000/api/UploadAliExcel"
							 :limit="1"
							 :on-exceed="handleExceed"
							 accept="*.xls,*.xlc"
							 :data="uploadexcel"
							:on-error="handleError" :on-success="handleSuccess"
							>
						 <el-button slot="trigger" size="small" type="primary"  style="margin-right: 20px">上传支付宝excel</el-button>
						&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="https://jingyan.baidu.com/article/72ee561a1c068ee16138dfd9.html" target="_blank">怎么下载支付宝账单？</a>
						<!-- <div slot="tip" class="el-upload__tip">只能上传excel文件，且不超过5MB</div>-->
						<!-- <div slot="tip" class="el-upload-list__item-name">{{fileName}}</div>-->
					</el-upload > 

					<br>
					<br>

						<el-upload
								 class="upload-demo2"
								 ref="upload"
                                 action="http://39.108.156.137:8000/api/UploadWechatExcel"
								 :limit="1"
								:on-exceed="handleExceed"
								accept="*.xls,*.xlc"
								 :data="uploadexcel"
								:on-error="handleError" :on-success="handleSuccess">
							 <el-button slot="trigger" size="small" type="primary"  style="margin-right: 20px">上传微信excel</el-button>
							<a href="https://jingyan.baidu.com/article/29697b91426b9eab20de3cfa.html" target="_blank">如何导出微信账单？</a>

					</el-upload>
					<br><br>


				</el-tab-pane>

				<br><br><br>


			</el-tabs>
		</div>
	</div>
</template>
<script>


//	import {postData, putData} from "../../api/api";
//	import ElTabPane from "element-ui/packages/tabs/src/tab-pane";
//	import ElUpload from "element-ui/packages/upload/src/index";
import {postData, putData} from "../../api/api";
import ElTabPane from "element-ui/packages/tabs/src/tab-pane";


export default {
	components: {ElTabPane},
	// components: {ElUpload, ElTabPane},
		data() {
			/*****检验两次密码是否一致***/

			var validatePass = (rule, value, callback) => {
				if (value === "") {
					callback(new Error("请输入密码"));
				} else {
					if (this.ruleForm.checknewpass !== "") {
						this.$refs.ruleForm.validateField("checknewpass");
					}
					callback();
				}
			};
			var validatePass2 = (rule, value, callback) => {
				if (value === "") {
					callback(new Error("请再次输入密码"));
				} else if (value !== this.ruleForm.newpass) {
					callback(new Error("两次输入密码不一致!"));
				} else {
					callback();
				}
			};
			return {
				uploadexcel:{
					phone_num:localStorage.getItem('phone_num')
				},

				uploadParm: {
					phone_num:localStorage.getItem("phone_num"),
				}, //图片的上传
				ruleForm: {},//修改密码的表单
				activeName: "first",
				loading: true,
				baseUrl: process.env.BASE_API,
				userlist: {
					phone:localStorage.getItem("phone_num"),
					full_name:''
				},//用户信息表单
				formLabelWidth: "150px",
				/***校验***/
				rules: {
					phone: [
						{
							required: true,
							message: "请输入电话号码"
						},
						{
							pattern: /^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$/,
							message: "手机格式不对"
						}
					],

					pass: [
						{
							required: true,
							trigger: "blur",
							message: "请输入密码"
						}
					],
					newpass: [
						{
							validator: validatePass,
							trigger: "blur"
						}
					],
					checknewpass: [
						{
							validator: validatePass2,
							trigger: "blur"
						}
					]
				}
			};
		},
		created() {

		},
		// computed: {
		// 	...mapGetters(["username"])
		// },
		methods: {



			newSubmitForm(){
				this.$refs.newupload.submit()
			},

			submitUpload(){
				alert('submit')
			},

			onUploadChange(file){

			},
			//获取个人用户的信息
			handleExceed(){
				alert('每次最多只能发送1个文件')
			},
			// 	//tab切换
			handleClick(tab, event) {
				console.log(tab, event);
			},
				//上传参数图片初始化

			// 	//上传之前
			// beforeupload(file) {
			// 	this.uploadParm.key = this.uploadParm.dir ;
			// 	 console.log(this.uploadParm)
			// },
			//图片上传上传成功

			handleSuccess(res, file, fileList) {
				this.$notify.success({
					title: '成功',
					message: `文件上传成功`
				});
			},
			// 文件上传失败时的钩子
			handleError(err, file, fileList) {
				this.$notify.error({
					title: '错误',
					message: `文件上传失败`
				});
			},
			//修改密码
			submitForm(ruleForm) {
				var parmas = {
					username: localStorage.getItem("phone_num"),
					oldpwd: this.ruleForm.pass,
					newpwd: this.ruleForm.newpass
				};
				if (parmas.oldpwd ==parmas.newpwd){
					this.$message({
						message:"不能和旧密码一致",
						type:"error"
					})
                    return 
				}
                if(this.ruleForm.newpass !=this.ruleForm.checknewpass){
                    this.$message({
                        message:'两次密码输入不一致',
                        type:'error'
                    });
                    return
                }

				postData(parmas).then(response => {
					if (response.code == 200) {
						this.$message({
							message: "修改成功",
							type: "success"
						});
					} else {
						this.$message({
							message: "修改失败" + response.msg,
							type: "error"
						});
					}
				});

			},
			// 	// 编辑提交的方法//昵称
			EditorUserClick() {
				this.$refs.EditorUserForms.validate(valid => {
					if (valid) {

						putData(this.userlist).then(response => {
							if (response.code == 200) {
								this.$message({
									message: "编辑成功",
									type: "success"
								});
							} else {
								alert(response.code)
								this.$message({
									message: "修改失败" + response.msg,
									type: "error"
								});
							}
						});
					}
				});
			}

		}
	};
</script>

<style>
	@import url("//unpkg.com/element-ui@2.8.2/lib/theme-chalk/index.css");
	.avatar-uploader .el-upload {
		border: 1px dashed #d9d9d9;
		border-radius: 3px;
		cursor: pointer;
		position: relative;
		overflow: hidden;
	}
	.avatar-uploader .el-upload:hover {
		border-color: #409EFF;
	}


	.avatar-uploader-icon {
		font-size: 28px;
		color: #8c939d;
		width: 100px;
		height: 100px;
		line-height: 80px;
		text-align: center;
	}
	.avatar {
		width: 178px;
		height: 178px;
		display: block;
	}
	.position{
		position: absolute;
		left: 100px;
		top: 20px;
	}
</style>
