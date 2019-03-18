import React from "react";
import ReactDOM from "react-dom";
import { Layout, Menu, Icon } from 'antd';
import { IntlProvider } from 'react-intl';
import { BrowserRouter as Router, Route, Link as RouterLink, Switch } from 'react-router-dom'

import TaskListPage from './pages/TaskListPage';
import styles from './index.less';

const { Header, Footer } = Layout;

const App = () => {
  return (
    <Router>
      <Layout className={styles.layout}>
        <Header className={styles.header}>
          <div className={styles.logo}/>
          <Menu
            theme="dark"
            mode="horizontal"
          >
            <Menu.Item key="1">
              <RouterLink to="/tasks">
                <p>TASKS</p>
              </RouterLink>
            </Menu.Item>
          </Menu>
        </Header>
        <Switch>
            <Route exact path='/' component={Tasks} />
            <Route path='/tasks' component={Tasks} />
            {/* <Route path='/tasks/:tasksId' component={Projects} />
            <Route path='/about' component={Info} /> */}
          </Switch>
        <Footer className={styles.footer}>
          Created by tayir-m
        </Footer>
      </Layout>
    </Router>
  );
};


function Tasks({ match }) {
  return (
    <TaskListPage />
  )
}

ReactDOM.render(
  <IntlProvider locale="cn">
    <App />
  </IntlProvider>
  , document.getElementById("app"));
if (module.hot) {
  module.hot.accept()
}