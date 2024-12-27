import { create } from 'zustand'
import { getHomeData } from '@/lib/service'


export const useHomeStore = create((set, get) => ({
      data:null,
      fetchData:async ()=>{
        const res = await getHomeData()
        set({data:res})
      },
      isFetched:()=>{
        return get().data != null
      },
      clearData:()=>{
        set({ data: null })
      }
}))