// import { ColorFactory } from 'antd/es/color-picker/color';

export function filterEmpty(data){
    return Object.entries(data).reduce((acc, [key, value]) => {
        if (value !== undefined && value !== null) {
          acc[key] = value;
        }
        return acc;
      }, {});
}

export function colorToHex(color){
  // if(color instanceof ColorFactory){
  //   return color.toHexString()
  // }
  return color
}

export const englishLettersValidater = (_, value) => {
  const regex = /^[A-Za-z0-9-]+$/;
  if (value && !regex.test(value)) {
    return Promise.reject(new Error('只能输入英语字母'));
  }
  return Promise.resolve();
};