import { create } from 'zustand'


const DEFAULT_ACTIVE_STATE = {
  category: "",
};

export const useActiveStore = create((set, get) => ({
  ...DEFAULT_ACTIVE_STATE,
  setCategoryID: (id) => {
    set({ category: `/category/?id=${id}` })
  },
  setCategory: (cat) => {
    set({ category: cat })
  },
}));
