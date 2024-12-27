export function limitString(str, len = 100) {
    if (str.length > len) {
        return str.slice(0, len) + '...';
    }
    return str;
}


export function isoDateToDateTime(isoDateString) {
    // 创建一个新的 Date 对象
    const date = new Date(isoDateString);

    // 获取各个时间组件的值
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份是从 0 开始的，所以要加 1
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');
    const formattedDate = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    return formattedDate
}