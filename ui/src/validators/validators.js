export const isValidDate = (v) => {
  try {
    if (typeof v !== "string") {
      return false;
    }

    return v.match(/\d{2}\/\d{2}\/\d{4}/) || false;
  } catch {
    return false;
  }
};
