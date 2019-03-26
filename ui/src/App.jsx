import React, { Component } from "react";
import ReactDOM from "react-dom";
import { Layout, Menu, LocaleProvider, Button } from 'antd';
import { IntlProvider, addLocaleData, FormattedMessage } from 'react-intl';
import { BrowserRouter as Router, Route, Link as RouterLink, Switch } from 'react-router-dom'
import browserLang from 'browser-lang';
import moment from 'moment';
import en from 'react-intl/locale-data/en';
import zh from 'react-intl/locale-data/zh';
import zhCN from 'antd/lib/locale-provider/zh_CN';
import enUS from 'antd/lib/locale-provider/en_US';
import zh_CN from './zh_CN'     // import defined messages in Chinese
import en_US from './en_US'     // import defined messages in English

import TaskListPage from './pages/TaskList';
import styles from './index.less';

const { Header, Footer } = Layout;

addLocaleData([...en, ...zh]);

class App extends Component {
  constructor(props) {
    super(props)
    const lang = window.localStorage.getItem('lang') || browserLang({ languages: ['zh', 'en'], fallback: 'en' });
    moment.locale(lang)
    this.state = {
      lang,
    }
  }

  changeLang = () => {
    const lang = this.state.lang === 'en' ? 'zh' : 'en';
    window.localStorage.setItem('lang', lang)
    this.setState({ lang })
  }
  render() {
    const reactIntlLocale = {
      'en': en_US,
      'zh': zh_CN,
    }
    const antdLocale = {
      'en': enUS,
      'zh': zhCN,
    }
    const { lang } = this.state;
    return (
      <LocaleProvider locale={antdLocale[lang]}>
        <IntlProvider
          locale={lang} messages={reactIntlLocale[lang]}
        >
          <Router>
            <Layout className={styles.layout}>
              <Header className={styles.header}>
                <div className={styles.menuContainer}>
                  <div className={styles.logo} >SQUEEZE</div>
                  <div className={styles.menu}>
                    <RouterLink to="/tasks">
                      <p><FormattedMessage id="app.tasks" /></p>
                    </RouterLink>
                  </div>
                </div>
                <div className={styles.langButton}>
                  <span onClick={this.changeLang} >{lang === 'zh' ? 'English' : '中文'}</span>
                </div>
              </Header>
              <Switch>
                <Route exact path='/' component={TaskListPage} />
                <Route path='/tasks' component={TaskListPage} />
                {/* <Route path='/tasks/:tasksId' component={Projects} />
              <Route path='/about' component={Info} /> */}
              </Switch>
              <Footer className={styles.footer}>
                <a href="https://github.com/agile6v/squeeze">GitHub</a>
              </Footer>
            </Layout>
          </Router >
        </IntlProvider>
      </LocaleProvider>
    );
  }
};


ReactDOM.render(<App />, document.getElementById("app"));

if (module.hot) {
  module.hot.accept()
}