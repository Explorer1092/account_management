<template>
    <section class="chart-container">
        <el-row>
            <el-col :span="12">
                <div id="chartColumn" style="width:100%; height:400px;"></div>
            </el-col>

            <el-col :span="12">
                <div id="chartPie" style="width:100%; height:400px;"></div>
            </el-col>

            <el-col :span="12">
                <div id="analysisChart" style="width:100%; height:400px;"></div>
            </el-col>


        </el-row>
    </section>
</template>

<script>
    import echarts from 'echarts'
    import {GetAnalysis, getMonthConsum, getWeekConsum, getWeekEarn} from "../../api/api";

    export default {
        data() {
            return {
                chartColumn: null,
                chartPie: null,
                analysisChart:null
            }
        },

        methods: {
            DrawColumnChart() {
                this.chartColumn = echarts.init(document.getElementById('chartColumn'));

                var params={
                    phone_num:localStorage.getItem("phone_num"),

                };
                getWeekConsum(params).then(data =>{
                    if (data.code ==200){
                        getWeekEarn(params).then(data1 =>{
                            if (data1.code ==200){
                                this.chartColumn.setOption({
                                    title: { text: '最近一周消费' },
                                    tooltip: {},
                                    xAxis: {
                                        data: ["6天前","5天前","4天前", "3天前", "2天前", "昨天",  "今天"]
                                    },
                                    yAxis: {},
                                    legend: {
                                        data: ['支出', '收入']
                                    },
                                    series: [{
                                        name: '支出',
                                        type: 'bar',
                                        data: [
                                            data.data[0].seven,data.data[0].six,data.data[0].five,data.data[0].four,data.data[0].three,data.data[0].two,data.data[0].one
                                        ]
                                    },
                                        {
                                            name:'收入',
                                            type:'bar',
                                            data:[
                                                data1.data[0].seven,data1.data[0].six,data1.data[0].five,data1.data[0].four,data1.data[0].three,data1.data[0].two,data1.data[0].one
                                            ]
                                        }
                                    ]
                                });
                            }
                            else{
                                console.log("DrawColumnChart-"+data.msg)
                            }
                        })

                    }
                });
            },


            DrawanalysisChart(){
                this.analysisChart = echarts.init(document.getElementById('analysisChart'));
                var params={
                    phone_num :localStorage.getItem("phone_num")
                }
                GetAnalysis(params).then(data =>{
                    if (data.code ==200){
                        this.analysisChart.setOption({
                            title: {
                                text: '本月账单分析',

                                x: 'center'
                            },
                            tooltip: {
                                trigger: 'item',
                                formatter: "{a} <br/>{b} : {c} ({d}%)"
                            },
                            legend: {
                                orient: 'vertical',
                                left: 'left',
                                data: ['饮食消费', '交通消费','生活消费','教育消费','转账','其他']
                            },
                            series: [
                                {
                                    name:'过去的一个月',
                                    type: 'pie',
                                    radius: '55%',
                                    center: ['50%', '60%'],
                                    data: [
                                        { value: data.data[0], name: '饮食消费' },
                                        { value: data.data[1], name: '交通消费' },
                                        { value: data.data[2], name: '生活消费' },
                                        { value: data.data[3], name: '教育消费' },
                                        { value: data.data[4], name: '转账' },
                                        { value: data.data[5], name: '其他' },
                                    ],
                                    itemStyle: {
                                        emphasis: {
                                            shadowBlur: 10,
                                            shadowOffsetX: 0,
                                            shadowColor: 'rgba(0, 0, 0, 0.5)'
                                        }
                                    }
                                }
                            ]
                        });
                    }else{
                       console.log("DrawanalysisChart-"+data.msg)
                    }
                })
            },

            DrawPieChart() {
                this.chartPie = echarts.init(document.getElementById('chartPie'));
                var params = {
                    phone_num: localStorage.getItem("phone_num")
                };
                getMonthConsum(params).then(data =>{
                    if(data.code==200){
                        this.chartPie.setOption({
                            title: {
                                text: '本月收支',

                                x: 'center'
                            },
                            tooltip: {
                                trigger: 'item',
                                formatter: "{a} <br/>{b} : {c} ({d}%)"
                            },
                            legend: {
                                orient: 'vertical',
                                left: 'left',
                                data: ['收入来源', '支出来源']
                            },
                            series: [
                                {
                                    name: '过去的一个月',
                                    type: 'pie',
                                    radius: '55%',
                                    center: ['50%', '60%'],
                                    data: [
                                        { value: data.data[1], name: '支出来源' },
                                        { value: data.data[0], name: '收入来源' },
                                    ],
                                    itemStyle: {
                                        emphasis: {
                                            shadowBlur: 10,
                                            shadowOffsetX: 0,
                                            shadowColor: 'rgba(0, 0, 0, 0.5)'
                                        }
                                    }
                                }
                            ]
                        });
                    }else{
                        console.log("DrawPieChart"+data.msg)
                    }
                });
            },
            drawCharts() {
                this.DrawColumnChart();
                this.DrawanalysisChart();
                this.DrawPieChart();
            },
        },

        mounted: function () {
            this.drawCharts()
        },
        updated: function () {
            this.drawCharts()
        }
    }
</script>

<style scoped>
    .chart-container {
        width: 100%;
        float: left;
    }
    /*.chart div {
        height: 400px;
        float: left;
    }*/

    .el-col {
        padding: 30px 20px;
    }
</style>
