import config from '@/lib/config'

export async function getHomeData(){
    const resp = await fetch(config.apiUrl+"/api/web/home")
    // if (!resp.ok) {
    //     throw new Error('Failed to fetch data')
    //   }
    const res = await resp.json()
    return res
}

export async function getStateData(){
    const resp = await fetch(config.apiUrl+"/api/web/state")
    const res = await resp.json()
    return res
}

export async function getArticleList(search){
    const resp = await fetch(config.apiUrl+`/api/web/article?${search}`)
    const res = await resp.json()
    return res
}

export async function getArticleData(id){
    const resp = await fetch(config.apiUrl+`/api/web/article/${id}`)
    const res = await resp.json()
    return res
}

export async function getCategoryData(id){
    const resp = await fetch(config.apiUrl+`/api/web/category/${id}`)
    const res = await resp.json()
    return res
}
export async function getTagData(id){
    const resp = await fetch(config.apiUrl+`/api/web/tag/${id}`)
    const res = await resp.json()
    return res
}

export async function getStatistic(id,type="views"){
    const resp = await fetch(config.apiUrl+`/api/web/statistic/${id}/${type}`,{
        method:"PUT"
    })
    const res = await resp.json()
    return res
}