import { Container } from 'unstated'
import ky from 'ky'
import { message } from 'antd'

import { getAppRoot } from '../lib/common'

require("babel-polyfill");

const api = ky.extend({ prefixUrl: getAppRoot() + '/api/' })

export default class TaskContainer extends Container {
  constructor (props) {
    super(props)

    this.state = {
      taskList: null,
      loading: false
    }
  }

  async fetchList () {
    this.setState({
      loading: true
    })

    try {
      const taskList = await api.get('list').json()

      this.setState({
        taskList,
        loading: false
      })
    } catch (err) {
      message.error(err.message)
      console.log('error: ', err)
    }
  }
}