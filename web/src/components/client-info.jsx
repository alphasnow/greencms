'use client'

import { v4 as uuidv4 } from 'uuid';
import {useClientStore,usePersistStore} from "@/stores"

export default function ClientInfo(){
    const store  = usePersistStore(useClientStore)
    if(store==undefined) return null
    
    const resetUUID = () => {
        store.setClientUUID(uuidv4())
    }

    return (<>
    <button onClick={resetUUID}>reset uuid</button>
    <p>{store.clientUUID}</p>
    <p>{store.updatedAt}</p>
    </>)
}