import React from 'react'
import Todo from './todo.jsx'
import {GET,POST} from './util/ajax'

export default class App extends React.Component{
  constructor(props){
    super(props)
    this.state = {
      todo_dat = [],
      sheet_dat = []
    }
  }

  // 初期設定
  componenntWillMount(){
    let json = GET("/api/todo/sheets")
    json.then( data => this.setState({sheet_dat: data}))
  }

  render() {
    return (
      <Todo
        sheet_meta={this.state.sheet_dat}
        todo_meta={this.state.todo_dat}
      />
    )
  }
}


//<Route path="/sheet/:id" component={Sheets} />  

/*
const Todos = props =>{
  console.log(props.match.params)
  return(
  <h1>HOME</h1>
)}
*/