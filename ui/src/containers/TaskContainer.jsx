import { Container } from 'unstated'
import ky from 'ky'
import { message } from 'antd'

import { getAppRoot } from '../lib/common'

require("babel-polyfill");

const api = ky.extend({ prefixUrl: getAppRoot() + '/api/' })

export default class TaskContainer extends Container {
  constructor(props) {
    super(props)

    this.state = {
      taskList: [],
      loading: false
    }
  }

  async fetchList() {
    this.setState({
      loading: true
    })

    try {
      const res = await api.get('list').json()
      const { data } = res;
      this.setState({
        taskList: data,
        loading: false
      })
    } catch (err) {
      message.error(err.message)
      console.log('error: ', err)
    }
  }
}