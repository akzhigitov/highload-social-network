import axios from "@/plugins/axios";
import Vuex from "vuex";
import Vue from "vue";

const state = {
    userID: null,
    token: null
};
const getters = {
    isAuthenticated: state => !!state.userID,
    token: state => state.token,
    userID: state => state.userID
};
const actions = {
    async Register({dispatch}, form) {
        await axios.post('register', form)
        let UserForm = new FormData()
        UserForm.append('username', form.username)
        UserForm.append('password', form.password)
        await dispatch('LogIn', UserForm)
    },
    async LogIn({commit}, User) {
        const response = await axios.post('/auth/sign-in', User)

        await commit('setToken', response.data.token)
        await commit('setUser', response.data.id)
    },
    async LogOut({commit}) {
        commit('logOut')
    }

};
const mutations = {
    setUser(state, userID) {
        state.userID = userID
    },
    setToken(state, token) {
        state.token = token
    },
    logOut(state) {
        state.userID = null
        state.token = null
    },
};
Vue.use(Vuex)
export default new Vuex.Store({state, getters, actions, mutations})