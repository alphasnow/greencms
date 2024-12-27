import { useEffect, useState } from "react"
import { useGlobalStore } from "./global"

export * from "./client"
export * from "./persist"
export * from "./home"
export * from "./global"
export * from "./active"

export function useStoreEffect(effect,deps=[]){
    const store = useGlobalStore()
    const [loaded,setLoaded] = useState(false)
    useEffect(()=>{
        const fetchData = async () => {
            if(store.isFetched() == false){
                await store.fetchData()
            }
            const clear = await effect()
            setLoaded(true)
            return clear
        }
        fetchData()
    },deps)
    return {loaded,store}
}