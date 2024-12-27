import ImageUploader from '@/components/FormField/ImageUploader';
import { Col, Form, Row } from 'antd';
import { ModalForm, ProFormDigit, ProFormText, ProFormSelect } from '@ant-design/pro-components';
import { useModel } from 'umi';
// import { accessOptions } from '@/access';

export default function CreateForm(props) {
  const model = useModel('AdminUser.model');

  return (
    <ModalForm
      title="添加"
      open={model.isShowCreate}
      onOpenChange={model.showCreate}
      onFinish={model.submitCreate}
      modalProps={{ destroyOnClose: true }}
      initialValues={{}}
    >
      <Form.Item
        label="头像"
        name="avatar_url"
        rules={[
          {
            required: true,
            message: '请上传头像',
          },
        ]}
        help="请上传 (jpg/png/webp) 类型的图片"
      >
        <ImageUploader path="admin-avatar" />
      </Form.Item>

      <ProFormText
        name="username"
        label="账号"
        rules={[{ required: true }]}
      />
      <ProFormText
        name="realname"
        label="姓名"
        rules={[{ required: true }]}
      />
      <ProFormText.Password
        name="password"
        label="密码"
        rules={[{ required: true }]}
      />
      <ProFormText.Password
        name="password_confirmed"
        label="确认密码"
        rules={[{ required: true }]}
      />
      <ProFormSelect
        request={async ()=>(await model.getOptions()).data.access}
        name="access"
        label="权限"
        rules={[{ required: true }]}
      />
    </ModalForm>
  );
}
