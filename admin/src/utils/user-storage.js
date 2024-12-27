
const CLIENT_UUID = "client_uuid"
const ACCESS_TOKEN = "access_token"

const stroage = (key) => {
    const get  = () => {
        return localStorage.getItem(key)
    }
    const set = (value) => {
        return localStorage.setItem(key,value)
    }
    const remove = () => {
        localStorage.removeItem(key)
    }
    return {get,set,remove}
}

export const uuidStorage = stroage(CLIENT_UUID)
export const tokenStorage = stroage(ACCESS_TOKEN)