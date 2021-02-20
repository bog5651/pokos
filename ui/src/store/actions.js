import g from "guark";

export default {
  async fetchClientList({ commit }) {
    try {
      commit("setIsClientListLoading", true);
      const clients = await g.call("getClients");

      commit("setClientList", clients);
      commit("setIsClientListLoading", false);

      return clients;
    } catch (e) {
      console.debug("fetchClientList", e);
      commit("setClientListError", e.message ? e.message : e);
      commit("setIsClientListLoading", false);

      return null;
    }
  },
  async fetchKkmList({ commit }) {
    try {
      commit("setIsKkmListLoading", true);
      const kkms = await g.call("getKKM");

      commit("setKkmList", kkms);
      commit("setIsKkmListLoading", false);

      return kkms;
    } catch (e) {
      console.debug("fetchKkmList", e);
      commit("setKkmListError", e.message ? e.message : e);
      commit("setIsKkmListLoading", false);

      return null;
    }
  },
  async fetchKkmModelsList({ commit }) {
    try {
      commit("setIsKkmModelListLoading", true);
      const kkmModels = await g.call("getModelsKKM");

      commit("setKkmModelList", kkmModels);
      commit("setIsKkmModelListLoading", false);

      return kkmModels;
    } catch (e) {
      console.debug("fetchKkmModelsList", e);
      commit("setKkmModelListError", e.message ? e.message : e);
      commit("setIsKkmModelListLoading", false);

      return null;
    }
  },
  async upsertClient({ commit }, { shouldCreate, id, name }) {
    try {
      commit("setIsClientListLoading", true);

      let client;

      if (shouldCreate) {
        client = await g.call("createClient", { name: name });
      } else {
        client = await g.call("updateClient", { client: { id: id, name: name } });
      }

      commit("upsertClientList", client);
      commit("setIsClientListLoading", false);

      return client;
    } catch (e) {
      console.debug("upsertClient", e);
      commit("setClientListError", e.message ? e.message : e);
      commit("setIsClientListLoading", false);

      return null;
    }
  },
  async deleteClient({ commit }, { id }) {
    try {
      commit("setIsClientListLoading", true);

      await g.call("deleteClient", { id: id });

      commit("removeClientListItem", id);
      commit("setIsClientListLoading", false);

      return { id: id };
    } catch (e) {
      console.debug("deleteClient", e);
      commit("setClientListError", e.message ? e.message : e);
      commit("setIsClientListLoading", false);

      return null;
    }
  },
  async upsertKkmModel({ commit }, { shouldCreate, id, name }) {
    try {
      commit("setIsKkmModelListLoading", true);

      let client;

      if (shouldCreate) {
        client = await g.call("createModelKKM", { name: name });
      } else {
        client = await g.call("updateModelKKM", { modelKkm: { id: id, name: name } });
      }

      commit("upsertKkmModelList", client);
      commit("setIsKkmModelListLoading", false);

      return client;
    } catch (e) {
      console.debug("upsertKkmModel", e);
      commit("setKkmModelListError", e.message ? e.message : e);
      commit("setIsKkmModelListLoading", false);

      return null;
    }
  },
  async deleteKkmModel({ commit }, { id }) {
    try {
      commit("setIsKkmModelListLoading", true);

      await g.call("deleteModelKKM", { id: id });

      commit("removeKkmModelListItem", id);
      commit("setIsClientListLoading", false);

      return { id: id };
    } catch (e) {
      console.debug("deleteKkmModel", e);
      commit("setKkmModelListError", e.message ? e.message : e);
      commit("setIsKkmModelListLoading", false);

      return null;
    }
  },
};
