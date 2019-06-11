import Login from './views/Login.vue'
import NotFound from './views/404.vue'
import Home from './views/Home.vue'
import Main from './views/Main.vue'
import Table from './views/nav1/Table.vue'
import Form from './views/nav1/Form.vue'
import user from './views/nav1/user.vue'
import echarts from './views/charts/echarts.vue'
import forum from './views/forum'


let routes = [
    {
        path: '/dist/',
        component: Login,
        name: '',
        hidden: true
    },
    {
        path: '/404/',
        component: NotFound,
        name: '',
        hidden: true
    },
    //{ path: '/main', component: Main },
    {
        path: '/dist/Home',
        component: Home,
        name: '导航一',
        iconCls: 'el-icon-message',//图标样式class
        children: [
            { path: '/dist/main', component: Main, name: '主页', hidden: true },
            { path: '/dist/table', component: Table, name: '账单展示' },
            { path: '/dist/form', component: Form, name: '实用工具' },
            { path: '/dist/user', component: user, name: '用户信息' },
        ]
    },

    {
        path: '/dist/Home',
        component: Home,
        name: 'Charts',
        iconCls: 'fa fa-bar-chart',
        children: [
            { path: '/dist/echarts', component: echarts, name: '图表展示' }
        ]
    },
    {
        path: '/dist/Forum',
        component:Home,
        name:'个人论坛',
        iconCls: 'el-icon-message',
        children:[{
            path:'/dist/forum',component:forum,name:'论坛中心'
        }]

    },
    {
        path: '*',
        hidden: true,
        redirect: { path: '/404' }
    }
];

export default routes;
