import babelpolyfill from 'babel-polyfill'
import Vue from 'vue'
import App from './App'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-default/index.css'

//import './assets/theme/theme-green/index.css'
import VueRouter from 'vue-router'
import store from './vuex/store'


//import NProgress from 'nprogress'
//import 'nprogress/nprogress.css'
import routes from './routes.js'
import Mock from './mock'
// Mock.bootstrap();
import Router from './routes.js'
import 'font-awesome/css/font-awesome.min.css'
import axios from 'axios';
import Vuex from 'vuex'

Vue.use(ElementUI)
Vue.use(VueRouter)
Vue.use(Vuex)

//NProgress.configure({ showSpinner: false });

const router = new VueRouter({
  mode:"history",
  routes
});

router.beforeEach((to,from,next) => {
  const isLogin=localStorage.getItem("phone_num");
  if (isLogin) {
    next()
  } else {
    //next('/error')
    if (to.path === '/dist/') { //这就是跳出循环的关键
      next()
    } else {
      next('/dist/')
    }
  }
});


//router.afterEach(transition => {
//NProgress.done();
//});

new Vue({
  //el: '#app',
  //template: '<App/>',
  router,
  store,
  //components: { App }
  render: h => h(App)
}).$mount('#app')


