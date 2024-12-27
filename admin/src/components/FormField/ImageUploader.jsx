
import React, { useState } from 'react';
import { LoadingOutlined, PlusOutlined } from '@ant-design/icons';
import { Form, message, Upload } from 'antd';
import { tokenStorage } from '@/utils/user-storage';

const getBase64 = (img, callback) => {
  const reader = new FileReader();
  reader.addEventListener('load', () => callback(reader.result));
  reader.readAsDataURL(img);
};
const beforeUpload = (file) => {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'|| file.type === 'image/webp';;
  if (!isJpgOrPng) {
    message.error('You can only upload JPG/PNG/WEBP file!');
  }
  const isLt2M = file.size / 1024 / 1024 < 2;
  if (!isLt2M) {
    message.error('Image must smaller than 2MB!');
  }
  return isJpgOrPng && isLt2M;
};

export default function ImageUploader({ value, onChange, path="admin",privacy="public" }) {
  // https://ant-design.antgroup.com/components/form-cn#formitemusestatus
  const [loading, setLoading] = useState(false);
  // const [imageUrl, setImageUrl] = useState(value);
  const { status, errors } = Form.Item.useStatus();

  // console.log(status,errors)
  const token = tokenStorage.get()
  const headers = { "Authorization": "Bearer " + token }
  const apiUrl = API_URL + "/api/admin/upload/form-file"

  // TODO: 自定义表单插件
  const handleChange = (info) => {
    // https://ant-design.antgroup.com/components/upload-cn#uploadfile
    // 'done' | 'uploading' | 'error' | 'removed'
    // console.log(info.file.status)
    if (info.file.status === 'uploading') {
      setLoading(true);
      return;
    }
    if (info.file.status === 'removed') {
      onChange('')
      return;
    }
    if (info.file.status === 'done') {
      setLoading(false);
      if (!info.file.response.success) {
        message.error(info.file.response.message)
        return
      }
      onChange(info.file.response.data.url)
      // // Get this url from response in real world.
      // getBase64(info.file.originFileObj, (url) => {
      //   setLoading(false);
      //   setImageUrl(url);
      // });
    }
  };

  const uploadButton = (
    <div>
      {loading ? <LoadingOutlined /> : <PlusOutlined />}
      <div
        style={{
          marginTop: 8,
        }}
      >
        上传
      </div>
    </div>
  );

  const fileList = value ? [{status: 'done',url: value}] : []

  return (
    <Upload
    defaultFileList={fileList}
      listType="picture-card"
      className="avatar-uploader"
      // showUploadList={false}
      maxCount={1}
      action={apiUrl}
      beforeUpload={beforeUpload}
      onChange={handleChange}
      headers={headers}
      data={{ "path": path,"privacy":privacy }}
    >
      {!value ? uploadButton : null}
    </Upload>
  );
}