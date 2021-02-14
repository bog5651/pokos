import g from "guark";
import Vue from "vue";
import VueRouter from "vue-router";
import App from "./App.vue";
import store from "./store";
import VueMaterial from "vue-material";
import "vue-material/dist/vue-material.min.css";
import "vue-material/dist/theme/default.css";

import KKMList from "./views/kkm-list.vue";
import KKMModelList from "./views/kkm-model-list.vue";
import ClientList from "./views/client-list.vue";

Vue.config.productionTip = false;

Vue.use(VueMaterial);
Vue.use(VueRouter);

const router = new VueRouter({
  routes: [
    {
      path: "/kkm",
      component: KKMList,
    },
    {
      path: "/kkm/models",
      component: KKMModelList,
    },
    {
      path: "/clients",
      component: ClientList,
    },
  ],
});
new Vue({
  store,
  router,
  render: (h) => h(App),
  created: () => g.hook("created"),
  mounted: () => g.hook("mounted"),
}).$mount("#app");
