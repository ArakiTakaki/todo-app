import React from 'react'
import View from './todo-view.jsx'
import Create from './todo-create.jsx'
import {GET} from './util/ajax'

// スタブ 
//import todo_list from './dammy/test-todos.json'
// スタブ

// 情報群の取得処理
// props
//    props.match.params["id"] sheet_id
export default class TodoController extends React.Component{

  constructor(props){
    super(props)
    this.state = {
      data: []
    }
  }

  componentWillReceiveProps(){
    const URI = CREATE_URI() + GET_ALL(this.props.match.params["id"])
    let request  = GET(URI)
    request.then( d => this.setState({ data : d }) )
  }

  render(){
   const sheet_id = this.props.match.params["id"]
    return (
      <div>
        <Create sheet_id={sheet_id}/>
        <View items={this.state.data}/>
      </div>
    )
  }
}

const CREATE_URI = item =>{
  return "/api/todo/todos?"
}
const GET_ALL = item =>{
  return "sheet_id=" + item
}