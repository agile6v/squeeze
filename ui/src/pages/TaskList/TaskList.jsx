import React, { Component } from 'react'
import { Table, message, Popconfirm, Modal } from 'antd';
import { FormattedMessage, injectIntl } from 'react-intl';
import moment from 'moment';

import ContentWrapper from '../../components/ContentWrapper';
import TaskModal from './Modal/TaskModal';
import { request } from '../../lib/common';

import styles from './TaskList.less';

class TaskList extends Component {
    constructor(props) {
        super(props)
        this.state = {
            taskModalVisible: false,
        }
    }

    componentDidMount() {
        this.props.taskListStore.fetchList()
    }

    toggleTaskModal = (taskModalVisible) => {
        this.setState({ taskModalVisible })
    }
    delete = async (ID) => {
        try {
            const res = await request.post('delete', {
                json: { ID }
            }).json()
            if (res.error) {
                message.error(res.error); return;
            }
            this.props.taskListStore.fetchList()
        } catch (err) {
            message.error(err.message)
            console.log('error: ', err)
        }
    }
    start = async (ID) => {
        try {
            const res = await request.post('start', {
                json: { ID }
            }).json()
            if (res.error) {
                message.error(res.error); return;
            }
            this.props.taskListStore.fetchList()
        } catch (err) {
            message.error(err.message)
            console.log('error: ', err)
        }
    }
    stop = async (ID) => {
        try {
            const res = await request.post('stop', {
                json: { ID }
            }).json()
            if (res.error) {
                message.error(res.error); return;
            }
            this.props.taskListStore.fetchList()
        } catch (err) {
            message.error(err.message)
            console.log('error: ', err)
        }
    }
    toDetails = async (record) => {
        console.log(JSON.parse(record.Response))
    }
    render() {
        const { state: { taskList, loading } } = this.props.taskListStore;
        const { taskModalVisible } = this.state;
        const { intl } = this.props;
        const resultMap = {
            0: <FormattedMessage id="tasklist.table.unfinished" />,
            1: <FormattedMessage id="tasklist.table.fail" />,
            2: <FormattedMessage id="tasklist.table.success" />,
        }
        const statustMap = {
            0: <FormattedMessage id="tasklist.table.notstarted" />,
            1: <FormattedMessage id="tasklist.table.resume" />,
            2: <FormattedMessage id="tasklist.table.running" />,
        }
        const deleteTitle = intl.formatMessage({ id: 'tasklist.table.deleteTitle' });
        const dataSource = taskList.map((task) => {
            let request;
            try {
                request = JSON.parse(task.Request)
            } catch (err) {
                message.error(err.message)
            }
            const { data } = request;
            const host = data.host || data.url
            return { ...task, protocol: data.protocol, host }
        })

        const columns = [{
            title: <FormattedMessage id="tasklist.table.Id" />,
            dataIndex: 'Id',
            key: 'Id',
            width: 80,
            sorter: (a, b) => a.Id - b.Id,
            // sortOrder: sortedInfo.columnKey === 'age' && sortedInfo.order,
        }, {
            title: <FormattedMessage id="tasklist.table.operation" />,
            dataIndex: 'operation',
            key: 'operation',
            width: 150,
            render: (value, record) => {
                return [record.Status === 1 || record.Status === 0 ? <Popconfirm key="1" title={<FormattedMessage id="tasklist.table.startTitle" />} onConfirm={() => { this.start(record.Id) }}>
                    <a ><FormattedMessage id="tasklist.table.start" /></a>
                </Popconfirm> : <Popconfirm key="1" title={<FormattedMessage id="tasklist.table.stopTitle" />} onConfirm={() => { this.stop(record.Id) }}>
                        <a ><FormattedMessage id="tasklist.table.stop" /></a>
                    </Popconfirm>,
                <span key="2"> | </span>,
                <a key="3" onClick={() => {
                    Modal.confirm({
                        title: deleteTitle,
                        content: `${intl.formatMessage({ id: 'tasklist.table.deleteContent' })}${record.Id}${intl.formatMessage({ id: 'tasklist.table.deleteContentTask' })}`,
                        onOk: () => {
                            this.delete(record.Id)
                        }
                    })
                }}><FormattedMessage id="tasklist.table.delete" /></a>,
                <span key="4"> | </span>,
                <a onClick={() => { this.toDetails(record) }} key="5"><FormattedMessage id="tasklist.table.details" /></a>]
            }
        }, {
            title: <FormattedMessage id="tasklist.table.protocol" />,
            dataIndex: 'protocol',
            key: 'protocol',
            width: 110,
            filters: [
                { text: 'HTTP', value: 'HTTP' },
                { text: 'HTTPS', value: 'HTTPS' },
                { text: 'UDP', value: 'UDP' },
                { text: 'TCP', value: 'TCP' },
                { text: 'HTTP2', value: 'HTTP2' },
                { text: 'WEBSOCKET', value: 'WEBSOCKET' },
            ],
            onFilter: (value, record) => { return record.protocol.indexOf(value) === 0 },
        }, {
            title: <FormattedMessage id="tasklist.table.host" />,
            dataIndex: 'host',
            key: 'host',
            filteredValue: null,
            onFilter: (value, record) => record.host.indexOf(value) === 0,
        }, {
            title: <FormattedMessage id="tasklist.table.status" />,
            dataIndex: 'Status',
            key: 'Status',
            width: 80,
            filters: [
                { text: <FormattedMessage id="tasklist.table.notstarted" />, value: 0 },
                { text: <FormattedMessage id="tasklist.table.resume" />, value: 1 },
                { text: <FormattedMessage id="tasklist.table.running" />, value: 2 },
            ],
            onFilter: (value, record) => { return record.Status === value },
            render: value => statustMap[value],
        }, {
            title: <FormattedMessage id="tasklist.table.result" />,
            dataIndex: 'Result',
            key: 'Result',
            width: 80,
            filters: [
                { text: <FormattedMessage id="tasklist.table.unfinished" />, value: 0 },
                { text: <FormattedMessage id="tasklist.table.fail" />, value: 1 },
                { text: <FormattedMessage id="tasklist.table.success" />, value: 2 },
            ],
            onFilter: (value, record) => { return record.Result === value },
            render: value => resultMap[value],
        }, {
            title: <FormattedMessage id="tasklist.table.createdTime" />,
            dataIndex: 'CreateAt',
            key: 'CreateAt',
            filteredValue: null,
            width: 180,
            onFilter: (value, record) => moment(record.CreateAt).format('YYYY-MM-DD HH:mm:ss').includes(value),
            // sorter: (a, b) => a.address.length - b.address.length,
            // sortOrder: sortedInfo.columnKey === 'address' && sortedInfo.order,
            render: (value) => {
                return moment(value).utc().format('YYYY-MM-DD HH:mm:ss')
            }
        }, {
            title: <FormattedMessage id="tasklist.table.updatedTime" />,
            dataIndex: 'UpdateAt',
            key: 'UpdateAt',
            filteredValue: null,
            width: 180,
            onFilter: (value, record) => moment(record.updatedTime).format('YYYY-MM-DD HH:mm:ss').includes(value),
            // sorter: (a, b) => a.address.length - b.address.length,
            render: (value) => {
                return moment(value).utc().format('YYYY-MM-DD HH:mm:ss')
            }
            // sortOrder: sortedInfo.columnKey === 'address' && sortedInfo.order,
        }];
        const header = [
            (<div key="1" className={styles.header}>
                <FormattedMessage id="tasklist.title" />
            </div>),
            (<div key="2" className={styles.addButton} onClick={() => { this.toggleTaskModal(true) }}>
                <FormattedMessage id="tasklist.new" />
            </div>)
        ]
        return (
            <ContentWrapper title="Tasks" header={header} >
                <Table
                    columns={columns}
                    size="small"
                    dataSource={dataSource}
                    rowKey="Id"
                    loading={loading}
                />
                {taskModalVisible && <TaskModal
                    fetchList={this.props.taskListStore.fetchList}
                    visible={taskModalVisible}
                    onCancel={() => { this.toggleTaskModal(false) }}
                />}
            </ContentWrapper>
        )
    }
}
export default injectIntl(TaskList)