export default {
  props: {
    action: {
      type: String,
      required: true,
    },
  },
  computed: {
    headerText() {
      switch (this.action) {
        case "view":
          return "Просмотр";
        case "create":
          return "Создание";
        case "delete":
          return "Удаление";
        case "update":
        default:
          return "Изменение";
      }
    },
    actionButtonText() {
      switch (this.action) {
        case "create":
          return "Создать";
        case "delete":
          return "Удалить";
        case "update":
        default:
          return "Сохранить";
      }
    },
  },
};
