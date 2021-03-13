<template>
  <div>
    <md-table
        v-model="matchedKkms"
        md-sort="id"
        md-sort-order="asc"
    >
      <md-table-toolbar>
        <h1 class="md-title">ККМ и Клиенты</h1>
        <md-field md-clearable class="md-toolbar-section-end">
          <md-input placeholder="Поиск по таблице" @input="onSearch"/>
        </md-field>
      </md-table-toolbar>

      <md-table-empty-state
          md-label="Не найдено"
          :md-description="emptyStateDescription">
        <md-button class="md-primary md-raised">Создать запись</md-button>
      </md-table-empty-state>

      <md-table-row slot="md-table-row" slot-scope="{ item }" :md-id="item.id">
        <md-table-cell md-label="ID" md-sort-by="id" md-numeric>{{ item.id }}</md-table-cell>
        <md-table-cell md-label="Имя клиента" md-sort-by="clientName">
          <router-link :to="`/clients/${item.clientId}/view`">
            {{ item.clientName }}
          </router-link>
        </md-table-cell>
        <md-table-cell md-label="Название модели" md-sort-by="modelName">
          <router-link :to="`/kkm/models/${item.modelId}/view`">
            {{ item.modelName }}
          </router-link>
        </md-table-cell>
        <md-table-cell md-label="Серийный номер" md-sort-by="serialNumber">{{ item.serialNumber }}</md-table-cell>
        <md-table-cell md-label="Дата регистрации" md-sort-by="registerDate">{{ item.registerDate }}</md-table-cell>
        <md-table-cell md-label="ОФД" md-sort-by="ofd">{{ item.ofd }}</md-table-cell>
        <md-table-cell md-label="Акцизный" md-sort-by="isExcise">{{ item.isExcise ? 'Да' : 'Нет' }}</md-table-cell>
        <md-table-cell md-label="Система НО" md-sort-by="systemNo">{{ item.systemNo }}</md-table-cell>
        <md-table-cell md-label="Вид работ" md-sort-by="type">{{ item.type }}</md-table-cell>
        <md-table-cell md-label="ФН" md-sort-by="fn">{{ item.fn }}</md-table-cell>
        <md-table-cell md-label="Адрес" md-sort-by="address">{{ item.address }}</md-table-cell>
        <md-table-cell md-label="Дата окончания ОФД" md-sort-by="endDateOfd">{{ item.endDateOfd }}</md-table-cell>
        <md-table-cell md-label="Дата окончания ФН" md-sort-by="endDateFn">{{ item.endDateFn }}</md-table-cell>
        <md-table-cell md-label="Дней с прохождения проверки" md-sort-by="inspectionDayCount">
          {{ item.inspectionDayCount }}
        </md-table-cell>
        <md-table-cell md-label="Комментарий" md-sort-by="inspectionDayCount">{{ item.comment }}</md-table-cell>
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
  name: 'KkmList',
  data() {
    return {
      searchTerm: '',
      debounceTimeout: null,
    }
  },
  computed: {
    ...mapState(["isKkmListLoading", "kkmListError"]),
    matchedKkms: {
      get() {
        return this.$store.getters.matchedKkms(this.searchTerm);
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
      this.$router.push({name: 'kkm-list-item-create'});
    },
    openUpdateForm(id) {
      this.$router.push({name: 'kkm-list-item', params: {id: id, action: "update"}});
    },
    openDeleteForm(id) {
      this.$router.push({name: 'kkm-list-item', params: {id: id, action: "delete"}});
    },
  }
}
</script>

<style lang="scss">
.custom-fab {
  margin-right: 25px;
}
</style>