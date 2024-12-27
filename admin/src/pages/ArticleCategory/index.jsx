import { PlusOutlined } from '@ant-design/icons';
import { PageContainer, ProTable } from '@ant-design/pro-components';
import { Button, Popconfirm, Space } from 'antd';
import { useModel } from 'umi';
import CreateForm from './CreateForm';
import EditForm from './EditForm';

export default function IndexPage() {
  const model = useModel('ArticleCategory.model');

  const tableToolBarRender = () => {
    return (
      <Button
        onClick={() => {
          model.showCreate(true);
        }}
        icon={<PlusOutlined />}
        type="primary"
      >
        添加
      </Button>
    );
  };
  const tableColumns = [
    // {
    //     dataIndex: 'index',
    //     valueType: 'indexBorder',
    //     width: 50,
    //   },
    {
      title: '序号',
      dataIndex: 'id',
      //fixed: 'left',
      width: 50,
      hideInSearch: true,
    },
    {
      title: '标题',
      dataIndex: 'title',
      copyable: true,
      ellipsis: true,
    },
    {
      title: '关键字',
      dataIndex: 'keywords',
      hideInSearch: true,
      ellipsis: true,
    },
    {
      title: '描述',
      dataIndex: 'description',
      hideInSearch: true,
      ellipsis: true,
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
      sorter: true,
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
            onClick={() => {
              model.setData(record);
              model.showEdit(true);
            }}
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
      <CreateForm />
      <EditForm />
    </PageContainer>
  );
}
