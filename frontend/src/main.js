import Vue from 'vue';
import App from './App';
import router from './router';
import store from './store';

import 'vuetify/dist/vuetify.min.css';
import 'font-awesome/css/font-awesome.css';

import Vuetify from 'vuetify';

import 'material-design-icons-iconfont/dist/material-design-icons.css';
import './styles/global.css';

import axios from "@/plugins/axios";
import VueAxios from 'vue-axios'
import {setupComponents} from './config/setup-components';

Vue.use(VueAxios, axios);
Vue.use(Vuetify);


setupComponents(Vue);
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
    el: '#app',
    store,
    router,
    vuetify: new Vuetify(),
    components: {App},
    render: h => h(App)
})
