<template>
  <div>
    <md-table
        v-model="matchedClients"
        md-card
        md-sort="id"
        md-sort-order="asc"
        md-fixed-header
    >
      <md-table-toolbar>
        <h1 class="md-title">Клиенты</h1>
        <md-field md-clearable class="md-toolbar-section-end">
          <md-input placeholder="Поиск по таблице" @input="onSearch"/>
        </md-field>
      </md-table-toolbar>
      <md-table-empty-state
          md-label="Не найдено"
          :md-description="emptyStateDescription">
        <md-button class="md-primary md-raised" @click="openAddForm">Создать клиента</md-button>
      </md-table-empty-state>
      <md-table-row slot="md-table-row" slot-scope="{ item }" :md-id="item.id">
        <md-table-cell md-label="ID" md-sort-by="id" md-numeric>{{ item.id }}</md-table-cell>
        <md-table-cell md-label="Имя" md-sort-by="name">{{ item.name }}</md-table-cell>
        <template>
          <md-table-cell md-label="Действия">
            <span class="md-layout md-alignment-center-center">
              <md-button class="md-primary" @click="openUpdateForm(item.id)">Изменить</md-button>
              <md-button class="md-accent" @click="openDeleteForm(item.id)">Удалить</md-button>
            </span>
          </md-table-cell>
        </template>
      </md-table-row>
    </md-table>
    <md-button class="md-fab md-primary custom-fab md-fab-bottom-right" @click="openAddForm">
      <md-icon>add</md-icon>
    </md-button>
  </div>
</template>

<script>
import {mapState} from "vuex";

export default {
  name: 'ClientList',
  data() {
    return {
      searchTerm: '',
      debounceTimeout: null,
    }
  },
  computed: {
    ...mapState(["isClientListLoading", "clientListError"]),
    matchedClients: {
      get() {
        return this.$store.getters.matchedClients(this.searchTerm);
      },
      set() {
      }
    },
    emptyStateDescription() {
      if (this.searchTerm === '') {
        return 'Записей не найдено';
      }

      return `Записей по запросу '${this.searchTerm}' не найдено`
    }
  },
  methods: {
    onSearch(searchTerm) {
      if (this.debounceTimeout) {
        clearTimeout(this.debounceTimeout);
      }

      this.debounceTimeout = setTimeout(() => {
        this.searchTerm = searchTerm;
      }, 250);
    },
    openAddForm() {
      this.$router.push({name: 'clients-list-item-create'});
    },
    openUpdateForm(id) {
      this.$router.push({name: 'clients-list-item', params: {id: id, action: "update"}});
    },
    openDeleteForm(id) {
      this.$router.push({name: 'clients-list-item', params: {id: id, action: "delete"}});
    },
  }
}
</script>

<style lang="scss">
.custom-fab {
  margin-right: 25px;
}
</style>