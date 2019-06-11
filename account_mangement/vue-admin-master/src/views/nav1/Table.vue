<template>
	<section>
		<!--工具条-->
		<el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
			<el-form :inline="true" :model="filters">
				<el-form-item>
					<el-input v-model="filters.date" placeholder="日期"></el-input>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" v-on:click="getBills">查询</el-button>
				</el-form-item>
				<el-form-item >
					<el-button type="primary" @click="AddBill">新增</el-button>
				</el-form-item>
			</el-form>
		</el-col>

		<!--列表-->
		<el-table v-bind:data="bill.slice((page-1)*pagesize,page*pagesize)"  highlight-current-row v-loading="listLoading" @selection-change="selsChange" style="width: 100%;">
			<el-table-column type="selection" width="55">
			</el-table-column>
			<el-table-column type="index" width="60" >
				<template scope="scope">
					<span>{{(page - 1) * pagesize + scope.$index + 1}}</span>
				</template>
			</el-table-column>
			<el-table-column prop="Time" label="交易时间" width="120" >
			</el-table-column>
			<el-table-column prop="Type" label="类型" width="100" >
			</el-table-column>
			<el-table-column prop="Otherside" label="交易对方" width="100">
			</el-table-column>
			<el-table-column prop="Goodsname" label="商品名称" width="120">
			</el-table-column>
			<el-table-column prop="Price" label="金额" min-width="120px">
			</el-table-column>
			<el-table-column prop="Classification" label="分类" min-width="180">
			</el-table-column>
			<el-table-column label="操作" width="150">
				<template scope="scope">
					<el-button  size ="small" type="primary" @click="modify(scope.$index, scope.row)">修改</el-button>
					<el-button type="danger"  size="small" @click="handleDel(scope.$index, scope.row)">删除</el-button>
				</template>
			</el-table-column>
		</el-table>

		<!--工具条-->
		<el-col :span="24" class="toolbar">
			<el-button type="danger" @click="batchRemove" :disabled="this.bill.length===0">批量删除</el-button>
			<el-pagination layout="total ,prev, pager,next" @current-change="handleCurrentChange" :current-page="page" :page-size="pagesize" :total="total" style="float:right;">
			</el-pagination>
		</el-col>


		<el-dialog class="el-dialog__body" title="新增账单" v-model="addFormVisible" :close-on-click-modal="true">
			<el-form :model="addForm" 	label-width="80px"  ref="editForm">
				<el-form-item label="日期" >
					<el-date-picker
							type="date"
							:picker-options="pickerOptions0"
							v-model="addForm.date"
							placeholder="选择日期">
					</el-date-picker>
				</el-form-item>

				<el-form-item label="时间点" >
					<el-time-picker type="time"  placeholder="选择时间" v-model="addForm.time"></el-time-picker>
				</el-form-item>
				<el-form-item label="商品描述">
					<el-input v-model="addForm.goodsname" auto-complete="on"></el-input>
				</el-form-item>
				<el-form-item label="类别" >
					<el-radio-group v-model="addForm.type">
						<el-radio class="radio" :label="AliPay">支付宝</el-radio>
						<el-radio class="radio" :label="Wechat">微信</el-radio>
					</el-radio-group>
				</el-form-item>
				<el-form-item label="资金流向" >
					<el-radio-group v-model="addForm.way">
						<el-radio class="radio" :label="Outflow" @change="">支出</el-radio>
						<el-radio class="radio" :label="Inflow">收入</el-radio>
					</el-radio-group>
				</el-form-item>

				<el-form-item label="交易对方">
					<el-input  v-model="addForm.otherside"></el-input>
				</el-form-item>

				<el-form-item label="金额" >
					<el-input  v-model="addForm.price"></el-input>
				</el-form-item>

				<el-form-item label="分类" >
					<el-select v-model="value" placeholder="请选择">
						<el-option
								v-for="item in options"
								:key="item.value"
								:label="item.label"
								:value="item.value">
						</el-option>
					</el-select>
				</el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button @click="cancel">取消</el-button>
				<el-button type="primary" @click.native="addSubmit" :loading="addLoading">提交</el-button>
			</div>
		</el-dialog>



		<el-dialog class="el-dialog__body" title="修改账单" v-model="modifyVisible" :close-on-click-modal="true" >
			<el-form :model="addForm"  label-width="80px"  ref="editForm">
				<el-form-item label="日期" >
					<el-date-picker
							type="date"
                            :picker-options="pickerOptions0"
							v-model="addForm.date"
							placeholder="选择日期">
					</el-date-picker>
				</el-form-item>

				<el-form-item label="时间点" >
					<el-time-picker type="time"  placeholder="选择时间" v-model="addForm.time"></el-time-picker>
				</el-form-item>
				<el-form-item label="商品描述">
					<el-input v-model="addForm.goodsname" auto-complete="on"></el-input>
				</el-form-item>
				<el-form-item label="类别" >
					<el-radio-group v-model="addForm.type">
						<el-radio class="radio" :label="AliPay">支付宝</el-radio>
						<el-radio class="radio" :label="Wechat">微信</el-radio>
					</el-radio-group>
				</el-form-item>
				<el-form-item label="资金流向" >
					<el-radio-group v-model="addForm.way">
						<el-radio class="radio" :label="Outflow" @change="">支出</el-radio>
						<el-radio class="radio" :label="Inflow">收入</el-radio>
					</el-radio-group>
				</el-form-item>

				<el-form-item label="交易对方">
					<el-input  v-model="addForm.otherside"></el-input>
				</el-form-item>

				<el-form-item label="金额" >
					<el-input  v-model="addForm.price"></el-input>
				</el-form-item>

				<el-form-item label="分类" >
					<el-select v-model="value" placeholder="请选择">
						<el-option
								v-for="item in options"
								:key="item.value"
								:label="item.label"
								:value="item.value">
						</el-option>
					</el-select>
				</el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button @click="cancel">取消</el-button>
				<el-button type="primary" @click.native="modifySubmit" :loading="addLoading">提交</el-button>
			</div>
		</el-dialog>



	</section>
</template>
<script>
	import axios from 'axios';
	import util from '../../common/js/util'

	//import NProgress from 'nprogress'
	import {getBillList, removeBill, batchRemoveBill, editBill, addBill, ModifyBill} from '../../api/api';
	import {index} from "element-ui";
	import ElRadioGroup from "element-ui/packages/radio/src/radio-group";

	export default {
		components: {ElRadioGroup},
		data() {
			return {
				oldtime:"",
				filters: {
					date: ''
				},

				Inflow:'收入',
				Outflow:'支出',

				AliPay:'支付宝',
				Wechat:'微信',

				pagesize:20,
				bill:[],

				bluePageIndex:1,
				total: 1,
				page: 1,
				listLoading: false,
				sels: [],//列表选中列

				addFormVisible: false,//新增界面是否显示
				addLoading: false,
				modifyVisible:false,

				//新增界面数据
				addForm: {
					phone_num: localStorage.getItem('phone_num'),
					date:'',
					time: '',
					type: '',
					otherside: '',
					goodsname: '',
					price: '',
					way: '',
					classification:''
				},

				modForm:{
					phone_num: localStorage.getItem('phone_num'),
					date:'',
					time: '',
					type: '',
					otherside: '',
					goodsname: '',
					price: '',
					way: '',
					classification:''
				},

				pickerOptions0: {
					disabledDate(time) {
						return time.getTime() > Date.now();
					}
				},

				options: [{
					value: '饮食消费',
					label: '饮食消费'
				}, {
					value: '交通消费',
					label: '交通消费'
				}, {
					value: '生活消费',
					label: '生活消费'
				}, {
					value: '教育消费',
					label: '教育消费'
				}, {
					value:	'转账',
					label:	'转账'
				},{
					value: '其他',
					label: '其他'
				}],
				value: ''

			}
		},


		methods: {


			getBills() {//获取账单列表
				var _this=this;
				let para = {
					phone_num: localStorage.getItem("phone_num"),
					date: _this.filters.date,
				};
				this.listLoading = true;
				//NProgress.start();

               axios.get('http://39.108.156.137:8000/api/GetBills',{
					params:{
						phone_num:localStorage.getItem("phone_num"),
						page:this.page,
						date:this.filters.date,
					}
				}).then(function (res) {
					if (res.data.code==200){
						_this.bill=(res.data.data)
						_this.total=res.data.totalPage

					}else {
						console.log(res.data.msg)
					}
				},function () {
					console.log("failed")
				})

				this.listLoading = false;
			},



			handleCurrentChange(val) {
				this.page = val;
				// this.getBills();
			},


			//删除
			handleDel: function (index, row) {
				this.$confirm('确认删除该记录吗?', '提示', {
					type: 'warning'
				}).then(() => {
					this.listLoading = true;
					//NProgress.start();

					let para = {
						phone_num:localStorage.getItem('phone_num'),
						time: row.Time,
						otherside: row.Otherside,
						type:row.Type,
						goodsname: row.Goodsname,
					};
					removeBill(para).then((res) => {
						this.listLoading = false;
						//NProgress.done();
						if (res.code ==200){
							this.$message({
								message: '删除成功',
								type: 'success'
							});
							this.getBills();
						}else{
							this.$message({
								message:res.msg,
								type:'error'
							})
						}
					});
				}).catch(() => {

				});
			},

			AddBill: function(){
				this.addFormVisible=true
			},


			cancel(){
				this.addLoading=false;
				this.addFormVisible = false;

				this.addForm.time='';
				this.addForm.type='';
				this.addForm.otherside='';
				this.addForm.goodsname='';
				this.addForm.price='';
				this.addForm.way='';
				this.addForm.date='';
			},



			//新增
			addSubmit: function (time) {
				this.$confirm('确认提交吗？', '提示', {}).then(() => {
					this.addLoading = true;
							//NProgress.start();

				try{
					let para = {
					date:this.addForm.date.toLocaleDateString(),
					time: this.addForm.time.toLocaleTimeString().substring(2,this.addForm.time.toLocaleTimeString().length),
					price: parseInt(this.addForm.price,10)
				};
				if (parseFloat(this.addForm.price).toString()== "NaN"){
					alert("price error!")
					this.cancel()
					return
				}
				}catch (e) {
					this.$message({
						message:'输入错误',
						type:'error'
					});
					this.cancel()
					return
				}
					let para = {
						phone_num: localStorage.getItem('phone_num'),
						date:this.addForm.date.toLocaleDateString(),
						time: this.addForm.time.toLocaleTimeString().substring(2,this.addForm.time.toLocaleTimeString().length),
						type: this.addForm.type,
						otherside: this.addForm.otherside,
						goodsname: this.addForm.goodsname,
						price: this.addForm.price,
						way: this.addForm.way,
						classification:this.value
					};
                    if(this.value=='' ||this.addForm.date =='' || this.addForm.time=='' || this.addForm.value=='' || this.addForm.way=='' || this.addForm.price=='' || this.addForm.type ==''){
                        this.$message({
                            message:'请完整输入',
                            type:'error'
                        });
                        this.cancel()
                        return
                    }
							addBill(para).then((res) => {
								this.addLoading = false;
								//NProgress.done();
								if (res.code==200){
									this.$message({
										message: '提交成功',
										type: 'success'
									});
                                    this.cancel();
									this.getBills();

								}else{
									this.$message({
										message:res.msg,
										type:'error'
									}),
                                    this.cancel()
								}
							},function () {
								console.log("request failed!")
							});
						});
					},

			selsChange: function (sels) {
				this.sels = sels;
			},

			modify(index, row){
				this.oldtime= row.Time,
				console.log(this.oldtime)
				this.$confirm('确认修改这条账单吗？', '提示', {
					type:'warning'
				}).then(()=>{
					this.modifyVisible=true
				})
			},

			modifySubmit(){
				this.$confirm('确认提交修改吗？', '提示', {
					type:'warning'
				}).then(()=> {
					try {
						let para = {
							date: this.addForm.date.toLocaleDateString(),
							time: this.addForm.time.toLocaleTimeString().substring(2, this.addForm.time.toLocaleTimeString().length),
							price: parseInt(this.addForm.price,10)
						};
						if (parseFloat(this.addForm.price).toString()== "NaN"){
							alert("price error!")
							this.cancel()
							return
						}
					} catch (e) {
						this.$message({
							message: '输入错误',
							type: 'error'
						});
						this.cancel()
						return
					}
					let para = {

						phone_num: localStorage.getItem('phone_num'),
						date: this.addForm.date.toLocaleDateString(),
						time: this.addForm.time.toLocaleTimeString().substring(2, this.addForm.time.toLocaleTimeString().length),
						type: this.addForm.type,
						otherside: this.addForm.otherside,
						goodsname: this.addForm.goodsname,
						price: this.addForm.price,
						way: this.addForm.way,
						classification: this.value,
						oldtime: this.oldtime,
					};
					if (this.value == '' || this.addForm.date == '' || this.addForm.time == '' || this.addForm.value == '' || this.addForm.way == '' || this.addForm.price == '' || this.addForm.type == '') {
						this.$message({
							message: '请完整输入',
							type: 'error'
						});
						this.cancel()
						return
					}
					ModifyBill(para).then((res) => {
						this.addLoading = false;
						//NProgress.done();
						if (res.code == 200) {
							this.$message({
								message: '提交成功',
								type: 'success'
							});
							this.cancel();
							this.getBills();
							this.modifyVisible=false;
						} else {
							this.$message({
								message: res.msg,
								type: 'error'
							}),
						this.cancel()
							this.modifyVisible=false;
						}
					}, function () {
						console.log("request failed!")
					});
				});
			},

			//批量删除
			batchRemove: function () {
				var times = this.sels.map(item => item.Time).toString();
				this.$confirm('确认删除选中记录吗？', '提示', {
					type: 'warning'
				}).then(() => {
					this.listLoading = true;
					//NProgress.start();
					let para = {
						phone_num:localStorage.getItem('phone_num'),
						time: times
					};
                    if(para.time ==''){
                        this.$message({
                            message:'请选择要删除的账单',
                            type:'error'
                        });
                        this.listLoading=false;
                        return
                    }
					batchRemoveBill(para).then((res) => {
						this.listLoading = false;
						//NProgress.done();
						if (res.code==200) {
							this.$message({
								message: '删除成功',
								type: 'success'
							});
							this.getBills();
						}else{
							this.$message({
								message:res.msg,
								type:'error'
							})
						}
					});
				}).catch(() => {

				});
			}
		},
		mounted() {
			this.getBills()
		},
	}

</script>

<style scoped>

	.el-dialog {
		 position: absolute;
		top: 50%;
		eft: 50%;
		margin: 0 !important;
		transform: translate(-50%, -50%);
		max-height: calc(100% - 30px);
		max-width: calc(100% - 30px);
		display: flex;
		flex-direction: column;
	}
	   .el-dialog__body {
		overflow: auto;
		color: #666666;
		}
	 

	.locations{
		position: absolute;
		left: 30px;
		top: 20px;
	}
</style>

