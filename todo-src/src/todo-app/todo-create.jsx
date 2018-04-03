import React from 'react'

// props
//    this.props.sheet_id
export default class TodoCreate extends React.Component{

  constructor(props){
    super(props)
    this.textParam = this.textParam.bind(this)
  }

  post(url,postParams){
    const option = {
      method: 'POST',
      body: JSON.stringify(postParams)
    }
    fetch( url,option)
      .then(res => {return res.json()} )
      .catch( e => error_log(e) )
  }

  render(){ 
    return 
      <section>
        <button onClick={this.textParam}>[+]</button>
        <input type="text" name="todo_content" ref="todo" />
      </section>
  }

  textParam(e){
    const url = '/api/todo/create_todo'
    const post = {
      sheet : this.props.sheet_id,
      data  : this.refs.todo.value
    }
    post(url,post)
  }

}
const error_log = (e) =>{
  console.log('todo-create:POST error')
  console.log(e)
}