import { message } from 'antd';
import { useRef, useState } from 'react';
import { getList, postDelete, postStore, postUpdate } from './service';
import { filterEmpty,colorToHex } from '@/utils/form-tool';

export default function Model() {
  // 这里面发生错误 不会直接报错 而是会引起使用model获得undefined
  const [data, setData] = useState(undefined);
  const [isShowEdit, showEdit] = useState(false);
  const [isShowCreate, showCreate] = useState(false);
  const tableRef = useRef();

  // 此处会引发错误
  // const {message} = useApp()

  const tableRequest = async (params, sort, filter) => {
    //console.log(params, sort, filter);
    let data = params
    // sort example: {updated_at: 'ascend'} {updated_at: 'descend'}
    if(sort){
      data = {...data,"sort":sort}
    }
    if(filter){
      data = {...data,"filter":filter}
    }
    return getList(data);
  };
  const tableDelete = async (record) => {
    await postDelete(record.id);
    message.success('删除成功');
    tableReload();
  };
  const tableReload = () => {
    tableRef.current.reload();
  };

  const submitCreate = async (formData) => {
    formData.color = colorToHex(formData.color)

    await postStore(formData);
    message.success('添加成功');

    showCreate(false);
    tableReload();
  };
  const submitEdit = async (formData) => {
    formData.color = colorToHex(formData.color)
    
    await postUpdate(data.id, formData);
    message.success('修改成功');

    showEdit(false);
    setData(undefined);
    tableReload();
  };

  return {
    tableRef,
    tableRequest,
    tableDelete,
    data,
    setData,
    isShowEdit,
    showEdit,
    submitEdit,
    isShowCreate,
    showCreate,
    submitCreate,
  };
}
