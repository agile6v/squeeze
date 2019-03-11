import React, { Component } from 'react'

export default class TaskList extends Component {
    constructor(props) {
        super(props)

        this.state = {
            taskList: [],
            loading: false,
        }
    }

    componentDidMount() {
        this.props.taskStore.fetchList()
    }

    render() {
        const { state: { taskList, loading } } = this.props.taskStore

        return (
            <div>{taskList}test</div>
        )
    }
}
