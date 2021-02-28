<template>
  <div class="md-layout">
    <div class="md-layout-item">
      <form novalidate class="md-layout md-alignment-top-center" @submit.prevent="validateListItem">
        <md-card class="md-layout-item md-size-50 md-small-size-100">
          <md-card-header>
            <div class="md-title">{{ headerText }} записи о ККМ</div>
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
                <md-field :class="getValidationClass('clientId')">
                  <label for="client">Клиент</label>
                  <md-select name="client" id="client" v-model="localItem.clientId" :disabled="isPending"
                             :readonly="['view', 'delete'].includes(action)" required>
                    <!--                    <md-option value="" disabled class="md-disabled">Выберите клиента</md-option>-->
                    <md-option v-for="opt in clientList" :key="opt.id" :value="opt.id">{{ opt.name }}</md-option>
                  </md-select>
                  <span class="md-error" v-if="!$v.localItem.clientId.required">Клиент должен быть выбран</span>
                </md-field>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-field :class="getValidationClass('modelId')">
                  <label for="model">Модель</label>
                  <md-select name="model" id="model" v-model="localItem.modelId" :disabled="isPending"
                             :readonly="['view', 'delete'].includes(action)">
                    <!--                    <md-option value="" disabled class="md-disabled">Выберите модель</md-option>-->
                    <md-option v-for="opt in kkmModelList" :key="opt.id" :value="opt.id">{{ opt.name }}</md-option>
                  </md-select>
                  <span class="md-error" v-if="!$v.localItem.modelId.required">Модель должна быть выбрана</span>
                </md-field>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-field :class="getValidationClass('serialNumber')">
                  <label for="serialNumber">Серийный номер</label>
                  <md-input name="serialNumber" id="serialNumber" autocomplete="off" v-model="localItem.serialNumber"
                            :disabled="isPending"
                            :readonly="['view', 'delete'].includes(action)"
                            required
                  />
                  <span class="md-error"
                        v-if="!$v.localItem.serialNumber.required">Серийный номер должен быть заполнен</span>
                </md-field>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-datepicker v-model="localItem.registerDate" md-immediately md-open-on-focus
                               required class="md-required">
                  <label>Дата регистрации</label>
                </md-datepicker>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-field :class="getValidationClass('ofd')">
                  <label for="ofd">ОФД</label>
                  <md-select name="ofd" id="ofd" v-model="localItem.ofd" :disabled="isPending"
                             :readonly="['view', 'delete'].includes(action)" required>
                    <!--                    <md-option value="" disabled class="md-disabled">Выберите ОФД</md-option>-->
                    <md-option v-for="opt in ofdList" :key="opt" :value="opt">{{ opt }}</md-option>
                  </md-select>
                  <span class="md-error" v-if="!$v.localItem.ofd.required">ОФД должен быть выбран</span>
                </md-field>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-switch v-model="localItem.isExcise" class="md-primary">Акцизный</md-switch>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-field :class="getValidationClass('systemNo')">
                  <label for="systemNo">Система НО</label>
                  <md-select name="systemNo" id="systemNo" v-model="localItem.systemNo" :disabled="isPending"
                             :readonly="['view', 'delete'].includes(action)" required>
                    <!--                    <md-option value="" disabled class="md-disabled">Система НО</md-option>-->
                    <md-option v-for="opt in noSystemList" :key="opt" :value="opt">{{ opt }}</md-option>
                  </md-select>
                  <span class="md-error" v-if="!$v.localItem.systemNo.required">Система НО должна быть выбрана</span>
                </md-field>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-field :class="getValidationClass('type')">
                  <label for="type">Тип НО</label>
                  <md-select name="type" id="type" v-model="localItem.type" :disabled="isPending"
                             :readonly="['view', 'delete'].includes(action)" required>
                    <!--                    <md-option value="" disabled class="md-disabled">Тип НО</md-option>-->
                    <md-option v-for="opt in noTypeList" :key="opt" :value="opt">{{ opt }}</md-option>
                  </md-select>
                  <span class="md-error" v-if="!$v.localItem.type.required">Тип НО должен быть выбран</span>
                </md-field>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-field :class="getValidationClass('fn')">
                  <label for="fn">ФН</label>
                  <md-input name="fn" id="fn" autocomplete="off" v-model="localItem.fn"
                            :disabled="isPending"
                            :readonly="['view', 'delete'].includes(action)"
                            required
                            type="number"
                  />
                  <span class="md-error" v-if="!$v.localItem.fn.required">ФН должен быть заполнен</span>
                </md-field>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-field :class="getValidationClass('address')">
                  <label for="address">Адрес</label>
                  <md-input name="address" id="address" autocomplete="off" v-model="localItem.address"
                            :disabled="isPending"
                            :readonly="['view', 'delete'].includes(action)"
                            required
                  />
                  <span class="md-error" v-if="!$v.localItem.address.required">Адрес должен быть заполнен</span>
                </md-field>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-datepicker v-model="localItem.endDateOfd" md-immediately md-open-on-focus
                               required class="md-required">
                  <label>Дата окончания ОФД</label>
                </md-datepicker>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-datepicker v-model="localItem.endDateFn" md-immediately md-open-on-focus
                               required class="md-required">
                  <label>Дата окончания ФН</label>
                </md-datepicker>
              </div>

              <div class="md-layout-item md-small-size-100" v-if="action === 'action'">
                <md-field :class="getValidationClass('inspectionDayCount')">
                  <label for="inspectionDayCount">Дней с момента ТО</label>
                  <md-input name="inspectionDayCount" id="inspectionDayCount" autocomplete="off"
                            v-model="localItem.inspectionDayCount"
                            :disabled="isPending"
                            :readonly="['view', 'delete'].includes(action)"
                            type="number"
                  />
                </md-field>
              </div>

              <div class="md-layout-item">
                <md-field :class="getValidationClass('comment')">
                  <label for="comment">Комментарий</label>
                  <md-textarea md-autogrow name="comment" id="comment" autocomplete="off" v-model="localItem.comment"
                               :disabled="isPending"
                               :readonly="['view', 'delete'].includes(action)"
                  />
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
import {minValue, required} from 'vuelidate/lib/validators';
import {isValidDate} from "@/validators/validators";
import {validationMixin} from "vuelidate";
import {format} from 'date-fns';

export default {
  name: "KkmListItem",
  mixins: [formTextMixin, validationMixin],
  data() {
    const dateFormat = this.$material.locale.dateFormat || 'yyyy-MM-dd';
    const templateDate = format(new Date(), dateFormat);

    return {
      localItem: {
        id: -1,
        clientId: -1,
        modelId: -1,
        serialNumber: '',
        registerDate: templateDate,
        ofd: '',
        isExcise: false,
        systemNo: '',
        type: '',
        fn: 0,
        address: '',
        endDateOfd: templateDate,
        endDateFn: templateDate,
        inspectionDayCount: 0,
        comment: '',
      },
      formState: null,
      resultText: '',
    }
  },
  validations: {
    localItem: {
      id: { required },
      clientId: { required: minValue(0) },
      modelId: { required: minValue(0) },
      serialNumber: { required },
      registerDate: { required, isValidDate },
      ofd: { required },
      isExcise: {}, // optional
      systemNo: { required },
      type: { required },
      fn: { required: minValue(0) },
      address: { required },
      endDateOfd: { required, isValidDate },
      endDateFn: { required, isValidDate },
      inspectionDayCount: {}, // calculated
      comment: {}, // optional
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
    ...mapState(["isKkmListLoading", "noSystemList", "noTypeList", "clientList", "kkmModelList", "ofdList"]),
    isPending() {
      return this.isKkmListLoading || this.formState === "pending";
    }
  },
  watch: {
    $route() {
      this.initView();
    },
    localItem(newVal, oldVal) {
      if (!oldVal) return;

      if (oldVal.clientId === newVal.clientId) {
        return;
      }

      console.log('changed:', newVal);
    }
  },
  created() {
    this.initView();
  },
  methods: {
    ...mapActions(["upsertKkm", "deleteKkm"]),
    initView() {
      this.formState = null;
      this.resultText = '';

      if (this.action === "create") {
        const dateFormat = this.$material.locale.dateFormat || 'yyyy-MM-dd';
        const templateDate = format(new Date(), dateFormat);

        this.localItem = {
          id: -1,
          clientId: -1,
          modelId: -1,
          serialNumber: '',
          registerDate: templateDate,
          ofd: '',
          isExcise: false,
          systemNo: '',
          type: '',
          fn: 0,
          address: '',
          endDateOfd: templateDate,
          endDateFn: templateDate,
          inspectionDayCount: 0,
          comment: '',
        };

        return;
      }

      this.localItem = this.$store.getters.kkmListItem(this.id);
    },
    async onComplete() {
      this.formState = 'pending';

      let fn;
      let messageTail;

      switch (this.action) {
        case "create":
          fn = () => this.upsertKkm({ shouldCreate: true, kkm: this.localItem });
          messageTail = "создана"
          break;
        case "delete":
          fn = () => this.deleteKkm({ id: this.id });
          messageTail = "удалена"
          break;
        case "update":
        default:
          fn = () => this.upsertKkm({ shouldCreate: false, id: this.id, kkm: this.localItem })
          messageTail = "обновлена"
      }

      const result = await fn();

      if (result === null) {
        this.formState = 'error';
        this.resultText = 'Произошла ошибка. Попробуйте повторить позже';
        return;
      }

      this.formState = 'success';
      this.resultText = `Модель ККМ ID=${ result.id } успешно ${ messageTail }`;
    },
    onCancel() {
      this.$router.back();
    },
    getValidationClass(fieldName) {
      const field = this.$v.localItem[fieldName];

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