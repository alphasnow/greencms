import { create } from 'zustand'
import { persist, createJSONStorage } from 'zustand/middleware'
import { getStateData } from '@/lib/service'

const DEFAULT_GLOBAL_STATE = {
    metas: {
        "title":"",
        "description":"",
        "keywords":"",
        "record_no":"",
        "logo":"",
        "analytics":"",
    },
    categories:[],
};

export const useGlobalStore = create(persist(
    (set, get) => ({
      ...DEFAULT_GLOBAL_STATE,
      fetchData:async ()=>{
        const res = await getStateData()
        set({metas:res.metas,categories:res.categories})
      },
      isFetched:()=>{
        return get().categories.length != 0
      },
      clearData:()=>{
        set({ data: {...DEFAULT_GLOBAL_STATE} })
      }
    }),
    {
      name: 'global-store',
      storage: createJSONStorage(() => sessionStorage),
    },
  ),
)