import { create } from 'zustand'
import { persist, createJSONStorage } from 'zustand/middleware'
import { v4 as uuidv4 } from 'uuid';

const DEFAULT_CLIENT_STATE = {
  clientUUID: uuidv4(),
  updatedAt:(new Date).getTime(),
};

export const useClientStore = create(persist(
    (set, get) => ({
      ...DEFAULT_CLIENT_STATE,
      setClientUUID: (uuid) => set({ uuid: uuid }),
      setClientIP: (ip) => set({ ip: ip }),
    }),
    {
      name: 'client-store',
      storage: createJSONStorage(() => localStorage),
    },
  ),
)