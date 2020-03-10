import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isLogin: localStorage.getItem('isLogin') || false,
    token: localStorage.getItem('token') || '',
    userId: localStorage.getItem('userId') || '',
  },
  getters: {
    isLogin: state => {
      return state.isLogin;
    },
    token: state => {
      return state.token;
    },
    userId: state => {
      return state.userId
    }
  },
  mutations: {
    setLoginStatus(state, status) {
      state.isLogin = status
    },
    setToken(state, token) {
      state.token = token
    },
    setUserId(state, user_id) {
      state.userId = user_id
    }
  },
  actions: {

  },
  modules: {
  }
})
