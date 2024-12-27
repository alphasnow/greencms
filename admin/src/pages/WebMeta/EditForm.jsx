import { ModalForm, ProFormSelect, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import { useModel } from 'umi';
import { getOptions } from './service';

export default function EditForm(props) {
  const model = useModel('WebMeta.model');

  return (
    <ModalForm
      title={`修改`}
      open={model.isShowEdit}
      onOpenChange={model.showEdit}
      onFinish={model.submitEdit}
      modalProps={{ destroyOnClose: true }}
      initialValues={model.data}
    >
      <ProFormText name="meta_name" label="名称" rules={[{ required: true }]} help="" />
      <ProFormText name="meta_key" label="数据键" rules={[{ required: true }]} help="" />
      <ProFormTextArea name="meta_value" label="数据值" rules={[{ required: true }]} help="" />
      <ProFormSelect
        request={async ()=>(await model.getOptions()).data.meta_group}
        name="meta_group"
        label="数据分组"
        rules={[{ required: true }]}
      />
    </ModalForm>
  );
}
