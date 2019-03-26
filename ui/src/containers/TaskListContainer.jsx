import { Container } from 'unstated'
import { message } from 'antd'

import { request } from '../lib/common';

export default class TaskContainer extends Container {
  constructor(props) {
    super(props)
    this.state = {
      taskList: [],
      loading: false
    }
  }

  fetchList = async () => {
    this.setState({
      loading: true
    })

    try {
      const res = await request.get('list').json()
      const { data } = res;
      this.setState({
        taskList: data,
        loading: false
      })
    } catch (err) {
      this.setState({ loading: false })
      message.error(err.message)
      console.log('error: ', err)
    }
  }
}