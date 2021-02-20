export default {
  setClientList(state, list) {
    state.clientList = list;
  },
  upsertClientList(state, item) {
    const idx = state.clientList.findIndex((c) => c.id === item.id);

    if (idx > -1) {
      state.clientList[idx] = item;
      return;
    }

    state.clientList = [...state.clientList, item];
  },
  removeClientListItem(state, itemId) {
    const idx = state.clientList.findIndex((i) => i.id === itemId);

    if (idx > -1) {
      state.clientList.splice(idx, 1);
    }
  },
  setKkmList(state, list) {
    state.kkmList = list;
  },
  upsertKkmList(state, item) {
    const idx = state.kkmList.findIndex((k) => k.id === item.id);

    if (idx > -1) {
      state.kkmList[idx] = item;
      return;
    }

    state.kkmList = [...state.kkmList, item];
  },
  removeKkmListItem(state, itemId) {
    const idx = state.kkmList.findIndex((i) => i.id === itemId);

    if (idx > -1) {
      state.kkmList.splice(idx, 1);
    }
  },
  setKkmModelList(state, list) {
    state.kkmModelList = list;
  },
  upsertKkmModelList(state, item) {
    const idx = state.kkmModelList.findIndex((m) => m.id === item.id);

    if (idx > -1) {
      state.kkmModelList[idx] = item;
      return;
    }

    state.kkmModelList = [...state.kkmModelList, item];
  },
  removeKkmModelListItem(state, itemId) {
    const idx = state.kkmModelList.findIndex((i) => i.id === itemId);

    if (idx > -1) {
      state.kkmModelList.splice(idx, 1);
    }
  },
  setIsClientListLoading(state, isLoading) {
    state.isClientListLoading = isLoading;
  },
  setIsKkmListLoading(state, isLoading) {
    state.isKkmListLoading = isLoading;
  },
  setIsKkmModelListLoading(state, isLoading) {
    state.isKkmModelListLoading = isLoading;
  },
  setClientListError(state, error) {
    state.clientListError = error;
  },
  setKkmListError(state, error) {
    state.kkmListError = error;
  },
  setKkmModelListError(state, error) {
    state.kkmModelListError = error;
  },
};
