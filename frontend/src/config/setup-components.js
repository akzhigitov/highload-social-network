// Core Components
import Toolbar from '../components/core/Toolbar.vue';
import Navigation from '../components/core/NavigationDrawer.vue';
import UserTreeView from '../components/UserTreeView.vue';
import Login from "@/pages/core/Login";

// import Registration from "../pages/core/Registration.vue";


function setupComponents(Vue) {

    Vue.component('toolbar', Toolbar);
    Vue.component('navigation', Navigation);
    Vue.component('user-tree-view', UserTreeView);
    Vue.component('login', Login);
}


export {
    setupComponents
}
