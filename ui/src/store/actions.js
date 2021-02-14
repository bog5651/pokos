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
        client = await g.call("updateClient", { id: id, name: name });
      }

      commit("upsertClientList", client);
      commit("setIsClientListLoading", false);

      return client.id;
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

      await new Promise((resolve) => {
        setInterval(() => resolve(1), 1500);
      });
      await g.call("deleteClient", { id: id });

      commit("removeClientListItem", id);
      commit("setIsClientListLoading", false);

      return true;
    } catch (e) {
      console.debug("deleteClient", e);
      commit("setClientListError", e.message ? e.message : e);
      commit("setIsClientListLoading", false);

      return null;
    }
  },
};
