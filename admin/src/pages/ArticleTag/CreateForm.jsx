import { ModalForm, ProFormText, ProFormColorPicker } from '@ant-design/pro-components';
import { useModel } from 'umi';
import { englishLettersValidater } from '@/utils/form-tool';

export default function CreateForm(props) {
  const model = useModel('ArticleTag.model');

  return (
    <ModalForm
      title="添加"
      open={model.isShowCreate}
      onOpenChange={model.showCreate}
      onFinish={model.submitCreate}
      modalProps={{ destroyOnClose: true }}
      initialValues={{color:"#64748b"}}
    >
      <ProFormText
        name="name"
        label="名称"
        rules={[{ required: true }]}
        help="名称,长度需要在10字以下"
      />
      <ProFormText
                name="slug"
                label="标识"
                rules={[{ required: true },{ validator: englishLettersValidater }]}
                help="仅限输入 字母/数字/-,示例: hot-read-1"
                fieldProps={{allowClear:true,}}
            />
            <ProFormColorPicker
                name="color"
                label="颜色"
                help="可选"
            />
    </ModalForm>
  );
}
