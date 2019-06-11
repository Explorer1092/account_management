import axios from 'axios';
import Qs from 'qs'

axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';


let base = 'http://39.108.156.137:8000/api';

export const getHealth = params =>{return axios.get(`${base}/health`,params).then(console.log(params),res=>res.data)};

export const registerUser = body =>{return axios.post(`${base}/Register`,Qs.stringify(body)).then(res=>res.data)};

export const requestLogin = body => { return axios.post(`${base}/Login`, Qs.stringify(body)).then(res => res.data); };

export const getBillList = params => { return axios.get(`${base}/GetBills`, { params: params }).then(res=>res.data) };

export const removeBill = body => { return axios.post(`${base}/RemoveBill`, Qs.stringify(body)).then(res=>res.data); };

export const batchRemoveBill = body => { return axios.post(`${base}/BatchRemoveBill`, Qs.stringify(body)).then(res=>res.data) };

export const addBill = body => { return axios.post(`${base}/AddBill`, body).then(res=>res.data); };

export const checkin =params =>{return axios.get(`${base}/Check_in`,{params}).then(res =>res.data)}

export const postData =body =>{return axios.post(`${base}/ChangePassword`,Qs.stringify(body)).then(res =>res.data)}

export const putData =body =>{return axios.post(`${base}/SetName`,Qs.stringify(body)).then(res =>res.data)}

export const getWeekConsum = params =>{return axios.get(`${base}/GetWeekConsum`,{params:params}).then(res =>res.data)}

export const getWeekEarn =params=>{return axios.get(`${base}/GetWeekEarn`,{params:params}).then(res =>res.data)}

export const getMonthConsum =params =>{return  axios.get(`${base}/GetMonthConsum`,{params:params}).then(res =>res.data)}

export const GetAnalysis =params =>{return axios.get(`${base}/GetAnalysis`,{params:params}).then(res =>res.data)}

export const ModifyBill =body =>{return axios.post(`${base}/ModifyBill`,body).then(res=>res.data); };
