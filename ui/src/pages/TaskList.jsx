import React, { Component } from 'react'
import { Table, Input, Icon } from 'antd';
import { IntlProvider, addLocaleData, FormattedMessage } from 'react-intl';

import ContentWrapper from '../components/ContentWrapper';

import styles from './TaskList.less';

export default class TaskList extends Component {
    constructor(props) {
        super(props)

        this.state = {
            taskList: [],
            loading: false,
        }
    }

    componentDidMount() {
        this.props.taskStore.fetchList()
    }

    render() {
        const { state: { taskList, loading } } = this.props.taskStore
        console.log(taskList)
        const columns = [{
            title: <FormattedMessage id="app.tasks" />,
            dataIndex: 'name',
            key: 'name',
            filters: [
                { text: 'Joe', value: 'Joe' },
                { text: 'Jim', value: 'Jim' },
            ],
            filteredValue: null,
            onFilter: (value, record) => record.name.includes(value),
            sorter: (a, b) => a.name.length - b.name.length,
            // sortOrder: sortedInfo.columnKey === 'name' && sortedInfo.order,
        }, {
            title: 'Age',
            dataIndex: 'age',
            key: 'age',
            sorter: (a, b) => a.age - b.age,
            // sortOrder: sortedInfo.columnKey === 'age' && sortedInfo.order,
        }, {
            title: 'Address',
            dataIndex: 'address',
            key: 'address',
            filters: [
                { text: 'London', value: 'London' },
                { text: 'New York', value: 'New York' },
            ],
            filteredValue: null,
            onFilter: (value, record) => record.address.includes(value),
            sorter: (a, b) => a.address.length - b.address.length,
            // sortOrder: sortedInfo.columnKey === 'address' && sortedInfo.order,
        }];
        const header = [
            (<div className={styles.header}><FormattedMessage id="app.tasks" /></div>),
            (<div className={styles.addButton}>new<Icon type="plus" /></div>)
        ]
        return (
            <ContentWrapper title="Tasks" header={header} >
                <Table
                    columns={columns}
                    size="small"
                    style={{ height: 1800 }}
                />
            </ContentWrapper>
        )
    }
}
