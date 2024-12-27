import { PlusOutlined } from '@ant-design/icons';
import { PageContainer, ProTable } from '@ant-design/pro-components';
import { Button, Popconfirm, Space } from 'antd';
import { useModel } from 'umi';
import {  history } from '@umijs/max';

export default function IndexPage() {
  const model = useModel('Article.model');

  const tableToolBarRender = () => {
    return (
      <Button
        onClick={() => history.push('/article/create')}
        icon={<PlusOutlined />}
        type="primary"
      >
        添加
      </Button>
    );
  };
  const tableColumns = [
    {
      title: '序号',
      dataIndex: 'id',
      hideInSearch: true,
      fixed: 'left',
      width: 50,
  },
  {
      title: '封面',
      dataIndex: 'image_url',
      hideInSearch: true,
      valueType: 'image',
  },
  {
      title: '标题',
      dataIndex: 'title',
      copyable: true,
      ellipsis: true,
  },
  {
    title: '分类',
    dataIndex: 'category_name',
},
  {
      title: '标签',
      dataIndex: 'tag_names',
      ellipsis: true,
      hideInSearch: true,
  },
  {
    title: '排序',
    dataIndex: 'sort',
    hideInSearch: true,
    sorter: true,
  },
  {
      title: '修改时间',
      dataIndex: 'updated_at',
      hideInSearch: true,
      valueType: 'dateTime',
      width: 200,
  },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      width: 150,
      render: (_, record) => (
        <Space.Compact size="small">
          <Button
            key="edit"
            onClick={() => history.push('/article/edit/'+record.id)}
          >
            修改
          </Button>
          <Popconfirm
            key="destroy"
            title={`确定删除 ${record['title']} 吗?`}
            okText="确定"
            cancelText="取消"
            onConfirm={() => model.tableDelete(record)}
          >
            <Button>删除</Button>
          </Popconfirm>
        </Space.Compact>
      ),
    },
  ];

  return (
    <PageContainer>
      <ProTable
        rowKey="id"
        actionRef={model.tableRef}
        columns={tableColumns}
        request={model.tableRequest}
        pagination={{ defaultPageSize: 10, showSizeChanger: true }}
        toolBarRender={tableToolBarRender}
        options={true}
      />
    </PageContainer>
  );
}
