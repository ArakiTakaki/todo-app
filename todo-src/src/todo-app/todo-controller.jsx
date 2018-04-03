import React from 'react'
import View from './todo-view.jsx'
import Create from './todo-create.jsx'

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
      data: [],
      sheet_id: this.props.match.params["id"]
    }
    this.serchTodos = this.serchTodos.bind(this)
  }

  componentWillReceiveProps(){
    const URI = CREATE_URI() + GET_ALL(this.state.sheet_id)
    this.serchTodos(URI)
  }

  render(){
   const sheet_id = this.state.sheet_id
    return (
      <div>
        <Create sheet_id={sheet_id}/>
        <View items={this.state.data}/>
      </div>
    )
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