import { Button, Form, Input, Row, Col, Card, Space, Select, message, Modal, App,InputNumber } from 'antd'
import { PageContainer, ProTable, ProCard } from '@ant-design/pro-components';
import ContentEditor from '@/components/FormField/ContentEditor';
import ImageUploader from '@/components/FormField/ImageUploader';
import { getTagsOptions, getCategoryOptions, postStore } from './service';
import { useEffect, useState } from 'react';
import { useAsyncEffect } from 'ahooks';
import TagSelecter from '@/components/FormField/TagSelecter';
import {  history } from '@umijs/max';

export default function CreatePage () {
    const { message, modal, notification } = App.useApp();
    const [categoryOptions, setCategoryOptions] = useState([])
    const [tagOptions, setTagOptions] = useState([])
    useEffect(() => {
        // https://devtrium.com/posts/async-functions-useeffect
        const fetchData = async () => {
            
            const res = await getCategoryOptions()
            setCategoryOptions(res.data)
            const res1 = await getTagsOptions()
            setTagOptions(res1.data)
        }
        fetchData()
    }, [])



    const initValues = {
        // author: "",
        // content: "<p>hello</p>",
        // description: "",
        // image: "http://127.0.0.1:8080/storage/uploads/admin/7657616792204731.png",
        // keywords: "",
        // origin_author: "",
        // origin_url:  "",
        // title: "hi",
    }

    const onFinish = async (values) => {

        const data = {...values}
        data["article_content"] = { "content": data["article_content"] }
        data["category_id"] = parseInt(data["category_id"])

        await postStore(data)
        // message.success("添加成功")

        // 刷新界面
        modal.success({
            title: '添加成功',
            afterClose: () => {
                history.back()
            }
        })
    }



    return (<PageContainer
        title="添加文章"
        extra={[
            <Button key="back" onClick={() => history.back()}>返回</Button>,
        ]}
    >
        {/* https://ant.design/components/form-cn#components-form-demo-layout */}
        <Form
            layout="vertical"
            initialValues={initValues}
            onFinish={onFinish}
            //onFinishFailed={onFinishFailed}
            autoComplete="off"
        >
            <Row gutter={[16, 16]}>
                <Col sm={{ span: 24 }} lg={{ span: 18 }}>
                    <Card
                        //title="正文" 
                        bordered={false}>
                        <Form.Item
                            label="标题"
                            name="title"
                            rules={[
                                {
                                    required: true,
                                    message: '请填写标题',
                                },
                            ]}
                        >
                            <Input placeholder='请填写标题' />
                        </Form.Item>
                        <Form.Item
                            label="正文"
                            name="article_content"
                            rules={[
                                {
                                    required: true,
                                    message: '请填写正文',
                                },
                            ]}
                        >
                            <ContentEditor path="article-editor" />
                        </Form.Item>
                    </Card>

                </Col>
                <Col sm={{ span: 24 }} lg={{ span: 6 }}>
                    <Card
                        //title="附属" 
                        bordered={false}>

                        <Form.Item
                            label="封面"
                            name="image_url"
                            rules={[
                                {
                                    required: true,
                                    message: '请上传封面图片',
                                },
                            ]}
                        >
                            <ImageUploader path="article-image" />
                        </Form.Item>
                        <Form.Item
                            label="分类"
                            name="category_id"
                            rules={[
                                {
                                    required: true,
                                    message: '请选择分类',
                                },
                            ]}
                        >
                            <Select
                                allowClear
                                options={categoryOptions}
                            />
                        </Form.Item>
                        <Form.Item
                            label="标签"
                            name="article_tags"
                        >
                            <TagSelecter options={tagOptions} />
                        </Form.Item>
                        <Form.Item
                            label="关键词"
                            name="keywords"
                        >
                            <Input />
                        </Form.Item>
                        <Form.Item
                            label="描述"
                            name="description"
                        >
                            <Input.TextArea />
                        </Form.Item>
                        <Form.Item
                            label="来源作者"
                            name="origin_author"
                        >
                            <Input />
                        </Form.Item>
                        <Form.Item
                            label="来源网址"
                            name="origin_url"
                        >
                            <Input />
                        </Form.Item>
                        <Form.Item
                            label="排序"
                            name="sort"
                        >
                            <InputNumber style={{ width: '100%' }} />
                        </Form.Item>
                    </Card>
                </Col>
                <Col span={24}>
                    <Card
                        bordered={false}
                        style={{ "textAlign": "right" }}
                    >
                        <Space>
                            <Button type="primary" htmlType="submit">
                                提交
                            </Button>
                            <Button htmlType="reset">
                                重置
                            </Button>
                        </Space>
                    </Card>
                </Col>
            </Row>
        </Form>
    </PageContainer>)
}