import Vue from "vue";
import App from "./App.vue";
import { BootstrapVue, BootstrapVueIcons } from "bootstrap-vue"
import router from "./router/router";
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;

Vue.use(BootstrapVue);
Vue.use(BootstrapVueIcons);

new Vue({
  router,
  vuetify,
  render: (h) => h(App)
}).$mount("#app");