import React from 'react'
import {POST} from './util/ajax'

const error_log = (e) =>{
  console.log('todo-create:POST error')
  console.log(e)
}

// props
//    this.props.sheet_id
export default class TodoCreate extends React.Component{

  constructor(props){
    super(props)
    this.reqest = this.reqest.bind(this)
    this.state = {
      review: true
    }

  }

  render(){ 
    return(
      <section>
        <button onClick={this.reqest}>[+]</button>
        <input type="text" name="todo_content" ref="todo" />
      </section>
      )
  }

  reqest(e){
    const url = '/api/todo/'+  this.props.sheet_id + '/create'
    const post = {
      "data"  : this.refs.todo.value
    }
    res = POST(url,post)
    res.then(data => console.log("POSTしたよ" + data))
    // let state  = !this.state.review
    // this.setState({review: state})
  }
}