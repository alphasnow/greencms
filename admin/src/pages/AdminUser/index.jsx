import { PlusOutlined } from '@ant-design/icons';
import { PageContainer, ProTable } from '@ant-design/pro-components';
import { Button, Popconfirm, Space } from 'antd';
import { useModel } from 'umi';
import CreateForm from './CreateForm';
import EditForm from './EditForm';

const RootAdminId = 1

export default function IndexPage() {
  const model = useModel('AdminUser.model');

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
    {
      title: '序号',
      dataIndex: 'id',
      hideInSearch: true,
      fixed: 'left',
      width: 50,
    },
    {
      title: '头像',
      dataIndex: 'avatar_url',
      hideInSearch: true,
      valueType: 'image',
    },
    {
      title: '账号',
      dataIndex: 'username',
      dataKey:"username",
    },
    {
      title: '姓名',
      dataIndex: 'realname',
      dataKey:"realname",
    },
    {
      title: '权限',
      dataIndex: 'access',
      dataKey:"access",
      valueType: 'select',
      request: async ()=>{
        const opts = (await model.getOptions()).data.access;
        return [...opts]
      }
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
      render: (_, record) => {
        if(record.id == RootAdminId){
          return null
        }
        return (        <Space.Compact size="small">
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
            title={`确定删除 ${record['username']} 吗?`}
            okText="确定"
            cancelText="取消"
            onConfirm={() => model.tableDelete(record)}
          >
            <Button>删除</Button>
          </Popconfirm>
        </Space.Compact>)
      },
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
