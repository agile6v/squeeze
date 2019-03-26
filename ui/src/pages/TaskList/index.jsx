import React, { Component } from 'react'
import { Provider, Subscribe } from 'unstated'

import TaskList from './TaskList'

import TaskListContainer from '../../containers/TaskListContainer'
const taskListContainer = new TaskListContainer()
export default class TaskListPage extends Component {
  render () {
    return (
      <Provider inject={[taskListContainer]}>
        <Subscribe to={[TaskListContainer]}>
          {(taskListStore) => (
              <TaskList taskListStore={taskListStore} />
          )}
        </Subscribe>
      </Provider >
    )
  }
}