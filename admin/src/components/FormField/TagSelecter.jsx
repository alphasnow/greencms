import {  Select,  } from 'antd'

// [{"id":16},{"id":36}] -> ["16","36"]
function toArrayId(arr){
    return arr.map(obj => obj.id.toString())
}
// ["16","36"] -> [{"id":16},{"id":36}]
function toObjectId(arr){
    return arr.map(item => ({id: Number(item)}))
}

export default function TagSelecter({ value, onChange, options }) {

    const op = options.map(obj => ({...obj, value: obj.value.toString()}));

    const change = (vals,items) => {
        const arr = toObjectId(vals)
        onChange(arr)
    }

    let def = []
    if(value){
        def = toArrayId(value)
    }

    return <Select
        mode="tags"
        onChange={change}
        options={op}
        defaultValue={def}
    />
}