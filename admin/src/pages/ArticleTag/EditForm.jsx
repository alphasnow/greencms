import { ModalForm, ProFormText, ProFormColorPicker } from '@ant-design/pro-components';
import { useModel } from 'umi';
import { englishLettersValidater } from '@/utils/form-tool';

export default function EditForm(props) {
  const model = useModel('ArticleTag.model');

  return (
    <ModalForm
      title={`修改`}
      open={model.isShowEdit}
      onOpenChange={model.showEdit}
      onFinish={model.submitEdit}
      modalProps={{ destroyOnClose: true }}
      initialValues={model.data}
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
