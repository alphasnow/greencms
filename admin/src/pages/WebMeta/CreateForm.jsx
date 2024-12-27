import { ModalForm, ProFormSelect, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import { useModel } from 'umi';

export default function CreateForm(props) {
  const model = useModel('WebMeta.model');

  return (
    <ModalForm
      title="添加"
      open={model.isShowCreate}
      onOpenChange={model.showCreate}
      onFinish={model.submitCreate}
      modalProps={{ destroyOnClose: true }}
      initialValues={{}}
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
