import ImageUploader from '@/components/FormField/ImageUploader';
import { Col, Form, Row } from 'antd';
import { ModalForm, ProFormDigit, ProFormText, ProFormTextArea,ProFormSelect } from '@ant-design/pro-components';
import { useModel } from 'umi';

export default function CreateForm(props) {
  const model = useModel('WebBanner.model');

  return (
    <ModalForm
      title="添加"
      open={model.isShowCreate}
      onOpenChange={model.showCreate}
      onFinish={model.submitCreate}
      modalProps={{ destroyOnClose: true }}
      initialValues={{ redirect_url: '/' }}
    >
      <Row gutter={16}>
        <Col span={24}>
          <Form.Item
            label="焦点图"
            name="image_url"
            rules={[
              {
                required: true,
                message: '请上传焦点图',
              },
            ]}
            help="请上传 (jpg/png/webp) 类型的图片"
          >
            <ImageUploader path="web-banner" />
          </Form.Item>
        </Col>
        <Col span={12}>
          <ProFormText name="title" label="标题" help="" />
          <ProFormSelect
        request={async ()=>(await model.getOptions()).data.banner_group}
        name="banner_group"
        label="分组"
        rules={[{ required: true }]}
      />
          <ProFormTextArea name="description" label="描述" help="" />
        </Col>
        <Col span={12}>
          <ProFormText name="redirect_url" label="链接" help="点击跳转链接" />
          <ProFormDigit name="sort" label="排序" help="数值越小越靠前" />
          <ProFormTextArea name="remark" label="备注" help="" />
        </Col>
      </Row>
    </ModalForm>
  );
}
