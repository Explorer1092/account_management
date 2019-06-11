<template>
	<div class="all-container">
		<div class="all-container-padding bg" >
			<el-tabs v-model="activeName" @tab-click="handleClick">
				<el-tab-pane label="个税计算器" name="first">
					&nbsp &nbsp &nbsp本月工资 &nbsp
					<el-input v-model="income"   style=" width:200px " placeholder="请输入" class="form-input"></el-input>&nbsp元
					<br>
					<br>

					&nbsp &nbsp &nbsp纳税期数 &nbsp
					<el-select v-model="value" placeholder="请选择" class="form-input">
					<el-option
							v-for="item in Month"
							:key="item.value"
							:label="item.label"
							:value="item.value">
					</el-option>
					</el-select>
					<br>
					<br>

					&nbsp &nbsp &nbsp累计 &nbsp &nbsp &nbsp &nbsp&nbsp
					<el-input v-model="income_total"  style=" width:200px " placeholder="必填" class="form-input"></el-input>&nbsp元
					&nbsp<el-button @click="outputs" type="primary" size="small"  style="background-color: #18c79c">自动获取</el-button>
					<br>
					<br>
					&nbsp &nbsp &nbsp专项附加扣除

					<el-select v-model="value2"  multiple  placeholder="请选择">
						<el-option
								v-for="item2 in plusList"
								:key="item2.value"
								:label="item2.label"
								:value="item2.value">
						</el-option>
					</el-select>

					<br><br>
					&nbsp &nbsp &nbsp五险一金 &nbsp
					<el-button @click="edit">填写</el-button>
					<el-dialog title="五险一金" v-model="viable" :close-on-click-modal="false" >
						<el-form v-model="insurance" label-width="80px">
							公积金 &nbsp
							<el-input v-model="insurance.Provident_fund" style="width: 100px"></el-input>%
							<br><br>
							医疗 &nbsp &nbsp &nbsp
							<el-input v-model="insurance.Medical" style="width: 100px"></el-input>%
							<br><br>
							养老 &nbsp &nbsp &nbsp
							<el-input v-model="insurance.PensionOld" style="width: 100px"></el-input>%
							<br><br>
							失业 &nbsp &nbsp &nbsp
							<el-input v-model="insurance.LoseJob" style="width:100px"></el-input>%
							<br><br>
							工伤 &nbsp &nbsp &nbsp
							<el-input v-model="insurance.Injury" style="width:100px"></el-input>%
							<br><br>
							生育 &nbsp &nbsp &nbsp
							<el-input v-model="insurance.Fertility" style="width:100px"></el-input>%
							<br><br>

							<el-button @click="addSubmit">提交</el-button>
						</el-form>
					</el-dialog>
					<br><br>
					&nbsp &nbsp &nbsp
					<el-button @click="compute" style="background-color:#f7ba2a;">开始计算</el-button>
					&nbsp &nbsp
					<el-button @click="cancel" type="primary">重置</el-button>


					<el-dialog title="" v-model="visable" ::close-on-click-modal="false">
						<el-form>
								<p class="word-v-middle">到手可得</p>
								<p class="money-word">{{result}}</p>
							<br>
								税前收入: &nbsp &nbsp &nbsp  &nbsp &nbsp{{income}}
								<br><br>
									个人所得税: &nbsp &nbsp &nbsp  &nbsp &nbsp {{result_tax}}
								<br><br>
									五险一金合计: &nbsp &nbsp &nbsp  &nbsp &nbsp {{this.resInsurrance}}

								<br><br> ----------------------------------------------------------
									<p class="parma">个税明细</p>

									累计应纳税所得额: &nbsp &nbsp &nbsp  &nbsp &nbsp {{this.NeedTax}}
								<br><br>
									适应税率: &nbsp &nbsp &nbsp  &nbsp &nbsp  {{this.rate}}
								<br><br>
									基准扣除:&nbsp &nbsp &nbsp  &nbsp &nbsp {{this.subNum}}

								<br><br>
									累计应缴纳税 		:&nbsp &nbsp &nbsp  &nbsp &nbsp {{this.totalSub}}
								<br><br>
									已缴税款 :&nbsp &nbsp &nbsp  &nbsp &nbsp{{this.haveSub}}

							<br><br> ----------------------------------------------------------
								<p class="parma">五险一金明细</p>

									公积金:&nbsp &nbsp &nbsp  &nbsp &nbsp {{this.insuranceNum.Provident_fund}}
								<br><br>
									医疗:&nbsp &nbsp &nbsp  &nbsp &nbsp{{this.insuranceNum.Medical}}
								<br><br>
									养老:&nbsp &nbsp &nbsp  &nbsp &nbsp {{this.insuranceNum.PensionOld}}
								<br><br>
									失业:&nbsp &nbsp &nbsp  &nbsp &nbsp {{this.insuranceNum.LoseJob}}
								<br><br>
									工伤:&nbsp &nbsp &nbsp  &nbsp &nbsp {{this.insuranceNum.Injury}}
								<br><br>
									生育:&nbsp &nbsp &nbsp  &nbsp &nbsp {{this.insuranceNum.Fertility}}

							<br>
						</el-form>

					</el-dialog>

				</el-tab-pane>
			</el-tabs>
		</div>
	</div>
</template>

<script>

	import ElRadioGroup from "element-ui/packages/radio/src/radio-group";
	export default {
		components: {ElRadioGroup},
		data() {
			return {
				value:1,
				value2:'',
				income_total:'',
				insurance:{
					Provident_fund:12,
					Medical:2,
					PensionOld:8,
					LoseJob:0.2,
					Injury:0,
					Fertility:0
				},
				viable:false,
				income:'',

				insuranceNum:{
					Provident_fund:0,
					Medical:0,
					PensionOld:0,
					LoseJob:0,
					Injury:0,
					Fertility:0
				},
				activeName:"first",
				Month:[
					{
						value:'1',
						label:'1月'
					},{
						value:'2',
						label:'2月'
					},
					{
						value:'3',
						label:'3月'
					},
					{
						value:'4',
						label:'4月'
					},
					{
						value:'5',
						label:'5月'
					},
					{
						value:'6',
						label:'6月'
					},
					{
						value:'7',
						label:'7月'
					},
					{
						value:'8',
						label:'8月'
					},
					{
						value:'9',
						label:'9月'
					},
					{
						value:'10',
						label:'10月'
					},
					{
						value:'11',
						label:'11月'
					},
					{
						value:'12',
						label:'12月'
					},
				],


				plusList:[
					{
						value:'子女教育',
						label:'子女教育   1000'
					},
					{
						value:'继续教育',
						label:'继续教育   400'
					},
					{
						value:'住房房贷利息',
						label:'住房房贷利息 1000'
					},
					{
						value:'住房租金',
						label:'住房租金 1000'
					},
					{
						value:'赡养老人',
						label:'赡养老人 2000'
					},
				],
				rate:0,
				subNum:0,
				eachMonth:[],
				result_tax:0,
				visable:false,
				resInsurrance:'',
				NeedTax:0,
				result:0,
				total:0,
				haveSub:0,
				totalSub:0
			}
		},
		methods: {
			outputs(){
				this.income_total=this.income*parseInt(this.value)
			},
			handleClick(tab, event) {
				console.log(tab, event);
			},
			edit(){
				this.viable=true
			},
			addSubmit(){
				this.viable=false
			},

			dataInit(){
				this.NeedTax=0;
				this.total=0;
				this.haveSub=0;
				this.totalSub=0;
				this.rate=0;
				this.subNum=0;
			},
			compute() {
				this.dataInit()

				var splice_res =this.value2.toString()

				var status =splice_res.indexOf('继续教育')
				var total =0;
				if (status != -1) {
					total = total + 400
				}
				var status =splice_res.indexOf('子女教育')
				if (status!=-1) {
					total =total+1000
				}
				var status =splice_res.indexOf('住房房贷利息')
				if (status!=-1) {
					total =total+1000
				}
				var status =splice_res.indexOf('住房租金')
				if (status!=-1) {
					total =total+1000
				}
				var status =splice_res.indexOf('赡养老人')
				if (status!=-1) {
					total =total+2000
				}
				this.total=total;
				var res_insurrance;
				var insurrance = parseFloat(this.insurance.PensionOld) + parseFloat(this.insurance.Medical) + parseFloat( this.insurance.Provident_fund) +parseFloat( this.insurance.LoseJob) +parseFloat( this.insurance.Fertility) +parseFloat( this.insurance.Injury)
				// if (this.income >25500){
				// 	res_insurrance=5639;
				// 	this.insuranceNum.Provident_fund=3048.12;
				// 	this.insuranceNum.Medical=508.02;
				// 	this.insuranceNum.PensionOld=2032.08;
				// 	this.insuranceNum.Fertility=0;
				// 	this.insuranceNum.LoseJob=50.8;
				// 	this.insuranceNum.Injury=0;
				//
				// }else {
				// 	res_insurrance = this.income * insurrance / 100;
				// }
				// this.resInsurrance=res_insurrance



				// var plusnum=total+res_insurrance
				// if (this.income-plusnum-5000 >0) {
				// 	this.NeedTax = this.income - total - res_insurrance - 5000
				// 	if (this.NeedTax <0){
				// 		this.NeedTax=0
				// 	}
				// }
				if (this.income <=2270){//2270以下直接就是全部收入，无需缴纳其他费用
					this.result=this.income;
					this.result_tax=0;
					this.resInsurrance=0;

				}
				if (this.income >2270 &&this.income<=3380){//2270~3380之间只用交公积金
					this.result =this.income-this.insurance.Provident_fund*this.income/100;
					this.insuranceNum.Provident_fund=this.insurance.Provident_fund*this.income/100;
					this.resInsurrance=this.insuranceNum.Provident_fund;
					this.insuranceNum.Medical=0;
					this.insuranceNum.PensionOld=0;
					this.insuranceNum.LoseJob=0;
					this.getInsurrenceNum()
				}


				if (this.income>3380&&this.income <=5000){//3380~5000只用交五险一金
					res_insurrance=this.income * insurrance / 100;
					this.resInsurrance=res_insurrance
					this.result=this.income-res_insurrance
					this.getInsurrenceNum()
				}

				if(this.income >5000 && this.income<=25500){ //这是收入大于5000
					this.resInsurrance=insurrance*this.income/100
					if(this.income -this.resInsurrance-total<=5000){
						this.getInsurrenceNum()
						this.result=this.income-this.resInsurrance
					}
					else{
						this.getInsurrenceNum()//ok
						this.computeTax()
						this.resInsurrance=insurrance*this.income/100
						this.result=this.income-this.result_tax-this.resInsurrance
					}
				}
				if(this.income >25500){
					this.resInsurrance=5639;
					this.insuranceNum.Provident_fund=3048.12;
					this.insuranceNum.Medical=508.02;
					this.insuranceNum.PensionOld=2032.08;
					this.insuranceNum.Fertility=0;
					this.insuranceNum.LoseJob=50.8;
					this.insuranceNum.Injury=0;
					this.computeTax()
					this.result=this.income-this.result_tax-this.resInsurrance
				}
				this.visable=true
			},


			getInsurrenceNum(){
				if(this.income<=25500) {
					this.insuranceNum.Provident_fund = this.income * this.insurance.Provident_fund / 100
					this.insuranceNum.LoseJob = this.income * this.insurance.LoseJob / 100
					this.insuranceNum.Fertility = this.income * this.insurance.Fertility / 100
					this.insuranceNum.Injury = this.income * this.insurance.Injury / 100
					this.insuranceNum.PensionOld = this.income * this.insurance.PensionOld / 100
					this.insuranceNum.Medical = this.income * this.insurance.Medical / 100
				}
			},
			 computeTax(){
				const month = parseInt(this.value);
				for( var i=1;i<=12;i++){
					this.getLevel(this.income*i-5000*i-this.total*i-this.resInsurrance*i)
					this.eachMonth[i]=((this.income*i-5000*i-this.total*i-this.resInsurrance*i)*this.rate)-this.subNum
				}
				this.getLevel(this.income*month-5000*month-this.total*month-this.resInsurrance*month)
				switch (month) {
					case 1:this.result_tax=this.eachMonth[1];break;
					case 2:this.result_tax=this.eachMonth[2]-this.eachMonth[1];break;
					case 3:this.result_tax=this.eachMonth[3]-this.eachMonth[2];break;
					case 4:this.result_tax=this.eachMonth[4]-this.eachMonth[3];break;
					case 5:this.result_tax=this.eachMonth[5]-this.eachMonth[4];break;
					case 6:this.result_tax=this.eachMonth[6]-this.eachMonth[5];break;
					case 7:this.result_tax=this.eachMonth[7]-this.eachMonth[6];break;
					case 8:this.result_tax=this.eachMonth[8]-this.eachMonth[7];break;
					case 9:this.result_tax=this.eachMonth[9]-this.eachMonth[8];break;
					case 10:this.result_tax=this.eachMonth[10]-this.eachMonth[9];break;
					case 11:this.result_tax=this.eachMonth[11]-this.eachMonth[10];break;
					case 12:this.result_tax=this.eachMonth[12]-this.eachMonth[11];break;
				}
				console.log(this.eachMonth)
				this.totalSub=this.eachMonth[month]
				if (month>1){
					this.haveSub=this.eachMonth[month-1]
				}
				 this.NeedTax = (this.income - this.total - this.resInsurrance - 5000)*month
			},

			getLevel(total){
				if (total<=36000){this.rate=0.03;this.subNum=0;return}
				if (total>36000&&total<=144000){this.rate=0.1;this.subNum=2520;return}
				if (total>144000&&total<=300000){this.rate=0.2;this.subNum=16920;return}
				if (total>300000&&total<=420000){this.rate=0.25;this.subNum=31920 ;return}
				if (total>420000&&total<=660000){this.rate=0.3;this.subNum=52920;return}
				if (total>660000&&total<=960000){this.rate=0.35;this.subNum=85920;return}
				if (total>960000){this.rate=0.45;this.subNum=181920;return}
				console.log("default problem is"+total)

				return
			},
			cancel(){
				this.income='';
				this.income_total='';
				this.value='';
				this.value2='';
				this.insurance.Provident_fund='12',
				this.insurance.Medical='2',
				this.insurance.PensionOld='8',
				this.insurance.LoseJob='0.2',
				this.insurance.Injury='0',
				this.insurance.Fertility='9'
			}
		}
	}

</script>

<style>
	.word-v-middle{
		margin-bottom: 0;
		font-size: 16px;
		min-height: 31px;
		display: flex;
		align-items: center;
		justify-content: center;
		height: 1px;
		margin-top: 5px;
		color: #87878a;
		white-space: normal;
		color: black;
	}

	.money-word{
		margin-bottom: 0;
		font-size: 35px;
		min-height: 31px;
		display: flex;
		align-items: center;
		justify-content: center;
		height: 1px;
		margin-top: 5px;
		color: #87878a;
		white-space: normal;
		color: orange;
	}

	.parma{
		font-size: 20px;
		color: crimson;
	}
</style>
