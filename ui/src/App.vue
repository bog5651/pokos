<template>
  <div id="app" class="page-container">
    <md-app md-waterfall md-mode="fixed" md-alignment="centered" class="md-app">
      <md-app-toolbar class="md-primary">
        <md-tabs class="md-primary" md-sync-route>
          <md-tab to="/clients" md-label="Клиенты"></md-tab>
          <md-tab to="/kkm" md-label="ККМ и клиенты" exact></md-tab>
          <md-tab to="/kkm/models" md-label="Модели ККМ" exact></md-tab>
        </md-tabs>
      </md-app-toolbar>
      <md-app-content>
        <router-view/>
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
import g from "guark"
import {mapActions} from 'vuex';

export default {
  name: "App",
  data() {
    return {
      isLoading: false,
    };
  },
  created() {
    this.isLoading = true;

    g.env().then(env => {
      // Disable right click menu on production, you can implement your own!
      if (!!env.dev_mode) {
        document.addEventListener('contextmenu', e => e.preventDefault());
      }
    })

    Promise.allSettled([
      this.fetchKkmList(),
      this.fetchClientList(),
      this.fetchKkmModelsList(),
    ]).then(() => this.isLoading = false);
  },
  methods: {
    ...mapActions(["fetchClientList", "fetchKkmList", "fetchKkmModelsList"]),
  },
}
</script>

<style lang="scss" scoped>
body, html {
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0;
  border: 0;
}

#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

img {
  -webkit-user-drag: none;
  -khtml-user-drag: none;
  -moz-user-drag: none;
  -o-user-drag: none;
  user-drag: none;
}

/** {*/
/*  user-select: none;*/
/*  -moz-user-select: none;*/
/*  -webkit-user-select: none;*/
/*  -ms-user-select: none;*/
/*  touch-callout: none;*/
/*  -webkit-touch-callout: none;*/
/*}*/

input {
  user-select: auto;
  -moz-user-select: auto;
  -webkit-user-select: auto;
  -ms-user-select: auto;
}

.md-app {
  width: 100%;
  height: 100%;
}
</style>