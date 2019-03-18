import React, { Component } from 'react'
import { Provider, Subscribe } from 'unstated'

import TaskList from './TaskList'

import TaskContainer from '../containers/TaskContainer'

export default class ProjectListPage extends Component {
  render () {
    return (
      <Provider>
        <Subscribe to={[TaskContainer]}>
          {(taskStore) => (
              <TaskList taskStore={taskStore} />
          )}
        </Subscribe>
      </Provider >
    )
  }
}