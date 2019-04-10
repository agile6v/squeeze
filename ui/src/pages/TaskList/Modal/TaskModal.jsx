import React, { Component } from 'react';
import { Modal, message } from 'antd';
import { FormattedMessage, injectIntl } from 'react-intl';

import { request } from '../../../lib/common';
import BaseForm from '../../../components/BaseForm';

import styles from './taskModal.less';

class TaskModal extends Component {
    constructor(props) {
        super(props)
        this.state = {
            loading: false,
            formKeys: [
                {
                    col: { span: 12 },
                    keys: ['protocol', 'requests', 'concurrency', 'disableKeepAlive', 'duration']
                },
                {
                    col: { span: 12 },
                    keys: ['url', 'method', 'contentType', 'timeout']
                },
            ],
        }
    }
    componentDidMount() {
        // this.form.setFieldsValue({ Protocol: 'http' })
        // console.log('我是TaskModal的didmount')
        // console.log(this.form)
    }
    submitForm = () => {
        this.form.validateFields((errors, values) => {
            if (!errors) {
                this.create(values)
            }
        })
    }
    create = async (values) => {
        this.setState({
            loading: true
        })
        try {
            const res = await request.post('create', {
                json: {
                    protocol: values.protocol,
                    data: { ...values, maxResults: 1000000 }
                }
            }).json()
            const { data } = res;
            this.props.fetchList()
        } catch (err) {
            message.error(err.message)
            console.log('error: ', err)
        }
        this.props.onCancel()
        this.setState({
            loading: false
        })
    }
    handleFormChange = (key, value) => {
        if (key === 'protocol') {
            if (value === 'HTTP') {
                this.setState({
                    formKeys: [
                        {
                            col: { span: 12 },
                            keys: ['protocol', 'requests', 'concurrency', 'disableKeepAlive', 'duration']
                        },
                        {
                            col: { span: 12 },
                            keys: ['url', 'method', 'contentType', 'timeout']
                        },
                    ]
                })
            } else if (value === 'WEBSOCKET') {
                this.setState({
                    formKeys: [
                        {
                            col: { span: 12 },
                            keys: ['protocol', 'host', 'concurrency', 'duration', 'body']
                        },
                        {
                            col: { span: 12 },
                            keys: ['scheme', 'path', 'requests', 'timeout']
                        },
                    ]
                })
            }
        }
    }
    render() {
        const { visible, onCancel, intl } = this.props;
        const { formKeys, loading } = this.state;
        const requiredMessage = intl.formatMessage({ id: 'taskform.required' });
        const formItems = {
            protocol: {
                type: 'combo',
                label: intl.formatMessage({ id: 'taskform.protocol.label' }),
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                rules: ['required'],
                requiredMessage,
                options: [
                    { label: 'HTTP', value: 'HTTP' },
                    { label: 'HTTPS', value: 'HTTPS' },
                    { label: 'UDP', value: 'UDP' },
                    { label: 'TCP', value: 'TCP' },
                    { label: 'HTTP2', value: 'HTTP2' },
                    { label: 'WEBSOCKET', value: 'WEBSOCKET' },
                ],
                props: {
                    placeholder: intl.formatMessage({ id: 'taskform.protocol.placeholder' }),
                    showSearch: true,
                },
                defaultValue: 'HTTP',
            },
            method: {
                type: 'combo',
                label: intl.formatMessage({ id: 'taskform.Method.label' }),
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                rules: ['required'],
                requiredMessage,
                options: [
                    { label: 'GET', value: 'GET' },
                    { label: 'POST', value: 'POST' },
                    { label: 'PUT', value: 'PUT' },
                    { label: 'PATCH', value: 'PATCH' },
                    { label: 'DELETE', value: 'DELETE' },
                    { label: 'HEAD', value: 'HEAD' },
                ],
                props: {
                    placeholder: intl.formatMessage({ id: 'taskform.Method.placeholder' }),
                    showSearch: true,
                },
            },
            url: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.URL.label' }),
                requiredMessage,
                rules: ['required'],
                placeholder: intl.formatMessage({ id: 'taskform.URL.placeholder' }),
            },
            requests: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'number',
                label: intl.formatMessage({ id: 'taskform.Requests.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Requests.placeholder' }),
            },
            concurrency: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'number',
                label: intl.formatMessage({ id: 'taskform.Concurrency.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Concurrency.placeholder' }),
            },
            timeout: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'number',
                label: intl.formatMessage({ id: 'taskform.Timeout.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Timeout.placeholder' }),
            },
            contentType: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'combo',
                label: intl.formatMessage({ id: 'taskform.ContentType.label' }),
                rules: ['required'],
                requiredMessage,
                options: [
                    { label: 'application/json', value: 'application/json' },
                    { label: 'application/x-www-form-urlencoded', value: 'application/x-www-form-urlencoded' },
                    { label: 'audio/mpeg', value: 'audio/mpeg' },
                    { label: 'image/gif', value: 'image/gif' },
                    { label: 'multipart/form-data', value: 'multipart/form-data' },
                    { label: 'multipart/mixed', value: 'multipart/mixed' },
                    { label: 'text/css', value: 'text/css' },
                    { label: 'video/mpeg', value: 'video/mpeg' },
                    { label: 'application/msword', value: 'application/msword' },
                    { label: 'application/vnd.ms-excel', value: 'application/vnd.ms-excel' },
                    { label: 'application/zip', value: 'application/zip' },
                ],
                placeholder: intl.formatMessage({ id: 'taskform.ContentType.placeholder' }),
            },
            disableKeepAlive: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'radio',
                label: intl.formatMessage({ id: 'taskform.DisableKeepAlive.label' }),
                rules: ['required'],
                requiredMessage,
                options: [
                    { label: 'true', value: true },
                    { label: 'false', value: false },
                ],
            },
            //websocket options 
            scheme: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Scheme.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Scheme.placeholder' }),
            },
            host: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Host.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Host.placeholder' }),
            },
            path: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Path.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Path.placeholder' }),
            },
            body: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Body.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Body.placeholder' }),
            },
            duration: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'number',
                label: intl.formatMessage({ id: 'taskform.Duration.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Duration.placeholder' }),
            },
        }
        return (
            <Modal
                title={<FormattedMessage id="taskmodal.title" />}
                visible={visible}
                onCancel={onCancel}
                maskClosable={false}
                width={820}
                className={styles.TaskModal}
                onOk={this.submitForm}
            >
                <BaseForm
                    getForm={(form) => { this.form = form }}
                    formKeys={formKeys}
                    formItems={formItems}
                    onChange={this.handleFormChange}
                />
            </Modal>
        )
    }
}

export default injectIntl(TaskModal)