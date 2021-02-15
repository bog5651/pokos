import FuzzySearch from "fuzzy-search";

export default {
  matchedClients(state) {
    const list = state.clientList || [];
    const searcher = new FuzzySearch(list, ["id", "name"], {
      caseSensitive: false,
      sort: true,
    });

    return (searchTerm) => {
      if (searchTerm === "") {
        return list;
      }

      const searchTermTrimmed = searchTerm.trim().toLowerCase();

      return searcher.search(searchTermTrimmed) || [];
    };
  },
  clientListItem(state) {
    const list = state.clientList || [];

    return (id) => {
      const found = list.find((c) => c.id === id);

      return found ? { ...found } : null;
    };
  },
  matchedKkmModels(state) {
    const list = state.kkmModelList || [];
    const searcher = new FuzzySearch(list, ["id", "name"], {
      caseSensitive: false,
      sort: true,
    });

    return (searchTerm) => {
      if (searchTerm === "") {
        return list;
      }

      const searchTermTrimmed = searchTerm.trim().toLowerCase();

      return searcher.search(searchTermTrimmed) || [];
    };
  },
  kkmModelListItem(state) {
    const list = state.kkmModelList || [];

    return (id) => {
      const found = list.find((c) => c.id === id);

      return found ? { ...found } : null;
    };
  },
  matchedKkms(state) {
    const list = state.kkmList || [];
    const searcher = new FuzzySearch(
      list.map((i) => ({ ...i, isExcise: i.isExcise ? "Да" : "Нет" })),
      [
        "id",
        "name",
        "clientName",
        "modelName",
        "serialNumber",
        "registerDate",
        "ofd",
        "isExcise",
        "systemNo",
        "type",
        "fn",
        "address",
        "endDateOfd",
        "endDateFn",
        "inspectionDayCount",
        "comment",
      ],
      {
        caseSensitive: false,
        sort: true,
      }
    );

    return (searchTerm) => {
      if (searchTerm === "") {
        return list;
      }

      const searchTermTrimmed = searchTerm.trim().toLowerCase();

      return searcher.search(searchTermTrimmed) || [];
    };
  },
  kkmListItem(state) {
    const list = state.kkmList || [];

    return (id) => {
      const found = list.find((c) => c.id === id);

      return found ? { ...found } : null;
    };
  },
};
