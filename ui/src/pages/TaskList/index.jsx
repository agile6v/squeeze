import React, { Component } from 'react'
import { Provider, Subscribe } from 'unstated'

import TaskList from './TaskList'

import TaskListContainer from '../../containers/TaskListContainer'

export default class ProjectListPage extends Component {
  render () {
    return (
      <Provider>
        <Subscribe to={[TaskListContainer]}>
          {(taskListStore) => (
              <TaskList taskListStore={taskListStore} />
          )}
        </Subscribe>
      </Provider >
    )
  }
}