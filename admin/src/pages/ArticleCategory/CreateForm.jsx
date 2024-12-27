import { ModalForm, ProFormDigit, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import { useModel } from 'umi';

export default function CreateForm(props) {
  const model = useModel('ArticleCategory.model');

  return (
    <ModalForm
      title="添加"
      open={model.isShowCreate}
      onOpenChange={model.showCreate}
      onFinish={model.submitCreate}
      modalProps={{ destroyOnClose: true }}
      initialValues={{}}
    >
      <ProFormText
        name="title"
        label="名称"
        rules={[{ required: true }]}
        help="分类标题,长度需要在10字以下"
      />
      <ProFormText name="keywords" label="关键字" help="多个关键字使用逗号隔开" />
      <ProFormTextArea name="description" label="描述" help="简短描述分类内容" />
      <ProFormDigit name="sort" label="排序" help="数值越小越靠前" />
    </ModalForm>
  );
}
