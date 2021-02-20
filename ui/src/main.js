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
import ClientListItem from "./views/client-list-item";
import KKMModelListItem from "./views/kkm-model-list-item";
import KKMListItem from "./views/kkm-list-item";
import ErrorPage from "./views/error-page";

Vue.config.productionTip = false;

Vue.use(VueMaterial);
Vue.use(VueRouter);

const router = new VueRouter({
  routes: [
    {
      path: "/kkm",
      name: "kkm-list",
      component: KKMList,
    },
    {
      path: "/kkm/create",
      name: "kkm-list-item-create",
      component: KKMListItem,
      props: {
        action: "create",
        id: -1,
      },
    },
    {
      path: "/kkm/:id/:action",
      name: "kkm-list-item",
      component: KKMListItem,
      props: (route) => ({
        action: route.params.action,
        id: +route.params.id,
      }),
    },
    {
      path: "/kkm/models",
      name: "kkm-model-list",
      component: KKMModelList,
    },
    {
      path: "/kkm/models/create",
      name: "kkm-model-list-item-create",
      component: KKMModelListItem,
      props: {
        action: "create",
        id: -1,
      },
    },
    {
      path: "/kkm/models/:id/:action",
      name: "kkm-model-list-item",
      component: KKMModelListItem,
      props: (route) => ({
        action: route.params.action,
        id: +route.params.id,
      }),
    },
    {
      path: "/clients",
      name: "clients-list",
      component: ClientList,
    },
    {
      path: "/clients/create",
      name: "clients-list-item-create",
      component: ClientListItem,
      props: {
        action: "create",
        id: -1,
      },
    },
    {
      path: "/clients/:id/:action",
      name: "clients-list-item",
      component: ClientListItem,
      props: (route) => ({
        action: route.params.action,
        id: +route.params.id,
      }),
    },
    {
      path: "*",
      component: ErrorPage,
      props: {
        error: "404",
        message: "Страница не найдена",
      },
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
