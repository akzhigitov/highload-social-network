import Vue from 'vue';
import VueRouter from 'vue-router';
import store from "../store";
import Home from '../pages/Home.vue';
import Users from '../pages/Users.vue';
import Login from '../pages/core/Login.vue';
import Error from '../pages/core/Error.vue';
import Registration from "../pages/core/Registration";

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/users',
        name: 'Users',
        component: Users
    },
    {
        path: '/users/:id',
        name: 'UserPage',
        component: Home
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: {
            allowAnonymous: true
        }
    },
    {
        path: '/registration',
        name: 'Registration',
        component: Registration,
        meta: {
            allowAnonymous: true
        }
    },
    {
        path: '/error',
        name: 'Error',
        component: Error,
        meta: {
            allowAnonymous: true
        }
    },
]

const router = new VueRouter({
    routes,
});

router.beforeEach((to, from, next) => {
    if (to.matched.some((record) => !record.meta.allowAnonymous)) {
        if (store.getters.isAuthenticated) {
            next();
            return;
        }
        next("/login");
    } else {
        next();
    }
});

export default router;
