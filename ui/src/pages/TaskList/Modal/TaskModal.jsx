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
                    keys: ['Protocol', 'Requests', 'Concurrency', 'DisableKeepAlive', 'Duration']
                },
                {
                    col: { span: 12 },
                    keys: ['URL', 'Method', 'ContentType', 'Timeout']
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
                    Protocol: values.Protocol,
                    Data: { ...values, MaxResults: 1000000 }
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
        if (key === 'Protocol') {
            if (value === 'http') {
                this.setState({
                    formKeys: [
                        {
                            col: { span: 12 },
                            keys: ['Protocol', 'Requests', 'Concurrency', 'DisableKeepAlive', 'Duration']
                        },
                        {
                            col: { span: 12 },
                            keys: ['URL', 'Method', 'ContentType', 'Timeout']
                        },
                    ]
                })
            } else if (value === 'websocket') {
                this.setState({
                    formKeys: [
                        {
                            col: { span: 12 },
                            keys: ['Protocol', 'Host', 'Concurrency', 'Duration', 'Body']
                        },
                        {
                            col: { span: 12 },
                            keys: ['Scheme', 'Path', 'Requests', 'Timeout']
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
            Protocol: {
                type: 'combo',
                label: intl.formatMessage({ id: 'taskform.protocol.label' }),
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                rules: ['required'],
                requiredMessage,
                options: [
                    { label: 'http', value: 'http' },
                    { label: 'https', value: 'https' },
                    { label: 'UDP', value: 'UDP' },
                    { label: 'TCP', value: 'TCP' },
                    { label: 'http2', value: 'http2' },
                    { label: 'websocket', value: 'websocket' },
                ],
                props: {
                    placeholder: intl.formatMessage({ id: 'taskform.protocol.placeholder' }),
                    showSearch: true,
                },
                defaultValue: 'http',
            },
            Method: {
                type: 'combo',
                label: intl.formatMessage({ id: 'taskform.Method.label' }),
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                rules: ['required'],
                requiredMessage,
                options: [
                    { label: 'get', value: 'get' },
                    { label: 'post', value: 'post' },
                    { label: 'put', value: 'put' },
                    { label: 'patch', value: 'patch' },
                    { label: 'delete', value: 'delete' },
                    { label: 'head', value: 'head' },
                ],
                props: {
                    placeholder: intl.formatMessage({ id: 'taskform.Method.placeholder' }),
                    showSearch: true,
                },
            },
            URL: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.URL.label' }),
                requiredMessage,
                rules: ['required'],
                placeholder: intl.formatMessage({ id: 'taskform.URL.placeholder' }),
            },
            Requests: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Requests.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Requests.placeholder' }),
            },
            Concurrency: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Concurrency.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Concurrency.placeholder' }),
            },
            Timeout: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Timeout.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Timeout.placeholder' }),
            },
            ContentType: {
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
            DisableKeepAlive: {
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
            Scheme: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Scheme.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Scheme.placeholder' }),
            },
            Host: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Host.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Host.placeholder' }),
            },
            Path: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Path.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Path.placeholder' }),
            },
            Body: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
                label: intl.formatMessage({ id: 'taskform.Body.label' }),
                rules: ['required'],
                requiredMessage,
                placeholder: intl.formatMessage({ id: 'taskform.Body.placeholder' }),
            },
            Duration: {
                labelCol: { span: 8 },
                wrapperCol: { span: 14 },
                type: 'text',
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