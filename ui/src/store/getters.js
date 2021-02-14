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
};
