import { PageContainer } from '@ant-design/pro-components';
import { Button, Form, Input, Row, Col, Card, Space, Select, message, Modal, App, InputNumber } from 'antd'
import { useModel } from 'umi';
import { Spin } from 'antd';
import ImageUploader from '@/components/FormField/ImageUploader';
import { request } from '@umijs/max';

export default function SettingPage() {
    const { initialState, loading, refresh, setInitialState } = useModel('@@initialState');
    const { currentUser } = initialState

    if (loading) {
        return <Spin />;
    }

    const formData = { ...currentUser, realname: currentUser.name, avatar_url: currentUser.avatar }
    const formSubmit = async (submitData) => {
        if (submitData.password != submitData.password_confirmed) {
            message.error("修改密码与确认密码不同")
            return
        }
        delete submitData.password_confirmed
        delete submitData.username

        await request('/api/admin/account/settings', {
            method: 'POST',
            data: submitData,
        });
        message.success("设置成功")
        refresh()

    }
    return (<PageContainer>
        <Card >
            <Form
                layout="vertical"
                initialValues={formData}
                onFinish={formSubmit}
                autoComplete="off"
            >
                <Form.Item
                    label="账号"
                    name="username"
                >
                    <Input disabled />
                </Form.Item>
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
                <Form.Item
                    label="姓名"
                    name="realname"
                    rules={[{ required: true }]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="修改密码"
                    name="password"
                >
                    <Input.Password placeholder='留空则不修改' />
                </Form.Item>
                <Form.Item
                    label="确认密码"
                    name="password_confirmed"
                >
                    <Input.Password placeholder='留空则不修改' />
                </Form.Item>

                <Space style={{ "textAlign": "right" }}>
                    <Button type="primary" htmlType="submit">
                        提交
                    </Button>
                    <Button htmlType="reset">
                        重置
                    </Button>
                </Space>
            </Form>
        </Card>

    </PageContainer>)
}