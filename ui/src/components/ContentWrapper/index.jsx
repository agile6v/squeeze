import React from 'react';
import { Layout, Breadcrumb } from 'antd';

import styles from './index.less';

export default function ({ children, header }) {
    return (
        <Layout.Content className={styles.content}>
            <div className={styles.header}>{header}</div>
            {children}
        </Layout.Content>
    )
}