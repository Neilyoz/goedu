import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isLogin: localStorage.getItem('isLogin') || false,
    token: localStorage.getItem('token') || '',
  },
  getters: {
    isLogin: state => {
      return state.isLogin;
    },
    token: state => {
      return state.token;
    }
  },
  mutations: {
    setLoginStatus(state, status) {
      state.isLogin = status
    },
    setToken(state, token) {
      state.token = token;
    }
  },
  actions: {

  },
  modules: {
  }
})
