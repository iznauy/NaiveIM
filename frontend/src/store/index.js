import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    name: ""
  },
  getters: {
    name(state) {
      return state.name;
    }
  },
  mutations : {
    setName(state, name) {
      state.name = name;
    }
  },
  actions: {
    updateName(context, name) {
      context.commit('setName', name)
    }
  }
})
