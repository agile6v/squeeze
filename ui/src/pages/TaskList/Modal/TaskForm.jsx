import React, { Component } from 'react';
import { Form, Input, Button, Checkbox, Row, Col } from 'antd';
import { FormattedMessage, injectIntl } from 'react-intl';

const formItemLayout = {
    labelCol: { span: 4 },
    wrapperCol: { span: 8 },
};
class TaskForm extends Component {
    state = {
        checkNick: false,
    };

    check = () => {
        this.props.form.validateFields(
            (err) => {
                if (!err) {
                    console.info('success');
                }
            },
        );
    }

    handleChange = (e) => {
        this.setState({
            checkNick: e.target.checked,
        }, () => {
            this.props.form.validateFields(['nickname'], { force: true });
        });
    }

    render() {
        const { getFieldDecorator } = this.props.form;
        const { intl } = this.props;
        return (
            <div>
                <Form>

                </Form>
            </div>
        );
    }
}

export default Form.create({ name: 'TaskForm' })(injectIntl(TaskForm));