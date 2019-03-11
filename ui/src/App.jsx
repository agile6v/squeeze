import React from "react";
import ReactDOM from "react-dom";
import { Layout, Menu, Icon } from 'antd';
import { IntlProvider } from 'react-intl';
import { BrowserRouter as Router, Route, Link as RouterLink, Switch } from 'react-router-dom'

import TaskListPage from './components/TaskListPage';
import './index.less';

const { Header, Content, Footer } = Layout;

const App = () => {
  return (
    <Router>
      <Layout className="layout">
        <Header className="header">
          <div className="logo" />
          <Menu
            theme="dark"
            mode="horizontal"
            defaultSelectedKeys={['1']}
          >
            <Menu.Item key="1">
              <RouterLink to="/tasks">
                <p>TASKS</p>
              </RouterLink>
            </Menu.Item>
          </Menu>
        </Header>
        <Content style={{ background: '#fff', padding: 24, margin: 0, minHeight: 600 }}>
          <Switch>
            <Route exact path='/' component={Tasks} />
            <Route path='/tasks' component={Tasks} />
            {/* <Route path='/tasks/:tasksId' component={Projects} />
            <Route path='/about' component={Info} /> */}
          </Switch>
        </Content>
        <Footer style={{ textAlign: 'center' }}>
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
  <IntlProvider locale="en">
    <App />
  </IntlProvider>
  , document.getElementById("app"));
if (module.hot) {
  module.hot.accept()
}