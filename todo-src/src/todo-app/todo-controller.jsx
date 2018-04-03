import React from 'react'
import View from './todo-view.jsx'

// スタブ 
//import todo_list from './dammy/test-todos.json'
// スタブ

// 情報群の取得処理
export default class TodoController extends React.Component{

  constructor(props){
    super(props)
    this.state = {
      data: []
    }
    this.serchTodos = this.serchTodos.bind(this)
  }

  componentWillReceiveProps(){
    const sheet_num = this.props.match.params["id"]
    const URI = CREATE_URI() + GET_ALL(sheet_num)
    this.serchTodos(URI)
  }

  render(){
    return <View items={this.state.data}/>
  }

  serchTodos(uri){
    fetch(
      uri,
      { method: 'GET' }
    ).then( res=> {
      return res.json()
    }).then( d =>{
      this.setState({data: d})
    }).catch( e => {
      console.log("todo_listのajax処理に失敗しました")
      console.log(e)
    })

  }
}

const CREATE_URI = item =>{
  return "/api/todo/todos?"
}
const GET_ALL = item =>{
  return "sheet_id=" + item
}