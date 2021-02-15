<template>
  <div class="md-layout">
    <div class="md-layout-item">
      <form novalidate class="md-layout md-alignment-top-center" @submit.prevent="validateListItem">
        <md-card class="md-layout-item md-size-50 md-small-size-100">
          <md-card-header>
            <div class="md-title">{{ headerText }} модели ККМ</div>
          </md-card-header>

          <md-progress-bar md-mode="indeterminate" v-if="isPending"/>

          <md-card-content>
            <div class="md-layout md-gutter">
              <div class="md-layout-item md-small-size-100" v-if="action !== 'create'">
                <md-field>
                  <label for="id">ID</label>
                  <md-input name="id" id="id" autocomplete="off" v-model="localItem.id" readonly required/>
                </md-field>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-field :class="getValidationClass('name')">
                  <label for="name">Название</label>
                  <md-input name="name" id="name" autocomplete="off" v-model="localItem.name" :disabled="isPending"
                            :readonly="['view', 'delete'].includes(action)" required/>
                  <span class="md-error" v-if="!$v.localItem.name.required">Название должно быть заполнено</span>
                </md-field>
              </div>
            </div>
          </md-card-content>

          <md-card-actions>
            <md-button type="submit" class="md-primary" :disabled="isPending">{{ actionButtonText }}</md-button>
            <md-button type="reset" :disabled="isPending" @click="onCancel">Отмена</md-button>
          </md-card-actions>
        </md-card>

        <md-snackbar :md-active="formState !== null">{{ resultText }}</md-snackbar>
      </form>
    </div>
  </div>
</template>

<script>
import {mapActions, mapState} from "vuex";
import formTextMixin from "@/mixins/formTextMixin";
import {required,} from 'vuelidate/lib/validators';
import {validationMixin} from "vuelidate";

export default {
  name: "KkmModelListItem",
  mixins: [formTextMixin, validationMixin],
  data() {
    return {
      localItem: {
        id: -1,
        name: "",
      },
      formState: null,
      resultText: '',
    }
  },
  validations: {
    localItem: {
      name: {
        required,
      }
    }
  },
  props: {
    id: {
      type: Number,
      required: true,
    },
    action: {
      type: String,
      required: true,
    }
  },
  computed: {
    ...mapState(["isKkmListLoading"]),
    isPending() {
      return this.isKkmListLoading || this.formState === "pending";
    }
  },
  watch: {
    $route() {
      this.initView();
    }
  },
  created() {
    this.initView();
  },
  methods: {
    ...mapActions(["upsertKkmModel", "deleteKkmModel"]),
    initView() {
      this.formState = null;
      this.resultText = '';

      if (this.action === "create") {
        this.localItem = {
          id: -1,
          name: "",
        };

        return;
      }

      this.localItem = this.$store.getters.kkmModelListItem(this.id);
    },
    async onComplete() {
      this.formState = 'pending';

      let fn;
      let messageTail;

      switch (this.action) {
        case "create":
          fn = () => this.upsertKkmModel({shouldCreate: true, name: this.localItem.name});
          messageTail = "создана"
          break;
        case "delete":
          fn = () => this.deleteKkmModel({id: this.id});
          messageTail = "удалена"
          break;
        case "update":
        default:
          fn = () => this.upsertKkmModel({shouldCreate: false, id: this.id, name: this.localItem.name})
          messageTail = "обновлена"
      }

      const result = await fn();

      if (result === null) {
        this.formState = 'error';
        this.resultText = 'Произошла ошибка. Попробуйте повторить позже';
        return;
      }

      this.formState = 'success';
      this.resultText = `Модель ККМ ID=${result.id} успешно ${messageTail}`;
    },
    onCancel() {
      this.$router.back();
    },
    getValidationClass(fieldName) {
      const field = this.$v.localItem[fieldName]

      if (field) {
        return {
          'md-invalid': field.$invalid && field.$dirty,
        }
      }
    },
    validateListItem() {
      if (this.action === 'view') {
        return;
      }

      this.$v.$touch()

      if (!this.$v.$invalid) {
        this.onComplete();
      }
    }
  },
};
</script>